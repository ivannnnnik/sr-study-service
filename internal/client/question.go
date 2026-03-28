package client

import (
	"context"
	"fmt"

	questionv1 "github.com/ivannnnnik/sr-proto/gen/go/question/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type QuestionClient struct{
	client questionv1.QuestionServiceClient
	conn *grpc.ClientConn
}

func NewQuestionClient(addr string)(*QuestionClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		return nil, fmt.Errorf("Fail connect to question-service: %w", err)
	}

	return &QuestionClient{
		client: questionv1.NewQuestionServiceClient(conn),
		conn: conn,
	}, nil
}

func (c *QuestionClient) GetQuestion(ctx context.Context, id string) (*questionv1.Question, error){
	response, err := c.client.GetQuestion(ctx, &questionv1.GetQuestionRequest{Id: id})
	if err != nil{
		return nil, fmt.Errorf("Get question: %w", err)
	}
	return response.Question, nil
}

func (c *QuestionClient) Close() error{
	return c.conn.Close()
}
