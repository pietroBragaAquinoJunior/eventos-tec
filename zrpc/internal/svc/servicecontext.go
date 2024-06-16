package svc

import (
	"eventos-tec/zrpc/internal/config"
	"eventos-tec/zrpc/internal/models"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	EventModel  models.EventModel
	CouponModel models.CouponModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("pgx", c.DataSourceName)
	return &ServiceContext{
		Config:      c,
		EventModel:  models.NewEventModel(conn),
		CouponModel: models.NewCouponModel(conn),
	}
}
