package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()
	engine.GET("/JSONP", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		context.JSONP(http.StatusOK, data)
	})
	engine.Run(":8080")
}
