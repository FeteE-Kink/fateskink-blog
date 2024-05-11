package auths

import (
	"context"
	"server/database"
	"server/exceptions"
	"server/models"
	"server/pkg/constants"
	"server/pkg/jsonWebToken"
	"server/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtTokenCheck(c *gin.Context) {
	jwtToken, tokenErr := extractBearerToken(c.GetHeader(constants.AuthorizationHeader))

	if tokenErr != nil {
		return
	}

	uid, parseError := parseToken(jwtToken)

	if parseError == nil {
		user := models.User{Id: uid}

		repo := repository.NewUserRepository(nil, database.Db)

		if err := repo.Find(&user); err == nil {
			c.Set(constants.ContextCurrentUser, user)
		}
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", exceptions.NewUnauthorizedError("Bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", exceptions.NewUnauthorizedError("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (uid int32, err error) {
	var userClaim models.UserClaims

	decodedErr := jsonWebToken.DecodeJwtToken(jwtToken, &userClaim)

	if decodedErr != nil {
		return 0, exceptions.NewUnauthorizedError("Bad jwt token")
	}

	return userClaim.Sub, nil
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constants.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
