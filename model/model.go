package model

import "time"

// Book model holding book data
type Book struct {
	ID             string
	Title          string
	Synopsis       string
	ISBN10         string
	ISBN13         string
	Category       string
	Language       string
	Publisher      string
	Edition        string
	SoldAmount     int
	CurrentAmount  int
	PaperbackPrice *float64
	EbookPrice     *float64
	AverageScore   *float64
	CreatedTime    *time.Time
	ModifiedTime   *time.Time
	Version        int //for optimistic locking
}

// Review model holding book's riview data
type Review struct {
	ID           string
	Score        int
	Description  string
	BookID       string
	CreatedTime  *time.Time
	ModifiedTime *time.Time
	Version      int //for optimistic locking
}