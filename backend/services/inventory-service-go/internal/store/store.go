package store

import (
	"fmt"
	"inventory-service-go/internal/model"
	"sync"
)

// Store defines the interface for data storage operations.
type Store interface {
	GetTicket(id string) (*model.Ticket, error)
	UpdateTicketStatus(id string, newStatus model.TicketStatus) (*model.Ticket, error)
}

// InMemoryStore is an in-memory implementation of the Store interface.
type InMemoryStore struct {
	tickets map[string]*model.Ticket
	mu      sync.RWMutex
}

// NewInMemoryStore creates and initializes an in-memory store with dummy data.
func NewInMemoryStore() *InMemoryStore {
	tickets := make(map[string]*model.Ticket)
	tickets["1"] = &model.Ticket{ID: "1", EventID: "101", Seat: "A1", Status: model.Available}
	tickets["2"] = &model.Ticket{ID: "2", EventID: "101", Seat: "A2", Status: model.Available}
	tickets["3"] = &model.Ticket{ID: "3", EventID: "102", Seat: "B1", Status: model.Sold}
	return &InMemoryStore{
		tickets: tickets,
	}
}

// GetTicket retrieves a ticket by its ID.
func (s *InMemoryStore) GetTicket(id string) (*model.Ticket, error) {
	s.mu.RLock() // Acquire a read lock
	defer s.mu.RUnlock()

	ticket, exists := s.tickets[id]
	if !exists {
		return nil, fmt.Errorf("ticket with ID %s not found", id)
	}
	return ticket, nil
}

// UpdateTicketStatus updates the status of a specific ticket.
func (s *InMemoryStore) UpdateTicketStatus(id string, newStatus model.TicketStatus) (*model.Ticket, error) {
	s.mu.Lock() // Acquire a write lock
	defer s.mu.Unlock()

	ticket, exists := s.tickets[id]
	if !exists {
		return nil, fmt.Errorf("ticket with ID %s not found", id)
	}

	// Basic state transition validation
	if ticket.Status == model.Sold {
		return nil, fmt.Errorf("cannot change status of a sold ticket")
	}

	ticket.Status = newStatus
	return ticket, nil
}
