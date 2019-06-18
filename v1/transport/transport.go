package transport

import "time"

type ReviewTransport struct {
	ID           string     `json:"id"`
	Score        int        `json:"score"`
	Description  string     `json:"description"`
	BookID       string     `json:"-"`
	CreatedTime  *time.Time `json:"created_time"`
	ModifiedTime *time.Time `json:"modified_time"`
	Version      int        `json:"version"`
}

type BookTransport struct {
	ID             string     `json:"id"`
	Title          string     `json:"title" validate:"required"`
	Synopsis       string     `json:"synopsis"`
	ISBN10         string     `json:"isbn10"`
	ISBN13         string     `json:"isbn13"`
	Category       string     `json:"category"`
	Language       string     `json:"language" validate:"required"`
	Publisher      string     `json:"publisher" validate:"required"`
	Edition        string     `json:"edition"`
	SoldAmount     int        `json:"sold_amount" validate:"gte=0"`
	CurrentAmount  int        `json:"current_amount" validate:"gte=0"`
	PaperbackPrice *float64   `json:"paperback_price" validate:"gte=0"`
	EbookPrice     *float64   `json:"ebook_price" validate:"gte=0"`
	AverageScore   *float64   `json:"average_score"`
	CreatedTime    *time.Time `json:"created_time"`
	ModifiedTime   *time.Time `json:"modified_time"`
	Version        int        `json:"version"`
}

type ResponseTransport struct {
	Total int             `json:"total"`
	Size  int             `json:"size"`
	Data  []BookTransport `json:"data"`
}

type FillBookTransport struct {
	Amount int `json:"amount"`
}

type SaleBookTransport struct {
	Amount int `json:"amount"`
}
