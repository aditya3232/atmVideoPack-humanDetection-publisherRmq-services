package handler

import (
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/constant"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/aditya3232/gatewatchApp-services.git/model/publisher_human_detection"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PublisherHumanDetectionHandler struct {
	publisherHumanDetectionService publisher_human_detection.Service
}

func NewPublisherHumanDetectionHandler(publisherHumanDetectionService publisher_human_detection.Service) *PublisherHumanDetectionHandler {
	return &PublisherHumanDetectionHandler{publisherHumanDetectionService}
}

// public message to rmqg
func (h *PublisherHumanDetectionHandler) CreateQueueHumanDetection(c *gin.Context) {
	var input publisher_human_detection.HumanDetectionInput

	err := c.ShouldBindWith(&input, binding.FormMultipart)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	humanDetection, err := h.publisherHumanDetectionService.CreateQueueHumanDetection(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, publisher_human_detection.FormatPublisherHumanDetection(humanDetection))
	log.Info("Queue human detection berhasil dibuat")
	c.JSON(http.StatusOK, response)
}
