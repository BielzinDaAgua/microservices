package port

type ShippingRequest struct {
	DeliveryItems []DeliveryItem
	PurchaseID    string
}

type DeliveryItem struct {
	ItemID   string
	Quantity int32
}

type ShippingResponse struct {
	DeliveryTime string
}

type ShippingService interface {
	CalculateDeliveryTime(request ShippingRequest) (ShippingResponse, error)
}
