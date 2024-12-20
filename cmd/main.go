package main

import (
	"github.com/MuhammedAshifVnr/user_service/internal/config"
	"github.com/MuhammedAshifVnr/user_service/internal/di"
)

func main() {
	config.LoadEnv()
	di.InitializeService()
}
