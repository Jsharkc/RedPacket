package orm

import (
	"github.com/Jsharkc/RedPacket/general"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBConn *gorm.DB
	err    error
)

// InitOrm - initial orm
func InitOrm() {
	DBConn, err = gorm.Open(general.Dialect, general.MysqlArg)
	if err != nil {
		panic(err)
	}
}

// InitTestOrm - initial orm for test
func InitTestOrm() {
	DBConn, err = gorm.Open(general.Dialect, general.MysqlTestArg)
	if err != nil {
		panic(err)
	}
}
