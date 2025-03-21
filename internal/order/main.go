package main

import (
	"Gorder/internal/common/config"
	"Gorder/internal/common/genproto/orderpb"
	"Gorder/internal/common/server"
	"Gorder/internal/order/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	go server.RunGrpcServer(serviceName, func(server *grpc.Server) {
		service := ports.NewGRPCServer()
		orderpb.RegisterOrderServiceServer(server, service)
	})

	server.RunHttpServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})

}
