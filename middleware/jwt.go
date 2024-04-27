package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "12345678"
)

var Token VoteJwt

type VoteJwt struct {
	Secret []byte
}

// Claim 自定义的数据结构，这里使用了结构体的组合
type Claim struct {
	jwt.RegisteredClaims
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Nick string `json:"nicks"`
	Role string `json:"role"`
}

func (j *VoteJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func (j *VoteJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}

// GetToken 颁发token access token 和 refresh token
func (j *VoteJwt) GetToken(id uint64, name, nick string, role string) (aToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := Claim{
		ID:               id,
		Name:             name,
		Nick:             nick,
		Role:             role,
		RegisteredClaims: rc,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	bytes := []byte(TokenIssuer)
	return claims.SignedString(bytes)

}

// ImJwtAuthMiddleware 用户单独聊天室的权限验证服务
func (j *VoteJwt) ImJwtAuthMiddleware(tokenStr string) (id, name, nick string, err error) {
	claim := &Claim{}
	if tokenStr == "" {
		fmt.Println("Im验证权限失败")
		return "", "", "", err
	}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		bytes := []byte(TokenIssuer)
		return bytes, nil
	})
	if err != nil {
		fmt.Println("Im验证权限失败")
		return "", "", "", err
	}
	if !token.Valid {
		fmt.Println("Im验证权限失败")
		return "", "", "", err
	}
	id = strconv.Itoa(int(claim.ID))
	return id, claim.Name, claim.Nick, nil
}
