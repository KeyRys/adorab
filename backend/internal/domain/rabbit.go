package domain

type Rabbit struct {
	ID       string `json:"id"`
	SellerID string `json:"seller_id"`

	Name   string  `json:"name"`
	Breed  string  `json:"breed"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
	Weight float64 `json:"weight"`
	Color  string  `json:"color"`
	Price  int64   `json:"price"`

	Purpose      string `json:"purpose"`
	HealthStatus string `json:"health_status"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`

	Status string `json:"status"`
}
type RabbitRelation struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	ChildID  string `json:"child_id"`
}

type RabbitRelationResponse struct {
	ParentName string `json:"parent_name"`
	ChildName  string `json:"child_name"`
}
