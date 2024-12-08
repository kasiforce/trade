package types

type AddrAddReq struct {
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"area"`
	Address   string `json:"detailArea"`
	IsDefault int    `json:"isDefault"`
}

type AddrInfoResp struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userID"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"area"`
	Address   string `json:"detailArea"`
	IsDefault int    `json:"isDefault"`
}

type AddrUpdateReq struct {
	ID        int    `url:"id" json:"id"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"area"`
	Address   string `json:"detailArea"`
	IsDefault int    `json:"isDefault"`
}

type AddrAddResp struct {
	ID int `json:"id"`
}

type UpdateDefaultReq struct {
	OldDefault int `json:"oldDefault"`
	NewDefault int `json:"newDefault"`
}
