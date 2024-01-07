package handler

import (
	"central_library/core/domain"
	"central_library/core/repo"
	"central_library/persistence"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userRepo repo.UserRepo
}

func NewUserHandler(repo repo.UserRepo) *UserHandler {
	return &UserHandler{userRepo: repo}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user domain.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := h.userRepo.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, *createdUser)
}

func (h *UserHandler) RecordBookIssue(ctx *gin.Context) {
	userIDstr := ctx.Param("id")

	userIDint, err := strconv.ParseUint(userIDstr, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userRepo.RecordBookIssue(uint(userIDint))
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
func (h *UserHandler) RecordBookReturn(ctx *gin.Context) {
	userIDstr := ctx.Param("id")

	userIDint, err := strconv.ParseUint(userIDstr, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.userRepo.RecordBookReturn(uint(userIDint))
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {

	users, err := h.userRepo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, users)
}
