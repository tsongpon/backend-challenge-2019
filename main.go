package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/handler"
	"github.com/tsongpon/backend-challenge-2019/repository"
	"github.com/tsongpon/backend-challenge-2019/service"
	v1handler "github.com/tsongpon/backend-challenge-2019/v1/handler"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	log.Info("starting server")
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "pingu123")

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":3306)/bookstore?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Error("database migration error", err.Error())
		panic(err.Error())
	}
	m.Steps(2)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	bookMysqlRepo := repository.NewMysqlBookRepository(db)
	bookService := service.NewBookService(bookMysqlRepo)
	bookHandler := v1handler.NewBookHandler(bookService)

	reviewMysqlRepo := repository.NewMysqlReviewRepository(db)
	reviewService := service.NewReviewService(reviewMysqlRepo)
	reviewHandler := v1handler.NewReviewHandler(reviewService)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/v1/books/:id", bookHandler.GetBook)
	e.GET("/v1/books", bookHandler.QueryBook)
	e.POST("/v1/books", bookHandler.CreateBook)
	e.PUT("/v1/books/:id", bookHandler.UpdateBook)
	e.DELETE("/v1/books/:id", bookHandler.DeleteBook)

	e.PUT("/v1/books/:id/fill", bookHandler.FillBook)
	e.PUT("/v1/books/:id/sale", bookHandler.SaleBook)

	e.GET("/v1/books/:book_id/reviews/:id", reviewHandler.GetReview)
	e.PUT("/v1/books/:book_id/reviews/:id", reviewHandler.UpdateReview)
	e.GET("/v1/books/:book_id/reviews", reviewHandler.GetBookReview)
	e.POST("/v1/books/:book_id/reviews", reviewHandler.CreateReview)
	e.DELETE("/v1/books/:book_id/reviews/:id", reviewHandler.DeleteReview)

	e.GET("/v1/reports/bestsallbook", bookHandler.GetBastSallBook)
	e.GET("/v1/reports/bestsallcategory", bookHandler.GetBastSallCategory)

	e.Logger.Fatal(e.Start(":5000"))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
