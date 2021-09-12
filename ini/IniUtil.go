package ini

import (
	"gopkg.in/ini.v1"
	"log"
)

//ini文件解析 工具类

var cfg *ini.File

func Init(filePath string) {
	var err error
	cfg, err = ini.Load(filePath)
	getErr("load config", err)
}

func getErr(msg string, err error) {
	if err != nil {
		log.Fatalf("无法读取客户端文件:%v err->%v\n", msg, err)
	}
}

func IniGetKey(SectionName string, KeyName string) string {
	return cfg.Section(SectionName).Key(KeyName).String()
}
