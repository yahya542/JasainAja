package models

import "time"

type Request struct {
	RequestID   int       `json:"request_id"`
	UserID      int       `json:"user_id"`
	ProviderID  int       `json:"provider_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
