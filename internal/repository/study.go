package repository

import (
	"context"
	"fmt"

	"github.com/ivannnnnik/sr-study-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type StudyRepository struct{
	db *sqlx.DB
}

func NewStudyRepository(db *sqlx.DB) *StudyRepository{
	return &StudyRepository{
		db: db,
	}
}

func (r *StudyRepository) Create(ctx context.Context, studyProgress *model.StudyProgress) error{

	query := `
	INSERT INTO study_progress(user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	RETURNING id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at;
	`

	err := r.db.QueryRowContext(ctx, query, studyProgress.UserID, studyProgress.QuestionID, studyProgress.EasyFactor, studyProgress.Interval, studyProgress.Repetitions, studyProgress.NextReviewAt, studyProgress.LastReviewAt).
	Scan(&studyProgress.ID, &studyProgress.UserID, &studyProgress.QuestionID, &studyProgress.EasyFactor, &studyProgress.Interval, &studyProgress.Repetitions, &studyProgress.NextReviewAt, &studyProgress.LastReviewAt)
	
	return err
}

func (r *StudyRepository) GetByID(ctx context.Context, id string) (*model.StudyProgress, error){
	query := `SELECT id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at FROM study_progress WHERE id = $1`

	var studyProgress model.StudyProgress
	err := r.db.GetContext(ctx, &studyProgress, query, id)
	if err != nil{
		return nil, err
	}

	return &studyProgress, nil
}

func (r *StudyRepository) List(ctx context.Context) ([]model.StudyProgress, error){
	query := `SELECT id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at FROM study_progress`

 	args := map[string]interface{}{}

	var studyProgresses []model.StudyProgress
	rows, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil{
        return nil, fmt.Errorf("listing study progresses: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
		var sp model.StudyProgress

		if err := rows.StructScan(&sp);err != nil{
			return nil, fmt.Errorf("scanning study progresses: %v", err)
		}
		studyProgresses = append(studyProgresses, sp)
	}

	return studyProgresses, rows.Err()
}


func (r *StudyRepository) Upsert(ctx context.Context, studyProgress *model.StudyProgress) error{
	// TODO update sql
	query := `
	INSERT INTO study_progress(user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT
	RETURNING id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at;
	`

	err := r.db.QueryRowContext(ctx, query, studyProgress.UserID, studyProgress.QuestionID, studyProgress.EasyFactor, studyProgress.Interval, studyProgress.Repetitions, studyProgress.NextReviewAt, studyProgress.LastReviewAt).
	Scan(&studyProgress.ID, &studyProgress.UserID, &studyProgress.QuestionID, &studyProgress.EasyFactor, &studyProgress.Interval, &studyProgress.Repetitions, &studyProgress.NextReviewAt, &studyProgress.LastReviewAt)
	
	return err
}


func (r *StudyRepository) ListByUser(ctx context.Context, userID int64) ([]model.StudyProgress, error){
	query := `SELECT id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at FROM study_progress`

 	args := map[string]interface{}{}

	var studyProgresses []model.StudyProgress
	rows, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil{
        return nil, fmt.Errorf("listing study progresses: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
		var sp model.StudyProgress

		if err := rows.StructScan(&sp);err != nil{
			return nil, fmt.Errorf("scanning study progresses: %v", err)
		}
		studyProgresses = append(studyProgresses, sp)
	}

	return studyProgresses, rows.Err()
}

func (r *StudyRepository) GetByUserAndQuestion(ctx context.Context, userID string, questionID string) (*model.StudyProgress, error){
	query := `SELECT id, user_id, question_id, easy_factor, interval, repetitions, category, difficulty, next_review_at, last_reviewed_at FROM study_progress WHERE id = $1`

	var studyProgress model.StudyProgress
	err := r.db.GetContext(ctx, &studyProgress, query, id)
	if err != nil{
		return nil, err
	}

	return &studyProgress, nil
}