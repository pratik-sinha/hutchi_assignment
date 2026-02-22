package main

import (
	"assignment/internal/server"
	"assignment/pkg/db"
)

func main() {
	// utils.LoadEnv()
	conn := db.ConnectPostGres()
	s := server.NewServer(conn)
	s.Run()
}
