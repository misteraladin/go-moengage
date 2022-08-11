package moengage

type Response struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Error   ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Attribute string `json:"attribute"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	RequestID string `json:"request_id"`
}
