package failure

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code        ErrorCode
	OriginalErr error
}

func (e *AppError) Error() string {
	if val, ok := errorMessage[e.Code]; ok {
		return val
	}
	return fmt.Sprintf("Lỗi hệ thống: %d\nVui lòng thử lại sau", e.Code)
}

type errorMessageMap map[ErrorCode]string

var errorMessage = errorMessageMap{
	ErrReusableCodeNotFound: "Mã quà tặng không tồn tại",
}

func (e *AppError) HTTPCode() int {
	if val, ok := httpCode[e.Code]; ok {
		return val
	}
	return http.StatusBadRequest
}

type httpCodeMap map[ErrorCode]int

var httpCode = httpCodeMap{
	ErrReusableCodeFailed: http.StatusInternalServerError,
}

type ErrorCode int

const (
	ErrReusableCodeGetByCodeBinding ErrorCode = 991001
	ErrReusableCodeNotFound         ErrorCode = 991002
	ErrReusableCodeFailed           ErrorCode = 991003
)
