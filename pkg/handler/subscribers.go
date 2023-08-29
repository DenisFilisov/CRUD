package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary get All Followers By News ID
// @Security ApiKeyAuth
// @Tags subscribers
// @Description getAllFollowersByNewsID
// @ID getAllFollowersByNewsID
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "News ID"
// @Success 200 {array} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/subscribers/{id} [get]
func (h *Handler) getAllFollowersByNewsID(c *gin.Context) {
	newsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	subscribers, err := h.services.Subscribers.GetAllSubscribersByNewsID(newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subscribers)
}

type subscribeToNews struct {
	NewsId int `json:"newsId" binding:"required" db:"news_id"`
}

// @Summary subscribe To News
// @Security ApiKeyAuth
// @Tags subscribers
// @Description subscribeToNews
// @ID subscribeToNews
// @Accept  json
// @Produce  json
// @Param input body subscribeToNews true "News ID"
// @Success 200 {object} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/subscribers/ [post]
func (h *Handler) subscribeToNews(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	var input subscribeToNews
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	news, err := h.services.Subscribers.SubscribeToNews(userId.(int), input.NewsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("User %d subscribed to news %s", userId, news.Description))
}

// @Summary Unsubscribe From News
// @Security ApiKeyAuth
// @Tags subscribers
// @Description UnsubscribeFromNews
// @ID UnsubscribeFromNews
// @Accept  json
// @Produce  json
// @Param   id path int true "News ID"
// @Success 200 {object} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/subscribers/{id} [delete]
func (h *Handler) UnsubscribeFromNews(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	newsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UnsubscribeFromNews(userId, newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("User with id=%d unsubscribed from news=%d", userId, newsId))
}
