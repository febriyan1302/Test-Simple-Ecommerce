package entity

var MessageSuccess string = "Success"
var MessageFailed string = "Failed"
var MessageNotFound string = "Not Found"

type Response struct {
	Code    int         `json:"code" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data"`
}
