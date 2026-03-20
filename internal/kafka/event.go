package kafka

import "time"

type StudyEvent struct{
	UserID int64 `json:"user_id"`
	QuestionID int64 `json:"question_id"`
	Quality int64 `json:"quality"`
	EaseFactor int64 `json:"ease_factor"`
	Interval int64 `json:"interval"`
	Timestamp time.Time `json:"timestamp"`
}