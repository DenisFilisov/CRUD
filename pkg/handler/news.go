package handler

import (
	"CRUD/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get All News
// @Security ApiKeyAuth
// @Tags news
// @Description GetAllNews
// @ID GetAllNews
// @Accept  json
// @Produce  json
// @Success 200 {array} model.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news [get]
func (h *Handler) getAllNews(c *gin.Context) {
	news, err := h.services.News.GetAllNews()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't find news")
		return
	}

	c.JSON(http.StatusOK, news)
}

// @Summary get News By Id
// @Security ApiKeyAuth
// @Tags news
// @Description getNewsById
// @ID getNewsById
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "News ID"
// @Success 200 {object} model.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/{id} [get]
func (h *Handler) getNewsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	news, err := h.services.News.FindNewsById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}

type postNewNews struct {
	Description string `json:"description" example:"new description"`
}

// @Summary PostNews
// @Security ApiKeyAuth
// @Tags news
// @Description postNewNews
// @ID postNewNews
// @Accept  json
// @Produce  json
// @Param input body postNewNews true "New News"
// @Success 200 {object} model.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/ [post]
func (h *Handler) postNewNews(c *gin.Context) {

	var news model.News
	if err := c.BindJSON(&news); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "wrong attributes in request")
		return
	}

	id, err := h.services.News.PostNews(news)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "can't post news")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type changeNewsById struct {
	Description string `json:"description" example:"new description"`
}

// @Summary change news with ID
// @Security ApiKeyAuth
// @Tags news
// @Description changeNewsById
// @ID changeNewsById
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "News ID"
// @Param input body changeNewsById true "New Description"
// @Success 200 {object} model.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/{id} [put]
func (h *Handler) changeNewsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var changedNews changeNewsById
	if err := c.BindJSON(&changedNews); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	news, err := h.services.UpdateNews(id, changedNews.Description)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, news)
}

// @Summary remove News
// @Security ApiKeyAuth
// @Tags news
// @Description removeNews
// @ID removeNews
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "News ID"
// @Success 200 {object} model.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/{id} [delete]
func (h *Handler) removeNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.RemoveNews(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("News with id=%d deleted", id))
}
