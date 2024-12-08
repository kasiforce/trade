package types

type CategoryInfo struct {
	CategoryID   int    `json:"categoryID"`
	CategoryName string `json:"categoryName"`
	Descriptions string `json:"description"`
}
type CategoryList struct {
	Categories []CategoryInfo `json:"categoryList"`
	Total      int            `json:"total"`
	PageNum    int            `json:"pageNum"`
}

type ShowCategoryReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"`
	PageNum     int    `form:"pageNum" json:"pageNum"`
	PageSize    int    `form:"pageSize" json:"pageSize"`
}

type CategoryID struct {
	CategoryID int `json:"categoryID"`
}
