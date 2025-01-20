package domain

type Role string

const (
	RoleParticipant Role = "participant"
	RoleStaff       Role = "staff"
	RoleAdmin       Role = "admin"
)
