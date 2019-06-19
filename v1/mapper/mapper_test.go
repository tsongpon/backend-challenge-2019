package mapper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/v1/transport"
)

func TestToBookTransport(t *testing.T) {
	now := time.Now()
	paperbackPrice := 1353.29
	ebookPrice := 1100.00
	m := model.Book{
		ID:             "a432eee1-be54-44e6-a5ef-8a0455306f4f",
		Title:          "The Go Programming",
		Synopsis:       "All you nee to known about golang",
		ISBN10:         "0321349601",
		ISBN13:         "978-0321349606",
		Language:       "Thai",
		Publisher:      "Addison-Wesley Professional",
		Category:       "Programming",
		Edition:        "1nd Edition, Kindle Edition",
		SoldAmount:     0,
		CurrentAmount:  10,
		PaperbackPrice: &paperbackPrice,
		EbookPrice:     &ebookPrice,
		CreatedTime:    &now,
		ModifiedTime:   &now,
		Version:        1,
	}
	tsp := ToBookTransport(m)

	assert.Equal(t, "a432eee1-be54-44e6-a5ef-8a0455306f4f", tsp.ID)
	assert.Equal(t, "The Go Programming", tsp.Title)
	assert.Equal(t, "All you nee to known about golang", tsp.Synopsis)
	assert.Equal(t, "0321349601", tsp.ISBN10)
	assert.Equal(t, &now, tsp.CreatedTime)
	assert.Equal(t, &now, tsp.ModifiedTime)
	assert.Equal(t, 1, tsp.Version)
}

func TestToBookModel(t *testing.T) {
	now := time.Now()
	paperbackPrice := 1353.29
	ebookPrice := 1100.00
	tsp := transport.BookTransport{
		ID:             "a432eee1-be54-44e6-a5ef-8a0455306f4f",
		Title:          "The Go Programming",
		Synopsis:       "All you nee to known about golang",
		ISBN10:         "0321349601",
		ISBN13:         "978-0321349606",
		Language:       "Thai",
		Publisher:      "Addison-Wesley Professional",
		Category:       "Programming",
		Edition:        "1nd Edition, Kindle Edition",
		SoldAmount:     0,
		CurrentAmount:  10,
		PaperbackPrice: &paperbackPrice,
		EbookPrice:     &ebookPrice,
		CreatedTime:    &now,
		ModifiedTime:   &now,
		Version:        1,
	}
	m := ToBookModel(tsp)

	assert.Equal(t, "a432eee1-be54-44e6-a5ef-8a0455306f4f", m.ID)
	assert.Equal(t, "The Go Programming", m.Title)
	assert.Equal(t, "All you nee to known about golang", m.Synopsis)
	assert.Equal(t, "0321349601", m.ISBN10)
	assert.Equal(t, &now, m.CreatedTime)
	assert.Equal(t, &now, m.ModifiedTime)
	assert.Equal(t, 1, m.Version)
}

func TestToReviewTransport(t *testing.T) {
	createdTime := time.Now().AddDate(0, -1, 0)
	modifiedTime := time.Now()
	modelVersion := 1
	rev := model.Review{
		ID:           "dfa4f5c3-1866-4807-963c-f8d991753769",
		Score:        4,
		Description:  "All you nee to known about golang",
		BookID:       "a432eee1-be54-44e6-a5ef-8a0455306f4f",
		CreatedTime:  &createdTime,
		ModifiedTime: &modifiedTime,
		Version:      modelVersion,
	}

	tsp := ToReviewTransport(rev)

	assert.Equal(t, "dfa4f5c3-1866-4807-963c-f8d991753769", tsp.ID)
	assert.Equal(t, 4, tsp.Score)
	assert.Equal(t, "All you nee to known about golang", tsp.Description)
	assert.Equal(t, "a432eee1-be54-44e6-a5ef-8a0455306f4f", tsp.BookID)
	assert.Equal(t, &createdTime, tsp.CreatedTime)
	assert.Equal(t, &modifiedTime, tsp.ModifiedTime)
	assert.Equal(t, 1, tsp.Version)
}

func TestToReviewModel(t *testing.T) {
	createdTime := time.Now().AddDate(0, -1, 0)
	modifiedTime := time.Now()
	modelVersion := 1
	tsp := transport.ReviewTransport{
		ID:           "dfa4f5c3-1866-4807-963c-f8d991753769",
		Score:        4,
		Description:  "All you nee to known about golang",
		BookID:       "a432eee1-be54-44e6-a5ef-8a0455306f4f",
		CreatedTime:  &createdTime,
		ModifiedTime: &modifiedTime,
		Version:      modelVersion,
	}

	rev := ToReviewModel(tsp)

	assert.Equal(t, "dfa4f5c3-1866-4807-963c-f8d991753769", rev.ID)
	assert.Equal(t, 4, rev.Score)
	assert.Equal(t, "All you nee to known about golang", rev.Description)
	assert.Equal(t, "a432eee1-be54-44e6-a5ef-8a0455306f4f", rev.BookID)
	assert.Equal(t, &createdTime, rev.CreatedTime)
	assert.Equal(t, &modifiedTime, rev.ModifiedTime)
	assert.Equal(t, 1, rev.Version)
}
