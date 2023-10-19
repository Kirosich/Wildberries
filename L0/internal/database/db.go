package database

import (
	"Wildber/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func Init() *Database {
	Database := &Database{}
	connStr := "user=postgres dbname=wildber password=272731 sslmode=disable"
	var err error
	Database.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return Database
}

func (db Database) Insert(uid string, info string) {
	query := `INSERT INTO json_table (uid, information) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, uid, info)
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}
	fmt.Println("Data Inserted successfully")
}

func (db Database) IsExist(uid string) bool {
	var count int

	query := `SELECT COUNT(*) FROM json_table WHERE uid = $1`
	row := db.DB.QueryRow(query, uid)
	if err := row.Scan(&count); err != nil {
		log.Fatalf("Row scan error: %v", err)
	}
	if count > 0 {
		return true
	}
	return false
}

func (db Database) SelectAll() ([]string, []models.Order, error) {
	query := `SELECT "uid", "information" FROM json_table`
	var uid []string
	var information []models.Order

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var one_uid string
		var one_information models.Order
		var jsonData []byte

		if err := rows.Scan(&one_uid, &jsonData); err != nil {
			return nil, nil, err
		}
		err = json.Unmarshal(jsonData, &one_information)
		if err != nil {
			fmt.Println("Message data Unmarshal error")
			return nil, nil, err
		}
		uid = append(uid, one_uid)
		information = append(information, one_information)
	}
	return uid, information, nil
}
