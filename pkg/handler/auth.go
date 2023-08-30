package handler

import (
	"CRUD/pkg/model"
	"github.com/DenisFilisov/cacheLib"
	"github.com/gin-gonic/gin"
	"net/http"
	time "time"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) singUp(g *gin.Context) {
	var input model.User

	if err := g.BindJSON(&input); err != nil {
		newErrorResponse(g, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorisation.CreateUser(input)
	if err != nil {
		newErrorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type singIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body singIn true "credentials"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) singIn(g *gin.Context) {
	var input singIn

	if err := g.BindJSON(&input); err != nil {
		newErrorResponse(g, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorisation.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}
	userId, expTime, err := h.services.Authorisation.ParseToken(token)

	duration := time.Unix(expTime, 0).Sub(time.Now())
	cacheLib.Set(string(userId), duration, token)

	g.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
