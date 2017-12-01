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
	"os"
	"log"
	//"os"
	"io"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	//address     = args[0] + "localhost:50051"
	defaultName = "steven"
)

func GetLogs(client pb.GreeterClient) {
	stream, err := client.SayHelloAgain(context.Background(), &pb.HelloRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Message)
	}
}

func main() {
	// Set up a connection to the server.
	address := os.Args[1] + ":50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(address)
		log.Fatalf("did not connect: %v" + address, err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }

	GetLogs(c)
	
	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.Message)
	//num := 0

	// for num < 5 {
	// 	r, err := c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: name})
	// 	if err != nil {
	// 			log.Fatalf("could not greet: %v", err)
	// 	}
	// 	log.Printf("Greeting: %s", r.Message)
	// 	num ++
	// }
}
