package auths

import (
	"context"
	"fmt"
	"server/app/database"
	"server/app/exceptions"
	"server/app/models"
	"server/app/pkg/constants"
	"server/app/pkg/helpers"
	"server/app/repository"
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

	c.Next()
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

	decodedErr := helpers.DecodeJwtToken(jwtToken, &userClaim)

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

func GinContextFormContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(constants.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("Cound not retieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func AuthAdminFormCtx(ctx context.Context) (*models.User, error) {
	gc, _ := GinContextFormContext(ctx)

	currentUser := gc.Value(constants.ContextCurrentUser)
	if currentUser == nil {
		return nil, exceptions.NewUnauthorizedError("unauthorized")
	}

	return currentUser.(*models.User), nil
}
