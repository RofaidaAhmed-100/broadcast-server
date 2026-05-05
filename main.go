package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("write start or connect")
        return
    }

    if os.Args[1] == "start" {
        startServer()
    } else if os.Args[1] == "connect" {
        connectClient()
    }
}