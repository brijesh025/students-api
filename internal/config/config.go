package config

type HTTPServer struct {
	Address string `yaml:"address"`
	Timeout int    `yaml:"timeout"`
}
 //env-default: "production"
type Config struct {
	Env string `yaml: "env" env: "ENV" env-required: "true"`  //struct tags
	StoragePath string `yaml: "storage_path" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}