package ports

import (
	"Gorder/internal/common/genproto/stockpb"
	"Gorder/internal/stock/app"
	"context"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
