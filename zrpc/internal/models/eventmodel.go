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
		ListUpcomingEvents(ctx context.Context, title string, city string, uf string, startDate string, endDate string, page int32, pageSize int32) ([]EventWithAddress, int32, int32, error)
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

func (m *customEventModel) ListUpcomingEvents(ctx context.Context, titulo string, cidade string, uf string, dataInicio string, dataFim string, pagina int32, tamanhoPagina int32) ([]EventWithAddress, int32, int32, error) {
	offset := (pagina - 1) * tamanhoPagina
	hoje := time.Now().Format("2006-01-02")
	if dataInicio == "" {
		dataInicio = time.Now().Format("2006-01-02")
	}
	if dataFim == "" {
		dataFim = time.Date(2100, time.November, 10, 23, 0, 0, 0, time.UTC).Format("2006-01-02")
	}
	query := fmt.Sprintf(`SELECT %s FROM %s e
    JOIN address a ON a.event_id = e.id
    WHERE e.date >= $1
    AND ($2::text IS NULL OR e.title LIKE '%%' || $2 || '%%')
    AND ($3::text IS NULL OR a.city LIKE '%%' || $3 || '%%')
    AND ($4::text IS NULL OR a.uf LIKE '%%' || $4 || '%%')
    AND ($5::date IS NULL OR e.date >= $5)
    AND ($6::date IS NULL OR e.date <= $6)
    LIMIT $7 OFFSET $8`, eventWithAddressRows, m.table)

	var resposta []EventWithAddress
	err := m.conn.QueryRowsCtx(ctx, &resposta, query, hoje, titulo, cidade, uf, dataInicio, dataFim, tamanhoPagina, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	// Adicionando os mesmos filtros ao COUNT
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM %s e
    JOIN address a ON a.event_id = e.id
    WHERE e.date >= $1
    AND ($2::text IS NULL OR e.title LIKE '%%' || $2 || '%%')
    AND ($3::text IS NULL OR a.city LIKE '%%' || $3 || '%%')
    AND ($4::text IS NULL OR a.uf LIKE '%%' || $4 || '%%')
    AND ($5::date IS NULL OR e.date >= $5)
    AND ($6::date IS NULL OR e.date <= $6)`, m.table)

	var totalRegistros int32
	err = m.conn.QueryRowCtx(ctx, &totalRegistros, countQuery, hoje, titulo, cidade, uf, dataInicio, dataFim)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPaginas := (totalRegistros + tamanhoPagina - 1) / tamanhoPagina

	return resposta, totalPaginas, pagina, nil
}
