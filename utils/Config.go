package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"

	hocon "github.com/go-akka/configuration"
)

// HoconConfig encapsulates application's configurations in HOCON format
type HoconConfig struct {
	File string        // config file
	Conf *hocon.Config // configurations
}

const (
	defaultConfigFile = "/config/application.conf"
)

func loadAppConfig(file string) *HoconConfig {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	defer os.Chdir(dir)

	config := HoconConfig{}
	log.Infof("Loading configurations from file [%s]", file)
	confDir, confFile := path.Split(file)
	os.Chdir(confDir)
	config.File = file
	config.Conf = hocon.LoadConfig(confFile)
	return &config
}

func InitAppConfig() {
	configFile := os.Getenv("APP_CONFIG")
	if configFile == "" {
		log.Infof("No environment APP_CONFIG found, fallback to [%s]", defaultConfigFile)
		_, b, _, _ := runtime.Caller(0)
		d := path.Join(path.Dir(b))
		configFile = filepath.Dir(d) + defaultConfigFile
	}

	AppConfig = loadAppConfig(configFile)
}
