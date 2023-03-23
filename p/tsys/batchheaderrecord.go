package tsys

import (
	"fmt"
	"time"

	"code.somebank.com/p/bytes"
)

// all code based upon TSYS technical reference data found within:
// EXTERNAL INTERFACE SPECIFICATIONS DATA
// CAPTURE RECORD FORMATS
// EIS 1081
// VERSION 13.2
// APRIL 21, 2015

type BatchHeaderRecord struct {
	RecordFormat          RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType       ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter      bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId             rune                     // 4 1 A/N X.25 Routing ID Z 4.341
	RecordType            *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243
	AcquirerBIN           uint64                   //10-15 6 NUM Acquirer BIN 4.2 (999995)
	AgentBankNumber       uint64                   //16-21 6 NUM Agent Bank Number 111111 4.22
	AgentChainNumber      uint64                   //22-27 6 NUM Agent Chain Number 000000 4.23
	MerchantNumber        uint64                   //28-39 12 NUM Merchant Number 4.189 (999999999911)
	StoreNumber           uint                     //40-43 4 NUM Store Number 4.284 (1515)
	TerminalNumber        uint                     //44-47 4 NUM Terminal Number 4.300 (5999)
	DeviceCode            DeviceCodeType           //48 1 A/N Device Code C, D, E, I, M,Q 4.100 (ReservedForThirdPartyDevelopers)
	IndustryCode          IndustryCodeType         //49 1 A/N Industry Code D, F, G, R 4.145 (UnknownOrUnsure)
	CurrencyCode          CurrencyCodeType         //50-52 3 NUM Currency Code 4.85 (UsdUnitedStatesUsDollar)
	LanguageIndicator     LanguageIndicatorType    //53-54 2 NUM Language Indicator 00 4.163 (English)
	TimeZoneDiff          *TimeZoneDifferential    //55-57 3 NUM Time Zone Differential 4.305 (705 EASTERN)
	BatchTransmissionDate string                   //58-61 4 NUM Batch Transmission Date MMDD 4.45
	BatchNumber           uint                     //62-64 3 NUM Batch Number 001 - 999 4.41 (001 - 999)
	BlockingIndicator     BlockingIndicatorType    //65 1 NUM Blocking Indicator 0 - Not Blocked 2 - Blocked

	// Table 3.2 Transaction header record - optional data groups
	// 0 Header Record

	// 1 Passenger Travel
	TravelAgencyCode string //66-73 8 A/N Travel Agency Code Left-justified/space-filled 4.318
	TravelAgencyName string //74-98 25 A/N Travel Agency Name Left-justified/space-filled 4.322

	// 2 Hotel, Auto Rental, and Card not Present
	MerchantLocalPhone     string //66-76 11 A/N Merchant Local Phone Num
	CardholderServicePhone string //999-9999999 4.187 77-87 11 A/N Cardholder Svc Phone Number 999-9999999 4.60

	// 4 Gen2 Authentication Genkey (Only for Gen2 Authentication participants) Gen 2 terminal authentication
	Genkey string //66-89 24 ASCII Genkey ASCII Representation of HEX 4.135

	// 5 Developer ID
	DeveloperId string //6 A/N 4.99 Developer ID
	VersionId   string //4 A/N 4.337 Version ID

	// 6 Gateway ID
	GatewayId string //10 A/N 4.134 Gateway ID
}

func NewBatchHeaderRecord(acquirerid uint64, agentbankno uint64, agentchainno uint64) *BatchHeaderRecord {
	h := BatchHeaderRecord{}

	h.RecordFormat = Settlement
	h.ApplicationType = SingleBatchPerConnection
	h.MessageDelimiter = bytes.MessageDelimiter
	h.RoutingId = 'Z'
	h.RecordType = NewOptionalDataGroupMap(HeaderRecord)

	h.AcquirerBIN = acquirerid
	h.AgentBankNumber = agentbankno
	h.AgentChainNumber = agentchainno

	dt := time.Now()
	h.BatchTransmissionDate = fmt.Sprintf("%02d%02d", dt.Month(), dt.Year())
	h.BlockingIndicator = NotBlocked

	h.RecordType.SetGroupBit(GroupMapBitType(5), true)
	h.DeveloperId = TSYS_DEVELOPER_ID
	h.VersionId = TSYS_DEVELOPER_VERSION

	return &h
}

func (h *BatchHeaderRecord) SetTerminalIdentification(merchantno uint64, storeno uint, terminalno uint, device DeviceCodeType, industry IndustryCodeType) {
	h.MerchantNumber = merchantno //28-39 12 NUM Merchant Number 4.189 (999999999911)
	h.StoreNumber = storeno       //40-43 4 NUM Store Number 4.284 (1515)
	h.TerminalNumber = terminalno //44-47 4 NUM Terminal Number 4.300 (5999)
	h.DeviceCode = device         //48 1 A/N Device Code C, D, E, I, M,Q 4.100 (ReservedForThirdPartyDevelopers)
	h.IndustryCode = industry     //49 1 A/N Industry Code D, F, G, R 4.145 eg AllTypesExceptHotelAndAuto
}

func (h *BatchHeaderRecord) SetTerminalLocale(currency CurrencyCodeType, language LanguageIndicatorType, timezone *TimeZoneDifferential) {

	h.CurrencyCode = currency      //50-52 3 NUM Currency Code 4.85 (UsdUnitedStatesUsDollar)
	h.LanguageIndicator = language //53-54 2 NUM Language Indicator 00 4.163 (English)
	h.TimeZoneDiff = timezone      //55-57 3 NUM Time Zone Differential 4.305 (705 EASTERN)
}

func (h *BatchHeaderRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", h.RecordFormat)          //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%d", h.ApplicationType)       //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", h.MessageDelimiter)      //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", h.RoutingId)             // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", h.RecordType)            //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%06d", h.AcquirerBIN)         //10-15 6 NUM Acquirer BIN 4.2 (999995)
	buffer.AppendFormat("%06d", h.AgentBankNumber)     //16-21 6 NUM Agent Bank Number 111111 4.22
	buffer.AppendFormat("%06d", h.AgentChainNumber)    //22-27 6 NUM Agent Chain Number 000000 4.23
	buffer.AppendFormat("%012d", h.MerchantNumber)     //28-39 12 NUM Merchant Number 4.189 (999999999911)
	buffer.AppendFormat("%04d", h.StoreNumber)         //40-43 4 NUM Store Number 4.284 (1515)
	buffer.AppendFormat("%04d", h.TerminalNumber)      //44-47 4 NUM Terminal Number 4.300 (5999)
	buffer.AppendFormat("%c", h.DeviceCode)            //48 1 A/N Device Code C, D, E, I, M,Q 4.100 (ReservedForThirdPartyDevelopers)
	buffer.AppendFormat("%c", h.IndustryCode)          //49 1 A/N Industry Code D, F, G, R 4.145 (UnknownOrUnsure)
	buffer.AppendFormat("%03d", h.CurrencyCode)        //50-52 3 NUM Currency Code 4.85 (UsdUnitedStatesUsDollar)
	buffer.AppendFormat("%02d", h.LanguageIndicator)   //53-54 2 NUM Language Indicator 00 4.163 (English)
	buffer.AppendFormat("%03s", h.TimeZoneDiff)        //55-57 3 NUM Time Zone Differential 4.305 (705 EASTERN)
	buffer.AppendFormat("%s", h.BatchTransmissionDate) //58-61 4 NUM Batch Transmission Date MMDD 4.45
	buffer.AppendFormat("%03d", h.BatchNumber)         //62-64 3 NUM Batch Number 001 - 999 4.41 (001 - 999)
	buffer.AppendFormat("%d", h.BlockingIndicator)     //65 1 NUM Blocking Indicator 0 - Not Blocked 2 - Blocked

	// //all batch header messages have same exact structure up to this point...(first 65 bytes only)

	// 1 Passenger Travel
	if true == h.RecordType.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%-8s", h.TravelAgencyCode)  //66-73 8 A/N Travel Agency Code Left-justified/space-filled 4.318
		buffer.AppendFormat("%-25s", h.TravelAgencyName) //74-98 25 A/N Travel Agency Name Left-justified/space-filled 4.322
	}

	// 2 Hotel, Auto Rental, and Card not Present
	if true == h.RecordType.GetGroupBit(GroupMapBitType(2)) {
		buffer.AppendFormat("%-11s", h.MerchantLocalPhone)     //66-76 11 A/N Merchant Local Phone Num
		buffer.AppendFormat("%-11s", h.CardholderServicePhone) //999-9999999 4.187 77-87 11 A/N Cardholder Svc Phone Number 999-9999999 4.60
	}

	// 4 Gen2 Authentication Genkey (Only for Gen2 Authentication participants) Gen 2 terminal authentication
	if true == h.RecordType.GetGroupBit(GroupMapBitType(4)) {
		buffer.AppendFormat("%-24s", h.Genkey) //66-89 24 ASCII Genkey ASCII Representation of HEX 4.135
	}

	// 5 Developer ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(5)) {
		buffer.AppendFormat("%-6s", h.DeveloperId) //6 A/N 4.99 Developer ID
		buffer.AppendFormat("%-4s", h.VersionId)   //4 A/N 4.337 Version ID
	}

	// 6 Gateway ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(6)) {
		buffer.AppendFormat("%-10s", h.GatewayId) //10 A/N 4.134 Gateway ID
	}
}

func (h *BatchHeaderRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	h.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (h *BatchHeaderRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", h.RecordFormat)                    //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", h.ApplicationType)              //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", h.MessageDelimiter)                    //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", h.RoutingId)                         // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", h.RecordType)                              //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%-32s%06d\n", "AcquirerBIN", h.AcquirerBIN)                            //10-15 6 NUM Acquirer BIN 4.2 (999995)
	buffer.AppendFormat("%-32s%06d\n", "AgentBankNumber", h.AgentBankNumber)                    //16-21 6 NUM Agent Bank Number 111111 4.22
	buffer.AppendFormat("%-32s%06d\n", "AgentChainNumber", h.AgentChainNumber)                  //22-27 6 NUM Agent Chain Number 000000 4.23
	buffer.AppendFormat("%-32s%012d\n", "MerchantNumber", h.MerchantNumber)                     //28-39 12 NUM Merchant Number 4.189 (999999999911)
	buffer.AppendFormat("%-32s%04d\n", "StoreNumber", h.StoreNumber)                            //40-43 4 NUM Store Number 4.284 (1515)
	buffer.AppendFormat("%-32s%04d\n", "TerminalNumber", h.TerminalNumber)                      //44-47 4 NUM Terminal Number 4.300 (5999)
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", h.DeviceCode)                        //48 1 A/N Device Code C, D, E, I, M,Q 4.100 (ReservedForThirdPartyDevelopers)
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "IndustryCode", h.IndustryCode)                    //49 1 A/N Industry Code D, F, G, R 4.145 (UnknownOrUnsure)
	buffer.AppendFormat("%-32s%03d (%[2]s)\n", "CurrencyCode", h.CurrencyCode)                  //50-52 3 NUM Currency Code 4.85 (UsdUnitedStatesUsDollar)
	buffer.AppendFormat("%-32s%02d (%[2]s)\n", "LanguageIndicator", h.LanguageIndicator)        //53-54 2 NUM Language Indicator 00 4.163 (English)
	buffer.AppendFormat("%-32s%03s %s\n", "TimeZoneDifferential", h.TimeZoneDiff)               //55-57 3 NUM Time Zone Differential 4.305 (705 EASTERN)
	buffer.AppendFormat("%-32s%-4s (MMDD)\n", "BatchTransmissionDate", h.BatchTransmissionDate) //58-61 4 NUM Batch Transmission Date MMDD 4.45
	buffer.AppendFormat("%-32s%03d\n", "BatchNumber", h.BatchNumber)                            //62-64 3 NUM Batch Number 001 - 999 4.41 (001 - 999)
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "BlockingIndicator", h.BlockingIndicator)          //65 1 NUM Blocking Indicator 0 - Not Blocked 2 - Blocked

	// 1 Passenger Travel
	if true == h.RecordType.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%-32s%-8s\n", h.TravelAgencyCode)  //66-73 8 A/N Travel Agency Code Left-justified/space-filled 4.318
		buffer.AppendFormat("%-32s%-25s\n", h.TravelAgencyName) //74-98 25 A/N Travel Agency Name Left-justified/space-filled 4.322
	}

	// 2 Hotel, Auto Rental, and Card not Present
	if true == h.RecordType.GetGroupBit(GroupMapBitType(2)) {
		buffer.AppendFormat("%-32s%-11s\n", h.MerchantLocalPhone)     //66-76 11 A/N Merchant Local Phone Num
		buffer.AppendFormat("%-32s%-11s\n", h.CardholderServicePhone) //999-9999999 4.187 77-87 11 A/N Cardholder Svc Phone Number 999-9999999 4.60
	}

	// 4 Gen2 Authentication Genkey (Only for Gen2 Authentication participants) Gen 2 terminal authentication
	if true == h.RecordType.GetGroupBit(GroupMapBitType(4)) {
		buffer.AppendFormat("%-32s%-24s\n", h.Genkey) //66-89 24 ASCII Genkey ASCII Representation of HEX 4.135
	}

	// 5 Developer ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(5)) {
		buffer.AppendFormat("%-32s%-6s\n", h.DeveloperId) //6 A/N 4.99 Developer ID
		buffer.AppendFormat("%-32s%-4s\n", h.VersionId)   //4 A/N 4.337 Version ID
	}

	// 6 Gateway ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(6)) {
		buffer.AppendFormat("%-32s%-10s\n", h.GatewayId) //10 A/N 4.134 Gateway ID
	}

	return buffer.String()
}

func ParseBatchHeaderRecord(s string) (*BatchHeaderRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	h := BatchHeaderRecord{}

	if 65 > len(s) {
		panic("minimum EIS1081 Batch Header Record length is 65")
	}
	//i++ //advance past the STX

	fmt.Sscanf(s[0:4], "%c%1d%c%c",
		&h.RecordFormat,     //1 1 A/N Record Format K 4.242
		&h.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&h.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&h.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	h.RecordType = ParseOptionalDataGroups(s[4:9]) //5-9 5 A/N Record Type H@@@@ 4.243

	fmt.Sscanf(s[9:54], "%06d%06d%06d%012d%04d%04d%c%c%03d%02d",
		&h.AcquirerBIN,       //10-15 6 NUM Acquirer BIN 4.2 (999995)
		&h.AgentBankNumber,   //16-21 6 NUM Agent Bank Number 111111 4.22
		&h.AgentChainNumber,  //22-27 6 NUM Agent Chain Number 000000 4.23
		&h.MerchantNumber,    //28-39 12 NUM Merchant Number 4.189 (999999999911)
		&h.StoreNumber,       //40-43 4 NUM Store Number 4.284 (1515)
		&h.TerminalNumber,    //44-47 4 NUM Terminal Number 4.300 (5999)
		&h.DeviceCode,        //48 1 A/N Device Code C, D, E, I, M,Q 4.100 (ReservedForThirdPartyDevelopers)
		&h.IndustryCode,      //49 1 A/N Industry Code D, F, G, R 4.145 (UnknownOrUnsure)
		&h.CurrencyCode,      //50-52 3 NUM Currency Code 4.85 (UsdUnitedStatesUsDollar)
		&h.LanguageIndicator) //53-54 2 NUM Language Indicator 00 4.163 (English)

	var tmzonediff uint = 0
	fmt.Sscanf(s[54:57], "%03d", &tmzonediff)
	h.TimeZoneDiff = NewTimeZoneDifferentialFromUint(tmzonediff) //55-57 3 NUM Time Zone Differential 4.305 (705 EASTERN)

	h.BatchTransmissionDate = s[57:61] //58-61 4 NUM Batch Transmission Date MMDD 4.45

	fmt.Sscanf(s[61:65], "%03d%1d",
		&h.BatchNumber,       //62-64 3 NUM Batch Number 001 - 999 4.41 (001 - 999)
		&h.BlockingIndicator) //65 1 NUM Blocking Indicator 0 - Not Blocked 2 - Blocked

	i := 65

	// 1 Passenger Travel
	if true == h.RecordType.GetGroupBit(GroupMapBitType(1)) {
		h.TravelAgencyCode = s[i : i+8] //66-73 8 A/N Travel Agency Code Left-justified/space-filled 4.318
		i += 8
		h.TravelAgencyName = s[i : i+25] //74-98 25 A/N Travel Agency Name Left-justified/space-filled 4.322
		i += 25
	}

	// 2 Hotel, Auto Rental, and Card not Present
	if true == h.RecordType.GetGroupBit(GroupMapBitType(2)) {
		h.MerchantLocalPhone = s[i : i+11] //66-76 11 A/N Merchant Local Phone Num
		i += 11
		h.CardholderServicePhone = s[i : i+11] //999-9999999 4.187 77-87 11 A/N Cardholder Svc Phone Number 999-9999999 4.60
		i += 11
	}

	// 4 Gen2 Authentication Genkey (Only for Gen2 Authentication participants) Gen 2 terminal authentication
	if true == h.RecordType.GetGroupBit(GroupMapBitType(4)) {
		h.Genkey = s[i : i+11] //66-89 24 ASCII Genkey ASCII Representation of HEX 4.135
		i += 24
	}

	// 5 Developer ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(5)) {
		h.DeveloperId = s[i : i+6] //6 A/N 4.99 Developer ID
		i += 6
		h.VersionId = s[i : i+4] //4 A/N 4.337 Version ID
		i += 4
	}

	// 6 Gateway ID
	if true == h.RecordType.GetGroupBit(GroupMapBitType(6)) {
		h.GatewayId = s[i : i+10] //10 A/N 4.134 Gateway ID
		i += 10
	}

	return &h, i
}
