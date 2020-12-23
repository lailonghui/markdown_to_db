package main

import (
	"flag"
	"markdown_to_db/config"
	"markdown_to_db/db"
	"markdown_to_db/internal"
)

var configFile string

func main() {
	resolveCmdParam()

	setup()

	internal.Create()
}

func setup() {
	config.Setup(configFile)
	db.Setup()
}

// 解析命令行参数
func resolveCmdParam() {
	flag.StringVar(&configFile, "f", "../config/config.yaml", "配置文件")
	flag.Parse()
}
