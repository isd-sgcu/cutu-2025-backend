package domain

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string
	Email         string
	Phone         string
	University    string
	SizeJersey    string
	FoodLimitation string
	InvitationCode *string
	State         string
	ImageURL      string
}
