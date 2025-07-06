package models

type Service struct {
	ServiceID       int    `json:"service_id"`
	ProviderID      int    `json:"provider_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes"`
	PriceTime       int    `json:"price_time"`
}