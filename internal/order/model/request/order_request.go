package request

type OrderRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
}
