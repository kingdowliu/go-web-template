package initialize

import (
	"github.com/kingdowliu/go-web-template/initialize/gorm"
	"github.com/kingdowliu/go-web-template/initialize/viper"
	"github.com/kingdowliu/go-web-template/initialize/zap"
)

func Initialize() {
	// viper初始化
	viper.InitViper()
	// gorm初始化
	gorm.InitMySQLGorm()
	// zap初始化
	zap.InitZap()
}
