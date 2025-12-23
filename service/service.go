package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go -package=mocks

type ExternalClient interface {
	Do(*http.Request) (*http.Response, error)
}

type UserService struct {
	client ExternalClient
}

func NewUserService(client ExternalClient) *UserService {
	return &UserService{client: client}
}

func (s *UserService) GetUserBalance(ctx context.Context, userID int) (int, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET",
		"https://api.example.com/balance"+strconv.Itoa(userID), nil)
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct{ Balance int }
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Balance, nil
}

//PS C:\Users\tla\GolandProjects\Learning_EM\service> go test -v
//=== RUN   TestGetUserBalance
//--- PASS: TestGetUserBalance (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/service        0.163s
//PS C:\Users\tla\GolandProjects\Learning_EM\service>
