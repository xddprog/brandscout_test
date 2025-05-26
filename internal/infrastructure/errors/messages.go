package apierrors

import (
	"fmt"
	"net/http"
)


type APIError struct {
	Code    int
	Message any
}


func (e *APIError) Error() string {
    switch msg := e.Message.(type) {
    case string:
        return msg
    case error:
        return msg.Error()
    default:
        return fmt.Sprintf("%v", msg)
    }
}


var (
	ErrItemNotFound = func (itemName string) *APIError {
		return &APIError{Code: http.StatusNotFound, Message: fmt.Sprintf("%s not found", itemName)}
	}
	ErrInternalServerError = APIError{Code: http.StatusInternalServerError, Message: "internal server error"}
	ErrEncodingError = APIError{Code: http.StatusInternalServerError, Message: "encoding error"}
)
