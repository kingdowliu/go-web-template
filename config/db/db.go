package db

type DBConfig struct {
	MySQL MySQLConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
