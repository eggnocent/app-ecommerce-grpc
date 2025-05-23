package main

import (
	"context"
	grpcmiddleware "github/eggnocent/app-grpc-eccomerce/internal/grpc-middleware"
	"github/eggnocent/app-grpc-eccomerce/internal/handler"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/pb/auth"
	"github/eggnocent/app-grpc-eccomerce/pb/product"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
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
	authMiddleware := grpcmiddleware.NewAuthMiddleware(cacheService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, cacheService)
	authHandler := handler.NewAuthHandler(authService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)
	product.RegisterProductServiceServer(serv, productHandler)

	if os.Getenv("ENVIRONTMENT") == "dev" {
		reflection.Register(serv)
		log.Println("reflection is register")
	}

	log.Println("server is running on port :50051")
	if err := serv.Serve(list); err != nil {
		log.Printf("server is error %v", err)
	}
}
