package handlers

import (
	"net/http"
	"strconv"

	"github.com/xddprog/internal/core/services"
	apierrors "github.com/xddprog/internal/infrastructure/errors"
	"github.com/xddprog/internal/utils"
)

type QuoteHandler struct {
	QuoteService *services.QuoteService
}

func NewQuoteHandler(quoteService *services.QuoteService) *QuoteHandler {
	return &QuoteHandler{QuoteService: quoteService}
}

func (h *QuoteHandler) SetupRoutes(server *http.ServeMux) {
	server.HandleFunc("GET /quotes/random", h.GetRandomQuote)
	server.HandleFunc("GET /quotes/{id}", h.GetQuoteByID)
	server.HandleFunc("POST /quotes", h.CreateQuote)
	server.HandleFunc("DELETE /quotes/{id}", h.DeleteQuote)
	server.HandleFunc("GET /quotes", h.GetAllQuotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.QuoteService.GetRandomQuote()
	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, quote)
}

func (h *QuoteHandler) GetQuoteByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteHTTPError(w, &apierrors.ErrEncodingError)
		return
	}

	quote, err := h.QuoteService.GetQuoteByID(convertId)
	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, quote)
}

func (h *QuoteHandler) GetQuoteByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := h.QuoteService.GetQuoteByAuthor(author)
	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, quotes)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteHTTPError(w, &apierrors.ErrEncodingError)
		return
	}
	err = h.QuoteService.DeleteQuote(convertId)
	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Quote deleted successfully"})
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.QuoteService.CreateQuote(r.Body)
	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, quote)
}

func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.QuoteService.GetAllQuotes()

	if author := r.URL.Query().Get("author"); author != "" {
		h.GetQuoteByAuthor(w, r)
		return
	}

	if err != nil {
		utils.WriteHTTPError(w, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, quotes)
}
