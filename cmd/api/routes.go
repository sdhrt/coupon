package main

import (
	"coupon/cmd/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register_routes functions register different routes / endpoints
func (app *application) register_routes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.POST("/register", app.register_user)
	router.POST("/login", app.login_user)

	coupon_route := router.Group("/coupon")
	// This middleware is used to check for access token in every route starting with /coupon
	coupon_route.Use(middleware.AuthMiddleware(app.config.jwt_secret))
	{
		coupon_route.POST("/get", app.get_coupons)
		coupon_route.POST("/create", app.create_coupon)
		coupon_route.POST("/delete", app.delete_coupon)
		coupon_route.POST("/redeem", app.redeem_coupon)
		coupon_route.POST("/redemptions", app.get_redemptions)
	}
}
