package service

import (
	"order-service/internal/order/model/domain"
	"order-service/internal/order/repository"
)

type OrderService interface {
	Create(order *domain.Order) error
}

type orderServiceImpl struct {
	repo repository.OrderRepository
}

func (o orderServiceImpl) Create(order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}
