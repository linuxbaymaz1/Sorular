package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const recordCount = 100000
const chunkSize = 2000

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("DB açılamadı: %v", err)
	}
	defer db.Close()

	pragmas := []string{
		"PRAGMA journal_mode = OFF;",
		"PRAGMA synchronous = 0;",
		"PRAGMA locking_mode = EXCLUSIVE;",
		"PRAGMA temp_store = MEMORY;",
	}
	for _, p := range pragmas {
		db.Exec(p)
	}

	createTables(db)

	fmt.Printf("%d Kayıt İçin SQL Testi Başlıyor...\n\n", recordCount)

	startInsert := time.Now()
	bulkInsert(db)
	fmt.Printf("[INSERT] %d kayıt RAM'e işlendi. Süre: %v\n", recordCount*2, time.Since(startInsert))

	startUpdate := time.Now()
	bulkUpdate(db)
	fmt.Printf("[UPDATE] Kayıtlar güncellendi. Süre: %v\n", time.Since(startUpdate))

	startDelete := time.Now()
	bulkDelete(db)
	fmt.Printf("[DELETE] Tüm kayıtlar silindi. Süre: %v\n", time.Since(startDelete))
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, email TEXT, status TEXT, score INTEGER);
	CREATE TABLE orders (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER, product TEXT, amount REAL, is_shipped BOOLEAN);`
	db.Exec(query)
}

func bulkInsert(db *sql.DB) {
	tx, _ := db.Begin()

	for i := 0; i < recordCount; i += chunkSize {
		var userBuilder strings.Builder
		var orderBuilder strings.Builder

		userBuilder.WriteString("INSERT INTO users (username, email, status, score) VALUES ")
		orderBuilder.WriteString("INSERT INTO orders (user_id, product, amount, is_shipped) VALUES ")

		for j := 0; j < chunkSize; j++ {
			userId := i + j + 1

			userBuilder.WriteString(fmt.Sprintf("('user_%d', 'user%d@lab.local', 'active', %d)", userId, userId, userId%100))
			orderBuilder.WriteString(fmt.Sprintf("(%d, 'Product', 99.90, false)", userId))

			if j < chunkSize-1 {
				userBuilder.WriteString(",")
				orderBuilder.WriteString(",")
			}
		}

		tx.Exec(userBuilder.String())
		tx.Exec(orderBuilder.String())
	}

	tx.Commit()
}

func bulkUpdate(db *sql.DB) {
	tx, _ := db.Begin()
	tx.Exec("UPDATE users SET status = 'inactive' WHERE score < 50")
	tx.Exec("UPDATE orders SET is_shipped = true, amount = 89.91 WHERE is_shipped = false")
	tx.Commit()
}

func bulkDelete(db *sql.DB) {
	tx, _ := db.Begin()
	tx.Exec("DELETE FROM orders")
	tx.Exec("DELETE FROM users")
	tx.Commit()
}
