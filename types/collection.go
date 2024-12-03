package types

type Collection struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
	CreateTime  string  `json:"createdTime"`
}

type CollectionListResp struct {
	Collections []Collection `json:"collectionList"`
	Total       int          `json:"total"`
}

type ShowCollectionReq struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
