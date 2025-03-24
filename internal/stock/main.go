package main

import (
	"Gorder/internal/common/config"
	"Gorder/internal/common/genproto/stockpb"
	"Gorder/internal/common/server"
	"Gorder/internal/stock/ports"
	"Gorder/internal/stock/service"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}
func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	switch serverType {
	case "grpc":
		server.RunGrpcServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
	default:
		panic("unexpected server type: " + serverType)
	}

}
