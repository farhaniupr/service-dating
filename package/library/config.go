package library

import (
	"fmt"
	"time"

	"github.com/farhaniupr/dating-api/resource/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	Env      string `yaml:"env"`
	Port     string `yaml:"port"`
	Database struct {
		Host            string `yaml:"host"`
		Database        string `yaml:"database"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		MaxIdleConns    int    `yaml:"maxIdleConns"`
		MaxOpenConns    int    `yaml:"maxOpenConns"`
		ConnMaxLifeTime string `yaml:"connMaxLifeTime"`
		ConnMaxIdleTime string `yaml:"connMaxIdleTime"`
	} `yaml:"database"`
	Jwt struct {
		Key string `yaml:"key"`
	} `yaml:"jwt"`
	Redis struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

var (
	EnvGlobal Env
)

// Get config
func ModuleConfig() Env {

	viper.SetConfigFile("config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("cannot read configuration")
	}

	err = viper.Unmarshal(&EnvGlobal)
	if err != nil {
		log.Println("environment can't be loaded: ", err.Error())
	}

	return EnvGlobal
}

func ParseTimeDuration(t string, defaultt time.Duration) time.Duration {
	timeDurr, err := time.ParseDuration(t)
	if err != nil {
		return defaultt
	}
	return timeDurr
}

func DBDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		EnvGlobal.Database.Username, EnvGlobal.Database.Password, EnvGlobal.Database.Host, EnvGlobal.Database.Database)
}

func MaxIdleConns() int {
	if !viper.IsSet("database.maxIdleConns") {
		return 3
	}
	return viper.GetInt("database.maxIdleConns")
}

func MaxOpenConns() int {
	if !viper.IsSet("database.maxOpenConns") {
		return 15
	}
	return viper.GetInt("database.maxOpenConns")
}

func DBPingInterval() time.Duration {
	cfg := viper.GetString("database.pingInterval")
	return ParseTimeDuration(cfg, constants.DefaultDBPingInterval)
}

func DBRetryAttempts() int {
	if viper.GetInt("database.retryAttempts") > 0 {
		return viper.GetInt("database.retryAttempts")
	}
	return constants.DefaultDBRetryAttempts
}

func ConnMaxLifeTime() time.Duration {
	time := viper.GetString("database.connMaxLifeTime")
	return ParseTimeDuration(time, constants.DefaultConnMaxLifeTime)
}

func ConnMaxIdleTime() time.Duration {
	time := viper.GetString("database.connMaxIdleTime")
	return ParseTimeDuration(time, constants.DefaultConnMaxIdleTime)
}
