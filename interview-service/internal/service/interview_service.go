package service

import (
	"interview/db"
)

type InterviewService struct {
	db *db.PrismaClient
}

func NewInterviewService(db *db.PrismaClient) *InterviewService {
	return &InterviewService{
		db: db,
	}
}

func (s *InterviewService) GetInterview(id string) {
	interview, err := s.db.Interview.FindUnique(db.Interview.ID.Equals(id)).Exec(ctx)
}
