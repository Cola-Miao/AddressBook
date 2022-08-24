package dataio

import (
	"AddressBook/dir"
	"AddressBook/models"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

// Save 将接收值模板化后存入目标路径
func Save(newC *models.Contact) {
	//创建文件
	f, err := os.Create(dir.AddSuffix(dir.Subdirectory(dir.DataDir, models.DataFolder, newC.Name), models.DocumentSuffix))
	if err != nil {
		fmt.Println("创建文件失败")
		log.Println(err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	//模板化并储存
	t, err := template.ParseFiles(models.TemplatePath)
	if err != nil {
		fmt.Println("信息模板化失败")
		log.Println(err)
		return
	}
	err = t.Execute(f, newC)
	if err != nil {
		fmt.Println("信息储存失败")
		log.Println(err)
	} else {
		fmt.Println("信息储存成功")
	}
}

// Load 从目标路径读取文件列表
func Load() *[]fs.FileInfo {
	fList, err := ioutil.ReadDir(dir.Subdirectory(dir.DataDir, models.DataFolder))
	if err != nil {
		fmt.Println("无法读取文档列表")
		log.Println(err)
		return nil
	}
	return &fList
}
