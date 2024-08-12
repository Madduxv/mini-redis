package server

import (
	"fmt"
	"net"
	"strings"

	"github.com/Madduxv/mini-redis/internal/protocol"
)

// StartServer starts the Redis server.
func StartServer() {
  ln, err := net.Listen("tcp", ":6379")
  if err != nil {
    fmt.Println("Error starting server:", err)
    return
  }
  defer ln.Close()

  srv := NewServer()

  fmt.Println("Server is running on port 6379...")
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error accepting connection:", err)
      continue
    }

    go handleConnection(conn, srv)
  }
}

func handleConnection(conn net.Conn, srv *Server) {
  defer conn.Close()
  buffer := make([]byte, 1024)
  for {
    n, err := conn.Read(buffer)
    if err != nil {
      fmt.Println("Error reading from connection:", err)
      return
    }
    fmt.Println(strings.TrimSpace(string(buffer[:n])))
    command, args, err := protocol.ParseRESP(strings.TrimSpace(string(buffer[:n])))

    fmt.Printf("Command: %s \r\n", command)
    fmt.Printf("Args: %s \r\n", args)

    switch command {
    case "HSET":
      if len(args) != 3 {
        conn.Write([]byte("\nERR wrong number of arguments for 'HSET' command\n"))
        continue
      }
      srv.HandleHSet(args[0], args[1], args[2])
      conn.Write([]byte("\r\nOK\r\n"))
    case "HGET":
      if len(args) != 2 {
        conn.Write([]byte("\nERR wrong number of arguments for 'HGET' command\n"))
        continue
      }
      value, exists := srv.HandleHGet(args[0], args[1])
      if exists {
        conn.Write([]byte(value + "\n"))
      } else {
        conn.Write([]byte("\r\n(nil)\r\n"))
      }
    default:
      conn.Write([]byte("\nERR unknown command '" + command + "'\n"))
    }

  }
}
