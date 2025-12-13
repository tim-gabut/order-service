package controller

import (
	"net/http"
	"order-service/internal/order/model/domain"
	"order-service/internal/order/model/request"
	"order-service/internal/order/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) Create(ctx *gin.Context) {
	var req request.OrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &domain.Order{
		Customer: req.CustomerName,
	}

	err := c.service.Create(order)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
