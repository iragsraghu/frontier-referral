package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")

	router.GET("/", homePage)
	router.GET("/api/v1/referrals", getAllDevices)
	router.POST("/api/v1/referrals", createDevice)
	router.GET("/api/v1/referral", getDevice)
	router.GET("/api/v1/referral/counts", getReferralCounts)

	if len(port) == 0 {
		port = "8080"
	}
	router.Run(":" + port)
	fmt.Println("Listening on port " + port)
}
