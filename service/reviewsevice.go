package service

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/repository"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository) *ReviewService {
	s := new(ReviewService)
	s.repo = reviewRepo
	return s
}

func (s *ReviewService) GetReview(id string) (*model.Review, error) {
	review, err := s.repo.GetReview(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) CreateRevirw(r model.Review) (*model.Review, error) {
	created, err := s.repo.CreateReview(r)
	if err != nil {
		log.Info("creave review error")
		return nil, err
	}
	fromDB, err := s.repo.GetReview(created.ID)
	if err != nil {
		return nil, err
	}
	return fromDB, nil
}

func (s *ReviewService) GetBookReviews(bookID string) ([]model.Review, error) {
	reviews, err := s.repo.GetReviewByBook(bookID)
	if err != nil {
		log.Error("get reviews error, bookID", bookID)
		return nil, err
	}
	return reviews, nil
}

func (s *ReviewService) UpdateReview(r model.Review) (*model.Review, error) {
	_, err := s.repo.GetReview(r.ID)
	if err != nil {
		return nil, err
	}
	updated, err := s.repo.UpdateReview(r)
	if err != nil {
		return nil, err
	}
	return s.repo.GetReview(updated.ID)
}

func (s *ReviewService) Delete(id string) error {
	if err := s.repo.DeleteReview(id); err != nil {
		log.Error(fmt.Sprintf("review book id %s error, %s", id, err.Error()))
		return err
	}
	return nil
}
