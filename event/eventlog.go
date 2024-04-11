package event

import (
	"context"
)

// EventLog is the main struct for the event log
type EventLog struct {
	cfg     Config
	storage Storage
}

// NewEventLog creates and initializes an instance of EventLog
func NewEventLog(cfg Config, storage Storage) *EventLog {
	return &EventLog{
		cfg:     cfg,
		storage: storage,
	}
}

// LogEvent is used to store an event for runtime debugging
func (e *EventLog) LogEvent(ctx context.Context, event *Event) error {
	return e.storage.LogEvent(ctx, event)
}
