package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}

// NOTE: golang secara default sudah memiliki kemampuan database pooling. database pooling adalah manajemen koneksi database sehingga kita tidak perlu melakukannya secara manual.
