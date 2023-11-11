package handler

import (
	"Booksiary/user-service/internal/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var input domain.UserRegistrationData
	var err error

	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.Registration.SignUp(input)
	if err != nil {
		NewErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (h *Handler) createUserCallback(ctx *gin.Context) {
	code, err := strconv.Atoi(ctx.Query("code"))
	if err != nil {
		NewErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		NewErrResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userUUID, err := h.service.SignUpCallback(code)
	if err != nil {
		NewErrResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "зарегестрирован",
		"uuid":   userUUID,
	})
}

func (h *Handler) createUserWithExternal(ctx *gin.Context) {
	var input string
	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	log.Print(input)
}

func getCode(ctx *gin.Context, keyCode string) (int, error) {
	code, err := strconv.Atoi(ctx.Param(keyCode))
	if err != nil {
		return 0, err
	}

	return code, nil
}
