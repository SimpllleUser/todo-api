package model

type HTTPError struct {
	Error string `json:"error"`
}

type BooleanResponse struct {
	Data bool `json:"data"`
}
