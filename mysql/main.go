package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:qwer1234@/omerdb1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")

	err = insertData(db)
	if err != nil {
		log.Fatal(err)
	}

	user, err := getUserByID(db, "1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	userLim, err := getUserByLimit(db, 3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userLim)

	user1 := User{
		Username: "Alice",
	}
	err = addData(db, user1)
	if err != nil {
		log.Fatal(err)
	}

	user2 := User{
		Username: "Bob",
	}
	err = addData(db, user2)
	if err != nil {
		log.Fatal(err)
	}

	users, err := fetchAllData(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}

func createTable(db *sql.DB) error {
	sql := `
	CREATE TABLE IF NOT EXISTS users (
            id INT NOT NULL AUTO_INCREMENT,
            username VARCHAR(50) NOT NULL,
            PRIMARY KEY (id)
        )
    `
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func insertData(db *sql.DB) error {
	sql := `
        INSERT INTO users (username)
        VALUES (?)
    `
	result, err := db.Exec(sql, "deneme")
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	log.Println("Data inserted successfully, ID:", id)
	return nil
}

func fetchAllData(db *sql.DB) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func getUserByID(db *sql.DB, id string) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User not found")
		}
		return user, err
	}
	return user, nil
}

func getUserByLimit(db *sql.DB, limit int) (User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM users limit ?", limit)
	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("User not found")
		}
		return user, err
	}
	return user, nil
}

func addData(db *sql.DB, user User) error {
	stmt, err := db.Prepare("INSERT INTO users(username) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username)
	if err != nil {
		return err
	}
	return nil
}
