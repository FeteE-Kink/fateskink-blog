package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id                int32
	Email             string
	EncryptedPassword string
	FullName          string
	LockVersion       int32
	Name              string
	// Avatar            *Attachment  `gorm:"polymorphic:Owner;polymorphicValue:User"`
	About     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   *string
	Phone     *string
	Birthday  *time.Time
}

type UserClaims struct {
	Sub int32
	jwt.RegisteredClaims
}
