package users

import (
	"github.com/gin-gonic/gin"

	"kindle-notes/apps/common"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (serializer *UserSerializer) Serialize(userModel UserModel) UserResponse {
	user := UserResponse{
		ID:    userModel.ID,
		Email: userModel.Email,
		Token: common.GenToken(userModel.ID),
	}
	return user
}
