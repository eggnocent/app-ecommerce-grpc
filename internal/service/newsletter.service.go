package service

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/newsletter"
	"time"

	"github.com/google/uuid"
)

type InewsletterService interface {
	SubscribeNewsLetter(ctx context.Context, request *newsletter.SubscribeNewsLetterRequest) (*newsletter.SubscribeNewsLetterResponse, error)
}

type newsletterService struct {
	newsletterRepository repository.INewsletterRepository
}

func (ns *newsletterService) SubscribeNewsLetter(ctx context.Context, request *newsletter.SubscribeNewsLetterRequest) (*newsletter.SubscribeNewsLetterResponse, error) {
	// cek ke database emailnya, udah terdaftar atau belum
	nenwsletterEntity, err := ns.newsletterRepository.GetNewsletterByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if nenwsletterEntity != nil {
		return &newsletter.SubscribeNewsLetterResponse{
			Base: utils.SuccessResponse("subscribe newsletter success"),
		}, nil
	}
	// kalo belum insert ke db
	newNewsletterEntity := entity.Newsletter{
		ID:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		CreatedAt: time.Now(),
		CreatedBy: "Public",
	}

	err = ns.newsletterRepository.CreateNewNewsletter(ctx, &newNewsletterEntity)
	if err != nil {
		return nil, err
	}
	return &newsletter.SubscribeNewsLetterResponse{
		Base: utils.SuccessResponse("subscribe newsletter success"),
	}, nil
}

func NewNewsletterService(newsletterRepository repository.INewsletterRepository) InewsletterService {
	return &newsletterService{
		newsletterRepository: newsletterRepository,
	}
}
