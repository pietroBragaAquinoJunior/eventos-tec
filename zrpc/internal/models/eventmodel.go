package models

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var (
	_                    EventModel = (*customEventModel)(nil)
	eventWithAddressRows            = "e.id as Id, e.title as Title, e.description as Description, e.imgurl as Imgurl, e.eventurl as Eventurl, e.remote as Remote, e.date as Date, a.city as Location"
)

type (
	// EventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventModel.
	EventModel interface {
		eventModel
		ListUpcomingEvents(ctx context.Context, page int32, pageSize int32) ([]EventWithAddress, int32, int32, error)
	}

	customEventModel struct {
		*defaultEventModel
	}

	EventWithAddress struct {
		Id          string    `db:"id"`
		Title       string    `db:"title"`
		Description string    `db:"description"`
		Imgurl      string    `db:"imgurl"`
		Eventurl    string    `db:"eventurl"`
		Remote      bool      `db:"remote"`
		Date        time.Time `db:"date"`
		Location    string    `db:"location"`
	}
)

// NewEventModel returns a model for the database table.
func NewEventModel(conn sqlx.SqlConn) EventModel {
	return &customEventModel{
		defaultEventModel: newEventModel(conn),
	}
}

func (m *customEventModel) ListUpcomingEvents(ctx context.Context, page int32, pageSize int32) ([]EventWithAddress, int32, int32, error) {
	offset := (page - 1) * pageSize
	hoje := time.Now().Format("2006-01-02")
	query := fmt.Sprintf("SELECT %s FROM %s e JOIN address a ON a.event_id = e.id WHERE e.date >= $1 LIMIT $2 OFFSET $3", eventWithAddressRows, m.table)
	var resp []EventWithAddress
	err := m.conn.QueryRowsCtx(ctx, &resp, query, hoje, pageSize, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	// Especificar claramente que 'event_id' é de 'a' (address) no COUNT
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s e JOIN address a on a.event_id = e.id WHERE e.date >= $1", m.table)
	var totalRecords int32
	err = m.conn.QueryRowCtx(ctx, &totalRecords, countQuery, hoje)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return resp, totalPages, page, nil
}
