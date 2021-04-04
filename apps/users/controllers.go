package users

import (
	"kindle-notes/apps/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	loginValidator := LoginValidator{}
	if err := c.ShouldBind(&loginValidator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenInfo, err := common.VerifyIdToken(loginValidator.IdToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := FindOne(&UserModel{Email: tokenInfo.Email})
	if err != nil {
		user = UserModel{Email: tokenInfo.Email}
		if err := SaveOne(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
	}

	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Serialize(user)})
}
