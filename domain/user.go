package domain

import "time"

type Role string
type Status string
type Education string

const (
	Member Role = "member"
	Staff  Role = "staff"
	Admin  Role = "admin"
)

const (
	StatusStudent       Status = "student"
	StatusAlumni        Status = "alumni"
	StatusGeneralPublic Status = "general_public"
)

const (
	EducationStudying  Education = "studying"
	EducationGraduated Education = "graduated"
)

type User struct {
	ID             string `gorm:"primaryKey"`
	Name           string
	Email          string
	Phone          string `gorm:"unique"` // Make phone unique
	University     string
	SizeJersey     string
	FoodLimitation string
	InvitationCode *string
	Status         Status
	GraduatedYear  *string
	Faculty        *string
	ImageURL       string
	LastEntered    *time.Time // Timestamp for the last QR scan
	Role           Role
	Education      Education
}

