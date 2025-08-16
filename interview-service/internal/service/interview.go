package service

import (
	"database/sql"
	"interview/internal/database"
	"log"
)

type GetAllParams struct {
	db     *sql.DB
	search string
}

func GetAll(params GetAllParams) []database.Interview {
	rows, err := params.db.Query("SELECT * FROM interviews")

	if params.search != "" {
		rows, err = params.db.Query("SELECT * FROM interviews WHERE company_name ILIKE $1", "%"+params.search+"%")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var interviews []database.Interview

	for rows.Next() {
		interview := database.Interview{}
		err = rows.Scan(
			&interview.ID,
			&interview.CompanyName,
			&interview.Time,
			&interview.Notes,
			&interview.VacancyURL,
			&interview.MeetingURL,
			&interview.UserID,
			&interview.IsCompleted,
			&interview.CreatedAt,
			&interview.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		interviews = append(interviews, interview)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return interviews
}
func GetByID(db *sql.DB, id int) database.Interview {
	var interview database.Interview
	err := db.QueryRow("SELECT * FROM interviews WHERE id = $1", id).Scan(
		&interview.ID,
		&interview.CompanyName,
		&interview.Time,
		&interview.Notes,
		&interview.VacancyURL,
		&interview.MeetingURL,
		&interview.UserID,
		&interview.IsCompleted,
		&interview.CreatedAt,
		&interview.UpdatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	return interview
}

func Create(db *sql.DB, interview database.Interview) {
	_, err := db.Exec(
		"INSERT INTO interviews (company_name, time, notes, vacancy_url, meeting_url, user_id, is_completed) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		interview.CompanyName,
		interview.Time,
		interview.Notes,
		interview.VacancyURL,
		interview.MeetingURL,
		interview.UserID,
		interview.IsCompleted,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(db *sql.DB, interview database.Interview) {
	_, err := db.Exec(
		"UPDATE interviews SET company_name = $1, time = $2, notes = $3, vacancy_url = $4, meeting_url = $5, user_id = $6, is_completed = $7 WHERE id = $8",
		interview.CompanyName,
		interview.Time,
		interview.Notes,
		interview.VacancyURL,
		interview.MeetingURL,
		interview.UserID,
		interview.IsCompleted,
		interview.ID,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM interviews WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
