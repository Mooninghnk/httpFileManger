package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func crtDir(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	crtDir(randToken())
}
