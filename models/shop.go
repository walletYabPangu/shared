package models

type Skin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SupplyTotal int64  `json:"supply_total"`
	SupplySold  int64  `json:"supply_sold"`
	PriceTon    int64  `json:"price_ton"` // nanoTON
	MediaURL    string `json:"media_url"`
	IsActive    bool   `json:"is_active"`
}
