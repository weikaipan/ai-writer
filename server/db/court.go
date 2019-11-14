package court

import (
)

type CourtInfo struct {
	Lat float32 `json:"Lat"`
	Lon	float32 `json:"Lon"`
}

func RetrieveCourtInfo() (CourtInfo, error) {
	return CourtInfo{
		Lat: 35.33,
		Lon: 37.22,
	},
	nil
}