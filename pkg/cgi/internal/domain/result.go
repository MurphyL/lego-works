package domain

type Result struct {
	Payload any    `json:"payload"`
	Message string `json:"message"`
}
