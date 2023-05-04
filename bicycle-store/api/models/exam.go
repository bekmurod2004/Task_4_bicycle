package models

type StigmaApi struct {
	Order_id   string `json:"order_id"`
	Promo_Code string `json:"promo_code"`
}

type SigmaSql struct {
	Order_id   int     `json:"order_id"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}
