package main

import "github.com/Madduxv/mini-redis/internal/server"

func main() {
  srv := server.New()
  srv.Start()

}

