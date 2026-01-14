package logger

import (
	"context"
	"log"

	"github.com/1107-adishjain/sentinel/pkg/models"
	"gorm.io/gorm"
)

// Worker listens to rate-limit events and persists them to DB.
type Worker struct {
	db     *gorm.DB
	events <-chan models.RateLimitEvent
}

// NewWorker wires DB and event channel together.
func NewWorker(db *gorm.DB, events <-chan models.RateLimitEvent) *Worker {
	return &Worker{
		db:     db,
		events: events,
	}
}

// Start begins the background loop.
// This should be started as a goroutine.
func (w *Worker) Start(ctx context.Context) {
	go func() {
		for event := range w.events {
			if err := w.db.Create(&event).Error; err != nil {
				log.Printf("Failed to store rate limit event: %v", err)
			}
		}
		log.Println("RateLimit logger worker exited cleanly")
	}()
}

