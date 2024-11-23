package fateskinkResolvers

import (
	"context"
	"server/app/graphql/gqlTypes/globalTypes"
	"server/app/graphql/inputs/globalInputs"
	"server/app/services"
)

func (r *Resolver) SignIn(ctx context.Context, args globalInputs.SignInInput) (*globalTypes.SignInType, error) {
	service := services.AuthService{
		Email:    args.Email,
		Password: args.Password,
		Ctx:      &ctx,
		Db:       r.Db,
	}

	err := service.Execute()

	if err != nil {
		return nil, err
	}

	return &globalTypes.SignInType{
		Token: service.Token,
	}, nil
}
