package service

import (
	"context"
	"github.com/UberionAI/Learning_EM/service/mocks"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetUserBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockExternalClient(ctrl)
	mockClient.EXPECT().
		Do(gomock.Any()).
		Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"balance": 1000}`)),
		}, nil)

	svc := NewUserService(mockClient)
	balance, err := svc.GetUserBalance(context.Background(), 123)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if balance != 1000 {
		t.Fatalf("got %d, want 1000", balance)
	}
}
