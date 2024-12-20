package models

// User represents the user entity
// Includes validation tags for input validation
import (
	"errors"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	FName         string  `gorm:"size:100"`
	City          string  `gorm:"size:100;index"`
	Phone         string  `gorm:"size:15;index"`
	Height        float64 `gorm:"type:decimal(5,2)"`
	Married       bool    `gorm:"index"`
	SearchVector  string  `gorm:"type:tsvector;->"`
}


// Validate ensures User fields meet required criteria
func (u *User) Validate() error {
	if u.FName == "" {
		return errors.New("first name is required")
	}
	if len(u.Phone) != 10 {
		return errors.New("phone must be 10 digits")
	}
	
	return nil
}
