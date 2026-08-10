package domain

type Product struct {
	ID       string  `json:"id"`
	SellerID string  `json:"seller_id"`
	Phone    string  `json:"phone"`
	Name     string  `json:"name"`
	Breed    string  `json:"breed"`
	Age      int     `json:"age"`
	Gender   string  `json:"gender"`
	Weight   float64 `json:"weight"`
	Color    string  `json:"color"`
	Price    int64   `json:"price"`

	Purpose      string `json:"purpose"`
	HealthStatus string `json:"health_status"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`
}
