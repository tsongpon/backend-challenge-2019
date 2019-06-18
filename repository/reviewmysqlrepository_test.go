package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tsongpon/backend-challenge-2019/model"
)

func TestCreateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	bookID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	rev := model.Review{
		Score:       5,
		Description: "Very good",
		BookID:      bookID,
	}

	mock.ExpectPrepare("INSERT INTO review (.+) ").ExpectExec().
		WithArgs(anyString{}, rev.Score, rev.Description, rev.BookID,
			anyTime{}, anyTime{}, 1).WillReturnResult((sqlmock.NewResult(0, 1)))

	repo := NewMysqlReviewRepository(db)
	created, err := repo.CreateReview(rev)

	assert.Nil(t, err, "should not get any error")
	assert.NotEqual(t, "", created.ID, "new id must be generated")
	assert.NotNil(t, created.CreatedTime, "created time must be returned")
	assert.NotNil(t, created.ModifiedTime, "modified time must be returned")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	revID := "3285919c-1db4-42b8-b8a6-3cd8771dfa52"
	rows := sqlmock.NewRows([]string{
		"id",
		"score",
		"description",
		"book_id",
		"createdtime",
		"modifiedtime",
		"version"}).
		AddRow(
			"3285919c-1db4-42b8-b8a6-3cd8771dfa52",
			4,
			"Good book",
			"a432eee1-be54-44e6-a5ef-8a0455306f4f",
			time.Now(),
			time.Now(),
			1)
	mock.ExpectQuery("^SELECT (.+) FROM review (.+)").
		WithArgs(revID).WillReturnRows(rows)

	repo := NewMysqlReviewRepository(db)
	res, err := repo.GetReview(revID)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, revID, res.ID, "book ID must be "+revID)
	assert.Equal(t, 4, res.Score, "score must be 4")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createdTime := time.Now().AddDate(0, -1, 0)
	modifiedTime := time.Now()
	modelVersion := 1
	revID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	rev := model.Review{
		ID:           revID,
		Score:        4,
		Description:  "All you nee to known about golang",
		BookID:       "a432eee1-be54-44e6-a5ef-8a0455306f4f",
		CreatedTime:  &createdTime,
		ModifiedTime: &modifiedTime,
		Version:      modelVersion,
	}

	mock.ExpectPrepare("UPDATE review (.+) ").ExpectExec().
		WithArgs(rev.Score, rev.Description, rev.BookID,
			anyTime{}, modelVersion+1, rev.ID, modelVersion).
		WillReturnResult((sqlmock.NewResult(1, 1)))

	repo := NewMysqlReviewRepository(db)
	updated, err := repo.UpdateReview(rev)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, &createdTime, updated.CreatedTime, "created time must not change")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRevireByBookID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	bookID := "3285919c-1db4-42b8-b8a6-3cd8771dfa52"
	rows := sqlmock.NewRows([]string{
		"id",
		"score",
		"description",
		"book_id",
		"createdtime",
		"modifiedtime",
		"version"}).
		AddRow(
			"a432eee1-be54-44e6-a5ef-8a0455306f4f",
			4,
			"Good!",
			bookID,
			time.Now(),
			time.Now(),
			1)
	mock.ExpectQuery(`^SELECT (.+) FROM review (.+)`).
		WithArgs(bookID).WillReturnRows(rows)

	repo := NewMysqlReviewRepository(db)
	reviews, err := repo.GetReviewByBook(bookID)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, 1, len(reviews), "should have only one review")
	assert.Equal(t, bookID, reviews[0].BookID, "review ust belong to given book id")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	revID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	mock.ExpectPrepare("DELETE FROM review (.+) ").ExpectExec().
		WithArgs(revID).WillReturnResult((sqlmock.NewResult(0, 1)))

	repo := NewMysqlReviewRepository(db)
	err = repo.DeleteReview(revID)

	assert.Nil(t, err, "Should not get any error")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
