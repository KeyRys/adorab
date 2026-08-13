package domain

type Cart struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type CartItem struct {
	ID        string `json:"id"`
	CartID    string `json:"cart_id"`
	ProductID string `json:"product_id"`
}

type CartItemWithProduct struct {
	ID      string  `json:"id"`
	Product Product `json:"product"`
}
