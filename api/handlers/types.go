package handlers

type ReceiptID struct {
	ID string `json:"id"`
}

type ReceiptPoints struct {
	Points int `json:"points"`
}

// Used after validation
type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        float64
	Items        []Item
	Points       int
}

// Used after validation
type Item struct {
	ShortDescription string
	Price            float64
}

type ReceiptRaw struct {
	Retailer     string    `json:"retailer" validate:"required"`
	PurchaseDate string    `json:"purchaseDate" validate:"required,purchaseDate"`
	PurchaseTime string    `json:"purchaseTime" validate:"required,purchaseTime"`
	Total        any       `json:"total" validate:"required,numeric"`
	Items        []ItemRaw `json:"items" validate:"required,min=1"`
}

type ItemRaw struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            any    `json:"price" validate:"required,numeric"`
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
