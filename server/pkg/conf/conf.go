package conf

import (
	"strings"
	"supersign/pkg/tools"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Server  server  `mapstructure:"SERVER"`
	Log     log     `mapstructure:"LOG"`
	Storage storage `mapstructure:"STORAGE"`
	Mysql   mysql   `mapstructure:"MYSQL"`
}

type server struct {
	URL          string `mapstructure:"URL"`
	MaxJob       int    `mapstructure:"MAXJOB"`
	RunMode      string `mapstructure:"RUNMODE"`
	ReadTimeout  int    `mapstructure:"READTIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITETIMEOUT"`
	HttpPort     int    `mapstructure:"HTTPPORT"`
	TLS          bool   `mapstructure:"TLS"`
	Crt          string `mapstructure:"CRT"`
	Key          string `mapstructure:"KEY"`
}

type log struct {
	Level string `mapstructure:"LEVEL"`
}

type storage struct {
	EnableOSS          bool   `mapstructure:"ENABLEOSS"`
	BucketName         string `mapstructure:"BUCKETNAME"`
	OSSEndpoint        string `mapstructure:"OSSENDPOINT"`
	OSSAccessKeyId     string `mapstructure:"OSSACCESSKEYID"`
	OSSAccessKeySecret string `mapstructure:"OSSACCESSKEYSECRET"`
}

type mysql struct {
	Enable      bool   `mapstructure:"ENABLE"`
	Dsn         string `mapstructure:"DSN"`
	MaxIdle     int    `mapstructure:"MAXIDLE"`
	MaxOpen     int    `mapstructure:"MAXOPEN"`
	MaxLifetime int    `mapstructure:"MAXLIFETIME"`
}

type apple struct {
	AppleDeveloperPath string
	UploadFilePath     string
	TemporaryFilePath  string
}

var (
	Server  server
	Log     log
	Storage storage
	Mysql   mysql
	Apple   = apple{
		AppleDeveloperPath: "data/apple_developer/",
		UploadFilePath:     "data/upload_file_path/",
		TemporaryFilePath:  "data/temporary_file_path/",
	}
	Path string
)

// Setup 配置文件设置
func Setup() {
	if Path != "" {
		viper.SetConfigFile(Path)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("default")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := setConfig(); err != nil {
		panic(err)
	}
	mkdir([]string{Apple.AppleDeveloperPath, Apple.UploadFilePath, Apple.TemporaryFilePath})
}

func mkdir(paths []string) {
	for _, path := range paths {
		err := tools.MkdirAll(path)
		if err != nil {
			panic(err)
		}
	}
}

// Reset 配置文件重设
func Reset() error {
	return setConfig()
}

// OnChange 配置文件热加载回调
func OnChange(run func()) {
	viper.OnConfigChange(func(in fsnotify.Event) { run() })
	viper.WatchConfig()
}

// setConfig 构造配置文件到结构体对象上
func setConfig() error {
	var config config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	Server = config.Server
	Log = config.Log
	Storage = config.Storage
	Mysql = config.Mysql
	return nil
}
