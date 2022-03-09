package errors

import (
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func Validation(err interface{}) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element ValidationErrorResponse
		element.FailedField = strings.ToLower(err.Field())
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, &element)
	}
	return errors
}

func NewError(err error, message string) *AppError {
	sentry.CaptureException(err)
	return &AppError{
		Error:   "internal_server_error",
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Error:   "bad_request",
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Error:   "not_found",
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func NewUnauthorizedError() *AppError {
	return &AppError{
		Error:   "unauthorized",
		Status:  http.StatusUnauthorized,
		Message: "unauthorized",
	}
}
