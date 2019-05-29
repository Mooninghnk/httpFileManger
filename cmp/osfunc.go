package osw

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

//RandToken creates a random string
func RandToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//CrtDir creates a folder
func CrtDir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

//getCurrentTime returns current time on the http server
func getCurentTime() string {
	return time.Now().String()
}
