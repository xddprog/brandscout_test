package repositories

import (
	"database/sql"

	"github.com/xddprog/internal/infrastructure/database/models"
)


type QuoteRepository struct {
	DB *sql.DB
}


func NewQuoteRepository(db *sql.DB) *QuoteRepository {
	return &QuoteRepository{DB: db}
}


func (r *QuoteRepository) GetRandomQuote() (*models.Quote, error) {
	query := "SELECT * FROM quotes ORDER BY RANDOM() LIMIT 1"
	row := r.DB.QueryRow(query)
	var quote models.Quote
	err := row.Scan(&quote.ID, &quote.Quote, &quote.Author)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}


func (r *QuoteRepository) GetQuoteByID(id int) (*models.Quote, error) {
	query := "SELECT * FROM quotes WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var quote models.Quote
	err := row.Scan(&quote.ID, &quote.Quote, &quote.Author)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}


func (r *QuoteRepository) GetQuoteByAuthor(author string) ([]models.Quote, error) {
	query := "SELECT * FROM quotes WHERE author = ?"
	rows, err := r.DB.Query	(query, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote

	for rows.Next() {
		var quote models.Quote
		err := rows.Scan(&quote.ID, &quote.Quote, &quote.Author)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}
	return quotes, nil
}


func (r *QuoteRepository) CreateQuote(quote models.CreateQuote) (*models.Quote, error) {
	query := "INSERT INTO quotes (quote, author) VALUES (?, ?) RETURNING *"
	var createdQuote models.Quote
	err := r.DB.QueryRow(query, quote.Quote, quote.Author).Scan(
		&createdQuote.ID, &createdQuote.Quote, &createdQuote.Author,
	)
	if err != nil {
		return nil, err
	}
	return &createdQuote, nil
}


func (r *QuoteRepository) DeleteQuote(id int) error {
	var deletedQuote models.Quote

	query := "DELETE FROM quotes WHERE id = ? RETURNING *"
	err := r.DB.QueryRow(query, id).Scan(&deletedQuote.ID, &deletedQuote.Quote, &deletedQuote.Author)
	if err != nil {
		return err
	}
	return nil
}


func (r *QuoteRepository) GetAllQuotes() ([]models.Quote, error) {
	query := "SELECT * FROM quotes"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var quotes []models.Quote

	for rows.Next() {
		var quote models.Quote
		err := rows.Scan(&quote.ID, &quote.Quote, &quote.Author)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return quotes, nil
}
