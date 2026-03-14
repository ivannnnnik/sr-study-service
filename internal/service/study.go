package service

import (
	"context"

	"github.com/ivannnnnik/sr-study-service/internal/model"
)


type studyRepo interface {
	Create(ctx context.Context, question *model.StudyProgress) error
	GetByID(ctx context.Context, id string) (*model.StudyProgress, error)
	List(ctx context.Context) ([]model.StudyProgress, error)
}

type StudyService struct{
	repo studyRepo
}

func NewStudyService(repo studyRepo) *StudyService{
	return &StudyService{
		repo: repo,
	}
}


func (svc *StudyService) Create(ctx context.Context, title, category, difficulty string) (*model.Question, error){
	
	questionModel := model.StudyProgress{
		Title: title,
		Category: category,
		Difficulty: difficulty,
	}
	
	err := svc.repo.Create(ctx, &questionModel)

	if err != nil{
		return nil, err
	}

	return &questionModel, nil
}

func (svc *QuestionService) GetQuestion(ctx context.Context, id string) (*model.Question, error){
	user, err := svc.repo.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}

	return user, nil

}

func (svc *QuestionService) List(ctx context.Context) ([]model.Question, error){
	questions, err := svc.repo.List(ctx)
	if err != nil{
		return nil, err
	}

	return questions, nil

}