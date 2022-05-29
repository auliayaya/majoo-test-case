package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//use mysql
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var AppConfig *APP
var Version string

type APP struct {
	Config *viper.Viper
	DB     *gorm.DB
}

func init() {
	AppConfig = LoadConfig()
}

func LoadConfig() *APP {
	a := new(APP)
	a.readConfig()
	a.initDB()
	return a

}

func (a *APP) readConfig() {
	v := viper.New()
	v.SetConfigName("app")
	v.AddConfigPath(".")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
	a.Config = v
}

func (a *APP) initDB() {

	// Mysql DB Configuration
	DSN := a.Config.GetString("MYSQL_DB_USERNAME") + ":" + a.Config.GetString("MYSQL_DB_PASSWORD") + "@tcp(" +
		a.Config.GetString("MYSQL_DB_HOST") + ":" + a.Config.GetString("MYSQL_DB_PORT") + ")/" +
		a.Config.GetString("MYSQL_DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	a.DB = db
}
