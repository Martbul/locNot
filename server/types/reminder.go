package types

type Remainder struct {
	UsedId       int    `json:"userId"`
	RemindingMsg string `json:"remindongMsg"`
	geoLocation
	validTo
}

func (r *Remainder) ValidateRemainder()
