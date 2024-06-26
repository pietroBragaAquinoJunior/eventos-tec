// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	eventFieldNames          = builder.RawFieldNames(&Event{}, true)
	eventRows                = strings.Join(eventFieldNames, ",")
	eventRowsExpectAutoSet   = strings.Join(stringx.Remove(eventFieldNames), ",")
	eventRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(eventFieldNames, "id"))
)

type (
	eventModel interface {
		Insert(ctx context.Context, data *Event) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Event, error)
		Update(ctx context.Context, data *Event) error
		Delete(ctx context.Context, id string) error
	}

	defaultEventModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Event struct {
		Id          string    `db:"id"`
		Title       string    `db:"title"`
		Description string    `db:"description"`
		Imgurl      string    `db:"imgurl"`
		Eventurl    string    `db:"eventurl"`
		Remote      bool      `db:"remote"`
		Date        time.Time `db:"date"`
	}
)

func newEventModel(conn sqlx.SqlConn) *defaultEventModel {
	return &defaultEventModel{
		conn:  conn,
		table: `"public"."event"`,
	}
}

func (m *defaultEventModel) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultEventModel) FindOne(ctx context.Context, id string) (*Event, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", eventRows, m.table)
	var resp Event
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultEventModel) Insert(ctx context.Context, data *Event) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7)", m.table, eventRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Title, data.Description, data.Imgurl, data.Eventurl, data.Remote, data.Date)
	return ret, err
}

func (m *defaultEventModel) Update(ctx context.Context, data *Event) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, eventRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.Title, data.Description, data.Imgurl, data.Eventurl, data.Remote, data.Date)
	return err
}

func (m *defaultEventModel) tableName() string {
	return m.table
}
