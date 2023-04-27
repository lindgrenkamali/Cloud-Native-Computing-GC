package structs

type ConfigFile struct {
	Database struct {
		File string `yaml:"file" envconfig:"DB_FILE"`
	} `yaml:"database"`
}

var Config ConfigFile