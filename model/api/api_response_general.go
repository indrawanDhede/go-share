package api

type ApiResponseGeneral struct {
	Total  int         `json:"total"`
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
