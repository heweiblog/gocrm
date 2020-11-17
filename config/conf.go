package config

import (
	"encoding/json"
	"log"
	"os"
)

type Listen struct {
	Ip   string
	Port int
}

type Mysql struct {
	Listen
	User     string
	Pass     string
	Database string
}

type Bind struct {
	Listen
	Conf string
}

type Ems struct {
	Ip string
}

type Out struct {
	Conf int
	Task int
}

type Config struct {
	Net     Listen
	Sql     Mysql
	Ybind   Bind
	Nap     Listen
	Ms      Ems
	Timeout Out
}

var (
	ConfFile = "/home/heweiwei/go/src/gocrm/config/gocrm.json"
	Version  = "1.0.0"
	Conf     *Config
)

func init() {
	// 读取配置出错直接抛出panic
	filePtr, err := os.Open(ConfFile)
	if err != nil {
		log.Panic("Open file failed [Err:%s]", err.Error())
	}
	defer filePtr.Close()

	Conf = new(Config)

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(Conf)
	if err != nil {
		log.Panic("Decoder failed", err.Error())
	}
}
