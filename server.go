package forzatelemetry

import (
	"fmt"
	"net"
)

// UDP server for listening telemetry data.
type Server struct {
	addr *net.UDPAddr
}

// Listens for UDP packets and invokes the callback function.
func (s *Server) Listen(f func(packet Packet)) error {
	conn, err := net.ListenUDP("udp", s.addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		f(buf[:n])
	}
}

// Creates a new server on a given port.
func NewServerPort(port int) (*Server, error) {
	return NewServerString(fmt.Sprintf(":%d", port))
}

// Creates a new server on a given address string.
func NewServerString(s string) (*Server, error) {
	addr, err := net.ResolveUDPAddr("udp", s)
	if err != nil {
		return nil, err
	}
	return NewServerAddr(addr)
}

// Creates a new server on a given address.
func NewServerAddr(addr *net.UDPAddr) (*Server, error) {
	return &Server{
		addr: addr,
	}, nil
}
