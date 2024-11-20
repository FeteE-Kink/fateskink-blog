package models

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	Id                int32
	Email             string
	EncryptedPassword string
	FullName          string
	Name              string
	Gender            *int32
	Avatar            *Attachment `gorm:"polymorphic:Owner;polymorphicValue:User"`
	About             *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Address           *string
	Phone             *string
	Birthday          *time.Time
}

type UserClaims struct {
	Sub int32
	jwt.RegisteredClaims
}

func (user *User) GenerateJwtClaims() (claims jwt.Claims) {
	return UserClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	if tx.Statement.Changed() {
// 		tx.Statement.SetColumn("lock_version", u.LockVersion+1)
// 	}

// 	return
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.assignDefaultData()
	return
}

func (u *User) assignDefaultData() (err error) {
	re := regexp.MustCompile(`(.*)@`)
	matches := re.FindStringSubmatch(u.Email)

	if len(matches) >= 2 {
		u.Name = matches[1]
	}

	return
}
