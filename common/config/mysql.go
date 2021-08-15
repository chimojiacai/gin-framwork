package config

/*
@Time : 2020-10-14 15:55
@Author : liyongzhen
@File : mysql
@Software: GoLand
*/
import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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

	//Db, err = gorm.Open("mysql", mysqlMUrl["dnsapp"])
	Db, err = gorm.Open(mysql.Open("root:liyongzhen@tcp(127.0.0.1:3306)/gen?charset=utf8&parseTime=True&loc=Asia%2FShanghai&readTimeout=1s&timeout=1s&writeTimeout=3s"), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("mysql connect error %v", err)
	}
	err = Db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{mysql.Open("root:liyongzhen@tcp(127.0.0.1:3306)/gen?charset=utf8&parseTime=True&loc=Asia%2FShanghai&readTimeout=1s&timeout=1s&writeTimeout=3s")},
		Replicas: []gorm.Dialector{mysql.Open("root:liyongzhen@tcp(127.0.0.1:3306)/jeans?charset=utf8&parseTime=True&loc=Asia%2FShanghai&readTimeout=1s&timeout=1s&writeTimeout=3s")},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
	}))
	if err != nil {
		panic("db config not found!")
	}

	if Db.Error != nil {
		logrus.Fatalf("database error %v", Db.Error)
	}
	if cast.ToBool(mysqlMUrl["debug"]) {
		Db = Db.Debug()
	}
	//Db.DB().SetMaxIdleConns(cast.ToInt(mysqlMUrl["maxidleconns"]))
	//Db.DB().SetMaxOpenConns(cast.ToInt(mysqlMUrl["maxopenconns"]))
	//Db.DB().SetConnMaxLifetime(time.Duration(cast.ToInt64(mysqlMUrl["maxconnmaxlifetime"])) * time.Second)
	admin := Admin{Id: 30}
	Db.First(&admin)
	logrus.Info("mysql init success!")
}

// 管理员表
type Admin struct {
	Id          uint64 `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键id
	Username    string `gorm:"column:username;NOT NULL"`             // 用户名
	Password    string `gorm:"column:password;NOT NULL"`             // 密码
	Role        string `gorm:"column:role;default:CUSTOM;NOT NULL"`  // 角色
	CreateTime  int64  `gorm:"column:create_time;NOT NULL"`          // 创建时间
	UpdateTime  int64  `gorm:"column:update_time;NOT NULL"`          // 更新时间
	LoginStr    string `gorm:"column:login_str;NOT NULL"`            // 登录秘钥
	RoleId      uint64 `gorm:"column:role_id;NOT NULL"`              // 角色唯一id
	AdminStatus int    `gorm:"column:admin_status;default:1"`        // 用户状态,1=正常,2=删除
}

func (m *Admin) TableName() string {
	return "admin"
}
