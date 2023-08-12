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

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"

	// NOTE: QueryContext digunakan untuk eksekusi query yang mengembalikan data, seperti SELECT.
	// NOTE: rows disini bentuknya adalah pointer jadi hanya bisa Next() terus untuk melakukan iterasi data hasil query.
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	// NOTE: defer rows.Close() digunakan untuk menutup koneksi database setelah selesai melakukan query.
	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
	}

	fmt.Println("Success get data")
}
