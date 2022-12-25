package main

import (
	"encoding/binary"
	"io"
	"net"
	"os"
)

type FPMHeader struct {
	Version     uint8
	MessageType uint8
	MessageLen  uint16
}

func handleConnection(conn net.Conn) {
	for {
		h := FPMHeader{}
		binary.Read(conn, binary.BigEndian, &h.Version)
		binary.Read(conn, binary.BigEndian, &h.MessageType)
		binary.Read(conn, binary.BigEndian, &h.MessageLen)

		if h.Version != 1 {
			panic("Unsupported FPM frame version")
		}

		if h.MessageType != 1 {
			panic("Unsupported FPM frame type")
		}

		n, err := io.CopyN(os.Stdout, conn, int64(h.MessageLen-4))
		if err != nil {
			panic(err)
		}

		if n != int64(h.MessageLen-4) {
			panic("Couldn't read entire message")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":2620")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		handleConnection(conn)
	}
}
