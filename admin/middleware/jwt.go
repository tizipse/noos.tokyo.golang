package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/herhe-com/framework/auth"
	"github.com/tizips/noos.tokyo/admin/constants"
	"github.com/tizips/noos.tokyo/admin/helper/authorize"
)

func Jwt() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {

		if token := c.GetHeader(constants.JwtOfAuthorization); len(token) > 0 {

			var claims jwt.RegisteredClaims

			refresh, err := auth.CheckJWT(&claims, string(token), constants.JwtOfIssuerWithAdmin)

			if err == nil {
				c.Set(constants.ContextOfIdWithAdmin, claims.Subject)
				c.Set(constants.ContextOfClaimsWithAdmin, claims)
			}

			if refresh {

				var refreshToken string

				if refreshToken, err = auth.RefreshJWT(ctx, &claims); err != nil {
					return
				}

				c.Set(constants.ContextOfIdWithAdmin, claims.Subject)
				c.Set(constants.ContextOfClaimsWithAdmin, claims)
				c.Header(constants.JwtOfAuthorization, refreshToken)

				if authorize.Check(c) {
					_ = auth.DoRoleOfRefresh(ctx, authorize.ID(c))
				}
			}
		}
	}
}
