package repository

import (
	"database/sql"
	"order-service/internal/order/model/domain"
)

type OrderRepository interface {
	Save(order *domain.Order) (*domain.Order, error)
	FindById(id string) (*domain.Order, error)
}

type orderRepositoryImpl struct {
	db *sql.DB
}

func (o orderRepositoryImpl) Save(order *domain.Order) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderRepositoryImpl) FindById(id string) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepositoryImpl{db: db}
}
