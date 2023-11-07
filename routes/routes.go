package routes

import (
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/connection"
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/handler"
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/middleware"
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/model/publisher_human_detection"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherHumanDetectionRepository := publisher_human_detection.NewRepository(connection.RabbitMQ())

	// Initialize services
	publisherHumanDetectionService := publisher_human_detection.NewService(publisherHumanDetectionRepository)

	// Initialize handlers
	publisherHumanDetectionHandler := handler.NewPublisherHumanDetectionHandler(publisherHumanDetectionService)

	// Configure routes
	api := router.Group("/publisher/atmvideopack/v1")

	humanDetectionRoutes := api.Group("/humandetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureHumanDetectionRoutes(humanDetectionRoutes, publisherHumanDetectionHandler)

}

func configureHumanDetectionRoutes(api *gin.RouterGroup, handler *handler.PublisherHumanDetectionHandler) {
	api.POST("/create", handler.CreateQueueHumanDetection)
}
