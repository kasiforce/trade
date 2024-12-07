package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/consts"
	"time"
)

var jwtSecret = []byte("TestSecret")

type Claims struct {
	ID   int
	Name string
	jwt.StandardClaims
}

// GenerateToken 传入用户id、用户名（或者管理员id、名字）生成唯一的token
func GenerateToken(id int, name string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	claims := Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "trade",
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	token = "Bearer " + token
	return token, err
}

// ParseToken 验证token是否正确、过期，同时解析出用户id/管理员id，并设置到context中。
// 当token验证通过，后续可通过id := c.GetInt("id")获取id
func ParseToken(c *gin.Context, token string) (err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			if claims.ExpiresAt > time.Now().Unix() {
				c.Set("id", claims.ID)
				//return GenerateToken(claims.ID, claims.Name)
				return nil
			}
			return errors.New("token expired")
		}
	}
	return err
}
