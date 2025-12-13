package response

type OrderResponse struct {
	Id          string  `json:"id"`
	Code        string  `json:"code"`
	Customer    string  `json:"customer"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}
