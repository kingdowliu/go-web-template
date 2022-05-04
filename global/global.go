package global

import (
	"github.com/kingdowliu/go-web-template/config/db"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB_CONFIG db.DBConfig // 数据库连接配置

	DB    *gorm.DB     // 数据库连接
	VIPER *viper.Viper // viper配置
)
