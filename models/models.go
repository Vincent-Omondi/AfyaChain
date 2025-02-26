package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID             int             `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName      string          `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName       string          `gorm:"type:varchar(100);not null" json:"last_name"`
	Gender         string          `gorm:"type:varchar(10);not null" json:"gender"`
	Email          string          `gorm:"type:varchar(100);unique;not null" json:"email"`
	Phone          string          `gorm:"type:varchar(15);not null" json:"phone"`
	Address        string          `gorm:"type:varchar(255)" json:"address"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID" json:"medical_records"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

type MedicalRecord struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID  int       `gorm:"not null" json:"patient_id"`
	RecordDate time.Time `gorm:"not null" json:"record_date"`
	Condition  string    `gorm:"type:text;not null" json:"condition"`
	Treatment  string    `gorm:"type:text" json:"treatment"`
	Notes      string    `gorm:"type:text" json:"notes"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
type AccessLog struct {
	PatientID int    `json:"patient_id"`
	UserID    int    `json:"user_id"`
	Action    string `json:"action"` // e.g., "view", "edit"
	Timestamp string `json:"timestamp"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthKey  string `json:"authKey"`
	UserType string `json:"userType"`
}
type BlockD struct {
	ID           uint      `gorm:"primaryKey"`
	PatientID    int       `json:"patient_id"`
	UserID       int       `json:"user_id"`
	Action       string    `json:"action"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previous_hash"`
	Hash         string    `json:"hash"`
	Pow          int       `json:"pow"`
	PreviousID   *uint     `gorm:"index"` // Foreign key to the previous block's ID
}

type Blockchain struct {
	gorm.Model
	Blocks []BlockD `gorm:"foreignKey:BlockchainID"`
}
type SignupForm struct {
	FirstName       string `json:"firstname" form:"firstname" binding:"required"`
	LastName        string `json:"lastname" form:"lastname" binding:"required"`
	Phone           string `json:"phone" form:"phone" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm-password" binding:"required"`
	Sex             string `json:"sex" form:"sex" binding:"required"`
	Country         string `json:"country" form:"country" binding:"required"`
	City            string `json:"city" form:"city" binding:"required"`
}
type Facility struct {
	gorm.Model
	FacilityName       string `json:"facility_name" gorm:"unique;not null"`
	RegistrationNumber string `json:"registration_number" gorm:"unique;not null"`
	PhoneNumber        string `json:"phone_number" gorm:"not null"`
	Email              string `json:"email" gorm:"unique;not null"`
	Password           string `json:"password" gorm:"not null"`
	Country            string `json:"country" gorm:"not null"`
	City               string `json:"city" gorm:"not null"`
	Address            string `json:"address" gorm:"not null"`
}
