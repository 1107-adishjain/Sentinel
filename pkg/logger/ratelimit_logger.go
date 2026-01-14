package logger

import (
	"github.com/1107-adishjain/sentinel/pkg/models"
)

// RateLimitLogger is responsible ONLY for accepting log events
// and pushing them to a buffered channel.
type RateLimitLogger struct {
	events chan models.RateLimitEvent
}

// NewRateLimitLogger initializes the buffered channel.
func NewRateLimitLogger(bufferSize int) *RateLimitLogger {
	return &RateLimitLogger{
		events: make(chan models.RateLimitEvent, bufferSize),
	}
}

// Events -> exposes the channel to the worker (read-only usage).
func (l *RateLimitLogger) Events() <-chan models.RateLimitEvent {
	return l.events
}

// Log pushes an event into the channel WITHOUT blocking.
// If buffer is full, event is dropped (by design).
func (l *RateLimitLogger) Log(event models.RateLimitEvent) {
	select {
	case l.events <- event: //pushes event to channel and proceed.
	default:
	}
}

// Close closes the events channel.
func (l *RateLimitLogger) Close() {
	close(l.events)
}
