package handler

import (
	"CRUD/pkg/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DenisFilisov/cacheLib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	time "time"
)

const (
	Cookie = "Cookie"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200
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
		if errors.Is(err, sql.ErrNoRows) {
			newErrorResponse(g, http.StatusBadRequest, "Username is taken")
			return
		}
		newErrorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type singIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type token struct {
	Token string `json:"token"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body singIn true "credentials"
// @Success 200 {object} token
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

	user, err := h.services.Authorisation.FindUserByUsernameAndPswd(input.Username, input.Password)
	if err != nil {
		newErrorResponse(g, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.Authorisation.GenerateTokens("", user)
	if err != nil {
		newErrorResponse(g, http.StatusInternalServerError, err.Error())
		return
	}
	userId, expTime, err := h.services.Authorisation.ParseToken(accessToken)

	duration := time.Unix(expTime, 0).Sub(time.Now())
	cacheLib.Set(fmt.Sprint(userId), duration, accessToken)

	var token token
	token.Token = accessToken
	g.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	g.JSON(http.StatusOK, token)
}

// @Summary refreshToken
// @Tags auth
// @Description refreshToken
// @RefreshToken string
// @Accept  json
// @Produce  json
// @Success 200 {object} token
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh-token [post]
func (h *Handler) refreshToken(g *gin.Context) {
	refreshToken := g.GetHeader(Cookie)

	if refreshToken == "" {
		newErrorResponse(g, http.StatusBadRequest, "Wrong refreshToken in Cookies")
		return
	}
	t := strings.Split(refreshToken, "=")
	if len(t) != 2 || t[0] != "refresh-token" {
		newErrorResponse(g, http.StatusBadRequest, "Wrong refreshToken in Cookies")
		return
	}

	accessToken, refreshToken, err := h.services.RefreshToken(t[1])
	if err != nil {
		newErrorResponse(g, http.StatusUnauthorized, "can't refresh token ")
		return
	}

	var token token
	token.Token = accessToken
	g.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	g.JSON(http.StatusOK, token)
}
