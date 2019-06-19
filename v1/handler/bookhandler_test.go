package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/query"
	"github.com/tsongpon/backend-challenge-2019/report"
)

type MockBookService struct {
	mock.Mock
}

func (s *MockBookService) Create(b model.Book) (*model.Book, error) {
	args := s.Called(b)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (s *MockBookService) GetBook(id string) (*model.Book, error) {
	args := s.Called(id)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (s *MockBookService) QueryBook(q query.BookQuery) ([]model.Book, error) {
	args := s.Called(q)
	return args.Get(0).([]model.Book), args.Error(1)
}

func (s *MockBookService) CountBook(q query.BookQuery) (int, error) {
	args := s.Called(q)
	return args.Get(0).(int), args.Error(1)
}

func (s *MockBookService) Update(b model.Book) (*model.Book, error) {
	args := s.Called(b)
	return args.Get(0).(*model.Book), args.Error(1)
}

func (s *MockBookService) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *MockBookService) FillBook(id string, amount int) error {
	args := s.Called(id, amount)
	return args.Error(0)
}

func (s *MockBookService) SaleBook(id string, amount int) error {
	args := s.Called(id, amount)
	return args.Error(0)
}

func (s *MockBookService) GetBestSallBooks() ([]report.BestSallerBook, error) {
	args := s.Called()
	return args.Get(0).([]report.BestSallerBook), args.Error(1)
}

func (s *MockBookService) GetBestSallCategory() ([]report.BestSallerCategory, error) {
	args := s.Called()
	return args.Get(0).([]report.BestSallerCategory), args.Error(1)
}

func TestGetBook(t *testing.T) {
	mockSev := new(MockBookService)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("147efcdb-5df0-4d3b-8a3b-cfca77fbd994")
	h := &BookHandler{mockSev}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
