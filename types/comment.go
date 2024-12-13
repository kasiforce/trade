package types

import "time"

// CommentInfo 表示评论信息(查询所有评论的返回结构)
type CommentInfo struct {
	CommentID       int       `json:"commentID"`             // 评价ID
	GoodsName       string    `json:"goodsName"`             // 商品ID
	CommentatorName string    `json:"commentatorName"`       // 评价者名
	CommentContent  string    `json:"commentContent"`        // 评价内容
	CommentTime     time.Time `json:"commentTime,omitempty"` // 评价时间
}

// CommentInfoByID 表示评论信息（获取发布的评价的返回结构）
type CommentInfoByID struct {
	CommentID         int       `json:"commentID"` // 评价ID
	GoodsID           int       `json:"goodsID"`   // 商品ID
	CommentatorID     int       `json:"commentatorID"`
	CommentatorName   string    `json:"commentatorName"`       // 评价者名
	CommentContent    string    `json:"commentContent"`        // 评价内容
	CommentTime       time.Time `json:"commentTime,omitempty"` // 评价时间
	CommentatorAvatar string    `json:"commentatorAvatar"`     //评论人头像
}

// ReceivedCommentInfo 表示收到的评价信息
type ReceivedCommentInfo struct {
	CommentID         int       `json:"commentID"`             // 评价ID
	GoodsID           int       `json:"goodsID"`               // 商品ID
	CommentatorID     int       `json:"commentatorID"`         // 评价者ID
	CommentatorName   string    `json:"commentatorName"`       // 评价者名
	CommentContent    string    `json:"commentContent"`        // 评价内容
	CommentTime       time.Time `json:"commentTime,omitempty"` // 评价时间
	CommentatorAvatar string    `json:"commentatorAvatar"`     //评论人头像
}

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
	CommentList interface{} `json:"commentList"`
	Total       int64       `json:"total"`
	PageNum     int         `json:"pageNum"`
}

// PostCommentReq 表示发布评论的请求
type PostCommentReq struct {
	GoodsID        int    `json:"goodsID" binding:"required"`        // 商品ID
	CommentatorID  int    `json:"commentatorID" binding:"required"`  // 评价者ID
	CommentContent string `json:"commentContent" binding:"required"` // 评价内容
}

// PostCommentResp 表示发布评论的返回信息
type PostCommentResp struct {
	Message string `json:"message"` // 返回消息
}
