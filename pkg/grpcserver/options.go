package grpcserver

import "net"

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.host = net.JoinHostPort("", port)
	}
}
