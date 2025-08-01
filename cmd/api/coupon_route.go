package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// handles /coupon/get route
//
//	returns coupons that is associated with user. user is fetched from the context
func (app *application) get_coupons(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find any coupons",
		})
		return
	}

	coupons, err := app.models.Coupons.Get_all_coupons(user_id.(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find any coupons",
		})
		return
	}

	fmt.Println("coupons")
	fmt.Println(coupons)
	c.JSON(http.StatusOK, gin.H{
		"message": "coupons returned",
		"coupons": coupons,
	})

}

// handle /coupon/create route
// Creates a coupon associatied with user. User is fetched from the context
//
//	 Takes
//			Coupon_code       string
//			Coupon_type       string
//			Coupon_value      string
//			Coupon_visibility string
//			Coupon_duration   string
func (app *application) create_coupon(c *gin.Context) {

	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "please login again",
		})
		return
	}

	var body_coupon struct {
		Coupon_code       string `json:"coupon_code"`
		Coupon_type       string `json:"coupon_type"`
		Coupon_value      string `json:"coupon_value"`
		Coupon_visibility string `json:"coupon_visibility"`
		Coupon_duration   string `json:"coupon_duration"`
	}

	err := c.BindJSON(&body_coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request when parsing body",
		})
		return
	}

	days, err := strconv.Atoi(body_coupon.Coupon_duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please add proper expiry date",
		})
		return
	}

	_, err = app.models.Coupons.Create_coupon(
		user_id.(string),
		body_coupon.Coupon_code,
		body_coupon.Coupon_type,
		body_coupon.Coupon_value,
		body_coupon.Coupon_visibility,
		time.Now().AddDate(0, 0, days),
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Coupon code already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Coupon has been created",
	})
	return

}

// handle /coupon/redeem_coupon route
// Redeems a coupon associatied with user. User is fetched from the context
//
//	 Takes
//			Coupon_code       string
func (app *application) redeem_coupon(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find any coupons",
		})
		return
	}

	var body_coupon struct {
		Coupon_code string `json:"coupon_code"`
	}

	err := c.BindJSON(&body_coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request when parsing body",
		})
		return
	}

	err = app.models.Coupons.Redeem_coupon(user_id.(string), body_coupon.Coupon_code)
	if err != nil {
		fmt.Println("redeem err")
		fmt.Println(err)
		if strings.Contains(err.Error(), "You have already redeemed this coupon") {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "You have already redeemed this coupon",
			})
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The coupon doesn't exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully redeemed coupon",
	})
}

// handle /coupon/delete route
// Deletes a coupon associatied with user. User is fetched from the context
//
//	 Takes
//			Coupon_id       string
func (app *application) delete_coupon(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find any coupons",
		})
		return
	}

	var body_coupon struct {
		Coupon_id string `json:"coupon_id"`
	}

	err := c.BindJSON(&body_coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request when parsing body",
		})
		return
	}

	err = app.models.Coupons.Delete_coupon(user_id.(string), body_coupon.Coupon_id)
	if err != nil {
		if strings.Contains(err.Error(), "No record deleted") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No coupon deleted",
			})
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted coupon succesfully",
	})
}

// handle /coupon/redemptions route
//
//	Returns redemptions of coupons associatied with user. User is fetched from the context
func (app *application) get_redemptions(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't find any coupons",
		})
		return
	}

	redemptions, err := app.models.Redemptions.GetAllRedemptions(user_id.(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Returned redemptions",
		"redemptions": redemptions,
	})
}
