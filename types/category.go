package types

type CategoryResp struct {
	CategoryID   int    `json:"categoryID"`
	CategoryName string `json:"categoryName"`
	Descriptions string `json:"description"`
}
