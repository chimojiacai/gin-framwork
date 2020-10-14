package config

/*
@Time : 2020-10-14 15:55
@Author : liyongzhen
@File : mysql
@Software: GoLand
*/
import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"time"
)

// 全局MySQL变量
var Db *gorm.DB

func MysqlInit() {
	var err error
	mysqlMUrl := viper.GetStringMap("mysql")
	if _, ok := mysqlMUrl["dnsapp"]; !ok {
		logrus.Fatal("mysql config not exits")
	}

	if mysqlMUrl["dnsapp"] == "" {
		logrus.Fatal("mysql config is empty")
	}

	Db, err = gorm.Open("mysql", mysqlMUrl["dnsapp"])

	if err != nil {
		logrus.Fatalf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		logrus.Fatalf("database error %v", Db.Error)
	}
	if cast.ToBool(mysqlMUrl["debug"]) {
		Db = Db.Debug()
	}
	Db.DB().SetMaxIdleConns(cast.ToInt(mysqlMUrl["maxidleconns"]))
	Db.DB().SetMaxOpenConns(cast.ToInt(mysqlMUrl["maxopenconns"]))
	Db.DB().SetConnMaxLifetime(time.Duration(cast.ToInt64(mysqlMUrl["maxconnmaxlifetime"])) * time.Second)
	logrus.Info("mysql init success!")
}
