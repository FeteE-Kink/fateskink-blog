package models

import (
	"regexp"
	"server/enums"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	Id                int32
	Email             string
	EncryptedPassword string
	FullName          string
	LockVersion       int32
	Name              string
	Gender            int32
	Role              enums.UserRole
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", u.LockVersion+1)
	}

	return
}

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
