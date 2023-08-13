package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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
	// rows disini bentuknya adalah pointer jadi hanya bisa Next() terus untuk melakukan iterasi data hasil query.
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

func TestQueryContext2(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=====================================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

	fmt.Println("Success get data")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// NOTE: contoh sql injection
	// username := "admin'; #"
	// password := "admin1"
	// script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	// rows, err := db.QueryContext(ctx, script)

	// NOTE: cara menghindari sql injection, yaitu menggunakan parameter query.
	username := "admin"
	password := "admin"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login with username:", username)
	} else {
		fmt.Println("Failed login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "adoding"
	password := "alalaadoding"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}
