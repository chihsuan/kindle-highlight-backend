package users

import (
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (serializer *UserSerializer) Serialize(userModel UserModel) UserResponse {
	user := UserResponse{
		ID:    userModel.ID,
		Email: userModel.Email,
	}
	return user
}
