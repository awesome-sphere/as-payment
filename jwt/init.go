package jwt

import (
	"github.com/awesome-sphere/as-payment/utils"
)

const BEARER_TOKEN = "Bearer "

var SECRET_KEY string
var ISSUER string

func InitializeJWTSettings() {
	SECRET_KEY = utils.GetenvOr("SECRET_KEY", "very_secret_key")
	ISSUER = utils.GetenvOr("ISSUER", "payment")
}
