package order_usecase

import (
	"github.com/Mugen-Builders/devolt/internal/domain/entity"
)

type DeleteOrderInputDTO struct {
	Id uint `json:"id"`
}

type DeleteOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewDeleteOrderUseCase(orderRepository entity.OrderRepository) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *DeleteOrderUseCase) Execute(input *DeleteOrderInputDTO) error {
	return u.OrderRepository.DeleteOrder(input.Id)
}
