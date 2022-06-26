package storage

import (
	"streaming/internal/config"
	"streaming/internal/dto"
	"streaming/internal/model"
	"streaming/internal/storage/repositories"
	"sync"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

type Store struct {
	logger         *zap.Logger
	userRepository UserRepository
	chatRepository ChatRepository
}

func New(cfg *config.AppConfig) *Store {
	db := InitDB(cfg)

	migrateDB(db)

	userRepo := repositories.NewUserRepository(db)
	chatRepo := repositories.NewChatRepository(db)

	return &Store{
		logger:         cfg.Logger(),
		userRepository: userRepo,
		chatRepository: chatRepo,
	}
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Chat{})
}

func InitDB(cfg *config.AppConfig) *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := cfg.PgConfig()

		db, err := gorm.Open(postgres.Open(databaseConfig.GetDBConnectionString()), &gorm.Config{})
		if err != nil {
			cfg.Logger().Fatal("Could not connect to database: %s", zap.Error(err))
		}

		cfg.Logger().Info("Successfully Connected to Database")

		instance = db
	})

	return instance
}

type Storage interface {
	CreateUser(in *dto.RegisterUser) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	FindUserByID(userID string) (*model.User, error)
	CreateChat(in *dto.CreateChat) (*model.Chat, error)
	FindAllChats() (*[]model.Chat, error)
	FindAllChatsByUserID(userID string) (*[]model.Chat, error)
	RemoveUserByUsername(username string) error
}

type UserRepository interface {
	SaveUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByID(userID string) (*model.User, error)
	DeleteUserByUsername(username string) error
}

type ChatRepository interface {
	SaveChat(chat *model.Chat) error
	GetAllChats() (*[]model.Chat, error)
	GetAllChatsByUserID(userID string) (*[]model.Chat, error)
}
