package db

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// go1.14/src/database/sql/sql.go
var errDBClosed = errors.New("sql: database is closed")

// re-use a db connection.
func Test_DB(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	forTypeAssert := db
	var arg interface{} = forTypeAssert
	_, ok := arg.(gorm.SQLCommon)
	if !ok {
		t.Fatal("db type is invalid.")
	}

	gDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatal(err)
	}
	var sqlDB *sql.DB = gDB.DB()

	if !reflect.DeepEqual(db, sqlDB) {
		t.Errorf("parse() got = %+v, but want = %+v", db, sqlDB)
	}

	db.Close()

	if err := db.Ping(); err.Error() != errDBClosed.Error() {
		t.Errorf("Ping() error: %+v", err)
	}

	if err := gDB.DB().Ping(); err.Error() != errDBClosed.Error() {
		t.Errorf("Ping() error: %+v", err)
	}

	if err := sqlDB.Ping(); err.Error() != errDBClosed.Error() {
		t.Errorf("Ping() error: %+v", err)
	}

	if !reflect.DeepEqual(db, sqlDB) {
		t.Errorf("parse() got = %+v, but want = %+v", db, sqlDB)
	}
}
