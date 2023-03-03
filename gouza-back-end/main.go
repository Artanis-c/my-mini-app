package main

import (
	"github.com/gin-gonic/gin"
	"gouza-back-end/src/domain/common/register"
)

func main() {
	webServer := gin.New()
	webServer.Use(gin.Recovery())
	register.UserRoute(webServer)
	err := webServer.Run(":9001")
	if err != nil {
		panic(err)
	}
}
