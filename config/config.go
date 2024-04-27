package config

type Server struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	MysqlUser MysqlUser `mapstructure:"mysql_user" json:"mysql_user" yaml:"mysql_user"`
	Mongo     Mongo     `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
}
