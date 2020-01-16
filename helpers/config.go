package helpers

import (
	"errors"
	"fmt"
	configer "github.com/Smilefish0/gener/config"
	"os"
)

// 从配置文件中获取 数据源连接字符串
func GetDatabaseDSN() string {
	dbConfig := configer.DatabaseConfig()
	if dbConfig.GetConnection() == "mysql" {
		return fmt.Sprintf("mysql://%v:%v@%v:%v/%v",
			dbConfig.GetUsername(),
			dbConfig.GetPassword(),
			dbConfig.GetHost(),
			dbConfig.GetPort(),
			dbConfig.GetDatabase(),
		)
		// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if dbConfig.GetConnection() == "postgres" {
		return fmt.Sprintf("pgsql://%v:%v@%v/%v?sslmode=disable",
			dbConfig.GetUsername(),
			dbConfig.GetPassword(),
			dbConfig.GetHost(),
			dbConfig.GetDatabase(),
		)
	} else if dbConfig.GetConnection() == "sqlite" {
		return fmt.Sprintf("sqlite://%v/%v", os.TempDir(),
			dbConfig.GetDatabase(),
		)
	} else {
		panic(errors.New("not supported database adapter"))
	}
}
