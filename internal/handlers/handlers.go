package handlers

import (
	"Hezzl_testing/internal/service/dto"
	"Hezzl_testing/pkg/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProjectGoodHandler struct {
	service HandlerInterface
}

type HandlerInterface interface {
	Create(projectId int, req dto.CreateGoodRequest) (*dto.CreateGoodResponse, error)
	Update(req dto.UpdateGoodRequest, id, projectId int) (*dto.UpdateGoodResponse, error)
	Delete(id, projectId int) (*dto.DeleteGoodResponse, error)
	Get(limit, offset int) (*dto.GetGoodsResponse, error)
	PatchPriority(req dto.ReprioritizeRequest, id, projectId int) (*dto.ReprioritizeResponse, *dto.UpdateGoodsResponse, error)
}

func NewHandler(service HandlerInterface) *ProjectGoodHandler {
	return &ProjectGoodHandler{service: service}
}

func (h *ProjectGoodHandler) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		getProjectId := c.Request.Form.Get("projectId")

		projectId, err := strconv.Atoi(getProjectId)
		if err != nil {
			log.Println(err)
			return
		}

		var req dto.CreateGoodRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println(err)
			return
		}

		good, err := h.service.Create(projectId, req)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, good)
	}
}

func (h *ProjectGoodHandler) UpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		getId := c.Request.Form.Get("id")
		getProjectId := c.Request.Form.Get("projectId")

		id, err := strconv.Atoi(getId)
		if err != nil {
			log.Println(err)
			return
		}

		projectId, err := strconv.Atoi(getProjectId)
		if err != nil {
			log.Println(err)
			return
		}

		var req dto.UpdateGoodRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println(err)
			return
		}

		good, err := h.service.Update(req, id, projectId)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, good)
	}
}

func (h *ProjectGoodHandler) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		getId := c.Request.Form.Get("id")
		getProjectId := c.Request.Form.Get("projectId")

		id, err := strconv.Atoi(getId)
		if err != nil {
			log.Println(err)
			return
		}

		projectId, err := strconv.Atoi(getProjectId)
		if err != nil {
			log.Println(err)
			return
		}

		deleted, err := h.service.Delete(id, projectId)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, deleted)
	}
}

func (h *ProjectGoodHandler) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		getLimit := c.Request.Form.Get("limit")
		getOffset := c.Request.Form.Get("offset")
		if getLimit == "" {
			getLimit = "10"
		}
		if getOffset == "" {
			getOffset = "1"
		}

		limit, err := strconv.Atoi(getLimit)
		if err != nil {
			log.Println(err)
			return
		}

		offset, err := strconv.Atoi(getOffset)
		if err != nil {
			log.Println(err)
			return
		}

		data, err := h.service.Get(limit, offset)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (h *ProjectGoodHandler) PriorityHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		getId := c.Request.Form.Get("id")
		getProjectId := c.Request.Form.Get("projectId")

		id, err := strconv.Atoi(getId)
		if err != nil {
			log.Println(err)
			return
		}

		projectId, err := strconv.Atoi(getProjectId)
		if err != nil {
			log.Println(err)
			return
		}

		var req dto.ReprioritizeRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println(err)
			return
		}

		priority, _, err := h.service.PatchPriority(req, id, projectId)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, priority)
	}
}
