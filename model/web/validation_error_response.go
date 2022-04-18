package web

type ValidationErrorResponse struct {
	Params  string `json:"params"`
	Message string `json:"message"`
}