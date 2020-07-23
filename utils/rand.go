package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

func RandInt(min, max int) int {
	diff := big.NewInt(int64(max - min))

	number, err := rand.Int(rand.Reader, diff)
	if err != nil {
		panic(err.Error())
	}

	return int(number.Int64()) + min
}

func RandBytes(length int) []byte {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	return bytes
}

func RandHexString(length int) string {
	return hex.EncodeToString(RandBytes(length))
}
