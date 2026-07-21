// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"river/goZero/internal/svc"
	"river/goZero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoZeroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoZeroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoZeroLogic {
	return &GoZeroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoZeroLogic) GoZero(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
