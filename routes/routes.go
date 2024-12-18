package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/api"
	"github.com/kasiforce/trade/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(middleware.Cors())
	router.StaticFS("/static", http.Dir("./static"))
	v1 := router.Group("")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//v1.GET("/admin/usersInfo/:id", api.ShowUserInfoHandler())
		v1.GET("/admin/usersInfo", api.ShowAllUserHandler())
		v1.POST("admin/usersInfo", api.AddUserHandler())
		v1.PUT("/admin/usersInfo/:id", api.UpdateUserHandler())
		v1.DELETE("/admin/usersInfo/:id", api.DeleteUserHandler())
		v1.GET("/admin/category", api.ShowCategoryHandler())
		v1.POST("/admin/category", api.AddCategoryHandler())
		v1.PUT("/admin/category/:id", api.UpdateCategoryHandler())
		v1.DELETE("/admin/category/:id", api.DeleteCategoryHandler())
		v1.GET("/home/category", api.ShowUserCategoryHandler())

		v1.DELETE("/address/:id", api.DeleteAddrHandler())
		v1.POST("/resetPsw", api.PwdUpdateHandler()) // 重置密码
		v1.PUT("/profiles/info/:id", api.UpdateHandler())
		v1.POST("/login", api.UserLoginHandler())
		v1.GET("/code", api.SendEmailCodeHandler())     //发送邮箱验证码
		v1.POST("/register", api.UserRegisterHandler()) //用户注册
		//管理员的增删改查
		v1.GET("/admin/adminInfo", api.ShowAllAdminHandler())
		v1.PUT("/admin/adminInfo/:id", api.UpdateAdminHandler())
		v1.POST("/admin/adminInfo", api.AddAdminHandler())
		v1.DELETE("/admin/adminInfo/:id", api.DeleteAdminHandler())

		//管理员登录
		v1.POST("/admin/login", api.AdminLoginHandler())
		//管理员查询所有商品
		v1.GET("/admin/product", api.AdminShowAllGoodsHandler())
		//删除商品
		v1.DELETE("/admin/product/:id", api.DeleteGoodsHandler())
		//管理员处理售后
		v1.POST("/admin/afterSale", api.UpdateRefundHandler())
		//退货信息
		v1.GET("/admin/afterSale", api.ShowAllrefundHandler())
		//查询所有评论
		v1.GET("/admin/comment", api.ShowAllCommentsHandler())
		//删除评论
		v1.DELETE("/admin/comment/:id", api.DeleteCommentHandler())
		//查询订单
		v1.GET("/admin/order", api.GetAllOrdersHandler())
		//商品列表
		v1.GET("/products", api.ShowAllGoodsHandler())
		//筛选商品
		v1.GET("/product/select", api.FilterGoodsHandler())
		//修改发布中商品详情
		v1.POST("/profiles/published", api.UpdateGoodsHandler())
		//用户删除商品
		v1.DELETE("/product/delete/:id", api.DeleteGoodsHandler())

		//查询所有公告
		v1.GET("/admin/announcement", api.ShowAllAnnouncementsHandler())
		//添加公告
		v1.POST("/admin/announcement", api.CreateAnnouncementHandler())
		//修改公告
		v1.PUT("/admin/announcement/:announcementID", api.UpdateAnnouncementHandler())
		//删除公告
		v1.DELETE("/admin/announcement/:announcementID", api.DeleteAnnouncementHandler())

		//支付宝支付
		v1.GET("/pay/aliPay", api.AlipayHandler)
		//支付宝消息通知回调接口
		v1.POST("/alipay/notify", api.AlipayNotifyHandler)
		//支付宝支付-成功
		v1.POST("/alipay/success",api.PaySuccessHandler())

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthToken())
		{
			authed.POST("/address", api.AddAddressHandler())
			authed.GET("/address", api.ShowAddrHandler())
			authed.PUT("/address/:id", api.UpdateAddrHandler())
			authed.PUT("/address/setDefault/:id", api.UpdateDefaultHandler())
			authed.GET("/profiles/introduction", api.ShowIntroductionHandler())
			authed.GET("/profiles/info", api.ShowUserByIDHandler())
			//获取发布的评价
			authed.GET("/profiles/comment/given", api.ShowCommentsByUserHandler())
			//根据用户ID获取收到的评价
			authed.GET("/profiles/comment/received", api.GetReceivedCommentsHandler())
			//用户商品查询 发布中和已售出
			authed.GET("/profiles/finished", api.IsSoldGoodsHandler())
			authed.GET("/profiles/published", api.PublishedGoodsHandler())
			//修改订单状态
			authed.POST("/orders/operate/:id", api.UpdateOrderStatusHandler())
			//修改订单地址
			authed.POST("/orders/address/:id", api.UpdateOrderAddressHandler())
			//生成订单
			authed.POST("/createOrder", api.CreateOrderHandler())
			//获取订单详情（以及支付结果）
			authed.GET("/order/:id", api.GetOrderDetailHandler())

			//获取-我买到的
			authed.GET("/orders/purchased", api.GetMyOrdersHandler())
			//获取-我卖出的
			authed.GET("/orders/selled", api.GetMySoldOrdersHandler())
			//获取商品详情
			authed.GET("/detail", api.IncreaseGoodsViewHandler(), api.ShowGoodsDetailHandler())
			//发布闲置
			authed.POST("/postProduct", api.CreateGoodsHandler())
			//更新收藏
			authed.PUT("/detail/:id", api.UpdateGoodsIsStarredHandler())
			//获取收藏
			authed.GET("/collection", api.ShowCollectionHandler())

			// SSE 推送公告
			v1.GET("/announcements/sse", api.SSEAnnouncementsHandler())

		}
	}
	return router
}
