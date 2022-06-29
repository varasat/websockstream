package main

import "websockstream/socket"

func main() {
	booter := socket.NewBooter()
	booter.Run()
}

