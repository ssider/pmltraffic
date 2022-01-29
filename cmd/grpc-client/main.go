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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pb "pml/traffic-test/pkg/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:9090", "grpc server address")
	delay := flag.Int("delay", 0, "server delay milliseconds")
	flag.Parse()
	// Set up a connection to the server.
	log.Printf("connect to %s\n", *addr)
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: os.Getenv("POD_NAME"), Delay: int64(*delay)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
}
