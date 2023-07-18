package database

import (
	"database/sql"
	"github.com/FabioRocha231/fullcyclejunho/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec(
		"INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)",
		order.ID, order.Price, order.Tax,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int

	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}