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
	StatusChulaStudent   Status = "chula_student"
	StatusAlumni         Status = "alumni"
	StatusGeneralPublic  Status = "general_public"
	StatusGeneralStudent Status = "general_student"
)

const (
	EducationStudying  Education = "studying"
	EducationGraduated Education = "graduated"
)

type User struct {
	ID             string     `json:"id" gorm:"primaryKey"`
	UID            string     `json:"uid" gorm:"unique"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone" gorm:"unique"` // Make phone unique
	University     string     `json:"university"`
	SizeJersey     string     `json:"sizeJersey"`
	FoodLimitation string     `json:"foodLimitation"`
	InvitationCode *string    `json:"invitationCode"`
	Age            *string    `json:"age"`
	ChronicDisease *string    `json:"chronicDisease"`
	DrugAllergy    *string    `json:"drugAllergy"`
	Status         Status     `json:"status"`
	GraduatedYear  *string    `json:"graduatedYear"`
	Faculty        *string    `json:"faculty"`
	ImageURL       string     `json:"imageUrl"`
	LastEntered    *time.Time `json:"lastEntered"` // Timestamp for the last QR scan
	RegisteredAt   time.Time  `json:"registeredAt"`
	Role           Role       `json:"role"`
	Education      Education  `json:"education"`
	IsAcroPhobia   bool       `json:"isAcroPhobia"`
}
