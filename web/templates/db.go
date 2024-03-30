package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DbInstance struct {
	db *sql.DB
}

type Store interface {
	preStart()
}

func NewDbInstance() *DbInstance {
	env := GetEnv()

	var connStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", env.User, env.Password, env.Db, env.Host, env.Port)
	fmt.Print(connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	return &DbInstance{db: db}
}

func (pq *DbInstance) preStart() {
	query := `CREATE TABLE IF NOT EXISTS users (
				username VARCHAR(50) PRIMARY KEY NOT NULL,
				full_name VARCHAR(255),
				bio TEXT,
				email TEXT UNIQUE NOT NULL,
				password TEXT NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS sessions (
			    session_id VARCHAR(50) PRIMARY KEY,
			    username VARCHAR(50) REFERENCES users(username) NOT NULL,
			    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			    valid BOOLEAN DEFAULT TRUE
			);

			CREATE TABLE IF NOT EXISTS pastes (
						paste_id VARCHAR(50) PRIMARY KEY,
						username VARCHAR(50) REFERENCES users(username) NOT NULL,
						content TEXT NOT NULL,
						lang VARCHAR(50) NOT NULL,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
					);
	`

	_, err := pq.db.Exec(query)

	if err != nil {
		panic(err)
	}
}
