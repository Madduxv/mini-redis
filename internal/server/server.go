package server

import (
    "fmt"
    "net"
)

type Server struct {
    // Define your server fields, like address and connections
    Addr string
}

func New() *Server {
    return &Server{
        Addr: "127.0.0.1:6379", // Example address
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
    // Handle client connection
    defer conn.Close()
    // Example: simple echo
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from connection:", err)
        return
    }
    conn.Write(buffer[:n])
}
