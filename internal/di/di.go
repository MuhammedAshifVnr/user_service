package di

import (
	"github.com/MuhammedAshifVnr/user_service/internal/config"
	"github.com/MuhammedAshifVnr/user_service/internal/repo"
	"github.com/MuhammedAshifVnr/user_service/internal/service"
)

func InitializeService() {
	DB, redis := config.InitDB()
	repo := repo.NewUserRepository(DB, redis)
	service := service.NewUserService(repo)
	config.RunGRPCServer(*service)
}
