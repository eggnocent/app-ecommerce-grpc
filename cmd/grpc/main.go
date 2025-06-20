package main

import (
	"context"
	grpcmiddleware "github/eggnocent/app-grpc-eccomerce/internal/grpc-middleware"
	"github/eggnocent/app-grpc-eccomerce/internal/handler"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/pb/auth"
	"github/eggnocent/app-grpc-eccomerce/pb/cart"
	"github/eggnocent/app-grpc-eccomerce/pb/newsletter"
	"github/eggnocent/app-grpc-eccomerce/pb/order"
	"github/eggnocent/app-grpc-eccomerce/pb/product"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	gocache "github.com/patrickmn/go-cache"
	"github.com/xendit/xendit-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	godotenv.Load()

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET")

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

	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(productRepository, cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(db, orderRepository, productRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	newsletterRepository := repository.NewNewsletterRepository(db)
	newsletterService := service.NewNewsletterService(newsletterRepository)
	newsletterHandler := handler.NewNewsletterHandler(newsletterService)

	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)
	product.RegisterProductServiceServer(serv, productHandler)
	cart.RegisterCartServiceServer(serv, cartHandler)
	order.RegisterOrderServiceServer(serv, orderHandler)
	newsletter.RegisterNewsletterServiceServer(serv, newsletterHandler)

	if os.Getenv("ENVIRONTMENT") == "dev" {
		reflection.Register(serv)
		log.Println("reflection is register")
	}

	log.Println("server is running on port :50051")
	if err := serv.Serve(list); err != nil {
		log.Printf("server is error %v", err)
	}
}
