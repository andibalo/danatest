package storage

import (
	"errors"
	"streaming/internal/voerrors"

	"gorm.io/gorm"
)

func (s *Store) RemoveUserByUsername(username string) error {

	err := s.userRepository.DeleteUserByUsername(username)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return voerrors.ErrNotFound
		}

		return err
	}

	return nil
}
