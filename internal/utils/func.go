package utils

import (
	"encoding/json"
	"log"
	"net/http"

	apierrors "github.com/xddprog/internal/infrastructure/errors"
)



func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if data == nil {
        return nil
    }
    if err := json.NewEncoder(w).Encode(data); err != nil {
        WriteHTTPError(w, err)
        return err
    }
    return nil
}


func WriteHTTPError(w http.ResponseWriter, err any) {
	var response map[string]any
	var statusCode int

	switch err := err.(type) {
	case *apierrors.APIError:
		statusCode = err.Code
		response = map[string]any{
			"error": err.Message,
		}

	case error:
		statusCode = http.StatusInternalServerError
		response = map[string]any{
			"error": err.Error(),
		}

	default:
		log.Printf("Internal Server Error: %v", err)
		statusCode = http.StatusInternalServerError
		response = map[string]any{
			"error":   "internal server error",
			"details": "An unexpected error occurred",
		}
	}

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
