package users

type LoginValidator struct {
	IdToken string `form:"idToken" json:"idToken" binding:"required"`
}
