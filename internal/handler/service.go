package handler

import (
	"context"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/service"
)

type IServiceHandler interface {
	HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error)
}

type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &service.HelloWorldResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("hello %s", request.Name),
		Base:    utils.SuccessResponse("success"),
	}, nil
}

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
