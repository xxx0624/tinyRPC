package tinyRPC

import (
	"errors"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

func (c *Client) TransformMethod(methodName string, methodPtr interface{}) {
	container := reflect.ValueOf(methodPtr).Elem()
	methodOutputNum := container.Type().NumOut()

	fn := func(reqData []reflect.Value) []reflect.Value {
		// build error handler
		errorHandler := func(err error) []reflect.Value {
			values := make([]reflect.Value, methodOutputNum)
			for i := 0; i < len(values)-1; i++ {
				values[i] = reflect.Zero(container.Type().Out(i))
			}
			values[len(values)-1] = reflect.ValueOf(&err).Elem()
			return values
		}

		transport := NewTransport(c.conn)
		// load arguments
		args := make([]interface{}, 0, len(reqData))
		for i := range reqData {
			args = append(args, reqData[i].Interface())
		}
		err := transport.Send(Data{Name: methodName, Args: args})
		if err != nil {
			return errorHandler(err)
		}

		// receive response from server
		resp, err := transport.Receive()
		if err != nil {
			return errorHandler(err)
		}
		if resp.Err != "" {
			return errorHandler(errors.New(resp.Err))
		}

		if len(resp.Args) == 0 {
			resp.Args = make([]interface{}, methodOutputNum)
		}
		values := make([]reflect.Value, methodOutputNum)
		for i := 0; i < methodOutputNum; i++ {
			if i == methodOutputNum-1 {
				values[i] = reflect.Zero(container.Type().Out(i))
			} else {
				if resp.Args[i] == nil {
					values[i] = reflect.Zero(container.Type().Out(i))
				} else {
					values[i] = reflect.ValueOf(resp.Args[i])
				}
			}
		}
		return values
	}

	container.Set(reflect.MakeFunc(container.Type(), fn))
}
