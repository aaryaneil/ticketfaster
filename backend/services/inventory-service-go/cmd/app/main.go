package main

import (
	"inventory-service-go/internal/handler"
	"inventory-service-go/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	inMemoryStore := store.NewInMemoryStore()
	apiHandler := handler.NewHandler(inMemoryStore)

	// Setup HTTP Router
	router := gin.Default()

	// Define API routes
	inventoryRoutes := router.Group("/inventory")
	{
		inventoryRoutes.GET("/tickets/:id", apiHandler.GetTicketByID)
		inventoryRoutes.POST("/tickets/:id/hold", apiHandler.HoldTicket)
		inventoryRoutes.POST("/tickets/:id/purchase", apiHandler.PurchaseTicket)
	}

	// Start the server
	router.Run(":3001")
}
