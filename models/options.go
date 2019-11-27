package models

type FindOptions struct {
	BlogID              int64
	UserID              int64
	OrderBy, SortBy     string
	Page, Limit, Offset int
}

func (op *FindOptions) Parse() {
	if op.Page <= 0 {
		op.Page = 1
	}

	if op.SortBy == "" {
		op.SortBy = "id"
	}

	if op.OrderBy == "" {
		op.OrderBy = "desc"
	}

	if op.Limit == 0 {
		op.Limit = 20
	}

	op.Offset = (op.Page - 1) * op.Limit

}
