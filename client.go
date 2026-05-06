package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func connectClient() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("error connecting to server")
        return
    }
    fmt.Println("connected to server")
    defer conn.Close()

    go receiveMessages(conn)

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        message := scanner.Text()
        conn.Write([]byte(message))
    }
}

func receiveMessages(conn net.Conn) {
    buffer := make([]byte, 1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("disconnected from server")
            return
        }
        fmt.Println(string(buffer[:n]))
    }
}