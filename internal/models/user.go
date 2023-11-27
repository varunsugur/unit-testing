package models

import (
	"gorm.io/gorm"
)

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	DOB      string `json:"dob" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	DOB          string `json:"dob"`
	PasswordHash string `json:"-"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerifyUser struct {
	UserEmail string `json:"userEmail" validate:"required"`
	// UserEmail string `json:"useremail" validate:"required"`
	DOB string `json:"dob" validate:"required"`
}

type ResetDetails struct {
	Email           string `json:"email" validate:"required"`
	NewPassword     string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	OTP             string `json:"otp" validate:"required"`
}
