package main

import (
	"Gorder/internal/common/genproto/stockpb"
	"Gorder/internal/common/server"
	"Gorder/internal/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")
	switch serverType {
	case "grpc":
		server.RunGrpcServer(serviceName, func(server *grpc.Server) {
			service := ports.NewGRPCServer()
			stockpb.RegisterStockServiceServer(server, service)
		})
	case "http":
	default:
		panic("unexpected server type: " + serverType)
	}

}
