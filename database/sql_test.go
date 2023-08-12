package database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// NOTE: menggunakan context agar bisa melakukan cancel saat ExecContext sedang berjalan.
	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('BUDI', 'BUDI')"

	// NOTE: ExecContext digunakan untuk eksekusi query yang tidak mengembalikan data, seperti INSERT, UPDATE, DELETE, dan lain-lain.
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}
