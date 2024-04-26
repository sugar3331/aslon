package global

import (
	"aslon/config"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	UserDB        *gorm.DB
	GlobalConfig  config.Server
	SnowflakeNode *snowflake.Node
	Viper         *viper.Viper
	Mogo          *mongo.Client
)
