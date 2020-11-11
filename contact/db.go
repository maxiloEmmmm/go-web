package contact

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	lib "github.com/maxiloEmmmm/go-tool"
	"log"
)

func InitDB(open func(string, string) (*sql.DB, error), mode string) {
	if mode == "" {
		mode = Config.App.Mode
	}

	if cfg, exist := Config.Database[mode]; exist {
		engine := cfg["engine"].(string)
		if !lib.InArray(&[]string{"mysql", "mssql", "sqlite3", "postgres"}, engine) {
			engine = "mysql"
		}

		db, err := open(engine, cfg["source"].(string))
		if err != nil {
			log.Fatalln(err)
		}

		if val, has := cfg["maxIdleConns"]; has {
			db.SetMaxIdleConns(val.(int))
		} else {
			db.SetMaxIdleConns(20)
		}

		if val, has := cfg["maxOpenConns"]; has {
			db.SetMaxOpenConns(val.(int))
		} else {
			db.SetMaxOpenConns(100)
		}

		if err := db.Ping(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(errors.New("db mode not find"))
	}
}

const (
	BoolFieldTrue  = 0
	BoolFieldFalse = 1
)

type BoolField struct {
	Bool bool
}

func (b *BoolField) Scan(value interface{}) error {
	if value == nil {
		b.Bool = false
	}

	val := sql.NullInt64{}
	err := val.Scan(value)
	if err != nil {
		return err
	}
	if val.Int64 == BoolFieldFalse {
		b.Bool = false
	} else {
		b.Bool = true
	}

	return nil
}

func (b BoolField) Value() (driver.Value, error) {
	if b.Bool {
		return int64(BoolFieldTrue), nil
	} else {
		return int64(BoolFieldFalse), nil
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
