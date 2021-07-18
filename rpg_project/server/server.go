package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//定义类对象
type World struct {
}

//绑定类方法
func (this *World) HelloWorld(name string,resp *string) error {
	*resp = name + "您好！"
	return nil
}

func main() {
	//1 注册RPC服务，绑定类对象方法 hello为服务名称
	/*err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc服务失败！", err)
	}*/
	RegisterService(new(World))
	//2 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("监听失败 ", err)
	}
	defer listener.Close()
	fmt.Println("开始监听......")
	//3 建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("链接失败 ", err)
	}
	defer conn.Close()
	fmt.Println("链接成功.....")
	// 4 绑定服务
	rpc.ServeConn(conn)

}
