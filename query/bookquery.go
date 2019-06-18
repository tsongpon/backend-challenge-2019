package query

type BookQuery struct {
	Limit  int
	Offset int
	SortBy string
	Title  string
}
