package main

import (
	"fmt"

	"github.com/xxx0624/tinyRPC"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func main() {
	addr := "127.0.0.1:8080"
	srv := tinyRPC.NewServer(addr)
	srv.Register("add", add)
	fmt.Println("service start...")
	go srv.Run()

	for {
	}
}
