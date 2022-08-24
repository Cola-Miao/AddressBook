package messtmpl

import "fmt"

func OutputMess() {
	fmt.Println()
	fmt.Println("*************ADDRESS BOOK*************")
	fmt.Println("a.新建联系人\nb.删除联系人\nc.更改信息\nd.查询目录")
	fmt.Println("输入 'exit' 退出程序")
	fmt.Println("**************************************")
	fmt.Println()
}

func PAC() {
	fmt.Println("**************************************")
	fmt.Println("输入回车继续")
	fmt.Scanln()
}
