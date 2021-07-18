package main

import (
	"context"
	"fmt"
	pb "gindemo/rpg_project/pb"
	"google.golang.org/grpc"
	"net"
)
type Childrent struct {

}

func (this *Childrent)SayHello(ctx context.Context,t *pb.Teacher) (*pb.Teacher,error)  {
 	t.Name+="is Sleeping"
 	return t,nil
}
func main()  {
	//初始化一个grpc对象
	grpcServer:=grpc.NewServer()
	//注册服务
	pb.RegisterSayNameServer(grpcServer,new(Childrent))
	//设置监听
	listener,err:=net.Listen("tcp","127.0.0.1:8800")
	if err!=nil{
		fmt.Println("listen err:",err)
		return
	}
	defer listener.Close()
	//启动服务
	grpcServer.Serve(listener)

}