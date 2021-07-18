package main

import (
	"gindemo/other_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context)  {
	c.HTML(200,"index/index.html",gin.H{
		"title":"第二课",
	})
}
func main()  {
	c:=gin.Default()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	c.LoadHTMLGlob("template/**/*")
	c.GET("/index",Hello)
	c.GET("/user",model.User)
	c.GET("/userone",model.UserOne)
	c.GET("/userlist",model.UserList)
	c.GET("/usermap",model.UserMap)
	c.GET("/user/:id",model.UserGet)
	c.GET("/usertest1",model.Test)
	c.GET("/add",model.AddHtml)
	c.POST("/addDo",model.AddData)
	c.Run(":8090")
}
