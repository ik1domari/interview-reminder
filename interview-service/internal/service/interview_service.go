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
