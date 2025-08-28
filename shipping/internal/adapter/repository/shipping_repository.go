package repository

import (
	"errors"
	"sync"
)

type ShippingRepository struct {
	mu         sync.RWMutex
	deliveries map[string]int // Maps purchase ID to total quantity of units
}

func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{
		deliveries: make(map[string]int),
	}
}

func (r *ShippingRepository) SaveDelivery(purchaseID string, quantity int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	r.deliveries[purchaseID] = quantity
	return nil
}

func (r *ShippingRepository) GetDelivery(purchaseID string) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quantity, exists := r.deliveries[purchaseID]
	if !exists {
		return 0, errors.New("delivery not found")
	}

	return quantity, nil
}
