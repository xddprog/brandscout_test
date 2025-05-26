package apierrors

import (
	"log"
	"database/sql"
)

func CheckDBError(err error, itemName string) *APIError {
	switch err {
	case sql.ErrNoRows:
		return ErrItemNotFound(itemName)
	default:
		log.Printf("Internal Server Error: %v", err)
		return &ErrInternalServerError
	}
}