package v1

import (
	"context"
	"log"
)

type Server struct {
	ServerName string
}

func (s *Server) ServiceImp(_ context.Context, req *Request) (*Response, error) {
	log.Printf("%s request haveth", s.ServerName)
	retVal := &Response{Message: req.Message}
	return retVal, nil
}

func (s *Server) mustEmbedUnimplementedServiceServer() {}
