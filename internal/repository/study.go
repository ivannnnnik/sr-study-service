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

func NewQuestionRepository(db *sqlx.DB) *StudyRepository{
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

	err := r.db.QueryRowContext(ctx, query, studyProgress.Title, studyProgress.Category, studyProgress.Difficulty).
	Scan(&studyProgress.ID, &studyProgress.Title, &studyProgress.Category, &studyProgress.Difficulty, &studyProgress.CreatedAt)
	
	return err
}

func (r *StudyRepository) GetByID(ctx context.Context, id string) (*model.Question, error){
	query := `SELECT id, title, category, difficulty, created_at FROM question WHERE id = $1`

	var question model.Question
	err := r.db.GetContext(ctx, &question, query, id)
	if err != nil{
		return nil, err
	}

	return &question, nil
}

func (r *StudyRepository) List(ctx context.Context) ([]model.Question, error){
	query := `SELECT id, title, category, difficulty, created_at FROM question`

 	args := map[string]interface{}{}

	var questions []model.Question
	rows, err := r.db.NamedQueryContext(ctx, query, args)
	if err != nil{
        return nil, fmt.Errorf("listing questions: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
		var q model.Question

		if err := rows.StructScan(&q);err != nil{
			return nil, fmt.Errorf("scanning question: %v", err)
		}
		questions = append(questions, q)
	}

	return questions, rows.Err()
}