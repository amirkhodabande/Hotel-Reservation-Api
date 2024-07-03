package types

type paginationQueryParam struct {
	Page  int64 `bson:"-" validate:"omitempty,numeric,min=0"`
	Limit int64 `bson:"-" validate:"omitempty,numeric,min=0"`
}

func (pagination *paginationQueryParam) GetPage() int64 {
	if pagination.Page == 0 {
		return 1
	}

	return pagination.Page
}

func (pagination *paginationQueryParam) GetLimit() int64 {
	if pagination.Limit == 0 {
		return 10
	}

	return pagination.Limit
}
