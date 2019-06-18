package repository

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/query"
)

type anyString struct{}
type anyTime struct{}

func (a anyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	bookID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"synopsis",
		"isbn10",
		"isbn13",
		"category",
		"language",
		"publisher",
		"edition",
		"soldamount",
		"currentamount",
		"paperbackprice",
		"ebookprice",
		"createdtime",
		"modifiedtime",
		"version",
		"averagescore"}).
		AddRow(
			"a432eee1-be54-44e6-a5ef-8a0455306f4f",
			"Java Concurrency in Practice",
			"Threads are a fundamental part of the Java platform",
			"0321349601",
			"978-0321349606",
			"Programming",
			"English",
			"Addison-Wesley Professional",
			"1nd Edition, Kindle Edition",
			0,
			100,
			1353.29,
			1210.5,
			time.Now(),
			time.Now(),
			1,
			4.5)
	mock.ExpectQuery("^SELECT (.+) FROM book b left join review r on b.id = r.book_id (.+)").
		WithArgs(bookID).WillReturnRows(rows)

	repo := NewMysqlBookRepository(db)
	res, err := repo.GetBook(bookID)
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Equal(t, bookID, res.ID, "bookID must be "+bookID)
	assert.Equal(t, "Java Concurrency in Practice", res.Title, "should book title Java Concurrency in Practice")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	paperbackPrice := 1353.29
	ebookPrice := 1100.00
	b := model.Book{
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
		Version:        1,
	}

	mock.ExpectPrepare("INSERT INTO book (.+) ").ExpectExec().
		WithArgs(anyString{}, b.Title, b.Synopsis, b.ISBN10, b.ISBN13,
			b.Language, b.Publisher, b.Category, b.Edition,
			b.SoldAmount, b.CurrentAmount, b.PaperbackPrice, b.EbookPrice,
			anyTime{}, anyTime{}, b.Version).WillReturnResult((sqlmock.NewResult(0, 1)))

	repo := NewMysqlBookRepository(db)
	created, err := repo.CreateBook(b)

	assert.Nil(t, err, "Should not get any error")
	assert.NotEqual(t, created.ID, "", "there is no generated id after created")
	assert.NotNil(t, created.ModifiedTime, "no mofified time afetr created")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	createdTime := time.Now().AddDate(0, -1, 0)
	modifiedTime := time.Now()
	modelVersion := 1
	bookID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	paperbackPrice := 1353.29
	ebookPrice := 1100.00
	b := model.Book{
		ID:             bookID,
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
		CreatedTime:    &createdTime,
		ModifiedTime:   &modifiedTime,
		Version:        modelVersion,
	}

	mock.ExpectPrepare("UPDATE book (.+) ").ExpectExec().
		WithArgs(b.Title, b.Synopsis, b.ISBN10, b.ISBN13,
			b.Language, b.Publisher, b.Category, b.Edition,
			b.SoldAmount, b.CurrentAmount, b.PaperbackPrice, b.EbookPrice,
			anyTime{}, modelVersion+1, b.ID, modelVersion).WillReturnResult((sqlmock.NewResult(1, 1)))

	repo := NewMysqlBookRepository(db)
	updated, err := repo.UpdateBook(b)

	assert.Nil(t, err, "Sould not get any error")
	assert.Equal(t, &createdTime, updated.CreatedTime, "created time must not change")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestQueryBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"synopsis",
		"isbn10",
		"isbn13",
		"category",
		"language",
		"publisher",
		"edition",
		"soldamount",
		"currentamount",
		"paperbackprice",
		"ebookprice",
		"createdtime",
		"modifiedtime",
		"version",
		"averagescore"}).
		AddRow(
			"a432eee1-be54-44e6-a5ef-8a0455306f4f",
			"Java Concurrency in Practice",
			"Threads are a fundamental part of the Java platform",
			"0321349601",
			"978-0321349606",
			"Programming",
			"English",
			"Addison-Wesley Professional",
			"1nd Edition, Kindle Edition",
			0,
			100,
			1353.29,
			1210.5,
			time.Now(),
			time.Now(),
			1,
			4.5)
	mock.ExpectQuery(`^SELECT (.+) FROM book b left join review r on b.id = r.book_id 
	WHERE title = 'Java Concurrency in Practice' (.+) ORDER BY category (.+)`).
		WithArgs(5, 0).WillReturnRows(rows)

	q := query.BookQuery{Limit: 5,
		Offset: 0,
		Title:  "Java Concurrency in Practice",
		SortBy: "category",
	}
	repo := NewMysqlBookRepository(db)
	books, err := repo.QueryBook(q)

	assert.Nil(t, err, "Should not get any error")
	assert.Equal(t, 1, len(books), "should have only one book return")
	assert.Equal(t, "a432eee1-be54-44e6-a5ef-8a0455306f4f", books[0].ID, "returned book not correct")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCountBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow(1)
	mock.ExpectQuery(`^SELECT COUNT(.+) FROM book WHERE title = 'Java Concurrency in Practice'`).
		WillReturnRows(rows)

	q := query.BookQuery{Limit: 5,
		Offset: 0,
		Title:  "Java Concurrency in Practice",
		SortBy: "category",
	}

	repo := NewMysqlBookRepository(db)
	c, err := repo.CountBook(q)

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, 1, c, "expected 1 from count")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	bookID := "a432eee1-be54-44e6-a5ef-8a0455306f4f"
	mock.ExpectPrepare("DELETE FROM book (.+) ").ExpectExec().
		WithArgs(bookID).WillReturnResult((sqlmock.NewResult(0, 1)))

	repo := NewMysqlBookRepository(db)
	err = repo.DeleteBook(bookID)

	assert.Nil(t, err, "should not get any error")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetBestSaller(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"title", "totalamount"}).
		AddRow("Java in action", 100).
		AddRow("Nodejs is the best", 1)
	mock.ExpectQuery("^SELECT (.+) FROM book (.+) ").WillReturnRows(rows)

	repo := NewMysqlBookRepository(db)
	res, err := repo.GetBestSaller()

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, "Java in action", res[0].Ttile, "Incorrect book title returned")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetBestSallerByCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"category", "totalamount"}).
		AddRow("Programming", 100).
		AddRow("Big Data", 1)
	mock.ExpectQuery("^SELECT category, (.+) FROM book (.+)").WillReturnRows(rows)

	repo := NewMysqlBookRepository(db)
	res, err := repo.GetBestSallerByCategory()

	assert.Nil(t, err, "should not get any error")
	assert.Equal(t, "Programming", res[0].Category, "first top saller category is Programming")
	assert.Equal(t, "Big Data", res[1].Category, "second top saller category is Big Data")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
