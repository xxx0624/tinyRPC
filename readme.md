# Tiny RPC
A tiny RPC implemented by Golang

## RPC
A program requests a service from a program located in another part of address space / another server.

## Process
1. Client side

    1. the client will include the function name and arguments separately
    2. encode these information
    3. send to the server side
2. Server side

    1. retrive the function name and arguments fromt the client side
    2. use reflect package to call the specified method
    3. encode the running result of the method, and send back to the client side


## Test
```
// start server
cd example
go run server.go

// start client and make a remote call
cd example
go run client.go
```

## Docker
```
cd /path/to/this/tinyRPC
docker build -t tinyrpc:test .
docker run -p 127.0.0.1:8080:8080 --name rpc -d tinyrpc:test
```

## TODO
1. Methods in Golang usually will include the error information in the return values. So in the implementation of the tinyRPC, all methods registered in the server side are required to return the error as the last one in the return list.

    Like:
`func methodExample(...) (..., error)`

2. Test Docker server

## Reference 
[rpc blog](https://blog.jiahonzheng.cn/2018/11/25/Golang%20%E7%AE%80%E6%98%93%20RPC%20%E6%A1%86%E6%9E%B6/)