package main

import "github.com/gin-gonic/gin"
//kaish
func main()  {
	r:=gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"name":c.Param("name"),
			"id":c.Param("id"),
		})
	})
	r.Run(":8080")
}
