package service

import (
	"context"
	"interview/db"
	"log"
)

type InterviewService struct {
	db *db.PrismaClient
}

func NewInterviewService(db *db.PrismaClient) *InterviewService {
	return &InterviewService{
		db: db,
	}
}

func (s *InterviewService) GetInterview(id string) (*db.InterviewModel, error) {
	ctx := context.Background()
	interview, err := s.db.Interview.FindUnique(db.Interview.ID.Equals(id)).Exec(ctx)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return interview, nil
}

func (s *InterviewService) CreateInterview(interview *db.InterviewModel) (*db.InterviewModel, error) {
	ctx := context.Background()
	createdInterview, err := s.db.Interview.CreateOne(
		db.Interview.CompanyName.Set(interview.CompanyName),
		db.Interview.UserID.Set(interview.UserID),
		db.Interview.IsCompleted.Set(interview.IsCompleted),
		db.Interview.Time.Set(interview.Time),
		db.Interview.Notes.Set(interview.Notes),
		db.Interview.VacancyURL.Set(interview.VacancyURL),
		db.Interview.MeetingURL.Set(interview.MeetingURL),
	).Exec(ctx)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return createdInterview, nil
}
