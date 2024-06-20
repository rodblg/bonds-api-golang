package bondApi

import "time"

type Bond struct {
	ID                       string
	Name                     string
	FaceValue                float64 //Original principal amount of the bond
	CurrentValue             float64 //
	Isin                     string  //International Securities Identification Number (ISIN)
	Issuer                   string  //Issuer of the bond
	InterestRate             float64
	InterestPaymentFrequency string
	MaturityDate             time.Time
	Description              string
}
