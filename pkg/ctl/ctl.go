package ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/e"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.Success
	if len(code) > 0 {
		status = code[0]
	}
	r := &Response{Code: status, Data: data, Msg: e.GetMsg(status)}
	return r
}

func RespError(ctx *gin.Context, err string, code ...int) *Response {
	status := e.Error
	if len(code) > 0 {
		status = code[0]
	}
	r := &Response{Code: status, Data: err, Msg: e.GetMsg(status)}
	return r
}
