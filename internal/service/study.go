package service

import (
	"context"
	"fmt"
	"math"
	"time"

	questionv1 "github.com/ivannnnnik/sr-proto/gen/go/question/v1"
	"github.com/ivannnnnik/sr-study-service/internal/model"
)

type studyRepo interface {
	GetByUserAndQuestion(ctx context.Context, userID, questionID string) (*model.StudyProgress, error)
	Upsert(ctx context.Context, sp *model.StudyProgress) error
	ListByUser(ctx context.Context, userID string) ([]model.StudyProgress, error)
}

type questionProvider interface{
	GetQuestion(ctx context.Context, id string) (*questionv1.Question, error)
}

type StudyService struct {
	repo studyRepo
	question questionProvider
}

func NewStudyService(repo studyRepo, question questionProvider) *StudyService {
	return &StudyService{repo: repo, question: question}
}

func (s *StudyService) SubmitAnswer(ctx context.Context, userID, questionID string, quality int32) (*model.StudyProgress, error) {
	sp, err := s.repo.GetByUserAndQuestion(ctx, userID, questionID)
	if err != nil {
		return nil, err
	}

	_, err = s.question.GetQuestion(ctx, questionID)
	if err != nil{
		return nil, fmt.Errorf("question not found: %w", err)
	}

	now := time.Now().UTC()

	if sp == nil {
		sp = &model.StudyProgress{
			UserID:     userID,
			QuestionID: questionID,
			EaseFactor: 2.5,
		}
	}

	sp.EaseFactor, sp.Interval, sp.Repetitions = calcSM2(sp.EaseFactor, sp.Interval, sp.Repetitions, quality)
	sp.LastReviewAt = now
	sp.NextReviewAt = now.AddDate(0, 0, int(sp.Interval))

	if err := s.repo.Upsert(ctx, sp); err != nil {
		return nil, err
	}

	return sp, nil
}

func (s *StudyService) GetProgress(ctx context.Context, userID string) ([]model.StudyProgress, error) {
	return s.repo.ListByUser(ctx, userID)
}

func calcSM2(ef float64, interval, reps, quality int32) (newEF float64, newInterval, newReps int32) {
	if quality >= 3 {
		switch reps {
		case 0:
			newInterval = 1
		case 1:
			newInterval = 6
		default:
			newInterval = int32(math.Round(float64(interval) * ef))
		}
		newReps = reps + 1
	} else {
		newInterval = 1
		newReps = 0
	}

	newEF = ef + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))
	if newEF < 1.3 {
		newEF = 1.3
	}
	return
}
