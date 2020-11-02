package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// CustomeClaims 证书结构
type CustomeClaims struct {
	user string
	iat  int64
	jwt.StandardClaims
}

// GenerateAssesToken  颁发一个有限期72小时的证书
func GenerateAssesToken(account string) (string, error) {
	customeClaims := &CustomeClaims{
		user: account,
		iat:  time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customeClaims)
	fmt.Println("secret", beego.AppConfig.String("secret"))
	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("secret")))
	return tokenString, err
}

// ParseAssesToken 解析 accessToken
func ParseAssesToken(accessToken string) (*jwt.Token, error) {
	jwtInfo, err := jwt.ParseWithClaims(accessToken, &CustomeClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		// fmt.Println(((*token).Claims))
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(beego.AppConfig.String("secret")), nil
	})
	fmt.Println("aaaaaa:", jwtInfo)
	return jwtInfo, err
}
