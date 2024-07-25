package requests

type Pagination struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
