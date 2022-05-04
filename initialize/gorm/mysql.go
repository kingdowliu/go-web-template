package gorm

import (
	"github.com/kingdowliu/go-web-template/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		return db
	}
}

func InitMySQLGorm() {
	global.DB = NewMySQLGorm()
}
