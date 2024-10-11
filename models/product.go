package models

type Product struct {
	ID         uint     `json:"id" gorm:"primary_key"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Quantity   int      `json:"quantity"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
