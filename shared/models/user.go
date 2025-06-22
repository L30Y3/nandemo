package models

type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Provider    string `json:"provider"`    // e.g. "google", "microsoft"
	ProviderID  string `json:"provider_id"` // ID from OAuth2 provider
}
