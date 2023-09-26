package model

type Paging struct {
	Page  int `json:"page"`
	Total int `json:"total"`
}

type HTTPError struct {
	Message string `json:"message"`
}
