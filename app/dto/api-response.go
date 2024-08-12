package dto

type ApiResponse[T any] struct {
	ResponseKey     string `json:"responseKey"`
	ResponseMessage string `json:"responseMessage"`
	Data            T      `json:"data"`
}
