package services

import (
	"encoding/json"
	"io"

	"github.com/xddprog/internal/core/repositories"
	"github.com/xddprog/internal/infrastructure/database/models"
	apierrors "github.com/xddprog/internal/infrastructure/errors"
)


type QuoteService struct {
	QuoteRepository *repositories.QuoteRepository
}


func NewQuoteService(quoteRepository *repositories.QuoteRepository) *QuoteService {
	return &QuoteService{QuoteRepository: quoteRepository}
}


func (s *QuoteService) GetRandomQuote() (*models.Quote, error) {
	quote, err := s.QuoteRepository.GetRandomQuote()
	if err != nil {
		return nil, apierrors.CheckDBError(err, "quote")
	}
	return quote, nil
}


func (s *QuoteService) GetQuoteByID(id int) (*models.Quote, error) {
	quote, err := s.QuoteRepository.GetQuoteByID(id)
	if err != nil {
		return nil, apierrors.CheckDBError(err, "quote")
	}
	return quote, nil
}


func (s *QuoteService) GetQuoteByAuthor(author string) ([]models.Quote, error) {
	quote, err := s.QuoteRepository.GetQuoteByAuthor(author)
	if err != nil {
		return nil, apierrors.CheckDBError(err, "quote")
	}
	return quote, nil
}


func (s *QuoteService) CreateQuote(quote io.ReadCloser) (*models.Quote, error) {
	var createQuote models.CreateQuote

	if err := json.NewDecoder(quote).Decode(&createQuote); err != nil {
		return nil, &apierrors.ErrEncodingError
	}

	createdQuote, err := s.QuoteRepository.CreateQuote(createQuote)
	if err != nil {
		return nil, apierrors.CheckDBError(err, "quote")
	}
	return createdQuote, nil
}


func (s *QuoteService) DeleteQuote(id int) error {
	err := s.QuoteRepository.DeleteQuote(id)
	if err != nil {
		return apierrors.CheckDBError(err, "quote")
	}
	return nil
}

func (s *QuoteService) GetAllQuotes() ([]models.Quote, error) {
	quotes, err := s.QuoteRepository.GetAllQuotes()
	if err != nil {
		return nil, apierrors.CheckDBError(err, "quote")
	}
	return quotes, nil
}
