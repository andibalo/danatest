package streaming

import (
	"errors"
	"log"
	"net/http"
	"streaming/internal/config"
	"streaming/internal/handlers"
	"streaming/internal/service"
	"streaming/internal/storage"
	"streaming/internal/ws"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.AppConfig) *Server {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	store := storage.New(cfg)

	chatService := service.NewChatService(cfg, store)
	userService := service.NewUserService(cfg, store)

	chatHandler := handlers.NewChatHandler(chatService)
	userHandler := handlers.NewUserHandler(userService)

	hub := ws.NewHub()

	go hub.Run()

	e.GET("/ws", func(c echo.Context) error {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)

		if !errors.Is(err, nil) {
			log.Println(err)
		}

		defer func() {
			delete(hub.Clients, ws)
			ws.Close()
			log.Println("Closed")
		}()

		hub.Clients[ws] = true
		log.Println("Connected")

		read(hub, ws)
		return nil
	})

	registerHandlers(e, &handlers.HealthCheck{}, chatHandler, userHandler)

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(":" + addr)
}

type Handler interface {
	AddRoutes(e *echo.Echo)
}

func registerHandlers(e *echo.Echo, handlers ...Handler) {
	for _, handler := range handlers {
		handler.AddRoutes(e)
	}
}

func read(hub *ws.Hub, client *websocket.Conn) {

	for {
		var message ws.Message
		err := client.ReadJSON(&message)

		if !errors.Is(err, nil) {
			log.Printf("error occuered %v", err)
			delete(hub.Clients, client)
			break
		}

		log.Println(message)

		hub.Broadcast <- message
	}
}
