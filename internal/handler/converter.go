package handler

import (
	"time"

	studyv1 "github.com/ivannnnnik/sr-proto/gen/go/study/v1"
	"github.com/ivannnnnik/sr-study-service/internal/model"
)

func ProgressToProto(sp *model.StudyProgress) *studyv1.StudyProgress {
	return &studyv1.StudyProgress{
		Id:             sp.ID,
		UserId:         sp.UserID,
		QuestionId:     sp.QuestionID,
		EaseFactor:     sp.EaseFactor,
		Interval:       sp.Interval,
		Repetitions:    sp.Repetitions,
		NextReviewAt:   sp.NextReviewAt.Format(time.RFC3339),
		LastReviewedAt: sp.LastReviewAt.Format(time.RFC3339),
	}
}

func ProgressListToProto(items []model.StudyProgress) []*studyv1.StudyProgress {
	result := make([]*studyv1.StudyProgress, 0, len(items))
	for i := range items {
		result = append(result, ProgressToProto(&items[i]))
	}
	return result
}
