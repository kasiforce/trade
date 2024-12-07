package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/api"
	"github.com/kasiforce/trade/middleware"
	"net/http"
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

		v1.PUT("/profiles/info/:id", api.UpdateHandler())
		v1.POST("/login", api.UserLoginHandler())

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
		//获取商品详情
		//v1.GET("/detail", api.ShowGoodsDetailHandler())
		//退货信息
		v1.GET("/admin/afterSale", api.ShowAllrefundHandler())

		//查询所有评论
		v1.GET("/admin/comment", api.ShowAllCommentsHandler())
		//删除评论
		v1.DELETE("/admin/comment/:id", api.DeleteCommentHandler())

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.AuthToken())
		{
			authed.POST("/address", api.AddAddressHandler())
			authed.GET("/address", api.ShowAddrHandler())
			authed.PUT("/address/:id", api.UpdateAddrHandler())
			authed.PUT("/address/setDefault/:id", api.UpdateDefaultHandler())
			authed.GET("/profiles/introduction", api.ShowIntroductionHandler())
			authed.GET("/profiles/info", api.ShowUserByIDHandler())
			//用户商品查询
			authed.GET("/profiles/finished", api.IsSoldGoodsHandler())
			authed.GET("/profiles/published", api.PublishedGoodsHandler())
			authed.GET("/orders/selled", api.IsSoldGoodsHandler())
			//authed.GET("/orders/purchased", api.IsPurchasedGoodsHandler())
		}
	}
	return router
}
