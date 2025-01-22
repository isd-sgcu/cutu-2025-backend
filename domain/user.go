package domain

type Role string

const (
	Student Role = "student"
	Staff   Role = "staff"
	Admin   Role = "admin"
)

type User struct {
	ID             string `gorm:"primaryKey"`
	Name           string
	Email          string
	Phone          string
	University     string
	SizeJersey     string
	FoodLimitation string
	InvitationCode *string
	State          string
	ImageURL       string
	IsEntered	   bool
	Role           Role `json:"-"` // exclude from JSON
}
