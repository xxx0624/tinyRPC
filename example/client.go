package main

import (
	"fmt"
	"log"
	"net"

	"github.com/xxx0624/tinyRPC"
)

func main() {
	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("connect %s error: %v\n", addr, err)
		return
	}

	cli := tinyRPC.NewClient(conn)
	var addFn func(int, int) (int, error)
	cli.TransformMethod("add", &addFn)
	answer, err := addFn(2, 20)
	fmt.Printf("answer is %v\n", answer)
	fmt.Printf("error is %v\n", err)
}
