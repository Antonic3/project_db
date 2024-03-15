package repository

import "project-db/model"

type IOrdersRepository interface {
	Create(newOrders model.Orders) (model.Orders, error)
	GetAll() ([]model.Orders, error)
	Update(updatedOrders model.Orders) (model.Orders, error)
	Delete(orders_id int) error
	GetByID(orders_id int) (model.Orders, error)
}
