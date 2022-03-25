package grpcserver

import (
	"github.com/swanden/rentateam/api/grpcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	defaultPort = "50051"
)

type Server struct {
	server *grpc.Server
	host   string
	notify chan error
}

func New(postServer grpcpb.PostsServer, opts ...Option) (*Server, error) {
	s := &Server{
		server: grpc.NewServer(),
		host:   net.JoinHostPort("", defaultPort),
	}

	for _, opt := range opts {
		opt(s)
	}

	reflection.Register(s.server)
	grpcpb.RegisterPostsServer(s.server, postServer)

	err := s.Start()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.host)
	if err != nil {
		return nil
	}

	go func() {
		s.notify <- s.server.Serve(lis)
		close(s.notify)
	}()

	return nil
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
