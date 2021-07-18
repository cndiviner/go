package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {
	//初始化consul配置
	consuConfig := api.DefaultConfig()
	//创建 consul对象
	consulClient, err := api.NewClient(consuConfig)
	if err != nil {
		fmt.Println("consulClient err:", err)
	}
	//注销服务
	consulClient.Agent().ServiceDeregister("bj38")

}
