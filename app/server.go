package main

import (
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	handleGenericError(err, "Failed to bind to port 6379")
	defer l.Close()

	conn, err := l.Accept()
	handleGenericError(err, "Error accepting connection: ")
	defer conn.Close()

	// data, err := io.ReadAll(conn)
	handleGenericError(err, "Error reading data:")

	// response := handleRedisCommands(string(data))

	response := handleRedisCommands("PING")

	_, err = conn.Write(response)
	handleGenericError(err)
}
