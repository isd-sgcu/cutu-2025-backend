package domain

type TokenResponse struct {
	UserID       string `json:"userId"`
	QrURL        string `json:"qrUrl"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
