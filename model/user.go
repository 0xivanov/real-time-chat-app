package model

import "net"

type User struct {
	ID   uint64
	Conn net.Conn
	// id
	// address
	// name
}
