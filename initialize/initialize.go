package initialize

import (
	"AddressBook/dir"
	"AddressBook/models"
	"fmt"
	"log"
	"os"
)

// DirInit 初始化文件路径
func DirInit() {
	dataDir := dir.Subdirectory(dir.DataDir, models.DataFolder)
	if !dir.PathExists(dataDir) {
		dir.MakeDir(dataDir)
	}
}

// LogInit 初始化日志设置
func LogInit() {
	logsDir := dir.Subdirectory(dir.DataDir, models.LogsFolder)
	if !dir.PathExists(logsDir) {
		dir.MakeDir(logsDir)
	}
	logPath := dir.AddSuffix(dir.Subdirectory(logsDir, "Log"), models.TextSuffix)
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println("*重要错误*无法创建日志文件", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
