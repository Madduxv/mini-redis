package server

import (
    "fmt"
    "net"
    "github.com/Madduxv/mini-redis/internal/protocol"
)

type Server struct {
    // Define server fields, like address and connections
    Addr string
}

func New() *Server {
    return &Server{
        Addr: "127.0.0.1:6379", // Redis runs on port 6379
    }
}

func (s *Server) Start() {
    listener, err := net.Listen("tcp", s.Addr)
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    fmt.Println("Server started on", s.Addr)

    // Accept connections and handle them
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go s.handleConnection(conn)
    }
}

func (s *Server) handleConnection(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from connection:", err)
        return
    }
    /* commands, parsed_data, err :=  */protocol.ParseRESP(buffer[:n])
    // conn.Write([]byte(commands))
    // conn.Write(buffer[:n])
}
