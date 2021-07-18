package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id int
	Name string
	Age int
	Addr string
}

func User(c *gin.Context )  {
	user_info:=UserInfo{
		Id:1,
		Name: "leifei",
		Age: 23,
		Addr: "北京",
	}
	c.HTML(200,"user/user.html",user_info)
}

func UserOne(c *gin.Context )  {
	//最大长度为3的数组；
	user_info:=[3]int{1,22,3}
	c.HTML(200,"user/userone.html",user_info)
}
//结构体数据
func UserList(c *gin.Context )  {
	//最大长度为3的数组；
	user_info:=[3]UserInfo{
		{
			Id:1,
			Name: "张三",
			Age: 23,
			Addr: "北京",
		},
		{
			Id:2,
			Name: "李四",
			Age: 25,
			Addr: "上海",
		},
		{
			Id:3,
			Name: "刘五",
			Age: 28,
			Addr: "广州",
		},
	}
	c.HTML(200,"user/userlist.html",user_info)
}
func UserMap(c *gin.Context )  {
	//最大长度为3的数组；
	user_info:=map[string]string{"name":"jfei","age":"21"}
	c.HTML(200,"user/usermap.html",user_info)
}
func UserGet(c *gin.Context )  {
	//最大长度为3的数组；
	user_info:=c.Param("id")
	c.HTML(200,"user/userget.html",user_info)
}

func Test(c *gin.Context )  {
	//最大长度为3的数组；
	//user_info:=c.Query("id")
	//获取ID的值，如果没有给默认值
	//user_info:=c.DefaultQuery("id","leifei")
	//获取数组
	user_info:=c.QueryArray("id")
	c.HTML(200,"user/userget.html",user_info)
}
func AddHtml(c *gin.Context ) {
	c.HTML(200,"user/add.html","添加操作")
}
func AddData(c *gin.Context ) {
	name:=c.PostForm("name")
	age:=c.PostForm("age")
	//如果没有值，给一个默认值
	//addr:=c.DefaultPostForm("addr","肖云路12号")
	love:=c.PostFormArray("love")
	fmt.Println(love)
	fmt.Println(name)
	fmt.Println(age)
	//c.String(200,"添加成功")
	//ajax 形式返回
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"成功","age":age})
	//c.HTML(200,"user/add.html",name)
}

