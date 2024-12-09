package types

type AdminInfo struct {
	AdminID   int    `json:"adminID"`
	Passwords string `json:"password"`
	AdminName string `json:"adminName"`
	Tel       string `json:"tel"`
	Mail      string `json:"mail"`
	Gender    int    `json:"gender"`
	Age       int    `json:"age"`
}

type AdminListResp struct {
	AdminList []AdminInfo `json:"adminList"` // 管理员列表
	Total     int         `json:"total"`     // 总记录数
	PageNum   int         `json:"pageNum"`   // 当前页码
}

type ShowAdminReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"` // 模糊搜索条件
	PageNum     int    `form:"pageNum" json:"pageNum"`         // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"`       // 每页记录数
}

type AdminLoginReq struct {
	Mail     string `form:"mail" json:"mail"`         // 邮箱
	Password string `form:"password" json:"password"` // 密码
}
