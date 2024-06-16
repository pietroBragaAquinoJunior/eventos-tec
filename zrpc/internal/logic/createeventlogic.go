package logic

import (
	"context"
	"eventos-tec/common/pb"
	"eventos-tec/zrpc/internal/models"
	"eventos-tec/zrpc/internal/svc"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEventLogic {
	return &CreateEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEventLogic) CreateEvent(in *__.CreateEventRequest) (*__.CreateEventResponse, error) {
	layout := "02-01-2006"

	parsedDate, err := time.Parse(layout, in.Event.Date)
	if err != nil {
		fmt.Println("Erro ao parsear a data:", err)
		return nil, err
	}

	m := &models.Event{Id: uuid.NewString(), Title: in.Event.Title, Date: parsedDate, Description: in.Event.Description, Imgurl: in.Event.Imgurl, Eventurl: in.Event.Eventurl, Remote: in.Event.Remote}

	_, err = l.svcCtx.EventModel.Insert(l.ctx, m)
	if err != nil {
		return nil, err
	}

	return &__.CreateEventResponse{Msg: "Inserido com sucesso."}, nil
}
