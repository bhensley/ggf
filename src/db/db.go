package db

import (
	"fmt"
	"ggf/src/utils"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type databasesByEnv struct {
	dev  string
	prod string
}

func New() *gorm.DB {
	// TODO: Consider reformatting port to coerce to int
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEnv("MYSQL_USERNAME", "root"),
		utils.GetEnv("MYSQL_PASSWORD", ""),
		utils.GetEnv("MYSQL_HOST", "localhost"),
		utils.GetEnv("MYSQL_PORT", "3306"),
		utils.GetEnv("MYSQL_DATABASE", getDbNameByEnv(utils.GetEnv("ENV", "dev"))))

	fmt.Println("DSN: ", dsn)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 10,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		fmt.Println("Storage Error: ", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {

}

func getDbNameByEnv(env string) string {
	dbs := databasesByEnv{
		dev:  "ggf_dev",
		prod: "ggf_prod",
	}

	return dbs.dev
}
