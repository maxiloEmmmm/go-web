package contact

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maxiloEmmmm/go-web/lib"
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
	// todo: test db field type for tinyint mediumInt
	if value == nil || value.(int64) == 1 {
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

func (b BoolField) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Bool)
}

func (b BoolField) UnmarshalJSON(data []byte) (err error) {
	tmp := false
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	b.Bool = tmp
	return
}

type JsonField struct {
	Interface map[string]interface{}
}

func (j *JsonField) Scan(value interface{}) (err error) {
	tmp := make(map[string]interface{})
	if value != nil {
		if err = json.Unmarshal(lib.Uint8sToBytes(value.([]uint8)), &tmp); err != nil {
			return
		}
		j.Interface = tmp
	}

	return
}

func (j JsonField) Value() (driver.Value, error) {
	return json.Marshal(j.Interface)
}

func (j JsonField) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Interface)
}

func (j JsonField) UnmarshalJSON(data []byte) (err error) {
	tmp := make(map[string]interface{}, 0)
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}

	j.Interface = tmp
	return
}

type JsonArrayField struct {
	Interface []interface{}
}

func (j *JsonArrayField) Scan(value interface{}) (err error) {
	tmp := make([]interface{}, 0)
	if value != nil {
		if err = json.Unmarshal(lib.Uint8sToBytes(value.([]uint8)), &tmp); err != nil {
			return
		}
		j.Interface = tmp
	}

	return
}

func (j JsonArrayField) Value() (driver.Value, error) {
	return json.Marshal(j.Interface)
}

func (j JsonArrayField) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Interface)
}

func (j JsonArrayField) UnmarshalJSON(data []byte) (err error) {
	tmp := make([]interface{}, 0)
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}

	j.Interface = tmp
	return
}
