// Copyright (c) 2023 Valters Jansons

package main

import (
	"flag"
	"io"
	"log"
	"net"

	pb "github.com/sigv/grpcopen/grpcopen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	addr = flag.String("addr", ":8088", "Address to listen on")
)

type server struct {
	pb.UnimplementedBaseServer
}

func (s *server) Foobar(stream pb.Base_FoobarServer) error {
	// right away, return gRPC error response
	return status.Errorf(codes.Unimplemented, "method Foobar not implemented")
}

func (s *server) Ping(stream pb.Base_PingServer) error {
	// keep waiting for unlimited number of ping requests in one stream
	for {
		in, err := stream.Recv()
		// if client closes connection, ok
		if err == io.EOF {
			return nil
		}
		// if other unexpected condition, error
		if err != nil {
			return err
		}

		// get request content, and send it back in response
		content := in.GetContent()
		log.Printf("recv ping: %v", content)
		if err := stream.Send(&pb.PingResponse{Content: content}); err != nil {
			return err
		}
	}
}

func main() {
	// parse cli arguments
	flag.Parse()

	// start listening
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// set up gRPC basics
	s := grpc.NewServer()
	pb.RegisterBaseServer(s, &server{})

	// Serve() is blocking; log before it
	log.Printf("listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
