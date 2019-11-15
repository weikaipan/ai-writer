package court

import (
	"strconv"
)

type CourtInfo struct {
	Lat float32 `json:"Lat"`
	Lon	float32 `json:"Lon"`
	name string `json:"name"`
}

type Response struct  {
	message string `json:"message"`
}

func RetrieveCourtInfo(name string) (CourtInfo, error) {
	return CourtInfo{
		Lat: 35.33,
		Lon: 37.22,
		name: name,
	},
	nil
}

func SpamCourtById(id int) (Response, error){
	return Response{
		message: strconv.FormatInt(int64(id), 10),
	},
	nil
}

func SaveNewCourt() (Response, error) {
	return Response{
		message: "success",
	},
	nil
}
