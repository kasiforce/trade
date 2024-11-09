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
	v1 := router.Group("/trade")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.GET("/admin/user/:id", api.ShowUserInfoHandler())
		v1.GET("/admin/user", api.ShowAllUserHandler())
		v1.POST("admin/user", api.AddUserHandler())
		v1.PUT("/admin/user/:id", api.UpdateUserHandler())
		v1.GET("/home/category", api.ShowCategoryHandler())
	}
	return router
}
