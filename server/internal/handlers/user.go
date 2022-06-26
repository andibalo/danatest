package handlers

import (
	"errors"
	"net/http"
	"streaming/internal/constants"
	"streaming/internal/request"
	"streaming/internal/response"
	"streaming/internal/service"
	"streaming/internal/voerrors"

	"github.com/labstack/echo/v4"
)

type User struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *User {

	return &User{
		userService: userService,
	}
}

func (h *User) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.UserAPIPath)

	r.POST("/", h.register)
	r.DELETE("/", h.removeUser)
}

func (h *User) register(c echo.Context) error {
	registerUserReq := &request.RegisterUserRequest{}

	if err := c.Bind(registerUserReq); err != nil {
		return err
	}

	err := registerUserReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedUserResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, respData, err := h.userService.FetchUserByUsername(registerUserReq.Username)

	if err != nil {
		if !errors.Is(err, voerrors.ErrNotFound) {
			return h.failedUserResponse(c, code, err, "")
		}
	}

	if respData != nil {
		return h.failedUserResponse(c, response.DuplicateUser, voerrors.ErrDuplicateUser, "")
	}

	code, err = h.userService.CreateUser(registerUserReq)

	if err != nil {
		return h.failedUserResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, nil)

	resp.SetResponseMessage("Successfully created user")

	return c.JSON(http.StatusOK, resp)
}

func (h *User) removeUser(c echo.Context) error {
	registerUserReq := &request.RegisterUserRequest{}

	if err := c.Bind(registerUserReq); err != nil {
		return err
	}

	err := registerUserReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedUserResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, err := h.userService.DeleteUserByUsername(registerUserReq.Username)

	if err != nil {
		return h.failedUserResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, nil)

	resp.SetResponseMessage("Successfully deleted user")

	return c.JSON(http.StatusOK, resp)
}

func (h *User) failedUserResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
