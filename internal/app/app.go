package app

import (
	"slog"

	"github.com/hyphypnotic/NISmentor/internal/app/config"
)

func Run() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		slog.Error("Failed to load config: %v", err)
		return
	}
}
