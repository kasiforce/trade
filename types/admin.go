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

type UpdateAdminReq struct {
	AdminID   int    `json:"adminID"`   // 管理员ID
	AdminName string `json:"adminName"` // 管理员名称
	Passwords string `json:"password"`  // 密码
	Tel       string `json:"tel"`       // 电话号码
	Mail      string `json:"mail"`      // 邮箱
	Gender    int    `json:"gender"`    // 性别
	Age       int    `json:"age"`       // 年龄
}

type AdminLoginReq struct {
	Mail     string `form:"mail" json:"mail"`         // 邮箱
	Password string `form:"password" json:"password"` // 密码
}
