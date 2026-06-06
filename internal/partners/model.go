package partners

import "time"

// Partner represents a person or organization that the system interacts with.
type Partner struct {
	ID        string    `json:"id" db:"id"` // UUID as string
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Roles     []string  `json:"roles" db:"roles"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// TableName returns the DB table name for Partner (optional helper).
func (Partner) TableName() string { return "partners" }
