package storage

import (
	"errors"
	"streaming/internal/model"
	"streaming/internal/voerrors"

	"gorm.io/gorm"
)

func (s *Store) FindUserByUsername(username string) (*model.User, error) {

	user, err := s.userRepository.GetUserByUsername(username)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, voerrors.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *Store) FindUserByID(userID string) (*model.User, error) {

	user, err := s.userRepository.GetUserByID(userID)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, voerrors.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *Store) FindAllChats() (*[]model.Chat, error) {

	posts, err := s.chatRepository.GetAllChats()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Store) FindAllChatsByUserID(userID string) (*[]model.Chat, error) {

	posts, err := s.chatRepository.GetAllChatsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
