package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ivannnnnik/sr-study-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type StudyRepository struct {
	db *sqlx.DB
}

func NewStudyRepository(db *sqlx.DB) *StudyRepository {
	return &StudyRepository{db: db}
}

func (r *StudyRepository) GetByUserAndQuestion(ctx context.Context, userID, questionID string) (*model.StudyProgress, error) {
	query := `SELECT id, user_id, question_id, ease_factor, interval, repetitions, next_review_at, last_reviewed_at
	          FROM study_progress WHERE user_id = $1 AND question_id = $2`

	var sp model.StudyProgress
	err := r.db.GetContext(ctx, &sp, query, userID, questionID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil // нет записи — не ошибка, первый ответ
	}
	if err != nil {
		return nil, fmt.Errorf("get progress: %w", err)
	}
	return &sp, nil
}

func (r *StudyRepository) Upsert(ctx context.Context, sp *model.StudyProgress) error {
	query := `
	INSERT INTO study_progress (user_id, question_id, ease_factor, interval, repetitions, next_review_at, last_reviewed_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (user_id, question_id) DO UPDATE SET
		ease_factor      = EXCLUDED.ease_factor,
		interval         = EXCLUDED.interval,
		repetitions      = EXCLUDED.repetitions,
		next_review_at   = EXCLUDED.next_review_at,
		last_reviewed_at = EXCLUDED.last_reviewed_at
	RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		sp.UserID, sp.QuestionID, sp.EaseFactor, sp.Interval, sp.Repetitions, sp.NextReviewAt, sp.LastReviewAt,
	).Scan(&sp.ID)
}

func (r *StudyRepository) ListByUser(ctx context.Context, userID string) ([]model.StudyProgress, error) {
	query := `SELECT id, user_id, question_id, ease_factor, interval, repetitions, next_review_at, last_reviewed_at
	          FROM study_progress WHERE user_id = $1`

	var items []model.StudyProgress
	err := r.db.SelectContext(ctx, &items, query, userID)
	if err != nil {
		return nil, fmt.Errorf("list by user: %w", err)
	}
	return items, nil
}
