package types

import (
	"net/http"
)


type HandlerInterface interface {
	SetupRoutes(server *http.ServeMux, baseUrl string)
}
