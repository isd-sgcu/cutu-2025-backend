package domain

type User struct {
	ID             string  `gorm:"primaryKey" json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	University     string  `json:"university"`
	SizeJersey     string  `json:"sizeJersey"`
	FoodLimitation string  `json:"foodLimitation"`
	InvitationCode *string `json:"invitationCode"`
	State          string  `json:"state"`
	ImageURL       string  `json:"imageURL"`
	Role           Role    `json:"role"`
}
