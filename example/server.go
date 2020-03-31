package example

func add(a, b int) int {
	return a + b
}

func main() {
	addr := "127.0.0.1:8080"
	srv := tinyRPC.NewServer(addr)
	srv.Register("add", add)
	go srv.Run()

	for {
	}
}
