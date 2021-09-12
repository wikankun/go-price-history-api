package entity

import "time"

type PriceHistory struct {
	ID        int       `json:"id,primary_key"`
	Item_ID   int       `json:"item_id"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
