package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:1722@tcp(localhost:3306)/go_test")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

// NOTE: golang secara default sudah memiliki kemampuan database pooling. database pooling adalah manajemen koneksi database sehingga kita tidak perlu melakukannya secara manual.
