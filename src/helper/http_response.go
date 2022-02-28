package helper

import (
	"net/http"

	"github.com/scale/src/domain"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func Response(code int, message string, data, errors interface{}) HttpResponse {
	res := HttpResponse{}
	res.Code = code
	res.Message = message
	res.Data = data
	res.Errors = errors

	return res
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.Error() {
	case domain.ErrNotFound.Error():
		return http.StatusNotFound
	case domain.ErrBadParamInput.Error():
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
