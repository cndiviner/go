package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct{
	Name string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	Age int `form:"age" binding:"required,gt=10"`
}
func main()  {
	r:=gin.Default()
	r.GET("/testing",testing)
	//r.POST("/testing",testing)
	r.Run(":8080")
}
func testing(c *gin.Context) {
	var person Person
	if err:=c.ShouldBind(&person); err==nil{
		c.String(http.StatusOK,"%v",person)
	}else{
		c.String(http.StatusInternalServerError,"person erro:%v",err)
	}
}
