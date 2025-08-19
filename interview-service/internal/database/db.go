package database

import (
	"database/sql"
	"log"
	"os"
)

type Interview struct {
	ID          int
	CompanyName string
	Time        string
	Notes       string
	VacancyURL  string
	MeetingURL  string
	UserID      int
	IsCompleted bool
	CreatedAt   string
	UpdatedAt   string
}

func New() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS interviews" +
		"id SERIAL PRIMARY KEY," +
		"company_name VARCHAR(255) NOT NULL," +
		"time TIMESTAMP NOT NULL," +
		"notes TEXT" +
		"vacancy_url VARCHAR(255)" +
		"meeting_url VARCHAR(255)" +
		"user_id INT NOT NULL" +
		"is_completed BOOLEAN NOT NULL DEFAULT FALSE" +
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
