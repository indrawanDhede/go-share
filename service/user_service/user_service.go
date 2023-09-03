package user_service

import (
	"context"
	"go_share/model/api/api_response"
)

type UserService interface {
	FindAll(ctx context.Context) []api_response.UserResponse
}
