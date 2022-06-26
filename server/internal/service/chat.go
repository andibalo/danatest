package service

import (
	"streaming/internal/dto"
	"streaming/internal/model"
	"streaming/internal/request"
	"streaming/internal/response"
	"streaming/internal/storage"

	"go.uber.org/zap"
)

type chatService struct {
	config  Config
	storage storage.Storage
}

func NewChatService(config Config, store storage.Storage) *chatService {

	return &chatService{
		config:  config,
		storage: store,
	}
}

func (s *chatService) CreateChat(createChatReq *request.CreateChatRequest, userID string) (code response.Code, err error) {

	s.config.Logger().Info("CreateChat: creating chat")

	chatIn := &dto.CreateChat{
		Message: createChatReq.Message,
		UserID:  userID,
	}

	_, err = s.storage.CreateChat(chatIn)

	if err != nil {
		s.config.Logger().Error("CreateChat: error creating chat", zap.Error(err))
		return response.ServerError, err
	}

	return response.Success, nil
}

func (s *chatService) FetchAllChats() (code response.Code, chats *[]model.Chat, err error) {

	s.config.Logger().Info("FetchAllChats: fetching all chats")

	chats, err = s.storage.FindAllChats()

	if err != nil {
		s.config.Logger().Error("FetchAllChats: error fetching all chats", zap.Error(err))
		return response.ServerError, nil, err
	}

	return response.Success, chats, nil
}

func (s *chatService) FetchAllChatsByUserID(userID string) (code response.Code, posts *[]model.Chat, err error) {

	s.config.Logger().Info("FetchAllChatsByUserID: fetching chats")

	posts, err = s.storage.FindAllChatsByUserID(userID)
	if err != nil {
		s.config.Logger().Error("FetchAllChatsByUserID: error fetching chats", zap.Error(err))
		return response.ServerError, nil, err
	}

	return response.Success, posts, nil
}
