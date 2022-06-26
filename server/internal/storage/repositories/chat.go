package repositories

import (
	"streaming/internal/model"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {

	return &ChatRepository{
		db: db,
	}
}

func (p *ChatRepository) SaveChat(chat *model.Chat) error {

	err := p.db.Create(chat).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *ChatRepository) GetAllChats() (*[]model.Chat, error) {

	var chats *[]model.Chat

	err := p.db.Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (p *ChatRepository) GetAllChatsByUserID(userID string) (*[]model.Chat, error) {

	var chats *[]model.Chat

	err := p.db.Where("user_id = ?", userID).Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}
