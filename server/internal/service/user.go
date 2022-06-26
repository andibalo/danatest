package service

import (
	"errors"
	"streaming/internal/dto"
	"streaming/internal/model"
	"streaming/internal/request"
	"streaming/internal/response"
	"streaming/internal/storage"
	"streaming/internal/voerrors"

	"go.uber.org/zap"
)

type userService struct {
	config Config
	store  storage.Storage
}

func NewUserService(config Config, store storage.Storage) *userService {

	return &userService{
		config: config,
		store:  store,
	}
}

func (s *userService) CreateUser(createUserReq *request.RegisterUserRequest) (code response.Code, err error) {

	s.config.Logger().Info("CreateUser: creating user")

	userIn := &dto.RegisterUser{
		Username: createUserReq.Username,
	}

	_, err = s.store.CreateUser(userIn)

	if err != nil {
		s.config.Logger().Error("CreateUser: error creating user", zap.Error(err))
		return response.ServerError, err
	}

	return response.Success, nil
}

func (s *userService) DeleteUserByUsername(username string) (code response.Code, err error) {

	s.config.Logger().Info("DeleteUserByUsername: deleting user")

	err = s.store.RemoveUserByUsername(username)

	if err != nil {
		s.config.Logger().Error("DeleteUserByUsername: error deleting user", zap.Error(err))
		return response.ServerError, err
	}

	return response.Success, nil
}

func (s *userService) FetchCurrentUser(userID string) (response.Code, *response.FetchUserResponse, error) {

	s.config.Logger().Info("FetchCurrentUser: finding user by id")
	user, err := s.store.FindUserByID(userID)

	if err != nil {
		if !errors.Is(err, voerrors.ErrNotFound) {
			s.config.Logger().Error("FetchCurrentUser: error finding user by email", zap.Error(err))
			return response.ServerError, nil, err
		}
	}

	resp := s.mapUserToUserResp(user)

	return response.Success, resp, nil
}

func (s *userService) FetchUserByUsername(username string) (response.Code, *response.FetchUserResponse, error) {

	s.config.Logger().Info("FetchUserByUsername: finding user by username")
	user, err := s.store.FindUserByUsername(username)

	if err != nil {
		if errors.Is(err, voerrors.ErrNotFound) {
			s.config.Logger().Error("FetchUserByUsername: user not found", zap.Error(err))
			return response.NotFound, nil, err
		}

		s.config.Logger().Error("FetchUserByUsername: error finding user by username", zap.Error(err))
		return response.ServerError, nil, err

	}

	resp := s.mapUserToUserResp(user)

	return response.Success, resp, nil
}

func (s *userService) mapUserToUserResp(user *model.User) *response.FetchUserResponse {

	return &response.FetchUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
