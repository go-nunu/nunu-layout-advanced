package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/internal/pkg/request"
	"github.com/go-nunu/nunu-layout-advanced/internal/pkg/response"
	"github.com/go-nunu/nunu-layout-advanced/internal/service"
	"net/http"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

type userHandler struct {
	*Handler
	userService service.UserService
}

// Register godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (h *userHandler) Register(ctx *gin.Context) {
	req := new(request.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.HandleError(ctx, http.StatusBadRequest, response.ErrBadRequest, nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		response.HandleError(ctx, http.StatusBadRequest, response.ErrBadRequest, nil)
		return
	}

	response.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, http.StatusBadRequest, response.ErrBadRequest, nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		response.HandleError(ctx, http.StatusUnauthorized, response.ErrUnauthorized, nil)
		return
	}

	response.HandleSuccess(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		response.HandleError(ctx, http.StatusUnauthorized, response.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		response.HandleError(ctx, http.StatusBadRequest, response.ErrBadRequest, nil)
		return
	}

	response.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req request.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, http.StatusBadRequest, response.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		response.HandleError(ctx, http.StatusInternalServerError, response.ErrInternalServerError, nil)
		return
	}

	response.HandleSuccess(ctx, nil)
}
