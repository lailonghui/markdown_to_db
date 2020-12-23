/*
@Time : 2020/12/23 16:17
@Author : lai
@Description :
@File : createTable
*/
package internal

import (
	"fmt"
	"io/ioutil"
	"markdown_to_db/config"
)

func Create() {
	tableFile := config.CONF_INSTANCE.FilePath
	data, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Printf("read tableFile err:%v\n", err)
		return
	}
	fmt.Println(string(data))
}
