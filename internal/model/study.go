package model

import "time"

type StudyProgress struct{
	ID           string    `db:"id"`
    UserID       string    `db:"user_id"`
    QuestionID       string    `db:"question_id"`

    EasyFactor     float64    `db:"easy_factor"`
    Interval     int64    `db:"interval"`
    Repetitions     int64    `db:"repetitions"`

    NextReviewAt    time.Time `db:"next_review_at"`
    LastReviewAt    time.Time `db:"last_reviewed_at"`
    CreatedAt    time.Time `db:"created_at"`
}