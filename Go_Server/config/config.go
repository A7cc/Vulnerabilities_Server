package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// 总配置文件
type config struct {
	Server       server `yaml:"server"`
	Db           db     `yaml:"db"`
	UploadDir    string `yaml:"uploadDir"`
	Log          string `yaml:"logpath"`
	StaticData   string `yaml:"staticData"`
	BackupsdbDir string `yaml:"backupsdbDir"`
	ZipDir       string `yaml:"zipDir"`
}

// 项目端口配置
type server struct {
	Port string `yaml:"port"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// 全局配置文件
var Config *config

// 初始化配置
func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	// 有错就down机
	if err != nil {
		panic(err)
	}
	// 绑定值

	if err = yaml.Unmarshal(yamlFile, &Config); err != nil {
		panic(err)
	}
}
