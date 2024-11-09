package types

type UserInfoResp struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	SchoolID int    `json:"schoolID"`
	Picture  string `json:"picture"`
	Tel      string `json:"tel"`
	Mail     string `json:"mail"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
}

type UserAddReq struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
	SchoolID int    `form:"schoolID" json:"schoolID"`
	Mail     string `form:"mail" json:"mail"`
}

type UserInfoUpdateReq struct {
	UserID   int    `url:"userID" json:"userID"`
	UserName string `form:"userName" json:"userName"`
	SchoolID int    `form:"schoolID" json:"schoolID"`
	Picture  string `form:"picture" json:"picture"`
	Tel      string `form:"tel" json:"tel"`
	Mail     string `form:"mail" json:"mail"`
	Gender   int    `form:"gender" json:"gender"`
	Status   int    `form:"status" json:"status"`
}
