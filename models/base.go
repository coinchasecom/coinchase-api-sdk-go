package models

type Response struct {
	Code int         `json:"code"`
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}
