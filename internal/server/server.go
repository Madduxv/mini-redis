package server

import (
  "fmt"
  "net"
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

func (s *Server) handleConnection(conn net.Conn, srv *Server) {
  defer conn.Close()
  buffer := make([]byte, 1024)
  n, err := conn.Read(buffer)
  if err != nil {
    fmt.Println("Error reading from connection:", err)
    return
  }
  command, args, err := protocol.ParseRESP(buffer[:n])

  switch command {
  case "HSET":
    if len(args) != 3 {
      conn.Write([]byte("ERR wrong number of arguments for 'HSET' command\n"))
    }
    srv.HandleHSet(args[0], args[1], args[2])
    conn.Write([]byte("OK\n"))
  case "HGET":
    if len(args) != 2 {
      conn.Write([]byte("ERR wrong number of arguments for 'HGET' command\n"))
    }
    value, exists := srv.HandleHGet(args[0], args[1])
    if exists {
      conn.Write([]byte(value + "\n"))
    } else {
      conn.Write([]byte("(nil)\n"))
    }
  default:
    conn.Write([]byte("ERR unknown command '" + command + "'\n"))
  }

}
