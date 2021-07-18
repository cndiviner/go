package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct{
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday"`
}
func main()  {
	r:=gin.Default()
	r.GET("/testing",testing)
	r.POST("/testing",testing)
	r.Run(":8080")
}
func testing(c *gin.Context) {
	var person Person
	if err:=c.ShouldBind(&person); err==nil{
		c.String(http.StatusOK,"%v",person)
	}else{
		c.String(http.StatusOK,"person erro:%s",err)
	}
}
