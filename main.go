package main

import "github.com/lkelly93/scheduler/internal/server"

//This starts a scheduler server. The binary created from go build in this repo
//will run a scheduler server but will not incase it in a container.
func main() {
	server := server.NewHTTPServer()
	server.Start(3000)
}
