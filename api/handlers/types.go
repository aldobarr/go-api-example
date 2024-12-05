package handlers

type ReceiptID struct {
	ID string `json:"id"`
}

type ReceiptPoints struct {
	Points int `json:"points"`
}

type Receipt struct {
	Retailer     string  `json:"retailer" validate:"required"`
	PurchaseDate string  `json:"purchaseDate" validate:"required,purchaseDate"`
	PurchaseTime string  `json:"purchaseTime" validate:"required,purchaseTime"`
	Total        float64 `json:"total" validate:"required,min=0.01"`
	Items        []Item  `json:"items" validate:"required,min=1"`
	Points       int
}

type Item struct {
	ShortDescription string  `json:"shortDescription" validate:"required"`
	Price            float64 `json:"price" validate:"required,min=0.01"`
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
