package main

import (
    "fmt"
    "net"
    "sync"
)

var clients []net.Conn
var mutex sync.Mutex

func startServer() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("error starting server")
        return
    }
    fmt.Println("server started on port 8080")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("error accepting connection")
            continue
        }
        fmt.Println("new client connected")
        mutex.Lock()
        clients = append(clients, conn)
        mutex.Unlock()
    }
}