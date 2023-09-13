// Copyright (c) 2023 Valters Jansons

package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/sigv/grpcopen/grpcopen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	defaultContent = "Lorem ipsum"
)

var (
	addr    = flag.String("addr", "localhost:8088", "Address to connect to")
	content = flag.String("content", defaultContent, "Ping request content")
)

func foobar(client pb.BaseClient, ctx context.Context) {
	// open stream for foobar
	log.Print("starting foobar")
	foobar_s, err := client.Foobar(ctx)
	if err != nil {
		log.Fatalf("failed to create foobar stream: %v", err)
	}

	// send foobar request and then close our sending side
	if err := foobar_s.Send(&pb.FoobarRequest{}); err != nil {
		log.Fatalf("failed to send foobar: %v", err)
	}
	// When the server receives the request, it will return a
	// error response, and RST_STREAM from its side.
	//
	// We sleep for a second, to ensure that the server does this.
	// Otherwise, we may be too fast with our CloseSend() which
	// would deliver our own END_STREAM.
	//
	// When sending request directly, client gets gRPC the error
	// response that is expected - a formatted gRPC error.
	//
	// Through HAProxy, this currently causes an HTTP 502 response
	// to be returned instead. Client does not expect that.
	time.Sleep(time.Second)
	foobar_s.CloseSend()

	// await all foobar responses
	for {
		_, err := foobar_s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.Unimplemented {
					log.Print("recv Unimplemented code (as expected)")
					break
				}
			}

			log.Fatalf("failed to recv foobar: %v", err)
		}

		log.Print("recv foobar")
	}
}

func ping(client pb.BaseClient, ctx context.Context) {
	// open stream for ping
	log.Print("starting ping")
	ping_s, err := client.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to create ping stream: %v", err)
	}

	// send ping request and then close our sending side
	if err := ping_s.Send(&pb.PingRequest{Content: *content}); err != nil {
		log.Fatalf("failed to send ping: %v", err)
	}
	time.Sleep(time.Second)
	ping_s.CloseSend()

	// await all ping responses
	for {
		r, err := ping_s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recv ping: %v", err)
		}

		log.Printf("recv pong: %s", r.GetContent())
	}
}

func main() {
	// parse cli arguments
	flag.Parse()

	// set up gRPC basics
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	client := pb.NewBaseClient(conn)

	// send actual messages!
	ping(client, ctx)
	foobar(client, ctx)
}
