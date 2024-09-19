package main

func handleRedisCommands(cmd string) []byte {
	switch cmd {
	case "PING":
		return []byte("+PONG\r\n")
	}
	return []byte("")
}
