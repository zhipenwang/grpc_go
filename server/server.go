package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_go/protofile/grpc_go/protofile"
	"io"
	"log"
	"net"
	"sync"
)

type Route struct {
	mu 	sync.Mutex
	info map[int32][]*protofile.Response
	socket map[int32]protofile.RpcStream_RouterChatServer
}

func main() {
	StreamServer()
}

func StreamServer()  {
	listen, err := net.Listen("tcp", ":8880")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protofile.RegisterRpcStreamServer(grpcServer, &Route{
		info: make(map[int32][]*protofile.Response),
		socket: make(map[int32]protofile.RpcStream_RouterChatServer),
	})
	grpcServer.Serve(listen)
}

func (r *Route) RouterSimple(ctx context.Context, request *protofile.Request) (*protofile.Response, error) {
	return &protofile.Response{
		Uid: 1,
		Name: "simple",
		Code: 200,
	}, nil
}

func (r *Route) RouterServer(request *protofile.Request, stream protofile.RpcStream_RouterServerServer) error  {
	err := stream.Send(&protofile.Response{
		Uid: 1,
		Name: "server1",
		Code: 200,
	})
	if err != nil {
		log.Fatalf("server send failed %v", err)
	}
	err = stream.Send(&protofile.Response{
		Uid: 1,
		Name: "server2",
		Code: 200,
	})
	if err != nil {
		log.Fatalf("server send failed %v", err)
	}
	return nil
}

func (r *Route) RouterClient(stream protofile.RpcStream_RouterClientServer) error {
	var name string
	for {
		request, err := stream.Recv()
		fmt.Println(err)
		if err == io.EOF {
			fmt.Println(name)
			err  = stream.SendAndClose(&protofile.Response{
				Uid: 1,
				Name: name,
				Code: 200,
			})
			fmt.Println(err)
			return err
		}

		if err != nil {
			return err
		}

		fmt.Println(request)

		name += request.Name + " "
	}
	return nil
}

func (r *Route) RouterChat(stream protofile.RpcStream_RouterChatServer) error {
	for {

		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Printf("client failed %v", err)
			return err
		}

		if in.Type == 1 {
			r.SetSocket(in.Uid, stream)
		}

		if in.Type == 2 {
			target_uid := int32(1)
			if in.Uid == 1 {
				target_uid = 2
			} else if in.Uid == 2 {
				target_uid = 1
			}
			r.ResponseToUid(in.Uid)
			r.SendToUid(target_uid, in.Uid, in.Name)
		}
	}
}

func (r *Route) SetSocket(uid int32, stream protofile.RpcStream_RouterChatServer)  {
	fmt.Println(uid, stream)
	r.socket[uid] = stream
}

func (r *Route) SendToUid(uid, from_uid int32, name string)  {
	if r.socket[uid] == nil {
		fmt.Printf("send to %d go pns\n", uid)
		return
	}
	stream := r.socket[uid]
	err := stream.Send(&protofile.Response{Uid: from_uid, Name: name, Code: 200})
	if err != nil {
		log.Fatalf("send failed %v", err)
	}
}

func (r *Route) ResponseToUid(uid int32)  {
	if r.socket[uid] == nil {
		fmt.Printf("response to %d go pns\n", uid)
		return
	}
	stream := r.socket[uid]
	err := stream.Send(&protofile.Response{Uid: uid, Name: "response", Code: 201})
	if err != nil {
		log.Fatalf("send failed %v", err)
	}
}