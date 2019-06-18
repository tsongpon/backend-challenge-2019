package repository

import (
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/query"
	"github.com/tsongpon/backend-challenge-2019/report"
)

// BookRepository define interface for book repository
type BookRepository interface {
	GetBook(string) (*model.Book, error)
	CreateBook(model.Book) (*model.Book, error)
	UpdateBook(model.Book) (*model.Book, error)
	QueryBook(query.BookQuery) ([]model.Book, error)
	CountBook(query.BookQuery) (int, error)
	DeleteBook(string) error
	GetBestSaller() ([]report.BestSallerBook, error)
	GetBestSallerByCategory() ([]report.BestSallerCategory, error)
}

// ReviewRepository define interface for review repository
type ReviewRepository interface {
	GetReview(string) (*model.Review, error)
	GetReviewByBook(string) ([]model.Review, error)
	CreateReview(model.Review) (*model.Review, error)
	UpdateReview(model.Review) (*model.Review, error)
	DeleteReview(string) error
}
