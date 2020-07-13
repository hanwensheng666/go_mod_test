package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func Init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse app.ini  err:%v", err)
		return
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	HTTPPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	// sec, err := Cfg.Section("app")
	// if err != nil {
	// 	log.Fatalf("Failed to parse app ,err:%v", err)
	// 	return
	// }
	PageSize = Cfg.Section("app").Key("PAGE_SIZE").MustInt(10)
	JwtSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
