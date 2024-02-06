package main

import (
	"fmt"

	"github.com/gordjw/chirpy/server"
)

func main() {
	fmt.Println("Chirpy starting up...")

	server.Run()
}
