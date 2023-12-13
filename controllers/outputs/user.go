package outputs

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserProfile struct {
	Email string `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	EmailVerified bool `json:"email_verified" binding:"required"`
}

type HashedUserPass struct {
	Uuid uuid.UUID `json:"uuid" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}