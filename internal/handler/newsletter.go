package handler

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/newsletter"
)

type newsletterHandler struct {
	newsletter.UnimplementedNewsletterServiceServer

	newsletterService service.InewsletterService
}

func (nh *newsletterHandler) SubscribeNewsLetter(ctx context.Context, request *newsletter.SubscribeNewsLetterRequest) (*newsletter.SubscribeNewsLetterResponse, error) {
	validationsErrors, err := utils.CheckValidation(request)
	if err != nil {
		return nil, err
	}

	if validationsErrors != nil {
		return &newsletter.SubscribeNewsLetterResponse{
			Base: utils.ValidationErrorResponse(validationsErrors),
		}, nil
	}

	req, err := nh.newsletterService.SubscribeNewsLetter(ctx, request)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewNewsletterHandler(newsletterService service.InewsletterService) *newsletterHandler {
	return &newsletterHandler{
		newsletterService: newsletterService,
	}
}
