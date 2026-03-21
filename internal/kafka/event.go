package kafka

import "time"

type StudyEvent struct{
	UserID string `json:"user_id"`
	QuestionID string `json:"question_id"`
	Quality int32 `json:"quality"`
	EaseFactor float64 `json:"ease_factor"`
	Interval int64 `json:"interval"`
	Timestamp time.Time `json:"timestamp"`
}