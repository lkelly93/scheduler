package main

import "github.com/lkelly93/scheduler/internal/server"

func main() {
    server := server.NewHTTPServer()
    server.Start(3000)
}
