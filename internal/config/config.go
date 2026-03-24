package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
	Timeout int    `yaml:"timeout"`
}

// env-default: "production"
type Config struct {
	Env         string     `yaml: "env" env: "ENV" env-required: "true" env-default` //struct tags
	StoragePath string     `yaml: "storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"` // these tags link to the same variables in other configuration files
}

func MustLoad() *Config{
	var configPath string=  os.Getenv(("CONFIG_PATH"))
	if (configPath == ""){
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags
		if (configPath == ""){
			log.Fatal("Config path is not set");
		}
	}

	_, e := os.Stat(configPath) // return the information of the passed file
	if(os.IsNotExist(e)){
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config;
	err := cleanenv.ReadConfig(configPath, &cfg);
	if(err!=nil){
		log.Fatalf("can not read config file: %s", err.Error())
	}
	return &cfg

}