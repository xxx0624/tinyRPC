package tinyRPC

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

type Server struct {
	addr string
	fns  map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		fns:  make(map[string]reflect.Value),
	}
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("listening on %s, error: %v\n", s.addr, err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		go func() {
			transport := NewTransport(conn)
			for {
				req, err := transport.Receive()
				if err != nil {
					log.Printf("receive data, error: %v\n", err)
					return
				}
				fn, ok := s.fns[req.Name]
				if !ok {
					// fn doesn'e exist
					e := fmt.Sprintf("func %s not exist", req.Name)
					if err = transport.Send(Data{Name: req.Name, Err: e}); err != nil {
						log.Printf("transport send error: %v\n", err)
					}
					return
				}

				args := make([]reflect.Value, len(req.Args))
				for i := range req.Args {
					args[i] = reflect.ValueOf(req.Args[i])
				}

				output := fn.Call(args)
				values := make([]interface{}, len(output)-1)
				for i := 0; i < len(output)-1; i++ {
					values[i] = output[i].Interface()
				}
				var e string
				if _, ok := output[len(output)-1].Interface().(error); !ok {
					e = ""
				} else {
					e = output[len(output)-1].Interface().(error).Error()
				}
				err = transport.Send(Data{Name: req.Name, Args: values, Err: e})
				if err != nil {
					log.Printf("transport send error: %v\n", err)
				}
			}
		}()
	}
}

func (s *Server) Register(name string, fn interface{}) {
	if _, ok := s.fns[name]; ok {
		return
	}
	s.fns[name] = reflect.ValueOf(fn)
}
