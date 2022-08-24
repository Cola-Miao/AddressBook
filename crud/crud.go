package crud

import (
	"AddressBook/dataio"
	"AddressBook/dir"
	"AddressBook/messtmpl"
	"AddressBook/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// New 将联系人信息存入结构体后调用SAVE
func New() error {
	newC := models.Contact{}
	fmt.Println("请输入联系人姓名与电话号码")
	_, err := fmt.Scanln(&newC.Name, &newC.Number)
	if err != nil {
		fmt.Println("输入信息有误")
		return err
	}
	if dir.PathExists(dir.AddSuffix(dir.Subdirectory(dir.DataDir, models.DataFolder, newC.Name), models.DocumentSuffix)) {
		fmt.Println("已存在同名文件")
		return nil
	}
	dataio.Save(&newC)
	return nil
}

// Delete 获取联系人姓名并删除对应文件
func Delete() error {
	var delC string
	fmt.Println("请输入待删除联系人姓名")
	_, err := fmt.Scanln(&delC)
	if err != nil {
		fmt.Println("输入信息有误")
		return err
	}
	err = os.Remove(dir.AddSuffix(dir.Subdirectory(dir.DataDir, models.DataFolder, delC), models.DocumentSuffix))
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println("删除成功")
	}
	return nil
}

// Change 未实现功能
func Change() error {
	var chaC models.Contact
	fmt.Println("请输入待更改联系人姓名")
	_, err := fmt.Scanln(&chaC.Name, &chaC.Number)
	if err != nil {
		fmt.Println("输入信息有误")
		return err
	}
	dataio.Save(&chaC)
	return nil
}

// Find 读取目标路径下所有.ctt文件并输出其内容
func Find() {
	var wg sync.WaitGroup
	fList := *dataio.Load()
	if len(fList) == 0 {
		fmt.Println("Forever Alone")
	}
	for _, fInfo := range fList {
		f, err := os.Open(dir.Subdirectory(dir.DataDir, models.DataFolder, fInfo.Name()))
		if err != nil {
			fmt.Println(fInfo.Name() + "无法正常打开")
		} else {
			wg.Add(1)
			go func() {
				defer func() {
					err := f.Close()
					if err != nil {
						log.Println(err)
					}
				}()
				cttInfo, err := ioutil.ReadAll(f)
				if err != nil {
					fmt.Println("文件内容获取失败")
					log.Println(err)
				} else {
					fmt.Println("**************************************\n" + string(cttInfo))
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	messtmpl.PAC()
}
