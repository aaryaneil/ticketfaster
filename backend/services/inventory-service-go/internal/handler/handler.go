package handler

import (
	"inventory-service-go/internal/model"
	"inventory-service-go/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store store.Store
}

func NewHandler(s store.Store) *Handler {
	return &Handler{store: s}
}

// GetTicketByID handles requests to fetch a single ticket.
func (h *Handler) GetTicketByID(c *gin.Context) {
	id := c.Param("id")
	ticket, err := h.store.GetTicket(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// HoldTicket handles requests to place a temporary hold on a ticket.
func (h *Handler) HoldTicket(c *gin.Context) {
	id := c.Param("id")

	// In a real app, you'd check the original status before updating
	ticket, err := h.store.UpdateTicketStatus(id, model.Held)
	if err != nil {
		// Using 409 Conflict for business logic failure (e.g., ticket already sold)
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// PurchaseTicket handles requests to mark a ticket as sold.
func (h *Handler) PurchaseTicket(c *gin.Context) {
	id := c.Param("id")
	ticket, err := h.store.UpdateTicketStatus(id, model.Sold)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}
