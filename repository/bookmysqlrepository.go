package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/bserror"
	"github.com/tsongpon/backend-challenge-2019/model"
	"github.com/tsongpon/backend-challenge-2019/query"
	"github.com/tsongpon/backend-challenge-2019/report"
)

type MysqlBookRepository struct {
	db *sql.DB
}

// NewMysqlBookRepository create new mysql repository
func NewMysqlBookRepository(db *sql.DB) *MysqlBookRepository {
	repo := new(MysqlBookRepository)
	repo.db = db
	return repo
}

// GetBook return book by given ID
func (r *MysqlBookRepository) GetBook(id string) (*model.Book, error) {
	sql := `SELECT 
				b.id, title, synopsis, isbn10, isbn13, language, category, publisher, 
				edition, soldamount, currentamount, paperbackprice, ebookprice, 
				b.createdtime, b.modifiedtime, b.version, avg(r.score) as averagescore
			FROM book b left join review r on b.id = r.book_id 
			WHERE b.id = ? GROUP BY b.id`
	var b model.Book
	err := r.db.QueryRow(sql, id).Scan(&b.ID, &b.Title, &b.Synopsis, &b.ISBN10, &b.ISBN13, &b.Language,
		&b.Category, &b.Publisher, &b.Edition, &b.SoldAmount, &b.CurrentAmount, &b.PaperbackPrice,
		&b.EbookPrice, &b.CreatedTime, &b.ModifiedTime, &b.Version, &b.AverageScore)
	if err != nil {
		log.Error(fmt.Sprintf("get book id %s error, %s", id, err.Error()))
		return nil, &bserror.NotFoundError{Msg: fmt.Sprintf("book id %s is not found", id)}
	}
	return &b, nil
}

// CreateBook create new book in database
func (r *MysqlBookRepository) CreateBook(b model.Book) (*model.Book, error) {
	sql := `INSERT INTO book (
			id, title, synopsis, isbn10, isbn13, language, publisher, category, edition, 
			soldamount, currentamount, paperbackprice, ebookprice, createdtime, modifiedtime, version
		) 
		values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		log.Error("prepare sql statement error, ", err.Error())
		return nil, err
	}
	defer stmt.Close()
	now := time.Now()
	b.ID = uuid.New().String()
	b.CreatedTime = &now
	b.ModifiedTime = &now
	_, err = stmt.Exec(b.ID, b.Title, b.Synopsis, b.ISBN10, b.ISBN13, b.Language, b.Publisher, b.Category, b.Edition,
		b.SoldAmount, b.CurrentAmount, b.PaperbackPrice, b.EbookPrice, b.CreatedTime, b.ModifiedTime, 1)

	if err != nil {
		log.Error("create book id ", b.ID, "error, ", err.Error())
		return nil, err
	}
	return &b, nil
}

func (r *MysqlBookRepository) UpdateBook(b model.Book) (*model.Book, error) {
	sql := `UPDATE book SET 
				title = ?,
				synopsis = ?,
				isbn10 = ?,
				isbn13 = ?,
				language = ?,
				publisher = ?,
				category = ?,
				edition = ?,
				soldamount = ?,
				currentamount = ?,
				paperbackprice = ?,
				ebookprice = ?,
				modifiedtime = ?,
				version = ?
			WHERE id = ? AND version = ?
			`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		log.Error("prepare sql statement error", err.Error())
		return nil, err
	}
	defer stmt.Close()

	nextVer := b.Version + 1
	res, err := stmt.Exec(b.Title, b.Synopsis, b.ISBN10, b.ISBN13, b.Language, b.Publisher, b.Category,
		b.Edition, b.SoldAmount, b.CurrentAmount, b.PaperbackPrice, b.EbookPrice, time.Now(), nextVer,
		b.ID, b.Version)

	if err != nil {
		log.Error(fmt.Sprintf("update book id %s error, %s", b.ID, err.Error()))
		return nil, err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return nil, &bserror.DataVersionError{Msg: "data conflict"}
	}

	return &b, nil
}

func (r *MysqlBookRepository) QueryBook(q query.BookQuery) ([]model.Book, error) {
	books := []model.Book{}
	sql := `SELECT 
				b.id, title, synopsis, isbn10, isbn13, language, publisher, category,
				edition, soldamount, currentamount, paperbackprice, ebookprice, 
				b.createdtime, b.modifiedtime, b.version, avg(r.score) as averagescore
			FROM book b left join review r on b.id = r.book_id `
	orderBy := "createdtime"
	if q.SortBy != "" {
		orderBy = q.SortBy
	}
	if q.SortBy != "" {
		orderBy = q.SortBy
	}
	order := " ORDER BY " + orderBy
	groupBy := " GROUP BY b.id"
	pagination := " LIMIT ? OFFSET ?"
	sql = sql + composeWhere(q) + groupBy + order + pagination
	result, err := r.db.Query(sql, q.Limit, q.Offset)
	if err != nil {
		log.Error("query books error", err.Error())
		return nil, err
	}

	for result.Next() {
		b := model.Book{}
		err := result.Scan(&b.ID, &b.Title, &b.Synopsis, &b.ISBN10, &b.ISBN13, &b.Language,
			&b.Publisher, &b.Category, &b.Edition, &b.SoldAmount, &b.CurrentAmount, &b.PaperbackPrice,
			&b.EbookPrice, &b.CreatedTime, &b.ModifiedTime, &b.Version, &b.AverageScore)
		if err != nil {
			log.Error("query books error", err.Error())
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *MysqlBookRepository) CountBook(q query.BookQuery) (int, error) {
	sql := "SELECT COUNT(id) as count FROM book" + composeWhere(q)
	var c int
	err := r.db.QueryRow(sql).Scan(&c)
	if err != nil {
		log.Error("count book error, ", err.Error())
		return c, err
	}
	return c, nil
}

func (r *MysqlBookRepository) DeleteBook(id string) error {
	sql := "DELETE FROM book WHERE id = ?"
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		log.Error(fmt.Sprintf("delete book id %s error, %s", id, err.Error()))
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error(fmt.Sprintf("delete book id %s error, %s", id, err.Error()))
		return err
	}
	return nil
}

func (r *MysqlBookRepository) GetBestSaller() ([]report.BestSallerBook, error) {
	rpts := []report.BestSallerBook{}
	sql := "SELECT title, MAX(soldamount) as totalamount FROM book GROUP BY title ORDER BY totalamount DESC"
	result, err := r.db.Query(sql)
	if err != nil {
		log.Error("query report error", err.Error())
		return nil, err
	}

	for result.Next() {
		each := report.BestSallerBook{}
		if err := result.Scan(&each.Ttile, &each.TotalSaleAmount); err != nil {
			return nil, err
		}
		rpts = append(rpts, each)
	}
	return rpts, nil
}

func (r *MysqlBookRepository) GetBestSallerByCategory() ([]report.BestSallerCategory, error) {
	rpts := []report.BestSallerCategory{}
	sql := "SELECT category, MAX(soldamount) as totalamount FROM book GROUP BY category ORDER BY totalamount DESC"
	result, err := r.db.Query(sql)
	if err != nil {
		log.Error("query report error", err.Error())
		return nil, err
	}

	for result.Next() {
		each := report.BestSallerCategory{}
		if err := result.Scan(&each.Category, &each.TotalSaleAmount); err != nil {
			return nil, err
		}
		rpts = append(rpts, each)
	}
	return rpts, nil
}

func composeWhere(q query.BookQuery) string {
	where := ""
	if q.Title != "" {
		where = " WHERE title = '" + q.Title + "'"
	}
	return where
}
