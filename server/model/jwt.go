
package model

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)


type MyClaims struct {
	Id int
	jwt.StandardClaims
}

// TokenExpireDuration 定义JWT的过期时间
const TokenExpireDuration = time.Hour * 24 * 30

// MySecret 自定义签名字段(私货)
var MySecret = []byte("ayanami")

// GenToken 生成JWT
func GenToken(user User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "chess",                                    // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
