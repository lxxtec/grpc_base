package main

import (
	"bufio"
	"context"
	"fmt"
	"gprc-base/protos"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// http2 默认是安全的
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	c := protos.NewEchoServiceClient(cli)
	// 使用流
	stream, err := c.GetCStreamEcho(context.Background())
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		if string(line) == "close" {
			res, _ := stream.CloseAndRecv()
			fmt.Println("all data: ", res.GetRes())
		}
		req := protos.EchoRequest{Req: string(line)}

		//res, err := c.GetUnaryEcho(context.Background(), &req)
		err = stream.Send(&req)
		if err != nil {
			panic(err)
		}
		//fmt.Println(res.GetRes())
	}
}
