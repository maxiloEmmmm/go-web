package contact

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	Db *gorm.DB
)

type ScopeFunction func(db *gorm.DB) *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		Config.Database.Username,
		Config.Database.Password,
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.Database))

	if err != nil {
		log.Fatalln(err)
	}

	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetMaxOpenConns(100)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return fmt.Sprintf("%s%s", Config.Database.Prefix, defaultTableName)
	}

	Db.SingularTable(true)

	Db.LogMode(true)

	if err := Db.DB().Ping(); err != nil {
		log.Fatalln(err)
	}
}

type BoolField struct {
	Bool bool
}

func (b *BoolField) Scan(value interface{}) error {
	if value == nil || value.(int) == 1 {
		b.Bool = false
	} else {
		b.Bool = true
	}

	return nil
}

func (b BoolField) Value() (driver.Value, error) {
	if b.Bool {
		return int64(0), nil
	} else {
		return int64(1), nil
	}
}

type JsonField struct {
	Interface interface{}
}

func (j *JsonField) Scan(value interface{}) error {
	if value == nil {
		j.Interface = struct{}{}
		return nil
	} else {
		return json.Unmarshal(value.([]byte), j.Interface)
	}
}

func (j JsonField) Value() (driver.Value, error) {
	return json.Marshal(j.Interface)
}
