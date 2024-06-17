package logic

import (
	"context"
	"time"

	"eventos-tec/common/pb"
	"eventos-tec/zrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEventsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEventsLogic {
	return &ListEventsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListEventsLogic) ListEvents(in *__.ListEventsRequest) (*__.ListEventsResponse, error) {
	page := in.GetPage()
	pageSize := in.GetPageSize()

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	modelEventList, totalPages, currentPage, err := l.svcCtx.EventModel.ListUpcomingEvents(l.ctx, in.Title, in.City, in.Uf, in.StartDate, in.EndDate, page, pageSize)
	if err != nil {
		return nil, err
	}
	var protoEventList []*__.EventWithLocationAndType
	for _, item := range modelEventList {
		var dataHora string
		if !item.Date.IsZero() {
			dataHora = item.Date.Format(time.RFC3339)
		} else {
			dataHora = ""
		}
		protoEventLocal := &__.EventWithLocationAndType{
			Title:       item.Title,
			Date:        dataHora,
			Banner:      item.Imgurl,
			Description: item.Description,
			Location:    item.Location,
		}
		if item.Remote == true {
			protoEventLocal.Type = "remote"
		} else {
			protoEventLocal.Type = "normal"
		}
		protoEventList = append(protoEventList, protoEventLocal)
	}

	return &__.ListEventsResponse{
		Events:      protoEventList,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
	}, nil
}
