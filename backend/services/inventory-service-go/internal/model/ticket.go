package model

// TicketStatus represents the possible states of a ticket.
type TicketStatus string

const (
	Available TicketStatus = "AVAILABLE"
	Held      TicketStatus = "HELD"
	Sold      TicketStatus = "SOLD"
)

// Ticket represents a single ticket in the inventory.
type Ticket struct {
	ID      string       `json:"id"`
	EventID string       `json:"eventId"`
	Seat    string       `json:"seat"`
	Status  TicketStatus `json:"status"`
}
