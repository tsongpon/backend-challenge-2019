package mapper

import (
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/v1/transport"
)

func ToBookModel(t transport.BookTransport) model.Book {
	m := model.Book{
		ID:             t.ID,
		Title:          t.Title,
		Synopsis:       t.Synopsis,
		ISBN10:         t.ISBN10,
		ISBN13:         t.ISBN13,
		Language:       t.Language,
		Publisher:      t.Publisher,
		Category:       t.Category,
		Edition:        t.Edition,
		SoldAmount:     t.SoldAmount,
		CurrentAmount:  t.CurrentAmount,
		PaperbackPrice: t.PaperbackPrice,
		EbookPrice:     t.EbookPrice,
		CreatedTime:    t.CreatedTime,
		ModifiedTime:   t.ModifiedTime,
		Version:        t.Version,
	}
	return m
}

func ToBookTransport(m model.Book) transport.BookTransport {
	t := transport.BookTransport{
		ID:             m.ID,
		Title:          m.Title,
		Synopsis:       m.Synopsis,
		ISBN10:         m.ISBN10,
		ISBN13:         m.ISBN13,
		Language:       m.Language,
		Publisher:      m.Publisher,
		Category:       m.Category,
		Edition:        m.Edition,
		SoldAmount:     m.SoldAmount,
		CurrentAmount:  m.CurrentAmount,
		PaperbackPrice: m.PaperbackPrice,
		EbookPrice:     m.EbookPrice,
		CreatedTime:    m.CreatedTime,
		ModifiedTime:   m.ModifiedTime,
		Version:        m.Version,
		AverageScore:   m.AverageScore,
	}
	return t
}

func ToReviewTransport(m model.Review) transport.ReviewTransport {
	t := transport.ReviewTransport{
		ID:           m.ID,
		Score:        m.Score,
		Description:  m.Description,
		BookID:       m.BookID,
		CreatedTime:  m.CreatedTime,
		ModifiedTime: m.ModifiedTime,
		Version:      m.Version,
	}
	return t
}

func ToReviewModel(t transport.ReviewTransport) model.Review {
	m := model.Review{
		ID:           t.ID,
		Score:        t.Score,
		Description:  t.Description,
		BookID:       t.BookID,
		CreatedTime:  t.CreatedTime,
		ModifiedTime: t.ModifiedTime,
		Version:      t.Version,
	}
	return m
}
