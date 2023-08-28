package models

type ResponseSuccess struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

type Response struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
