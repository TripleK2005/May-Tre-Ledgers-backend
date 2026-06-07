package items

import "time"

// ItemType constants
const (
	ItemTypeMaterial = "MATERIAL"
	ItemTypeProduct  = "PRODUCT"
)

// Item represents a material or finished good stored in inventory.
type Item struct {
	ID        string    `json:"id" db:"id"`     // UUID as string
	Type      string    `json:"type" db:"type"` // MATERIAL or PRODUCT
	Name      string    `json:"name" db:"name"`
	Unit      string    `json:"unit" db:"unit"`
	Size      *string   `json:"size,omitempty" db:"size"` // nullable
	Price     int64     `json:"price" db:"price"`         // reference price
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (Item) TableName() string { return "items" }
