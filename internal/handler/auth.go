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

func (sh *authHandler) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &auth.LoginResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	// register

	req, err := sh.authService.Login(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (sh *authHandler) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &auth.LogoutResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	// register

	req, err := sh.authService.Logout(ctx, request)
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
