package storage

import (
	"streaming/internal/dto"
	"streaming/internal/model"
)

func (s *Store) CreateUser(in *dto.RegisterUser) (*model.User, error) {

	user := &model.User{
		Username: in.Username,
	}

	err := s.userRepository.SaveUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateChat(in *dto.CreateChat) (*model.Chat, error) {

	chat := &model.Chat{
		Message: in.Message,
		UserID:  in.UserID,
	}

	err := s.chatRepository.SaveChat(chat)

	if err != nil {
		return nil, err
	}

	return chat, nil
}
