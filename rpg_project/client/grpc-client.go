package main

import (
	"context"
	"fmt"
	pb "gindemo/rpg_project/pb"
	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return
	}
	defer grpcConn.Close()
	grpcClient := pb.NewSayNameClient(grpcConn)
	//调用远程服务,context.TODO() 表示空对象
	var teacher pb.Teacher
	teacher.Name = "leifei"
	teacher.Age = 19
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t, err)

}
