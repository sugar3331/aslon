package config

type MysqlUser struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Dbname   string `mapstructure:"db_name" json:"db_name" yaml:"db_name"` // 数据库名
}
