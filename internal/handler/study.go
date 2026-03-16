package handler

import (
	"context"

	studyv1 "github.com/ivannnnnik/sr-proto/gen/go/study/v1"
	"github.com/ivannnnnik/sr-study-service/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type studyService interface {
	SubmitAnswer(ctx context.Context, userID, questionID string, quality int32) (*model.StudyProgress, error)
	GetProgress(ctx context.Context, userID string) ([]model.StudyProgress, error)
}

type StudyHandler struct {
	studyv1.UnimplementedStudyServiceServer
	service studyService
}

func NewStudyHandler(svc studyService) *StudyHandler {
	return &StudyHandler{service: svc}
}

func (h *StudyHandler) SubmitAnswer(ctx context.Context, req *studyv1.SubmitAnswerRequest) (*studyv1.SubmitAnswerResponse, error) {
	if req.UserId == "" || req.QuestionId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id and question_id required")
	}
	if req.Quality < 0 || req.Quality > 5 {
		return nil, status.Error(codes.InvalidArgument, "quality must be 0-5")
	}

	sp, err := h.service.SubmitAnswer(ctx, req.UserId, req.QuestionId, req.Quality)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "submit answer: %v", err)
	}

	return &studyv1.SubmitAnswerResponse{
		Progress: ProgressToProto(sp),
	}, nil
}

func (h *StudyHandler) GetProgress(ctx context.Context, req *studyv1.GetProgressRequest) (*studyv1.GetProgressResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id required")
	}

	items, err := h.service.GetProgress(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get progress: %v", err)
	}

	return &studyv1.GetProgressResponse{
		Items: ProgressListToProto(items),
	}, nil
}
