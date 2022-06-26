package service

import (
	"streaming/internal/model"
	"streaming/internal/request"
	"streaming/internal/response"

	"go.uber.org/zap"
)

type ChatService interface {
	CreateChat(createPostReq *request.CreateChatRequest, userID string) (code response.Code, err error)
	FetchAllChats() (code response.Code, posts *[]model.Chat, err error)
	FetchAllChatsByUserID(userID string) (code response.Code, posts *[]model.Chat, err error)
}

type UserService interface {
	FetchCurrentUser(userID string) (response.Code, *response.FetchUserResponse, error)
	CreateUser(createUserReq *request.RegisterUserRequest) (code response.Code, err error)
	FetchUserByUsername(username string) (response.Code, *response.FetchUserResponse, error)
	DeleteUserByUsername(username string) (code response.Code, err error)
}

type Config interface {
	Logger() *zap.Logger
}
