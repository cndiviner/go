package main

import (
	"context"
	"fmt"
	pb "gindemo/rpg_project/pb"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	//初始化 consul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象，可以重新定义IP/PORT,也可以使用默认的
	consulClient, err := api.NewClient(consulConfig)
	//服务发现，从consul上，获取健康的服务
	services, _, err := consulClient.Health().Service("grpc And Consul", "grcp", true, nil)
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	//链接服务
	//grpcConn, _ := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	//使用服务发现的consul上的IP/port 来与服务建立链接
	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())
	//初始化grpc 客户端
	grpcClient := pb.NewHelloClient(grpcConn)
	var project pb.Project
	project.Name = "霞光里小区"
	project.ProjectId = 23
	//调用远程函数
	p, err := grpcClient.SayHello(context.TODO(), &project)
	fmt.Println(p, err)
}
