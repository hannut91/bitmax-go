package types

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type OrderResponse struct {
	COID    string `json:"coid"`
	Action  string `json:"action"`
	Success bool   `json:"success"`
}
