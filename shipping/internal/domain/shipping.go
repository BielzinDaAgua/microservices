package domain

type Shipping struct {
	DeliveryItems []DeliveryItem
	PurchaseID    string
}

type DeliveryItem struct {
	ItemID   string
	Quantity int
}

func (s *Shipping) CalculateDeliveryTime() int {
	totalUnits := 0
	for _, item := range s.DeliveryItems {
		totalUnits += item.Quantity
	}

	// Example logic for delivery time based on total units
	if totalUnits <= 5 {
		return 1 // 1 day
	} else if totalUnits <= 10 {
		return 2 // 2 days
	} else {
		return 3 // 3 days
	}
}
