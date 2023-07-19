package usecase

import "github.com/FabioRocha231/fullcyclejunho/internal/entity"

type OrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

// {"id": "1", "price": 10.0, "tax": 0.1}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{OrderRepository: orderRepository}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	if err := order.CalculateFinalPrice(); err != nil {
		return nil, err
	}

	if err := c.OrderRepository.Save(order); err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
