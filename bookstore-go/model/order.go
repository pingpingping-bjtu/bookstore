package model

import "time"

type Order struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	OrderNo     string     `json:"order_no"`
	TotalAmount int        `json:"total_amount"`
	Status      string     `json:"status"`
	IsPaid      bool       `json:"is_paid"`
	PaymentTime *time.Time `json:"payment_time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
