package database

import (
	"gorm.io/gorm"
	"log"
)

// 查询有哪些宿主机
func GetBeacons(db *gorm.DB, value *[]Beacon) {
	result := db.Find(value)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
}

// 查询一个宿主机的所有Data
func GetBeaconData(db *gorm.DB, value *[]BeaconData, uuid string) {
	//result := db.Find(value)
	result := db.Where("beacon_id = ?", uuid).Find(value)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}
}

// 查询是否存在某宿主机
func ExistQuerry(db *gorm.DB, uuid string) bool {
	beacon := &Beacon{DeviceID: uuid}

	result := db.Find(beacon)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}
