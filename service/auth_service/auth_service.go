package auth_service

import (
	"context"
	"go_share/model/api/api_request"
	"go_share/model/api/api_response"
)

type AuthService interface {
	Register(ctx context.Context, request api_request.AuthRegisterRequest) (api_response.AuthRegisterResponse, error)
	Login(ctx context.Context, request api_request.AuthLoginRequest, channelLogin chan<- interface{})
}
