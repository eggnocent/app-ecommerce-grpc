package main

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/handler"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/pb/auth"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
	grpcmiddleware "github/eggnocent/app-grpc-eccomerce/pkg/grpc-middleware"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	gocache "github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	godotenv.Load()
	list, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Panicf("Error when listening: %v", err)
	}

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("connected to database...")
	cacheService := gocache.New(time.Hour*24, time.Hour)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, cacheService)
	authHandler := handler.NewAuthHandler(authService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)

	if os.Getenv("ENVIRONTMENT") == "dev" {
		reflection.Register(serv)
		log.Println("reflection is register")
	}

	log.Println("server is running on port :50051")
	if err := serv.Serve(list); err != nil {
		log.Printf("server is error %v", err)
	}
}
