package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	cfg            *ini.File
	AppSetting     = &app{}
	DBSetting      = &db{}
	VersionSetting = &version{}
	ServerSetting  = &server{}
)

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("settting.Setup, fail to load conf file:%v", err)
	}
	sectionMapTo("server", ServerSetting)
	sectionMapTo("app", AppSetting)
	sectionMapTo("database", DBSetting)
	sectionMapTo("version", VersionSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout*time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout*time.Second
}

func sectionMapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.sectionMapTo err: %v", err)
	}
}
