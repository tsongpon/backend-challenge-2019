package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tsongpon/backend-challenge-2019/model"
)

// start mocking book review repository //
type MockReviewRepository struct {
	mock.Mock
}

func (m *MockReviewRepository) GetReview(id string) (*model.Review, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Review), args.Error(1)
}
func (m *MockReviewRepository) GetReviewByBook(bookID string) ([]model.Review, error) {
	args := m.Called(bookID)
	return args.Get(0).([]model.Review), args.Error(1)
}
func (m *MockReviewRepository) CreateReview(rev model.Review) (*model.Review, error) {
	args := m.Called(rev)
	return args.Get(0).(*model.Review), args.Error(1)
}
func (m *MockReviewRepository) UpdateReview(rev model.Review) (*model.Review, error) {
	args := m.Called(rev)
	return args.Get(0).(*model.Review), args.Error(1)
}
func (m *MockReviewRepository) DeleteReview(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
// end mocking book review repository //

func TestGetReview(t *testing.T) {
	now := time.Now()
	mockReview := model.Review{
		ID:           "63ce552e-b750-4665-acf3-568a2e844a83",
		Score:        5,
		Description:  "Good!",
		BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
		CreatedTime:  &now,
		ModifiedTime: &now,
		Version:      1,
	}
	mockRepo := new(MockReviewRepository)
	mockRepo.On("GetReview", "63ce552e-b750-4665-acf3-568a2e844a83").Return(&mockReview, nil)

	sev := NewReviewService(mockRepo)

	rev, err := sev.GetReview("63ce552e-b750-4665-acf3-568a2e844a83")

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, "63ce552e-b750-4665-acf3-568a2e844a83", rev.ID, "review ID should be the one that return from repo")
	assert.Equal(t, 5, rev.Score, "score should be the one that return from repo")
	assert.Equal(t, mockReview.CreatedTime, rev.CreatedTime, "CreatedTime should be the one that return from repo")
	assert.Equal(t, mockReview.ModifiedTime, rev.ModifiedTime, "ModifiedTime should be the one that return from repo")
	assert.Equal(t, 1, rev.Version, "Version should be the one that return from repo")

	mockRepo.AssertExpectations(t)
}

func TestCreateReview(t *testing.T) {
	now := time.Now()
	review := model.Review{
		Score:       5,
		Description: "Good!",
		BookID:      "52b3637a-6984-401e-84d2-2aa6b9d55665",
	}
	createdReview := model.Review{
		ID:           "63ce552e-b750-4665-acf3-568a2e844a83",
		Score:        5,
		Description:  "Good!",
		BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
		CreatedTime:  &now,
		ModifiedTime: &now,
		Version:      1,
	}
	mockRepo := new(MockReviewRepository)
	mockRepo.On("CreateReview", review).Return(&createdReview, nil)
	mockRepo.On("GetReview", "63ce552e-b750-4665-acf3-568a2e844a83").Return(&createdReview, nil)

	sev := NewReviewService(mockRepo)

	created, err := sev.CreateRevirw(review)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, "63ce552e-b750-4665-acf3-568a2e844a83", created.ID, "review ID should be the one that return from repo")
	assert.Equal(t, 5, created.Score, "score should be the one that return from repo")
	assert.Equal(t, created.CreatedTime, &now, "CreatedTime should be the one that return from repo")
	assert.Equal(t, created.ModifiedTime, &now, "ModifiedTime should be the one that return from repo")
	assert.Equal(t, 1, created.Version, "Version should be the one that return from repo")

	mockRepo.AssertExpectations(t)
}

func TestGetBookReviews(t *testing.T) {
	now := time.Now()
	mockReviews := []model.Review{
		model.Review{
			ID:           "63ce552e-b750-4665-acf3-568a2e844a83",
			Score:        5,
			Description:  "Good!",
			BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
			CreatedTime:  &now,
			ModifiedTime: &now,
			Version:      1,
		},
		model.Review{
			ID:           "f5b47970-4d64-42e7-97ab-37dc4273c542",
			Score:        3,
			Description:  "Can be better",
			BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
			CreatedTime:  &now,
			ModifiedTime: &now,
			Version:      1,
		},
	}
	mockRepo := new(MockReviewRepository)
	mockRepo.On("GetReviewByBook", "52b3637a-6984-401e-84d2-2aa6b9d55665").Return(mockReviews, nil)

	sev := NewReviewService(mockRepo)

	revs, err := sev.GetBookReviews("52b3637a-6984-401e-84d2-2aa6b9d55665")

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, 2, len(revs), "should get 2 review for this book")
	assert.Equal(t, "52b3637a-6984-401e-84d2-2aa6b9d55665", revs[0].BookID, "Incorrect book")
	assert.Equal(t, "52b3637a-6984-401e-84d2-2aa6b9d55665", revs[1].BookID, "Incorrect book")

	mockRepo.AssertExpectations(t)
}

func TestUpdateReview(t *testing.T) {
	someTimeAgo := time.Now().AddDate(0, -1, 0)
	now := time.Now()
	toUpdateReview := model.Review{
		ID:           "63ce552e-b750-4665-acf3-568a2e844a83",
		Score:        5,
		Description:  "Good!",
		BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
		CreatedTime:  &someTimeAgo,
		ModifiedTime: &someTimeAgo,
		Version:      1,
	}
	updatedReview := model.Review{
		ID:           "63ce552e-b750-4665-acf3-568a2e844a83",
		Score:        5,
		Description:  "Good!",
		BookID:       "52b3637a-6984-401e-84d2-2aa6b9d55665",
		CreatedTime:  &someTimeAgo,
		ModifiedTime: &now,
		Version:      2,
	}
	mockRepo := new(MockReviewRepository)
	mockRepo.On("GetReview", "63ce552e-b750-4665-acf3-568a2e844a83").Return(&toUpdateReview, nil)
	mockRepo.On("UpdateReview", toUpdateReview).Return(&updatedReview, nil)

	sev := NewReviewService(mockRepo)
	updated, err := sev.UpdateReview(toUpdateReview)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, "63ce552e-b750-4665-acf3-568a2e844a83", updated.ID, "review ID should be the one that return from repo")
	assert.NotEqual(t, toUpdateReview.ModifiedTime, updatedReview.ModifiedTime, "ModifiedTime should be change after update")
	assert.Equal(t, toUpdateReview.CreatedTime, toUpdateReview.CreatedTime, "CreatedTime should not be change after update")

	mockRepo.AssertExpectations(t)
}

func TestDeleteReview(t *testing.T) {
	mockRepo := new(MockReviewRepository)
	mockRepo.On("DeleteReview", "63ce552e-b750-4665-acf3-568a2e844a83").Return(nil)

	sev := NewReviewService(mockRepo)
	err := sev.Delete("63ce552e-b750-4665-acf3-568a2e844a83")
	assert.Nil(t, err, "should not get any error")
}

func TestDeleteReviewWithErrorFromRepo(t *testing.T) {
	mockRepo := new(MockReviewRepository)
	mockRepo.On("DeleteReview", "63ce552e-b750-4665-acf3-568a2e844a83").Return(errors.New("there is something wrong"))

	sev := NewReviewService(mockRepo)
	err := sev.Delete("63ce552e-b750-4665-acf3-568a2e844a83")
	assert.NotNil(t, err, "should get error whne repository return error")
}
