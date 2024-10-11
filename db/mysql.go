package db

import (
	"cloudplatform/config"
	"cloudplatform/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlogger "log"
	"os"
	"time"
)

var MysqlClient *gorm.DB

func GetMysql(conf *config.Mysql) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		conf.User,
		conf.Password,
		conf.Address, conf.Name,
		conf.Scheme)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(
			stdlogger.New(os.Stdout, "\r\n", stdlogger.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		),
	})
	if err != nil {
		log.Errorf("connect mysql  error:%s\r\n", err.Error())
		//todo 是否返回错误
		return err
	}
	MysqlClient = cli
	return nil
}
