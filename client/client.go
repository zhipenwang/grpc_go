package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_go/protofile/grpc_go/protofile"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8880", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn failed is %v", err)
	}
	defer conn.Close()

	client := protofile.NewRpcStreamClient(conn)

	RouterSimple(client)
	RouterServer(client)
	RouterClient(client)
	RouterChat(client)
}

func RouterSimple(client protofile.RpcStreamClient)  {
	response, err:= client.RouterSimple(context.Background(), &protofile.Request{
		Uid: 1,
	})
	if err != nil {
		log.Fatalf("request simple failed %v", err)
	}
	fmt.Println(response)
}

func RouterServer(client protofile.RpcStreamClient)  {
	stream, err := client.RouterServer(context.Background(), &protofile.Request{
		Uid: 1,
	})
	if err != nil {
		log.Fatalf("request server failed %v", err)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			fmt.Printf("response server failed %v", err)
			return
		}
		fmt.Println(response)
	}
}

func RouterClient(client protofile.RpcStreamClient) {
	stream, err := client.RouterClient(context.Background())
	if err != nil {
		log.Fatalf("request client failed %v", err)
	}

	request_info := []*protofile.Request{}
	request_info = append(request_info, &protofile.Request{Uid: 1, Name: "client1"})
	request_info = append(request_info, &protofile.Request{Uid: 1, Name: "client2"})
	for _, v := range request_info {
		err := stream.Send(v)
		if err != nil {
			log.Fatalf("client send failed %v", err)
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("response client failed %v", err)
	}
	fmt.Println(response)
}

func RouterChat(client protofile.RpcStreamClient)  {

	stream, err := client.RouterChat(context.Background())
	if err != nil {
		log.Fatalf("failed routerchat method %v", err)
	}

	waitChan := make(chan int)

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitChan)
				return
			}

			if err != nil {
				stream.CloseSend()
				log.Fatalf("failed to %v", err)
			}
			if in.Code == 200 {
				log.Printf("receive uid=%d, name=%s, code=%d", in.Uid, in.Name, in.Code)
			} else if in.Code == 201 {
				log.Printf("response uid=%d, name=%s, code=%d", in.Uid, in.Name, in.Code)
			}

		}
	}()


	uid_int, _ := strconv.Atoi(os.Args[1])
	uid := int32(uid_int)

	stream.Send(&protofile.Request{Uid: uid, Type: 1})

	request_info := []*protofile.Request{}
	request_info = append(request_info, &protofile.Request{Uid: uid, Name: "test1", Type: 2})
	request_info = append(request_info, &protofile.Request{Uid: uid, Name: "test2", Type: 2})
	request_info = append(request_info, &protofile.Request{Uid: uid, Name: "test3", Type: 2})


	for _, request_i := range request_info {
		time.Sleep(time.Second*5)
		if err := stream.Send(request_i); err != nil {
			log.Fatalf("failed send is %v", err)
		}
	}

	for {
		select {

		}
	}

	<-waitChan
}
