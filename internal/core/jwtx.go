package core

import (
	"errors"
	jwt1 "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenExpired     = "Token已过期"
	TokenNotValidYet = "Token不再有效"
	TokenMalformed   = "Token非法"
	TokenInvalid     = "Token无效"
)

type CustomClaims struct {
	jwt1.StandardClaims
	Mobile string
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(ctx.JWTSecret)}
}

func (j *JWT) GenerateToken(text string, duration time.Duration) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		StandardClaims: jwt1.StandardClaims{
			NotBefore: now.Unix(),
			ExpiresAt: now.Add(duration).Unix(),
		},
		Mobile: text,
	}
	at := jwt1.NewWithClaims(jwt1.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(ctx.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) ParseToken(tokenStr string) (string, error) {
	token, err := jwt1.ParseWithClaims(tokenStr, &CustomClaims{}, func(tokenStr *jwt1.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if result, ok := err.(jwt1.ValidationError); ok {
			if result.Errors&jwt1.ValidationErrorMalformed != 0 {
				return "", errors.New(TokenMalformed)
			} else if result.Errors&jwt1.ValidationErrorExpired != 0 {
				return "", errors.New(TokenExpired)
			} else if result.Errors&jwt1.ValidationErrorNotValidYet != 0 {
				return "", errors.New(TokenNotValidYet)
			} else {
				return "", errors.New(TokenInvalid)
			}
		}
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Mobile, nil
	}
	return "", errors.New(TokenInvalid)
}

func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	duration := time.Hour * 24 * 7 // 一周
	token, err := jwt1.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt1.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt1.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(duration).Unix()
		return j.GenerateToken(claims.Mobile, duration)
	}
	return "", errors.New(TokenInvalid)
}
