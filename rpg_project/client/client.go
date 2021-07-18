package main

import (
	"fmt"
	"net/rpc"
)

func main01()  {
	//1.建立链接
	conn,err:=rpc.Dial("tcp","127.0.0.1:8800")
	if err!=nil{
		fmt.Println("客户端连接服务器失败",err)
		return
	}
	defer conn.Close()
	//2。调用远程函数
	var result string
	err=conn.Call("hello.HelloWorld","李白",&result)
	if err!=nil{
		fmt.Println("调用远程函数失败",err)
		return
	}
	fmt.Println(result)
}

func main()  {
	myClient:=ClientInit("127.0.0.1:8800")
	var result string
	err:=myClient.HelloWorld("皮卡丘",&result)
	if err!=nil{
		fmt.Println("调用远程函数失败",err)
		return
	}
	fmt.Println(result)

}