package tsys

import (
	"fmt"
)

type TripLegData struct {
	CarrierCodeAbbreviation    string           //2 A/N (for example, UA, AA, and TW) 4.36
	ServiceClass               string           //1 or 2 A/N 4.175
	StopOverCode               StopOverCodeType //1 A/N 4.189
	DestinationCityAirportCode AirportCodeType  //3 A/N
}

func (t *TripLegData) String() string {

	return fmt.Sprintf("%2s%2s%c%3s",
		t.CarrierCodeAbbreviation,    //2 A/N
		t.ServiceClass,               //1 or 2 A/N
		t.StopOverCode,               //1 A/N
		t.DestinationCityAirportCode) //3 A/N

}

func ParseTripLegData(s string) *TripLegData {

	t := TripLegData{}

	t.CarrierCodeAbbreviation = s[0:2]                    // 2 A/N
	t.ServiceClass = s[2:4]                               // 1 or 2 A/N
	t.StopOverCode = StopOverCodeType(s[4])               // 1 A/N
	t.DestinationCityAirportCode = AirportCodeType(s[5:]) // 3 A/N

	return &t
}
