package fateskinkResolvers

import (
	"context"
	"server/app/graphql/gqlTypes/fateskinkTypes"
	fateskinkInputs "server/app/graphql/inputs/feateskinkInputs"
	"server/app/pkg/auths"
	"server/app/services"
)

func (r *Resolver) ArticleCrate(ctx context.Context, args struct {
	Input fateskinkInputs.ArticleCreateInput
}) (*fateskinkTypes.ResponseMessageType, error) {
	user, err := auths.AuthAdminFormCtx(ctx)
	if err != nil {
		return nil, err
	}

	service := services.ArticleCreateService{
		Input: &args.Input,
		User:  user,
		Db:    r.Db,
	}
	if err := service.Execute(); err != nil {
		return nil, err
	}

	return &fateskinkTypes.ResponseMessageType{
		Message: "Create Successfully!",
	}, nil
}
