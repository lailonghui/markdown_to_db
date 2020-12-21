package main

import (
	"db_doc_generator/config"
	"db_doc_generator/db"
	"db_doc_generator/internal"
	"flag"
)

var configFile string

func main() {
	resolveCmdParam()

	setup()

	internal.Gen()
}

func setup() {
	config.Setup(configFile)
	db.Setup()
}

// 解析命令行参数
func resolveCmdParam() {
	flag.StringVar(&configFile, "f", "config/config.yaml", "配置文件")
	flag.Parse()
}
