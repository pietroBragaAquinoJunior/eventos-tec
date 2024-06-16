package models

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EventModel = (*customEventModel)(nil)

type (
	// EventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventModel.
	EventModel interface {
		eventModel
		ListEvents(ctx context.Context, page int32, pageSize int32) ([]Event, int32, int32, error)
	}

	customEventModel struct {
		*defaultEventModel
	}
)

// NewEventModel returns a model for the database table.
func NewEventModel(conn sqlx.SqlConn) EventModel {
	return &customEventModel{
		defaultEventModel: newEventModel(conn),
	}
}

func (m *customEventModel) ListEvents(ctx context.Context, page int32, pageSize int32) ([]Event, int32, int32, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", eventRows, m.table, pageSize, offset)

	var resp []Event
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, 0, 0, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	var totalRecords int32
	err = m.conn.QueryRowCtx(ctx, &totalRecords, countQuery)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalRecords + pageSize - 1) / pageSize

	return resp, totalPages, page, nil
}
