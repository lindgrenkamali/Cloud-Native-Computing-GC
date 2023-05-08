package dbcontext

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Database struct {
		File     string `yaml:"file" envconfig:"DB_FILE"`
		Username string `yaml:"sql_user" envconfig:"DB_USERNAME"`
		Password string `yaml:"sql_pass" envconfig:"DB_PASSWORD"`
		Name     string `yaml:"sql_database" envconfig:"DB_DATABASE"`
		Server   string `yaml:"sql_server" envconfig:"DB_SERVER"`
		Port     int    `yaml:"sql_port" envconfig:"DB_PORT"`
	} `yaml:"database"`
}

func GetConfig() {
	fileName := "Config.yml"

	//e := os.Getenv("RUNENVIRONMENT")

	if true {
		fileName = "production" + fileName

	} else {
		fileName = "dev" + fileName
	}

	f, _ := os.Open(fileName)
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	var config ConfigFile
	decoder.Decode(&config)
	Config = config
}

var Config ConfigFile
