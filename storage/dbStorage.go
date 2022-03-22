package storage

import (
	"database/sql"
	"fmt"

	"github.com/NuVeS/PasswordMangerBot/config"
	_ "github.com/lib/pq"
)

type dbStorage struct {
	db *sql.DB
}

var instance *dbStorage

func NewInstance() *dbStorage {
	if instance == nil {
		configuration := config.NewInstance()
		sqlcode := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configuration.Database.Host,
			configuration.Database.Port, configuration.Database.User, configuration.Database.Password, configuration.Database.Dbname)
		db, err := sql.Open("postgres", sqlcode)
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		instance = &dbStorage{db: db}
	}
	return instance
}

func (storage dbStorage) Get(id string, title string) string {
	get := `SELECT password FROM accounts WHERE id = $1 and title = $2`
	rows, err := storage.db.Query(get, id, title)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var password string
	for rows.Next() {
		err = rows.Scan(&password)
		if err != nil {
			panic(err)
		}
	}

	if len(password) > 0 {
		return fmt.Sprintf("%s %s", title, password)
	} else {
		return ""
	}
}

func (storage dbStorage) Set(id string, title string, password string) bool {
	insert := `INSERT INTO "accounts"("id", "username", "password") VALUES ($1, $2, $3)`
	_, err := storage.db.Exec(insert, id, title, password)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (storage dbStorage) GetAll(id string) map[string]string {
	get := `SELECT title, password FROM accounts WHERE id = $1`
	rows, err := storage.db.Query(get, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var passwords = make(map[string]string)
	for rows.Next() {
		var title string
		var password string
		err = rows.Scan(&title, &password)
		if err != nil {
			panic(err)
		}
		passwords[title] = password
	}
	return passwords
}

func (storage dbStorage) Update(id string, title string, password string) bool {
	update := `UPDATE accounts SET title=$1, password=$2 WHERE id=$3 AND title = $1`
	_, e := storage.db.Exec(update, title, password, id)
	if e != nil {
		return false
	} else {
		return true
	}
}

func (storage dbStorage) Remove(id string, title string) bool {
	deleteStmt := `DELETE FROM accounts WHERE id = $1 AND username=$2`
	_, e := storage.db.Exec(deleteStmt, id, title)
	if e != nil {
		return false
	} else {
		return true
	}
}
