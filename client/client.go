package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dat4/grpc-test/mygrpc"
	"google.golang.org/grpc"
)

func openDoor(client pb.DoorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()
	stream, err := client.OpenDoor(ctx)
	var door *pb.Door
	for i := 0; i < 10; i++ {
		door = &pb.Door{Open: "ok"}
		if err := stream.Send(door); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

func main() {
	fmt.Println("vim-go")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewDoorServiceClient(conn)
	openDoor(client)
}
