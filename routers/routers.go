package routers

import (
	"github.com/202lp2/go2/apis"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apis.ItemsIndex)
	}

	return r
}
