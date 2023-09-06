package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// setting type for getting IDs
type ids struct {
	Ids []int `json:"ids"`
}

/*
*Get all items by ID's array
 */
func (h *Handler) GetItems(ctx *gin.Context) {
	var ids ids
	/*
	*Check input array of ID's
	 */
	if err := ctx.BindJSON(&ids); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Invalid items IDs")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid items IDs"})
	}

	repoItems, err := h.Repo.GetItems(ids.Ids) //get repository with items data
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Interanal error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": repoItems}) //return response with data in json format
}
