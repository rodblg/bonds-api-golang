package bondApi

import "time"

type User struct {
	ID             string    `json:"id,omitempty"`
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	PurchasedBonds []Bond    `json:"purchased_bonds,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
