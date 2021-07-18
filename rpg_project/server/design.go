package main

import "net/rpc"
//创建接口，在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}
//调用该方法时，需要给i传参数，参数应该是实现helloworld方法的类对象
func RegisterService(i MyInterface)  {
	rpc.RegisterName("hello",i)
}

