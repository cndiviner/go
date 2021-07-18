package main

import "net/rpc"

//定义类
type MyClient struct {
	c *rpc.Client
}

func ClientInit(addr string) MyClient {
	conn, _ := rpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

func (this *MyClient) HelloWorld(a string, b *string) error {
	return this.c.Call("hello.HelloWorld", a, b)
}
