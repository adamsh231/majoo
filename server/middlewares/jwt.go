package middlewares

import (
	"errors"
	"fmt"
	"github.com/adamsh231/majoo/packages/helper"
	jwtPkg "github.com/adamsh231/majoo/packages/jwt"
	"github.com/adamsh231/majoo/packages/messages"
	"github.com/adamsh231/majoo/server/http/handlers"
	"github.com/adamsh231/majoo/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

type JwtMiddleware struct {
	*usecase.Contract
}

func (jwtMiddleware JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}
	handler := handlers.Handler{UcContract: jwtMiddleware.Contract}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		helper.LogOnly(messages.Unauthorized, "middleware-jwt-checkHeader")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			helper.LogOnly(messages.UnexpectedSigningMethod, "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(jwtMiddleware.JwtCredential.TokenSecret)
		return secret, nil
	})
	if err != nil {
		helper.LogOnly(err.Error(), "middleware-jwt-checkClaims")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		helper.LogOnly(messages.ExpiredToken, "middleware-jwt-checkTokenLiveTime")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//jwe roll back encrypted id
	jweRes, err := jwtMiddleware.JweCredential.Rollback(claims.Payload)
	if err != nil {
		helper.LogOnly(err.Error(), "pkg-jwe-rollback")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}
	if jweRes == nil {
		helper.LogOnly(messages.Unauthorized, "pkg-jwe-resultNil")
		return handler.SendResponse(ctx, handlers.ResponseWithOutMeta, nil, nil, errors.New(messages.Unauthorized), http.StatusUnauthorized)
	}

	//TODO: CARI TAU CARA KERJA
	//set id to uce case contract
	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	roleID := fmt.Sprintf("%v", jweRes["roleID"])
	jwtMiddleware.Contract.UserID = claims.Id
	jwtMiddleware.Contract.RoleID = helper.StringToInt(roleID)

	return ctx.Next()
}
