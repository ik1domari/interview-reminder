package models

import (
	"time"
)

type Interview struct {
	ID          string    `json:"id"`
	Company     string    `json:"company"`
	UserID      string    `json:"user_id"`
	IsCompleted bool      `json:"is_completed"`
	Time        time.Time `json:"time"`
	Notes       string    `json:"notes"`
	VacancyURL  string    `json:"vacancy_url"`
	MeetingURL  string    `json:"meeting_url"`
}
