package product

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
	Status   string  `json:"status"`
}