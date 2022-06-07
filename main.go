package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/api/v1/referrals", getAllDevices)
	router.POST("/api/v1/referrals", createDevice)
	router.GET("/api/v1/referral", getDevice)
	router.GET("/api/v1/referral/counts", getReferralCounts)

	router.Run()
}
