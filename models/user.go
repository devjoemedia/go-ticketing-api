package models

type User struct {
	ID            uint           `json:"id"`
	Name          string         `json:"name"          gorm:"not null"`
	Email         string         `json:"email"         gorm:"uniqueIndex;not null"`
	Password      string         `json:"-"             gorm:"not null"` // never serialized
	RefreshTokens []RefreshToken `json:"-"            gorm:"foreignKey:UserID"`
}

// ─── Request / Response shapes ───────────────────────────
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type RegisterRequest struct {
	Name     string `json:"name"     validate:"required,min=2,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
