package rest

type Result struct {
	Payload any    `json:"payload"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
