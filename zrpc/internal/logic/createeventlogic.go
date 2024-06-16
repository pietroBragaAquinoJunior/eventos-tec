package logic

import (
	"context"
	"encoding/base64"
	"eventos-tec/common/pb"
	"eventos-tec/zrpc/internal/models"
	"eventos-tec/zrpc/internal/svc"
	"github.com/google/uuid"
	"os"
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
	id := uuid.NewString()
	fileName := uuid.NewString() + ".jpg"
	filePath := l.svcCtx.Config.ImageStorePath + fileName

	// Verifica se o diretório existe, se não, cria o diretório
	if _, err := os.Stat(l.svcCtx.Config.ImageStorePath); os.IsNotExist(err) {
		if err := os.MkdirAll(l.svcCtx.Config.ImageStorePath, 0755); err != nil {
			return nil, err
		}
	}

	// Decodifica a string Base64 para bytes
	data, err := base64.StdEncoding.DecodeString(in.Event.Image)
	if err != nil {
		l.Logger.Errorf("Erro ao decodificar a imagem Base64: %v", err)
		return nil, err
	}

	// Salva o arquivo de imagem
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		l.Logger.Errorf("Erro ao salvar a imagem: %v", err)
		return nil, err
	}

	// Parseia a data do evento
	layout := "02-01-2006"
	parsedDate, err := time.Parse(layout, in.Event.Date)
	if err != nil {
		l.Logger.Errorf("Erro ao parsear a data: %v", err)
		return nil, err
	}

	// Cria um novo modelo de evento
	m := &models.Event{
		Id:          id,
		Title:       in.Event.Title,
		Date:        parsedDate,
		Description: in.Event.Description,
		Eventurl:    in.Event.Eventurl,
		Remote:      in.Event.Remote,
		Imgurl:      filePath,
	}

	// Insere o evento no banco de dados
	if _, err = l.svcCtx.EventModel.Insert(l.ctx, m); err != nil {
		l.Logger.Errorf("Erro ao inserir o evento no banco de dados: %v", err)
		return nil, err
	}

	return &__.CreateEventResponse{Msg: "Evento inserido com sucesso, ID: " + id}, nil
}
