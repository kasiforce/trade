package types

// CommentInfo 表示评论信息
type CommentInfo struct {
	CommentatorID   int64   `json:"commentatorID"`         // 评价者ID
	CommentatorName string  `json:"commentatorName"`       // 评价者名
	CommentContent  string  `json:"commentContent"`        // 评价内容
	CommentID       int64   `json:"commentID"`             // 评价ID
	CommentTime     *string `json:"commentTime,omitempty"` // 评价时间
	GoodsID         int64   `json:"goodsID"`               // 商品ID
}

// CommentListResp 表示评论列表的响应
type CommentListResp struct {
	Comments []CommentInfo `json:"comments"` // 评论列表
	Total    int           `json:"total"`    // 总记录数
	PageNum  int           `json:"pageNum"`  // 当前页码
}

// ShowCommentsReq 表示查询评论列表的请求
type ShowCommentsReq struct {
	GoodsID       int64  `form:"goodsID" json:"goodsID"`             // 商品ID
	Commentator   string `form:"commentator" json:"commentator"`     // 评价者名模糊查询
	PageNum       int    `form:"pageNum" json:"pageNum"`             // 当前页码
	PageSize      int    `form:"pageSize" json:"pageSize"`           // 每页记录数
	CommentatorID int    `form:"commentatorID" json:"commentatorID"` // 评价者ID
}

// CreateCommentReq 表示创建评论的请求
type CreateCommentReq struct {
	CommentatorID   int64  `json:"commentatorID"`   // 评价者ID
	CommentatorName string `json:"commentatorName"` // 评价者名
	CommentContent  string `json:"commentContent"`  // 评价内容
	GoodsID         int64  `json:"goodsID"`         // 商品ID
}

// UpdateCommentReq 表示更新评论的请求
type UpdateCommentReq struct {
	CommentID      int64  `json:"commentID"`      // 评价ID
	CommentContent string `json:"commentContent"` // 评价内容
}
