package events

type ProductPriceChangedEvent struct {
	ProductID int
	NewPrice  float64
}
