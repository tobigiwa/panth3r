package http

import (
	"log/slog"
	"panth3rWaitlistBackend/internal/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	repository store.Service
	logger     *slog.Logger
}

func NewApplication(db *mongo.Collection, logger *slog.Logger) *Application {
	return &Application{
		repository: store.NewStore(db),
		logger:     logger,
	}
}
