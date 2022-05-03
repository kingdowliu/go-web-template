package db_config

import "fmt"

//
// MySQLConfig
// @Description:
//
type MySQLConfig struct {
	Host     string // 地址
	Port     int    // 端口
	DBName   string // 数据库名
	User     string // 用户名
	PassWord string // 密码
	Other    string // 其他项
}

//
// Dst
// @Description: 生成dsn数据库连接字符串
// @receiver mysqlConfig
// @return string
//
func (mysqlConfig *MySQLConfig) Dst() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", mysqlConfig.User, mysqlConfig.PassWord, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName, mysqlConfig.Other)
	return dsn
}
