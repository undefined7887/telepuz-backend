package rand

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

func Int(min, max int) int {
	diff := big.NewInt(int64(max - min))

	number, err := rand.Int(rand.Reader, diff)
	if err != nil {
		panic(err.Error())
	}

	return int(number.Int64()) + min
}

func Hex(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	return hex.EncodeToString(bytes)
}
