package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

// The MarketSpecificDataField sub-fields contain market specific data for hotel and auto rental transactions
// The format of this four-character field is: “PIDD.”
type MarketSpecificDataField struct {
	PrestigiousIndicator         PrestigiousPropertyType //P 1 A/N
	MarketSpecificDataIdentifier MarketSpecificDataType  //I 1 A/N
	StayOrRentalDuration         int                     //DD 3-4 NUM
}

func NewMarketSpecificDataField(prestigiousind PrestigiousPropertyType, marketid MarketSpecificDataType, rentalduration int) *MarketSpecificDataField {
	return &MarketSpecificDataField{prestigiousind, marketid, rentalduration}
}

func (m *MarketSpecificDataField) Copy(src *MarketSpecificDataField) {
	m.PrestigiousIndicator = src.PrestigiousIndicator
	m.MarketSpecificDataIdentifier = src.MarketSpecificDataIdentifier
	m.StayOrRentalDuration = src.StayOrRentalDuration
}

func ParseMarketSpecificDataField(s string) *MarketSpecificDataField {

	m := MarketSpecificDataField{AutoRentalOrNonParticipatingProperty, OtherIndustries, 0}

	if 4 > len(s) {
		return &m
		//panic("MarketSpecificData must be at least 4 characters in length")
	}

	fmt.Sscanf(s, "%c%c%02d", &m.PrestigiousIndicator, &m.MarketSpecificDataIdentifier, &m.StayOrRentalDuration)

	return &m
}

func (m *MarketSpecificDataField) String() string {
	if 0 >= m.StayOrRentalDuration {
		return ""
	}
	return fmt.Sprintf("%c%c%02d", m.PrestigiousIndicator, m.MarketSpecificDataIdentifier, m.StayOrRentalDuration)
}

func (m *MarketSpecificDataField) VerboseString() string {

	if 0 >= m.StayOrRentalDuration {
		return ""
	}

	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "PrestigiousIndicator", m.PrestigiousIndicator)
	buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "MarketSpecificDataIdentifier", m.MarketSpecificDataIdentifier)
	buffer.AppendFormat("%8s%-32s%02d (DD - Days)\n", " ", "StayOrRentalDuration", m.StayOrRentalDuration)

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}
