package entrance

import (
	"AddressBook/crud"
	"AddressBook/messtmpl"
	"fmt"
	"log"
	"os"
)

// Entrance 程序主选单
func Entrance() {
	var key string
	messtmpl.OutputMess()

	_, err := fmt.Scanln(&key)
	if err != nil {
		log.Println(err)
	} else {
		switch key {
		case "a":
			err = crud.New()
			if err != nil {
				fmt.Println("创建联系人失败")
			}
		case "b":
			err = crud.Delete()
			if err != nil {
				fmt.Println("删除联系人失败")
			}
		case "c":
			err = crud.Change() //Can't work Now
			if err != nil {
				fmt.Println("更改信息失败")
			}
		case "d":
			crud.Find()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Unreasonable input!")
		}
	}
}
