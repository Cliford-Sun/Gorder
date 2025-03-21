package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHttpServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		// TODO Warning log
	}
	RunHttpServerOnAddr(addr, wrapper)

}
func RunHttpServerOnAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")

	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
