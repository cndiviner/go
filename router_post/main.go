package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r:=gin.Default()
	r.POST("/test", func(c *gin.Context) {
		firstName:=c.PostForm("first_name")
		lastNmae:=c.DefaultPostForm("last_name","default_last_name")
		c.String(http.StatusOK,"%s,%s",firstName,lastNmae)
	})
	r.Run(":8080")
}
