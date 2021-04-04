package users

type LoginValidator struct {
	IdToken string `form:"idToken" json:"idToken" binding:"required"`
}

type CreateValidator struct {
	Email string `form:"email" json:"email" binding:"required"`
}
