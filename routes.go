package main

import (
	"fmt"
	"frontier-referral/entity"
	"frontier-referral/referral_code"
	"frontier-referral/repository"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	repo repository.DeviceRepository = repository.NewRepository()
)

// homePage is a handler function for the GET / route
func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Referral API",
	})
}

// Get all devices from firestore database and return as JSON
func getAllDevices(c *gin.Context) {
	records, err := repo.FindAll()
	if err != nil {
		log.Fatalf("Failed to get data: %v", err)
	}
	c.JSON(200, gin.H{
		"message": "Successfully fetched all devices",
		"data":    records,
	})
}

// Create a new device and return as JSON
func createDevice(c *gin.Context) {
	var record entity.Device
	device_id := c.PostForm("device_id")
	referrer_id := c.PostForm("referral_code")

	if device_id == "" || referrer_id == "" {
		c.JSON(400, gin.H{
			"message": "Device id or referral code is empty",
		})
		return
	}
	fmt.Println("First")
	unique_id := referral_code.RandomString()
	record.DeviceID = device_id
	record.UniqueID = unique_id
	record.ReferrerID = referrer_id

	existing_device, err := repo.FindDevice(device_id)
	if err != nil {
		log.Fatalf("Failed to get data: %v", err)
	}

	if existing_device != nil && existing_device.DeviceID == device_id {
		c.JSON(400, gin.H{
			"message": "Device already exists",
		})
		return
	} else {
		existing_referrer, err := repo.FindByReferrer(referrer_id)
		if err != nil {
			log.Fatalf("Failed to get data: %v", err)
		}
		if existing_referrer.UniqueID == referrer_id {
			repo.Update(referrer_id, device_id)
		}
		repo.Save(&record)
		c.JSON(200, gin.H{
			"message": "Successfully created a new device",
			"data":    record,
		})
	}

}

// Get a device by device_id and return as JSON
func getDevice(c *gin.Context) {
	device_id := c.PostForm("device_id")
	if device_id == "" {
		c.JSON(400, gin.H{
			"message": "Device id is empty",
		})
		return

	}
	record, err := repo.FindDevice(device_id)
	if err != nil {
		log.Fatalf("Failed to get data: %v", err)
	}

	if record == nil {
		c.JSON(404, gin.H{
			"message": "Device not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully fetched device",
		"data":    record,
	})
}

// Get referral counts for a device and return as JSON
func getReferralCounts(c *gin.Context) {
	device_id := c.PostForm("device_id")
	if device_id == "" {
		c.JSON(400, gin.H{
			"message": "Device id is empty",
		})
		return

	}
	record, err := repo.FindDevice(device_id)
	if err != nil {
		log.Fatalf("Failed to get data: %v", err)
	}

	if record == nil {
		c.JSON(404, gin.H{
			"message": "Device not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":        "Successfully fetched referral counts",
		"Referee Counts": len(record.RefereeIDs),
	})
}
