package application

import (
	"errors"
	"time"

	"shipping/internal/domain"
	"shipping/internal/port"
)

type ShippingService struct {
	repo port.ShippingRepository
}

func NewShippingService(repo port.ShippingRepository) *ShippingService {
	return &ShippingService{repo: repo}
}

func (s *ShippingService) CalculateDeliveryTime(deliveryItems []domain.DeliveryItem, purchaseID string) (time.Duration, error) {
	if len(deliveryItems) == 0 {
		return 0, errors.New("no delivery items provided")
	}

	totalUnits := 0
	for _, item := range deliveryItems {
		totalUnits += item.Quantity
	}

	// Example logic for delivery time calculation based on total units
	var deliveryTime time.Duration
	if totalUnits <= 5 {
		deliveryTime = 24 * time.Hour // 1 day
	} else if totalUnits <= 20 {
		deliveryTime = 48 * time.Hour // 2 days
	} else {
		deliveryTime = 72 * time.Hour // 3 days
	}

	// Here you would typically update the order status or notify the Order microservice
	// after successful payment, which is handled elsewhere in the application.

	return deliveryTime, nil
}
