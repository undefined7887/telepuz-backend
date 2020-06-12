package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/undefined7887/telepuz-backend/cache"
	"github.com/undefined7887/telepuz-backend/format"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/services/auth"
	"github.com/undefined7887/telepuz-backend/services/users"
	"github.com/undefined7887/telepuz-backend/services/base/endpoint"
	"github.com/vmihailenco/msgpack/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

var clientsPool = cache.NewPool()
var userPool = cache.NewPool()

var authService = auth.NewService(userPool)
var usersService = users.NewService(userPool)

func main() {
	listenWebSockets(os.Args[1])
}

func listenWebSockets(addr string) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(writer, req, nil)
		if err != nil {
			fmt.Println("Failed to accept connection:", err.Error())
			return
		}

		go handleConn(conn)
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Failed to start websocket server:", err.Error())
	}
}

type Client struct {
	Id   string
	Conn *websocket.Conn
}

func (c *Client) GetId() string {
	return c.Id
}

func handleConn(conn *websocket.Conn) {
	fmt.Println("New client:", conn.RemoteAddr())

	client := &Client{Id: rand.Hex(format.IdLength), Conn: conn}
	clientsPool.Add(client)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to receive request:", err.Error())
			return
		}

		decoder := msgpack.NewDecoder(bytes.NewReader(msg)).UseJSONTag(true)

		reqInfo := &endpoint.RequestInfo{}
		if err := decoder.Decode(reqInfo); err != nil {
			fmt.Println("Failed to unmarshal request info:", err.Error())
			continue
		}

		// TODO: Perform check on method name

		method := getMethod(reqInfo.MethodName)
		if method == nil {
			fmt.Println("Failed to get method:", reqInfo.MethodName)
			continue
		}

		req := method.NewRequest()
		if err := decoder.Decode(req); err != nil {
			fmt.Println("Failed to unmarshal request:", err.Error())
			continue
		}

		fmt.Printf("New request:\n%s\n%s\n", reqInfo, req)
		go handleRequest(client, method, req, reqInfo)
	}
}

func getMethod(name string) endpoint.Method {
	service := strings.Split(name, ".")[0]

	switch service {
	case "auth":
		return authService.GetMethod(name)

	case "users":
		return usersService.GetMethod(name)

	default:
		return nil
	}
}

func handleRequest(client *Client, method endpoint.Method, req *endpoint.Request, reqInfo *endpoint.RequestInfo) {
	startTimePoint := time.Now()

	res := method.Call(req)
	resInfo := &endpoint.ResponseInfo{MethodName: reqInfo.MethodName}

	buffer := bytes.NewBuffer(nil)
	encoder := msgpack.NewEncoder(buffer).UseJSONTag(true)

	if err := encoder.Encode(resInfo); err != nil {
		fmt.Println("Failed to marshall response info:", err.Error())
		return
	}

	if err := encoder.Encode(res); err != nil {
		fmt.Println("Failed to marshall response:", err.Error())
		return
	}

	if err := client.Conn.WriteMessage(websocket.BinaryMessage, buffer.Bytes()); err != nil {
		fmt.Println("Failed to send response:", err.Error())
		return
	}

	fmt.Printf("Sent response (time=%d ms):\n%s\n%s\n", time.Since(startTimePoint).Milliseconds(), resInfo, res)
}
