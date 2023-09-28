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

// @BasePath /api/v1

// GetItems godoc
// @Summary get all items
// @Schemes
// @Description return all items by id from array
// @Param ids body ids true "ids"
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {object} ResponseItems
// @Failure 500 {object} ResponseError
// @Failure 400 {object} ResponseError
// @Router /app/get-items [post]
func (h *Handler) GetItems(ctx *gin.Context) {
	var ids ids
	/*
	*Check input array of ID's
	 */
	if err := ctx.BindJSON(&ids); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Invalid items IDs")

		ctx.JSON(http.StatusBadRequest, ResponseError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid items IDs"})
	}

	repoItems, err := h.Repo.GetItems(ids.Ids) //get repository with items data
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Interanal error")
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": ResponseItems{Data: repoItems}}) //return response with data in json format
}
