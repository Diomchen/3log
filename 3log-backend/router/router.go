// 路由管理
package router

import "github.com/gin-gonic/gin"

type Option func(router *gin.Engine)

var options = make([]Option, 0)

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	router := gin.Default()
	for _, opt := range options {
		opt(router)
	}
	return router
}
