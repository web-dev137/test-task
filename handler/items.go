package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ids struct {
	Ids []int `json:"ids"`
}

func (h *Handler) GetItems(ctx *gin.Context) {
	var ids ids
	if err := ctx.BindJSON(&ids); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Invalid items IDs")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid items IDs"})
	}
	repoItems, err := h.Repo.GetItems(ids.Ids)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Interanal error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": repoItems})
}
