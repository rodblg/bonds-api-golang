package bondApi

import (
	"net/http"
	"time"
)

type Bond struct {
	ID                       string    `json:"id,omitempty"`
	Name                     string    `json:"name"`
	FaceValue                float64   `json:"face_value"`    //Original principal amount of the bond
	CurrentValue             float64   `json:"current_value"` //
	Isin                     string    `json:"isin"`          //International Securities Identification Number (ISIN)
	Issuer                   string    `json:"issuer"`        //Issuer of the bond
	InterestRate             float64   `json:"interest_rate"`
	InterestPaymentFrequency string    `json:"interest_payment_frequency"`
	MaturityDate             time.Time `json:"maturity_date"`
	Description              string    `json:"description"`
	Buyer                    string    `json:"buyer"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type BondResponse struct {
	*Bond
}

func NewBondResponse(bond *Bond) *BondResponse {
	resp := &BondResponse{
		Bond: bond,
	}

	return resp
}

func (rd *BondResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
