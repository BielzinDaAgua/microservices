package domain

type Delivery struct {
	Items      []DeliveryItem
	PurchaseID string
}

type DeliveryItem struct {
	Name     string
	Quantity int
}

func (d *Delivery) CalculateDeliveryTime() int {
	totalUnits := 0
	for _, item := range d.Items {
		totalUnits += item.Quantity
	}

	// Example logic for delivery time calculation
	if totalUnits <= 5 {
		return 1 // 1 day for 5 or fewer items
	} else if totalUnits <= 10 {
		return 2 // 2 days for 6 to 10 items
	} else {
		return 3 // 3 days for more than 10 items
	}
}
