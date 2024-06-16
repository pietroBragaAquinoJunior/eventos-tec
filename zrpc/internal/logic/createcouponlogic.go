package logic

import (
	"context"
	"eventos-tec/zrpc/internal/models"
	"github.com/google/uuid"
	"time"

	"eventos-tec/common/pb"
	"eventos-tec/zrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCouponLogic {
	return &CreateCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCouponLogic) CreateCoupon(in *__.CreateCouponRequest) (*__.CreateCouponResponse, error) {

	// Parseia a data do evento
	layout := "02-01-2006"
	parsedDate, err := time.Parse(layout, in.Coupon.Date)
	if err != nil {
		l.Logger.Errorf("Erro ao parsear a data: %v", err)
		return nil, err
	}

	couponId := uuid.NewString()

	_, err = l.svcCtx.CouponModel.Insert(l.ctx, &models.Coupon{Id: couponId, EventId: in.EventId, Code: in.Coupon.Code, Discount: int64(in.Coupon.Discount), ValidUntil: parsedDate})
	if err != nil {
		return nil, err
	}

	return &__.CreateCouponResponse{CouponId: couponId}, nil
}
