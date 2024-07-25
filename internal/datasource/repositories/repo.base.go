package repositories

import (
	"back-end/internal/http/datatransfers/requests"
	"back-end/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var consoleLog = logger.ConsoleLog()

const (
	SortTypeDesc = -1
	SortTypeAsc  = 1
)

type OptionsQuery interface {
	SetOnlyFields(fieldNames ...string)
	SetPagination(pagination *requests.Pagination)
	QueryOnlyField() interface{}
	QueryPaginationLimit() *int
	QueryPaginationSkip() *int
	QuerySort() bson.D
	AddSortKey(map[string]int)
}

type optionsQuery struct {
	pagination *requests.Pagination
	sort       bson.D
	onlyFields []string
}

func NewOptions() OptionsQuery {
	return &optionsQuery{
		sort: make(bson.D, 0),
	}
}

func (o *optionsQuery) SetOnlyFields(fieldNames ...string) {
	o.onlyFields = fieldNames
}

func (o *optionsQuery) SetPagination(pagination *requests.Pagination) {
	o.pagination = pagination
}

func (o *optionsQuery) QueryOnlyField() interface{} {
	if len(o.onlyFields) < 1 {
		return nil
	}
	result := make(bson.M)
	for _, fieldName := range o.onlyFields {
		result[fieldName] = 1
	}
	return result
}

func (o *optionsQuery) QueryPaginationLimit() *int {
	if o.pagination == nil {
		return nil
	}
	return &o.pagination.PageSize
}

func (o *optionsQuery) QueryPaginationSkip() *int {
	if o.pagination == nil {
		return nil
	}
	skip := (o.pagination.PageIndex - 1) * o.pagination.PageSize
	return &skip
}

func (o *optionsQuery) AddSortKey(sorts map[string]int) {
	for sortBy, sortType := range sorts {
		if sortType != SortTypeAsc && sortType != SortTypeDesc {
			sortType = SortTypeDesc
		}
		o.sort = append(o.sort, bson.E{Key: sortBy, Value: sortType})
	}
}

func (o *optionsQuery) QuerySort() bson.D {
	return o.sort
}
