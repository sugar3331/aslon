package initialize

import (
	"aslon/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormTalMysql 初始化Mysql数据库
func GormTalMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GlobalConfig.MysqlUser.Username, global.GlobalConfig.MysqlUser.Password, global.GlobalConfig.MysqlUser.Path, global.GlobalConfig.MysqlUser.Port,
		global.GlobalConfig.MysqlUser.Dbname)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("mysql_user db init err :", err)
		return nil
	} else {
		fmt.Println("mysql_user db init success")
		return db
	}
}
