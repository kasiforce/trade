package types

type UserInfo struct {
	UserID     int    `json:"userID"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	SchoolName string `json:"schoolName"`
	Picture    string `json:"picture"`
	Tel        string `json:"tel"`
	Mail       string `json:"mail"`
	Gender     int    `json:"gender"`
	Status     int    `json:"status"`
}
type UserListResp struct {
	UsersList []UserInfo `json:"usersList"`
	Total     int        `json:"total"`
	PageNum   int        `json:"pageNum"`
}
type ShowUserReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"`
	PageNum     int    `form:"pageNum" json:"pageNum"`
	PageSize    int    `form:"pageSize" json:"pageSize"`
}

// IntroductionResp 获取个人介绍
type IntroductionResp struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	Picture  string `json:"avatarUrl"`
	Gender   int    `json:"gender"`
	School   string `json:"school"`
}

type UpdateUserReq struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	Tel      string `json:"tel"`
	Gender   int    `json:"gender"`
}

type UserLoginReq struct {
	Mail     string `form:"mail" json:"mail"`
	Password string `form:"password" json:"password"`
}

type UserRegisterReq struct {
	Mail       string `form:"mail" json:"mail"`
	Password   string `form:"password" json:"password"`
	SchoolName string `form:"schoolName" json:"schoolName"`
	Code       string `form:"code" json:"code"`
}

type UserWithToken struct {
	UserID     int    `json:"userID"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	SchoolName string `json:"schoolName"`
	Picture    string `json:"picture"`
	Tel        string `json:"tel"`
	Mail       string `json:"mail"`
	Gender     int    `json:"gender"`
	Status     int    `json:"status"`
	Token      string `json:"token"`
}

type MailCodeReq struct {
	Mail string `form:"mail" json:"mail"`
}

type PwdUpdateReq struct {
	Mail     string `form:"mail" json:"mail"`
	Password string `form:"password" json:"password"`
	Code     string `form:"code" json:"code"`
}

type UserID struct {
	UserID int `json:"userID"`
}

type ID struct {
	ID int `json:"id" form:"id"`
}
