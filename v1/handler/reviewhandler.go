package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tsongpon/backend-challenge-2019/service"
	"github.com/tsongpon/backend-challenge-2019/v1/mapper"
	"github.com/tsongpon/backend-challenge-2019/v1/transport"
)

type ReviewHandler struct {
	service *service.ReviewService
}

func NewReviewHandler(s *service.ReviewService) *ReviewHandler {
	h := new(ReviewHandler)
	h.service = s
	return h
}

func (h *ReviewHandler) GetReview(c echo.Context) error {
	review, err := h.service.GetReview(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapper.ToReviewTransport(*review))
}

func (h *ReviewHandler) UpdateReview(c echo.Context) error {
	bookID := c.Param("book_id")
	rt := transport.ReviewTransport{}
	if err := c.Bind(&rt); err != nil {
		return err
	}
	rt.BookID = bookID
	updated, err := h.service.UpdateReview(mapper.ToReviewModel(rt))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapper.ToReviewTransport(*updated))
}

func (h *ReviewHandler) GetBookReview(c echo.Context) error {
	reviews, err := h.service.GetBookReviews(c.Param("book_id"))
	if err != nil {
		return err
	}
	rts := []transport.ReviewTransport{}
	for _, e := range reviews {
		rts = append(rts, mapper.ToReviewTransport(e))
	}
	return c.JSON(http.StatusOK, rts)
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	bookID := c.Param("book_id")
	rt := transport.ReviewTransport{}
	if err := c.Bind(&rt); err != nil {
		return err
	}
	rt.BookID = bookID
	created, err := h.service.CreateRevirw(mapper.ToReviewModel(rt))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, mapper.ToReviewTransport(*created))
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
