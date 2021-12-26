package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

// device_info表
type DeviceInfo struct {
	//gorm.Model
	//ID         uint   `gorm:"-;primary_key;AUTO_INCREMENT"`
	DeviceID   string `gorm:"primaryKey"`
	DeviceName string
}

// data表
type Data struct {
	//gorm.Model
	Timestamp       int64 `gorm:"primaryKey"`
	KeystrokeRecord string
}

// 连接到sqlite3数据库
func sqlite3Conn(path string) {
	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// 初始化数据库，如果数据库不存在则创建，如果存在则从中获取设备信息
func initDatabase() {
	sqlite3Conn(filepath.Join(HOMEDIR, DB))

	db.AutoMigrate(&Data{})
	db.AutoMigrate(&DeviceInfo{})

	var deviceinfo DeviceInfo
	result := db.First(&deviceinfo)
	if result.RowsAffected == 0 {
		
		deviceID := uuid.NewString()
		deviceName, err := os.Hostname()
		if err != nil {
			log.Println(err.Error())
		}

		db.Create(&DeviceInfo{DeviceID: deviceID, DeviceName: deviceName})
	} else {
		DeviceID = deviceinfo.DeviceID
		DeviceName = deviceinfo.DeviceName
	}
	return
}
