package models

import (
	"time"
)

type RefreshToken struct {
	UserID    uint      `json:"user_id"    gorm:"not null;index"`
	Token     string    `json:"token"      gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	Revoked   bool      `json:"revoked"    gorm:"default:false"`
	// Useful for security audits
	UserAgent string `json:"user_agent" gorm:"size:500"`
	IPAddress string `json:"ip_address" gorm:"size:100"`
}

// IsValid checks expiry and revocation in one place
func (rt *RefreshToken) IsValid() bool {
	return !rt.Revoked && time.Now().Before(rt.ExpiresAt)
}
