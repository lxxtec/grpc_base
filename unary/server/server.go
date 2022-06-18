package main

import (
	"fmt"
	"gprc-base/protos"
	"io"
	"net"
	"strings"

	"google.golang.org/grpc"
)

// type EchoServiceServer interface {
// 	GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
// 	mustEmbedUnimplementedEchoServiceServer()
// }

// echoService 是对上面接口的实现
type echoService struct {
	protos.UnimplementedEchoServiceServer
}

// func (e *echoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
// 	res := "received:" + req.GetReq()
// 	fmt.Println(res)
// 	return &protos.EchoResponse{Res: res}, nil
// }

func (e *echoService) GetCStreamEcho(stream protos.EchoService_GetCStreamEchoServer) error {
	str := []string{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := strings.Join(str, " ")
			return stream.SendAndClose(&protos.EchoResponse{Res: res})
		}
		str = append(str, req.GetReq())
		fmt.Println("received: ", req.GetReq())
	}
}

// grpc服务端编程
// 1. 编写proto消息文件
// 2. 生成模板代码
// 3. 重写模板
func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, new(echoService))
	// 监听
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on localhost:8080")
	defer lis.Close()
	rpcs.Serve(lis)
}
