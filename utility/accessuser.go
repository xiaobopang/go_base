package utility

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"go-base/internal/consts"
)

func GetAccessUser(ctx context.Context) string {
	// Get request object from context
	r := g.RequestFromCtx(ctx)

	return r.GetCtxVar(consts.CtxAccessUserKey).String()
}

func GetAccessUserMail(ctx context.Context) string {
	// Get request object from context
	r := g.RequestFromCtx(ctx)

	return r.GetCtxVar(consts.CtxAccessUserMailKey).String()
}
