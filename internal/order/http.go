package main

import (
	"Gorder/internal/order/app"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {

}
func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {

}
