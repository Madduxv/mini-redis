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
  conn.Write([]byte("\r\n"))
  buffer := make([]byte, 1024)
  for {
    n, err := conn.Read(buffer)
    if err != nil {
      fmt.Println("Error reading from connection:", err)
      return
    }
    // fmt.Println(strings.TrimSpace(string(buffer[:n])))
    command, args, err := protocol.ParseRESP(strings.TrimSpace(string(buffer[:n])))

    // fmt.Printf("Command: %s \r\n", command)
    // fmt.Printf("Args: %s \r\n", args)

    switch command {
    case "HSET":
      if len(args) != 3 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HSET' command\r\n"))
        continue
      }
      srv.HandleHSet(args[0], args[1], args[2])
      conn.Write([]byte("OK\r\n"))

    case "HSETLIST":
      if len(args) != 3 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HSETLIST' command\r\n"))
        continue
      }
      srv.HandleHSetList(args[0], args[1], strings.Split(strings.TrimSpace(string(args[2])), ", "))
      conn.Write([]byte("OK\r\n"))

    case "HGET":
      if len(args) != 2 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HGET' command\r\n"))
        continue
      }
      value, exists := srv.HandleHGet(args[0], args[1])
      if exists {
        conn.Write([]byte(value + "\r\n"))
      } else {
        conn.Write([]byte("(nil)\r\n"))
      }

    case "HGETLIST":
      if len(args) != 2 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HGETLIST' command\r\n"))
        continue
      }
      value, exists := srv.HandleHGetList(args[0], args[1])
      if exists {
        conn.Write([]byte(strings.Join(value, ",") + "\r\n"))
      } else {
        conn.Write([]byte("(nil)\r\n"))
      }

    case "HREMOVE":
      if len(args) != 1 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HREMOVELIST' command\r\n"))
        continue
      }
      srv.HandleHRemove(args[0])
      conn.Write([]byte("OK\r\n"))

    case "HREMOVESTRINGFIELD":
      if len(args) != 2 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HREMOVESTRINGFIELD' command\r\n"))
        continue
      }
      if srv.HandleHRemoveStringField(args[0], args[1]) {
        conn.Write([]byte("OK\r\n"))
        continue
      } else {
        conn.Write([]byte("\r\nEntry does not exist\r\n"))
      }

    case "HREMOVELISTFIELD":
      if len(args) != 2 {
        conn.Write([]byte("\r\nERR wrong number of arguments for 'HREMOVELISTFIELD' command\r\n"))
        continue
      }
      if srv.HandleHRemoveListField(args[0], args[1]) {
        conn.Write([]byte("OK\r\n"))
        continue
      } else {
        conn.Write([]byte("\r\nEntry does not exist\r\n"))
      }

    default:
      conn.Write([]byte("\r\nERR unknown command '" + command + "'\r\n"))
    }

  }
}
