package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/bserror"
	"github.com/tsongpon/backend-challenge-2019/model"
)

type MysqlReviewRepository struct {
	db *sql.DB
}

func NewMysqlReviewRepository(db *sql.DB) *MysqlReviewRepository {
	repo := new(MysqlReviewRepository)
	repo.db = db
	return repo
}

func (r *MysqlReviewRepository) CreateReview(review model.Review) (*model.Review, error) {
	sql := `INSERT INTO review (
		id, score, description, book_id, createdtime, modifiedtime, version
	) values(?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		log.Error("prepare sql statement error, ", err.Error())
		return nil, err
	}
	defer stmt.Close()
	now := time.Now()
	review.ID = uuid.New().String()
	review.CreatedTime = &now
	review.ModifiedTime = &now
	_, err = stmt.Exec(review.ID, review.Score, review.Description,
		review.BookID, review.CreatedTime, review.ModifiedTime, 1)

	if err != nil {
		log.Error("create book id ", review.ID, "error, ", err.Error())
		return nil, err
	}
	return &review, nil
}

func (r *MysqlReviewRepository) UpdateReview(review model.Review) (*model.Review, error) {
	sql := `UPDATE review SET 
				score = ?,
				description = ?,
				book_id =?,
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

	nextVer := review.Version + 1
	res, err := stmt.Exec(review.Score, review.Description, review.BookID, time.Now(), nextVer, review.ID, review.Version)

	if err != nil {
		log.Error(fmt.Sprintf("update review id %s error, %s", review.ID, err.Error()))
		return nil, err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return nil, &bserror.DataVersionError{Msg: "data conflict"}
	}

	return &review, nil
}

func (r *MysqlReviewRepository) GetReview(id string) (*model.Review, error) {
	sql := `SELECT 
				id, score, description, book_id, createdtime, modifiedtime, version
			FROM review 
			WHERE id = ?`
	var review model.Review
	err := r.db.QueryRow(sql, id).Scan(&review.ID, &review.Score, &review.Description,
		&review.BookID, &review.CreatedTime, &review.ModifiedTime, &review.Version)
	if err != nil {
		log.Error(fmt.Sprintf("get review id %s error, %s", id, err.Error()))
		return nil, &bserror.NotFoundError{Msg: fmt.Sprintf("review id %s is not found", id)}
	}
	return &review, nil
}

func (r *MysqlReviewRepository) GetReviewByBook(bookID string) ([]model.Review, error) {
	reviews := []model.Review{}
	sql := `SELECT 
				id, score, description, book_id, createdtime, modifiedtime, version
			FROM review 
			WHERE book_id = ?`
	result, err := r.db.Query(sql, bookID)
	if err != nil {
		log.Error("query review error", err.Error())
		return nil, err
	}

	for result.Next() {
		rev := model.Review{}
		err := result.Scan(&rev.ID, &rev.Score, &rev.Description, &rev.BookID,
			&rev.CreatedTime, &rev.ModifiedTime, &rev.Version)
		if err != nil {
			log.Error("query reviews error", err.Error())
			return nil, err
		}
		reviews = append(reviews, rev)
	}
	return reviews, nil
}

func (r *MysqlReviewRepository) DeleteReview(id string) error {
	sql := "DELETE FROM review WHERE id = ?"
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		log.Error(fmt.Sprintf("delete review id %s error, %s", id, err.Error()))
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error(fmt.Sprintf("delete review id %s error, %s", id, err.Error()))
		return err
	}
	return nil
}
