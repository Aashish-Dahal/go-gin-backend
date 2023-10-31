package paginations

type FeedPagination struct {
	Pagination
	Keyword2 string `form:"keyword2"`
}
