package main

import (
	"fmt"
	"os"

	"gormlogrus/gorm_logger"
	"gormlogrus/model/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	/*
		//初始化日志库
		file, err := os.OpenFile("gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if nil != err {
			return
		}
	*/
	/*
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             2 * time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			})
	*/

	dsn := "hoicee:hoicee_cc@tcp(127.0.0.1:43306)/cqi?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_logger.New(os.Stdout),
	})
	if nil != err {
		fmt.Println(err)
		return
	}

	//初始化日志库
	file, err := os.OpenFile("./gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if nil != err {
		return
	}

	sdb := gormDB.Session(&gorm.Session{
		//DryRun: true,
		Logger: gorm_logger.New(file),
	})

	appKey := db.AppKey{}
	result := sdb.First(&appKey)
	if nil != result.Error {
		fmt.Printf("gormDB error\n")
	} else {
		//stmt := result.Statement

		//fmt.Printf("%s\n", stmt.SQL.String())
		//fmt.Printf("%v\n", stmt.Vars)
		//fmt.Printf("%v\n", appKey)

		//sql := gormDB.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
		//fmt.Printf("sql %s\n", sql)

		fmt.Println("Done")
	}
}
