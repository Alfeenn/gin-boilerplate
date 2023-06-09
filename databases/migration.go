package databases

import (
	"fmt"

	"github.com/Alfeenn/article/app"
	"github.com/Alfeenn/article/databases/migrate"
	"github.com/Alfeenn/article/helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Tables() []interface{} {
	return []interface{}{
		&migrate.UserTable{},
	}
}

func MigrationDB() *gorm.DB {
	err := godotenv.Load(".env")
	helper.PanicIfErr(err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{ // use existing connection
		Conn: app.NewDB(),
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Migration DB Error: ", err.Error())
	}

	return gormDB
}
