package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		ctx.JSON(http.StatusUnauthorized, result)
		ctx.Abort()
	}
}

type PayloadParams struct {
	ID        uint      `json:"id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func AuthContextToID(ctx context.Context) (string, error) {
	getUser := ctx.Value("authorization_payload")
	var result PayloadParams
	JSONToStruct(getUser, &result)
	if result.ID == 0 {
		return "", errors.New("User must be authenticated")
	}
	return fmt.Sprint(result.ID), nil
}

func JSONToStruct(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
