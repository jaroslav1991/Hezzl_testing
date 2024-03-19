package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type badResponse struct {
	Code    uint     `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func WriteErrorResponse(c *gin.Context) {
	response := badResponse{
		Code:    3,
		Message: "errors.good.NotFound",
		Details: []string{},
	}
	c.JSON(http.StatusNotFound, response)
}
