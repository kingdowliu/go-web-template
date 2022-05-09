package gorm

import (
	"github.com/kingdowliu/go-web-template/global"
	demoModel "github.com/kingdowliu/go-web-template/models/demo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewMySQLGorm() *gorm.DB {
	m := global.DB_CONFIG.MySQL
	if m.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dst(),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil
	} else {
		db.AutoMigrate(&demoModel.User{})
		return db
	}
}

func InitMySQLGorm() {
	global.DB = NewMySQLGorm()
}
