package service

import (
	"github.com/awesome-sphere/as-payment/jwt"
	"github.com/gin-gonic/gin"
)

func PayOrder(c *gin.Context) {
	is_valid, _ := jwt.AuthorizeToken(c)
	if is_valid {

	}
}
