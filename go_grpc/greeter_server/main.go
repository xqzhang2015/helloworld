/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"go.uber.org/atomic"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	inflight atomic.Int32
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	inflight.Inc()
	defer inflight.Dec()
	select {
	case <-time.After(100 * time.Millisecond):
		//
	}

	// log.Printf("Received: %v", in.GetName())
	// return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	return &pb.HelloReply{Message: "H"}, nil
}

func main() {
	port := flag.String("port", ":50051", "the port")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		for {
			select {
			case <-time.After(3 * time.Second):
				log.Printf("inflight: %d", inflight.Load())
			}
		}
	}()

	// s := grpc.NewServer(grpc.MaxConcurrentStreams(100))
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
