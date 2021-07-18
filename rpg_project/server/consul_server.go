package main

import (
	"context"
	"fmt"
	pb "gindemo/rpg_project/pb"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

type Children struct {
}

func (this *Children) SayHello(ctx context.Context, p *pb.Project) (*pb.Project, error) {
	p.Name = "hello" + p.Name
	return p, nil
}
func main() {

	//把grpc服务，注册到consul上
	consulConfig := api.DefaultConfig()
	api.DefaultConfig()
	//创建 consul 对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	//告诉consul ,即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grcp", "consul"},
		Name:    "grpc And Consul",
		Address: "127.0.0.1",
		Port: 8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//注册grpc服务到consul上
	consulClient.Agent().ServiceRegister(&reg)

	//----------------------------------//
	//初始化
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))
	//设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listen err:", err)
	}
	defer listener.Close()
	//启动服务
	grpcServer.Serve(listener)
}
