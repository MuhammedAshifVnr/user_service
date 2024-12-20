package config

import (
	"log"
	"net"

	gp "github.com/MuhammedAshifVnr/user_service/internal/grpc"
	"github.com/MuhammedAshifVnr/user_service/internal/models"
	"github.com/MuhammedAshifVnr/user_service/internal/service"
	pb "github.com/MuhammedAshifVnr/user_service/proto"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	viper.AutomaticEnv()
}

func InitDB() (*gorm.DB, *redis.Client) {
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB :%v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	db.AutoMigrate(&models.User{})
	return db, rdb
}

func RunGRPCServer(userService service.UserService) {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := gp.NewUserHandler(userService)
	pb.RegisterUserServiceServer(grpcServer, server)

	log.Println("gRPC server running on :5001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
