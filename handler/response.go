package handler

type ResponseItems struct {
	Data interface{} `json:"items"`
}

type ResponseError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
