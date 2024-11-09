package routes

import (
	"server/app/database"
	"server/app/pkg/auths"
	"server/app/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func RegisterGraphQLRouter(r *gin.Engine) {
	r.POST("/fateskinkGql",
		auths.JwtTokenCheck,
		auths.GinContextToContextMiddleware(),
		initializers.FateskinkGqlHandler(database.Db),
	)
}
