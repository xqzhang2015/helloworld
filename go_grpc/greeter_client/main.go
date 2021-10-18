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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address = "abc.com:50051"
	// address     = "192.168.0.32:50051,192.168.0.32:50053"
	defaultName = "world"
)

var (
	total atomic.Int32
)

func main() {
	concurrency := flag.Int("c", 1000, "an int")
	name := flag.String("name", "xiaoqun", "the name")
	flag.Parse()

	log.Printf("concurrency: %d, name: %s", *concurrency, *name)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), /*grpc.WithBlock(),*/
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithCodec(MyCodec{}),
		// grpc.WithBalancer(grpc.RoundRobin(resolver.NewPseudoResolver([]string{
		// 	"192.168.0.32:50051",
		// 	"192.168.0.32:50053",
		// }))),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 12000*time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: *name}
	in, _ := proto.Marshal(req)

	var wg sync.WaitGroup
	wg.Add(*concurrency)
	for c := 0; c < *concurrency; c++ {
		go func() {
			defer wg.Done()

			for n := 0; n < 10000; n++ {
				out := new(MyStruct)
				err = conn.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out)
				if err != nil {
					log.Printf("could not greet: %v", err)
				}
				total.Inc()
			}
		}()
	}
	wg.Wait()
	log.Printf("total: %d", total.Load())

	// err = conn.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out)
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

	// reply := new(pb.HelloReply)
	// proto.Unmarshal(out.GetData(), reply)
	// log.Printf("Greeting: %s", reply.GetMessage())
}
