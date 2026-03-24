package sql

import (
	"gorm.io/gorm"
)

type sql struct {
	DB *gorm.DB
}

var Sql = &sql{}

func (d *sql) Startup() {
	// db, err := gorm.Open(sqlite.Open("seo-tools.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// db.AutoMigrate(&m.Site{}, &m.Ahrefs{}, &m.Majestic{}, &m.Semrush{}, &m.Pbn{}, &m.Link{}, &m.Account{}, &m.Required{}, &m.Contacts{}, &m.Group{})
	// d.DB = db
}
