package controllers

import "apiproject_new/models"

type ResponseStruct struct {
	Code int `json:"code"` //状态码
	Message map[string]interface{} `json:"message"` //反馈信息
	user models.User

}
