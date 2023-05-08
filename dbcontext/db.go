package dbcontext

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() {
	if Config.Database.File == "" {
		var url string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			Config.Database.Username, Config.Database.Password, Config.Database.Server, Config.Database.Port, Config.Database.Name)
		DB, _ = gorm.Open(mysql.Open(url), &gorm.Config{})

	} else {
		DB, _ = gorm.Open(sqlite.Open(Config.Database.File), &gorm.Config{})
	}
}
