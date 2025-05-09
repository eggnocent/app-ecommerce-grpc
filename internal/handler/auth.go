package handler

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/auth"
)

type authHandler struct {
	auth.UnsafeAuthServiceServer

	authService service.IAuthService
}

func (sh *authHandler) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &auth.RegisterResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	// register

	req, err := sh.authService.Register(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewAuthHandler(authService service.IAuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}
