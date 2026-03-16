package model

import "time"

type StudyProgress struct{
	ID           string    `db:"id"`
    UserID       string    `db:"user_id"`
    QuestionID       string    `db:"question_id"`

    EaseFactor     float64    `db:"ease_factor"`
    Interval     int32    `db:"interval"`
    Repetitions     int32    `db:"repetitions"`

    NextReviewAt    time.Time `db:"next_review_at"`
    LastReviewAt    time.Time `db:"last_reviewed_at"`
}