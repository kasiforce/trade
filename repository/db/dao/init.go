package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	conf "github.com/kasiforce/trade/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"strings"
	"time"
)

// GORM 提供了自动迁移（AutoMigrate）功能，可以根据你的 Go 结构体（structs）自动创建或更新数据库表结构

var _db *gorm.DB

func InitMySQL() {
	msql := conf.Config.MySql["default"]
	pathRead := strings.Join([]string{msql.UserName, ":", msql.Password, "@tcp(", msql.DbHost, ":", msql.DbPort, ")/", msql.DbName, "?charset=" + msql.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{msql.UserName, ":", msql.Password, "@tcp(", msql.DbHost, ":", msql.DbPort, ")/", msql.DbName, "?charset=" + msql.Charset + "&parseTime=true"}, "")
	//conn := strings.Join([]string{msql.UserName, ":", msql.Password,
	//	"@tcp(", msql.DbHost, ":", msql.DbPort, ")/", msql.DbName, "?charset=", msql.Charset, "&parseTime=True"}, "")
	var ormLogger = logger.Default
	if gin.Mode() == "debug" {
		ormLogger = ormLogger.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 不根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //设置表名都为单数形式，因此有些复数表名要显式定义
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(pathWrite)},                      // 写操作
			Replicas: []gorm.Dialector{mysql.Open(pathRead), mysql.Open(pathRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                    // sources/replicas 负载均衡策略
		}))

	_db = _db.Set("gorm:table_options", "charset=utf8mb4")
}

//创建新的数据库连接，同时绑定上下文来控制数据库操作

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
