package migrate

import (
	"os"

	"github.com/Alfeenn/article/model"
)

type UserTable struct {
	model.User `gorm:"embedded"`
}

func (UserTable) TableName() string {
	return os.Getenv("DBNAME") + ".user"
}
