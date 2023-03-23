package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

// all code based upon TSYS technical reference data found within:
// EXTERNAL INTERFACE SPECIFICATIONS DATA
// CAPTURE RECORD FORMATS
// EIS 1081
// VERSION 13.2
// APRIL 21, 2015

type BatchParameterRecord struct {
	RecordFormat           RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType        ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter       bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId              rune                     // 4 1 A/N X.25 Routing ID Z 4.341
	RecordType             *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243
	CountryCode            CountryCodeType          //10-12 3 NUM Country Code 840 4.78
	CityCode               string                   //13-21 9 A/N City Code Left-justified/space-filled 4.74
	MerchantCategoryCode   MerchantCategoryType     //22-25 4 NUM Merchant Category Code 5999 4.183
	MerchantName           string                   //26-50 25 A/N Merchant Name Left-justified/space-filled 4.50.1
	MerchantCity           string                   //51-63 13 A/N Merchant City Left-justified/space-filled 4.50.2
	MerchantState          string                   //64-65 2 A/N Merchant State 4.50.3
	MerchantLocationNumber string                   //66-70 5 A/N Merchant Location Number 00001 4.188
	TerminalIdNumber       uint64                   //71-78 8 NUM Terminal ID Number 4.299
}

func NewBatchParameterRecord(country CountryCodeType,
	zipcode string,
	category MerchantCategoryType,
	name string,
	city string,
	state string,
	locationNumber string,
	terminalIdNo uint64) *BatchParameterRecord {

	p := BatchParameterRecord{}

	p.RecordFormat = Settlement
	p.ApplicationType = SingleBatchPerConnection
	p.MessageDelimiter = bytes.MessageDelimiter
	p.RoutingId = 'Z'
	p.RecordType = NewOptionalDataGroupMap(ParameterRecord)

	p.CountryCode = country                   //10-12 3 NUM Country Code 840 4.78
	p.CityCode = zipcode                      //13-21 9 A/N City Code Left-justified/space-filled 4.74
	p.MerchantCategoryCode = category         //22-25 4 NUM Merchant Category Code 5999 4.183
	p.MerchantName = name                     //26-50 25 A/N Merchant Name Left-justified/space-filled 4.50.1
	p.MerchantCity = city                     //51-63 13 A/N Merchant City Left-justified/space-filled 4.50.2
	p.MerchantState = state                   //64-65 2 A/N Merchant State 4.50.3
	p.MerchantLocationNumber = locationNumber //66-70 5 A/N Merchant Location Number 00001 4.188
	p.TerminalIdNumber = terminalIdNo         //71-78 8 NUM Terminal ID Number 4.299

	return &p
}

func (p *BatchParameterRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", p.RecordFormat)             //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%d", p.ApplicationType)          //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", p.MessageDelimiter)         //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", p.RoutingId)                // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", p.RecordType)               //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%03d", p.CountryCode)            //10-12 3 NUM Country Code 840 4.78
	buffer.AppendFormat("%-9s", p.CityCode)               //13-21 9 A/N City Code Left-justified/space-filled 4.74
	buffer.AppendFormat("%04d", p.MerchantCategoryCode)   //22-25 4 NUM Merchant Category Code 5999 4.183
	buffer.AppendFormat("%-25s", p.MerchantName)          //26-50 25 A/N Merchant Name Left-justified/space-filled 4.50.1
	buffer.AppendFormat("%-13s", p.MerchantCity)          //51-63 13 A/N Merchant City Left-justified/space-filled 4.50.2
	buffer.AppendFormat("%-2s", p.MerchantState)          //64-65 2 A/N Merchant State 4.50.3
	buffer.AppendFormat("%05d", p.MerchantLocationNumber) //66-70 5 A/N Merchant Location Number 00001 4.188
	buffer.AppendFormat("%08d", p.TerminalIdNumber)       //71-78 8 NUM Terminal ID Number 4.299
}

func (p *BatchParameterRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	p.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (p *BatchParameterRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", p.RecordFormat)                   //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", p.ApplicationType)             //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", p.MessageDelimiter)                   //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", p.RoutingId)                        // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", p.RecordType)                             //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%-32s%03d (%[2]s)\n", "CountryCode", p.CountryCode)                   //10-12 3 NUM Country Code 840 4.78
	buffer.AppendFormat("%-32s%-9s\n", "CityCode", p.CityCode)                                 //13-21 9 A/N City Code Left-justified/space-filled 4.74
	buffer.AppendFormat("%-32s%04d (%[2]s)\n", "MerchantCategoryCode", p.MerchantCategoryCode) //22-25 4 NUM Merchant Category Code 5999 4.183
	buffer.AppendFormat("%-32s%-25s\n", "MerchantName", p.MerchantName)                        //26-50 25 A/N Merchant Name Left-justified/space-filled 4.50.1
	buffer.AppendFormat("%-32s%-13s\n", "MerchantCity", p.MerchantCity)                        //51-63 13 A/N Merchant City Left-justified/space-filled 4.50.2
	buffer.AppendFormat("%-32s%-2s\n", "MerchantState", p.MerchantState)                       //64-65 2 A/N Merchant State 4.50.3
	buffer.AppendFormat("%-32s%05d\n", "MerchantLocationNumber", p.MerchantLocationNumber)     //66-70 5 A/N Merchant Location Number 00001 4.188
	buffer.AppendFormat("%-32s%08d\n", "TerminalIdNumber", p.TerminalIdNumber)                 //71-78 8 NUM Terminal ID Number 4.299

	return buffer.String()
}

func ParseBatchParameterRecord(s string) (*BatchParameterRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	p := BatchParameterRecord{}

	i := 0
	if 65 > len(s) {
		panic("minimum EIS1081 Batch Parameter Record length is 65")
	}
	//i++ //advance past the STX if not already removed

	fmt.Sscanf(s[i:i+4], "%c%1d%c%c",
		&p.RecordFormat,     //1 1 A/N Record Format K 4.242
		&p.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&p.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&p.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341
	i += 4

	p.RecordType = ParseOptionalDataGroups(s[i : i+5]) //5-9 5 A/N Record Type H@@@@ 4.243
	i += 5

	fmt.Sscanf(s[i:i+3], "%03d", &p.CountryCode) //10-12 3 NUM Country Code 840 4.78
	i += 3
	p.CityCode = s[i : i+9] //13-21 9 A/N City Code Left-justified/space-filled 4.74
	i += 9

	fmt.Sscanf(s[i:i+4], "%04d", &p.MerchantCategoryCode) //22-25 4 NUM Merchant Category Code 5999 4.183
	i += 4

	p.MerchantName = s[i : i+25] //26-50 25 A/N Merchant Name Left-justified/space-filled 4.50.1
	i += 25

	p.MerchantCity = s[i : i+13] //51-63 13 A/N Merchant City Left-justified/space-filled 4.50.2
	i += 13

	p.MerchantState = s[i : i+2] //64-65 2 A/N Merchant State 4.50.3
	i += 2

	fmt.Sscanf(s[i:i+13], "%05d%08d",
		&p.MerchantLocationNumber, //66-70 5 A/N Merchant Location Number 00001 4.188
		&p.TerminalIdNumber)       //71-78 8 NUM Terminal ID Number 4.299
	i += 13

	return &p, i
}
