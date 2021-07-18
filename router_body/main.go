package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	r:=gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyByts,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		c.Request.Body=ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName:=c.PostForm("first_name")
		lastNmae:=c.DefaultPostForm("last_name","default_last_name")
		c.String(http.StatusOK,"%s,%s,$s",firstName,lastNmae,string(bodyByts))

		c.String(http.StatusOK,string(bodyByts))
	})
	r.Run(":8800")
}
