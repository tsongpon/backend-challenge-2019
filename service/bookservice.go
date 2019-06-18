package service

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/bserror"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/query"
	"github.com/tsongpon/backend-challenge-2019/report"
	"github.com/tsongpon/backend-challenge-2019/repository"
)

type BookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) *BookService {
	s := new(BookService)
	s.bookRepo = bookRepo
	return s
}

func (s *BookService) Create(b model.Book) (*model.Book, error) {
	log.Info(fmt.Sprintf("create new book, title %s", b.Title))
	created, err := s.bookRepo.CreateBook(b)
	if err != nil {
		log.Error("create book error", err.Error())
		return nil, err
	}
	fromDB, err := s.bookRepo.GetBook(created.ID)
	if err != nil {
		return nil, err
	}
	return fromDB, nil
}

func (s *BookService) GetBook(id string) (*model.Book, error) {
	if b, err := s.bookRepo.GetBook(id); err == nil {
		return b, nil
	} else {
		log.Error(fmt.Sprintf("get book id %s error, %s", id, err.Error()))
		return nil, err
	}
}

func (s *BookService) QueryBook(q query.BookQuery) ([]model.Book, error) {
	if books, err := s.bookRepo.QueryBook(q); err != nil {
		log.Error("error while listing books", err.Error())
		return nil, err
	} else {
		return books, nil
	}
}

func (s *BookService) CountBook(q query.BookQuery) (int, error) {
	return s.bookRepo.CountBook(q)
}

func (s *BookService) Update(b model.Book) (*model.Book, error) {
	_, err := s.bookRepo.GetBook(b.ID)
	if err != nil {
		return nil, err
	}
	if updated, err := s.bookRepo.UpdateBook(b); err != nil {
		log.Error(fmt.Sprintf("update book id %s error, %s", b.ID, err.Error()))
		return nil, err
	} else {
		return s.bookRepo.GetBook(updated.ID)
	}
}

func (s *BookService) Delete(id string) error {
	if err := s.bookRepo.DeleteBook(id); err != nil {
		log.Error(fmt.Sprintf("delete book id %s error, %s", id, err.Error()))
		return err
	}
	return nil
}

func (s *BookService) FillBook(id string, amount int) error {
	b, err := s.bookRepo.GetBook(id)
	if err != nil {
		return err
	}
	b.CurrentAmount = b.CurrentAmount + amount
	if _, err := s.bookRepo.UpdateBook(*b); err != nil {
		return err
	}
	return nil
}

func (s *BookService) SaleBook(id string, amount int) error {
	b, err := s.bookRepo.GetBook(id)
	if err != nil {
		return err
	}
	if b.CurrentAmount < amount {
		msg := fmt.Sprintf("insufficient stock, only %d items left", b.CurrentAmount)
		err := &bserror.InsufficientStockError{Msg: msg}
		return err
	}
	b.CurrentAmount = b.CurrentAmount - amount
	b.SoldAmount = b.SoldAmount + amount
	if _, err := s.bookRepo.UpdateBook(*b); err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBestSallBooks() ([]report.BestSallerBook, error) {
	return s.bookRepo.GetBestSaller()
}

func (s *BookService) GetBestSallCategory() ([]report.BestSallerCategory, error) {
	return s.bookRepo.GetBestSallerByCategory()
}
