package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine().Run()
}

func ginEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/haha", printHahahaHandler)
	return r
}

func printHahahaHandler(c *gin.Context) {
	s := printHahaha()
	c.JSON(http.StatusOK, s)

	//若有Header的話 可以用gin裡面自帶的 c.GETHEADER("HeaderName")
	//c.JSON可以自行轉成JSON格式
	//http裡面有Code 比如http.StatusOK = code 200
}

func printHahaha() string {
	s := "haha"
	fmt.Println(s)
	return s
}
