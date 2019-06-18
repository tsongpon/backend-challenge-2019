package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tsongpon/backend-challenge-2019/bserror"
	"github.com/tsongpon/backend-challenge-2019/query"
	"github.com/tsongpon/backend-challenge-2019/service"
	"github.com/tsongpon/backend-challenge-2019/v1/mapper"
	"github.com/tsongpon/backend-challenge-2019/v1/transport"
)

const (
	defaultLimit  = 5
	defaultOffset = 0
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	h := new(BookHandler)
	h.service = s
	return h
}

func (h *BookHandler) GetBook(c echo.Context) error {
	b, err := h.service.GetBook(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapper.ToBookTransport(*b))
}

func (h *BookHandler) QueryBook(c echo.Context) error {
	var limit int
	var offset int
	var err error
	if limit, err = strconv.Atoi(c.QueryParam("size")); err != nil {
		limit = defaultLimit
	}
	if offset, err = strconv.Atoi(c.QueryParam("offset")); err != nil {
		offset = defaultOffset
	}
	sort := c.QueryParam("sort")
	title := c.QueryParam("title")
	q := query.BookQuery{Limit: limit, Offset: offset, Title: title, SortBy: sort}
	books, err := h.service.QueryBook(q)
	if err != nil {
		return err
	}
	total, err := h.service.CountBook(q)
	if err != nil {
		return err
	}
	bts := []transport.BookTransport{}
	for _, e := range books {
		bts = append(bts, mapper.ToBookTransport(e))
	}
	return c.JSON(http.StatusOK, transport.ResponseTransport{Data: bts, Size: len(bts), Total: total})
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	t := transport.BookTransport{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	created, err := h.service.Create(mapper.ToBookModel(t))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, mapper.ToBookTransport(*created))
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	t := transport.BookTransport{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	t.ID = c.Param("id")
	if err := c.Validate(t); err != nil {
		return err
	}
	updated, err := h.service.Update(mapper.ToBookModel(t))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapper.ToBookTransport(*updated))
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *BookHandler) FillBook(c echo.Context) error {
	id := c.Param("id")
	t := transport.FillBookTransport{}
	if err := c.Bind(&t); err != nil {
		return &bserror.BadParameterError{Msg: "invalid payload, please check"}
	}
	if err := h.service.FillBook(id, t.Amount); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *BookHandler) SaleBook(c echo.Context) error {
	id := c.Param("id")
	t := transport.SaleBookTransport{}
	if err := c.Bind(&t); err != nil {
		return &bserror.BadParameterError{Msg: "invalid payload, please check"}
	}
	if t.Amount <= 0 {
		return &bserror.BadParameterError{Msg: "amount must more than 0"}
	}
	if err := h.service.SaleBook(id, t.Amount); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *BookHandler) GetBastSallBook(c echo.Context) error {
	rpt, err := h.service.GetBestSallBooks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rpt)
}

func (h *BookHandler) GetBastSallCategory(c echo.Context) error {
	rpt, err := h.service.GetBestSallCategory()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rpt)
}
