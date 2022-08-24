package dir

import (
	"AddressBook/models"
	"fmt"
	"log"
	"os"
	"os/user"
)

var DataDir = GetAppDir()

// PathExists 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// GetAppDir 获取用户文档路径
func GetAppDir() string {
	var AppDir string
	userInfo, err := user.Current()
	if err != nil {
		fmt.Println("Fatal:获取文档路径失败")
		log.Fatalln(err)
	}
	AppDir = Subdirectory(userInfo.HomeDir, "Documents", models.AppFolder)
	return AppDir
}

// MakeDir 创建目录
func MakeDir(dataDir string) {
	err := os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		fmt.Println("Fatal：无法创建目录")
		log.Fatal(err)
	}
}

func Subdirectory(directory string, subdirectory ...string) string {
	path := directory
	for _, sub := range subdirectory {
		path += "\\" + sub
	}
	return path
}

func AddSuffix(documentPath, suffix string) string {
	return documentPath + "." + suffix
}
