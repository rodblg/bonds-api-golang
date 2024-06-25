package bondApi

import (
	"net/http"
	"time"
)

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

type UserResponse struct {
	*User
}

func NewUserResponse(user *User) *UserResponse {
	resp := &UserResponse{
		User: user,
	}

	return resp
}

func (rd *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
