package mysql

import (
	"go_web/conf"
	"go_web/model"
	"strings"

	logging "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 声明数据库
var MysqlDB *gorm.DB

func Init() {
	// var builder strings.Builder
	s := []string{conf.MysqlUser, ":", conf.MysqlPassword, "@tcp(", conf.MysqlHost, ":", conf.MysqlPort, ")/", conf.MysqlDataBase, "?charset=", conf.MysqlCharset, "&parseTime=True&loc=Local"}
	dsn := strings.Join(s, "")

	mysqlLogger := logger.Default.LogMode(logger.LogLevel(conf.MysqlLogMode))

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         conf.MysqlDefaultStringSize,
		DisableDatetimePrecision:  conf.MysqlDisableDatetimePrecision,
		DontSupportRenameIndex:    conf.MysqlDontSupportRenameIndex,
		DontSupportRenameColumn:   conf.MysqlDontSupportRenameColumn,
		SkipInitializeWithVersion: conf.MysqlSkipInitializeWithVersion,
	}), &gorm.Config{
		Logger: mysqlLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: conf.MysqlSingularTable, // 表明不加S
			// TablePrefix: "test_", // 指定表名的前缀 test_
		},
	})

	if err != nil {
		logging.Info(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logging.Info(err)
	}

	sqlDB.SetMaxIdleConns(conf.MysqlMaxIdleConns)       // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(conf.MysqlMaxOpenconns)       // 设置数据库连接最大打开数
	sqlDB.SetConnMaxLifetime(conf.MysqlConnMaxLifetime) // 设置可重用连接的最长时间

	err = db.AutoMigrate(&model.User{}) // 自动迁移

	if err != nil {
		logging.Info(err)
	}

	MysqlDB = db

}
