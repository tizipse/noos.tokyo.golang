package authorize

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/tizips/noos.tokyo/admin/constants"
	"github.com/tizips/noos.tokyo/model"
	"time"
)

func MakeJWT(user model.SysUser) (token string, err error) {

	lifetime := facades.Cfg.GetInt("jwt.lifetime")

	now := carbon.Now()

	claims := jwt.RegisteredClaims{
		Issuer:    auth.Issuer(constants.JwtOfIssuerWithAdmin),
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(now.ToStdTime()),
		NotBefore: jwt.NewNumericDate(now.ToStdTime()),
		ExpiresAt: jwt.NewNumericDate(now.AddDays(lifetime).ToStdTime()),
	}

	return auth.MakeJWT(claims)
}

func CheckBlacklistOfJwt(ctx context.Context, c *app.RequestContext) bool {

	if Claims(c) == nil {
		return false
	}

	return auth.CheckBlacklist(ctx, BlacklistOfKeyWithJwt(Claims(c).ID)...)
}

func BlacklistOfJwt(ctx context.Context, c *app.RequestContext) bool {

	if Claims(c) == nil {
		return false
	}

	now := carbon.Now()

	expires := time.Duration(Claims(c).ExpiresAt.Unix()+12*60*60-now.Timestamp()) * time.Second

	return auth.Blacklist(ctx, now.Timestamp(), expires, BlacklistOfKeyWithJwt(Claims(c).ID)...)
}

func BlacklistOfKeyWithJwt(id string) []any {
	return []any{"jwt", id}
}
