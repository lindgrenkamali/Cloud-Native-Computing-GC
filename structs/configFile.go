package structs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Database struct {
		File string `yaml:"file" envconfig:"DB_FILE"`
	} `yaml:"database"`
}

func GetConfig() {
	fileName := "Config.yml"

	e := os.Getenv("RUNENVIRONMENT")

	if len(e) > 0 {

		fileName = "dev" + fileName

	} else {
		fileName = "production" + fileName
	}

	f, _ := os.Open(fileName)
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	var config ConfigFile
	decoder.Decode(&config)
	Config = config
}

var Config ConfigFile