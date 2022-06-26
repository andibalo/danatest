package handlers

import (
	"net/http"
	"streaming/internal/constants"
	"streaming/internal/request"
	"streaming/internal/response"
	"streaming/internal/service"
	"streaming/internal/voerrors"

	"github.com/labstack/echo/v4"
)

type Chat struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) *Chat {
	return &Chat{
		chatService: chatService,
	}
}

func (h *Chat) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.ChatAPIPath)

	r.POST("/", h.createChat)
	r.GET("/", h.getAllChats)
	r.GET("/:id", h.getAllPostsByUserID)
}

func (h *Chat) createChat(c echo.Context) error {
	createChatReq := &request.CreateChatRequest{}

	if err := c.Bind(createChatReq); err != nil {
		return err
	}

	err := createChatReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedChatResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, err := h.chatService.CreateChat(createChatReq, "USER")

	if err != nil {
		return h.failedChatResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, nil)

	return c.JSON(http.StatusOK, resp)
}

func (h *Chat) getAllChats(c echo.Context) error {

	code, chats, err := h.chatService.FetchAllChats()

	if err != nil {
		return h.failedChatResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, chats)

	return c.JSON(http.StatusOK, resp)
}

func (h *Chat) getAllPostsByUserID(c echo.Context) error {

	userId := c.Param("id")

	code, chats, err := h.chatService.FetchAllChatsByUserID(userId)

	if err != nil {
		return h.failedChatResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, chats)

	return c.JSON(http.StatusOK, resp)
}

func (h *Chat) failedChatResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
	if code == "" {
		code = voerrors.MapErrorsToCode(err)
	}

	resp := response.Wrapper{
		ResponseCode: code,
		Status:       code.GetStatus(),
		Message:      code.GetMessage(),
	}

	if errorMsg != "" {
		resp.SetResponseMessage(errorMsg)
	}

	return c.JSON(voerrors.MapErrorsToStatusCode(err), resp)
}
