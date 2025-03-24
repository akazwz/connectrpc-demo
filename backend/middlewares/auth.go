package middlewares

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"

	"opendrive/gen/auth/v1/authv1connect"
	"opendrive/utils"
)

var publicAuthServiceRegisterProcedures = []string{
	authv1connect.PublicAuthServiceRegisterProcedure,
	authv1connect.PublicAuthServiceLoginProcedure,
	authv1connect.PublicAuthServiceRefreshTokenProcedure,
}

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			for _, procedure := range publicAuthServiceRegisterProcedures {
				if req.Spec().Procedure == procedure {
					return next(ctx, req)
				}
			}
			bearerToken := req.Header().Get("Authorization")
			if bearerToken == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			token := strings.Split(bearerToken, " ")[1]
			if token == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}
			verifyToken, err := utils.VerifyToken(token)
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}
			ctx = context.WithValue(ctx, "user_id", verifyToken.UserID)
			return next(ctx, req)
		}
	}
	return interceptor
}
