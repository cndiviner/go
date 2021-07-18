package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{
			"title":"第二课",
		})
	})
	r.Run(":8080")

}
