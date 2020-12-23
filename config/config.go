package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var CONF_INSTANCE *Conf

// 配置信息
type Conf struct {
	DbConf       `yaml:"db"`
	FilePath     string `yaml:"filePath"`
	GenerateConf `yaml:"generate"`
}

type DbConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	SslMode  string `yaml:"sslMode"`
	Timezone string `yaml:"timezone"`
}

type ModuleConf struct {
	Name   string   `taml:"name"`
	Tables []string `yaml:"tables"`
}

type GenerateConf struct {
	Modules []ModuleConf `yaml:"modules"`
}

func Setup(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	//fmt.Println(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &CONF_INSTANCE)
	//fmt.Println(CONF_INSTANCE)
	if err != nil {
		log.Fatal(err)
	}
}
