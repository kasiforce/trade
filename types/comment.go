package types

import "time"

// CommentInfo 表示评论信息
type CommentInfo struct {
	CommentID       int       `json:"commentID"`             // 评价ID
	GoodsName       int       `json:"goodsName"`             // 商品ID
	CommentatorName string    `json:"commentatorName"`       // 评价者名
	CommentContent  string    `json:"commentContent"`        // 评价内容
	CommentTime     time.Time `json:"commentTime,omitempty"` // 评价时间
}

// CommentListResp 表示评论列表的响应
//type CommentListResp struct {
//	Comments []CommentInfo `json:"comments"` // 评论列表
//	Total    int           `json:"total"`    // 总记录数
//	PageNum  int           `json:"pageNum"`  // 当前页码
//}

// ShowCommentsReq 表示查询评论列表的请求
type ShowCommentsReq struct {
	SearchQuery string `form:"searchQuery" json:"searchQuery"`
	PageNum     int    `form:"pageNum" json:"pageNum"`   // 当前页码
	PageSize    int    `form:"pageSize" json:"pageSize"` // 每页记录数
}

// CreateCommentReq 表示创建评论的请求
type CreateCommentReq struct {
	CommentatorID   int    `json:"commentatorID"`   // 评价者ID
	CommentatorName string `json:"commentatorName"` // 评价者名
	CommentContent  string `json:"commentContent"`  // 评价内容
	GoodsID         int    `json:"goodsID"`         // 商品ID
}

// UpdateCommentReq 表示更新评论的请求
//type UpdateCommentReq struct {
//	CommentID      int    `json:"commentID"`      // 评价ID
//	CommentContent string `json:"commentContent"` // 评价内容
//}

// CommentListResp 服务返回结构
type CommentListResp struct {
	CommentList interface{} `json:"item"`
	Total       int64       `json:"total"`
	PageNum     int         `json:"pageNum"`
}
