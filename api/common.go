package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	//conf "github.com/kasiforce/trade/config"
	"github.com/kasiforce/trade/pkg/ctl"
	"strings"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.Response {
	var validErr validator.ValidationErrors
	if errors.As(err, &validErr) {
		var ErrMsg []string
		for _, e := range validErr {
			eMsg := fmt.Sprintf("Validation failed on field %s with tag %s", e.Field(), e.Tag())
			ErrMsg = append(ErrMsg, eMsg)
		}
		errorMsg := strings.Join(ErrMsg, "; ")
		return ctl.RespError(ctx, errorMsg)
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return ctl.RespError(ctx, "JSON类型不匹配")
	}

	return ctl.RespError(ctx, err.Error())

}
