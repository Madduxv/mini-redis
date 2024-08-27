package server

import (
	"fmt"
	"github.com/Madduxv/mini-redis/internal/protocol"
	"net"
	"strconv"
	"strings"
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

		case "HREM":
			if len(args) != 1 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'HREM' command\r\n"))
				continue
			}
			srv.HandleHRem(args[0])
			conn.Write([]byte("OK\r\n"))

		case "HDEL":
			if len(args) != 2 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'HDEL' command\r\n"))
				continue
			}
			if srv.HandleHDel(args[0], args[1]) {
				conn.Write([]byte("OK\r\n"))
				continue
			} else {
				conn.Write([]byte("\r\nEntry does not exist\r\n"))
			}

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

		case "RPUSH":
			if len(args) != 3 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'RPUSH' command\r\n"))
				continue
			}
			srv.HandleRPush(args[0], args[1], args[2])
			conn.Write([]byte("OK\r\n"))

		case "SADD":
			if len(args) != 2 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'SADD' command\r\n"))
				continue
			}
			srv.HandleSAdd(args[0], args[1])
			conn.Write([]byte("OK\r\n"))

		case "SREM":
			if len(args) != 2 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'SREM' command\r\n"))
				continue
			}
			removed := srv.HandleSRem(args[0], args[1])
			if removed == 1 {
				conn.Write([]byte("1\r\n"))
			} else {
				conn.Write([]byte("0\r\n"))
			}

		case "SGET":
			if len(args) != 1 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'SGET' command\r\n"))
				continue
			}
			value, exists := srv.HandleSGet(args[0])
			if exists {
				conn.Write([]byte(strings.Join(value, ",") + "\r\n"))
			} else {
				conn.Write([]byte("(nil)\r\n"))
			}

		case "LRANGE":
			if len(args) != 4 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'LRANGE' command: Expected 4, but found " + string(len(args)) + "\r\n"))
				continue
			}
			start, err := strconv.Atoi(args[2])
			end, err1 := strconv.Atoi(args[3])
			if err != nil || err1 != nil {
				conn.Write([]byte("\r\nERR Invalid Format: Either start or end value is not a number\r\n"))
			}
			value := srv.HandleLRange(args[0], args[1], start, end)
			conn.Write([]byte("\r\n" + strings.Join(value, ",") + "\r\n"))

		case "LCLEAR":
			if len(args) != 2 {
				conn.Write([]byte("\r\nERR wrong number of arguments for 'LRANGE' command: Expected 4, but found " + string(len(args)) + "\r\n"))
				continue
			}
			conn.Write([]byte(srv.HandleLClear(args[0], args[1]) + "\r\n"))

		default:
			conn.Write([]byte("\r\nERR unknown command '" + command + "'\r\n"))
		}

	}
}
