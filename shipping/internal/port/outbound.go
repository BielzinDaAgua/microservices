package port

type ShippingService interface {
	CalculateDeliveryTime(purchaseID string, items []DeliveryItem) (int, error)
}

type DeliveryItem struct {
	ItemID   string
	Quantity int
}
