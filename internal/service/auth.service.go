package service

import (
	"context"
	"errors"
	"github/eggnocent/app-grpc-eccomerce/internal/entity"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"github/eggnocent/app-grpc-eccomerce/pb/auth"
	"os"
	"time"

	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error)
	Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error)
	ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error)
	GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
	cacheService   *gocache.Cache
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not match, try again!"),
		}, nil
	}
	// cek email apakah sudah ada?
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("user already exist"),
		}, nil
	}

	// hashing password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	// insert data ke db

	newUser := entity.User{
		Id:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		Password:  string(hashedPass),
		RoleCode:  entity.UserRoleCustomer,
		CreatedAt: time.Now(),
		CreatedBy: &request.FullName,
	}

	err = as.authRepository.InsertUser(ctx, &newUser)

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User is Successfuly registered"),
	}, nil
}

func (as *authService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	// cek apakah email ada
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.LoginResponse{
			Base: utils.BadRequestResponse("user doesnt not exist"),
		}, nil
	}
	// cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
		}
		return nil, err
	}
	// generate jwt
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtentity.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.RoleCode,
	})

	secretKey := os.Getenv("JWT_SECRET")
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Base:        utils.SuccessResponse("login successfull"),
		AccessToken: accessToken,
	}, nil
	// kirim response
}

func (as *authService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	// ambil metadata
	jwtToken, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tokenClaims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// blacklist token di cache hingga masa berlaku habis
	as.cacheService.Set(jwtToken, "", time.Until(time.Unix(tokenClaims.ExpiresAt.Unix(), 0)))

	return &auth.LogoutResponse{
		Base: utils.SuccessResponse("logout success"),
	}, nil
}

func (as *authService) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	// cek apakah pass confirm sama

	if request.NewPassword != request.NewPasswordConfirmation {
		return &auth.ChangePasswordResponse{
			Base: utils.BadRequestResponse("Password is not match, try again!"),
		}, nil
	}
	// cek apakah old pass sama
	jwtToken, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := jwtentity.GetClaimsFromToken(jwtToken)
	if err != nil {
		return nil, err
	}

	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if user == nil {
		return &auth.ChangePasswordResponse{
			Base: utils.BadRequestResponse("user doens't exist"),
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return &auth.ChangePasswordResponse{
				Base: utils.BadRequestResponse("old password not match"),
			}, nil
		}
		return nil, err
	}
	// update new pass ke db
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	if err != nil {
		return nil, err
	}

	err = as.authRepository.UpdateUserPassword(ctx, user.Id, string(hashedPass), user.FullName)
	if err != nil {
		return nil, err
	}
	// kirim resp
	return &auth.ChangePasswordResponse{
		Base: utils.SuccessResponse("change password success"),
	}, nil
}

func (as *authService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	// get data token

	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// ambil data dari db
	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.GetProfileResponse{
			Base: utils.BadRequestResponse("user doesnt exist"),
		}, nil
	}

	// kirim resp
	return &auth.GetProfileResponse{
		Base:        utils.SuccessResponse("get profile success"),
		UserId:      claims.Subject,
		FullName:    claims.FullName,
		Email:       claims.Email,
		RoleCode:    claims.Role,
		MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}

func NewAuthService(authRepository repository.IAuthRepository, cacheService *gocache.Cache) IAuthService {
	return &authService{
		authRepository: authRepository,
		cacheService:   cacheService,
	}
}
