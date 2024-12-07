package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System           `yaml:"system"`
	MySql  map[string]*MySql `yaml:"mysql"`
	Oss    *Oss              `yaml:"oss"`
	Email  *Email            `yaml:"email"`
	Redis  *Redis            `yaml:"redis"`
	//Cache     *Cache            `yaml:"cache"`
	//PhotoPath *LocalPhotoPath   `yaml:"photoPath"`
}

type System struct {
	AppEnv      string `yaml:"appEnv"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"host"`
	UploadModel string `yaml:"uploadModel"`
}

type MySql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Email struct {
	ValidEmail string `yaml:"validEmail"`
	SmtpHost   string `yaml:"smtpHost"`
	SmtpEmail  string `yaml:"smtpEmail"`
	SmtpPass   string `yaml:"smtpPass"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

type Oss struct {
	BucketName  string `yaml:"bucketName"`
	AccessKey   string `yaml:"accessKey"`
	SecretKey   string `yaml:"SecretKey"`
	Endpoint    string `yaml:"endPoint"`
	EndpointOut string `yaml:"endpointOut"`
	QiNiuServer string `yaml:"qiNiuServer"`
}

type LocalPhotoPath struct {
	PhotoHost   string `yaml:"photoHost"`
	ProductPath string `yaml:"productPath"`
	AvatarPath  string `yaml:"avatarPath"`
}

type Cache struct {
	CacheType    string `yaml:"cacheType"`
	CacheExpires int64  `yaml:"cacheExpires"`
	CacheWarmUp  bool   `yaml:"cacheWarmUp"`
	CacheServer  string `yaml:"cacheServer"`
}

//读取配置文件

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

//获取缓存过期时间
//func GetExpiresTime() int64 {
//	if Config.Cache.CacheExpires == 0 {
//		return int64(30 * time.Minute)
//	}
//
//	if Config.Cache.CacheExpires == -1 {
//		return -1
//	}
//	return int64(time.Duration(Config.Cache.CacheExpires) * time.Minute)
//}
