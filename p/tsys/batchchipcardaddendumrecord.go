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

type BatchChipCardAddendumRecord struct {
	RecordFormat     RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType  ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId        rune                     //4 1 A/N X.25 Routing ID Z 4.341
	RecordType       *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243

	// Key:
	// The "V", "M", "A", and "D" fields indicate card brand usage. A "V" means Visa. A "M"
	// means MasterCard. A "A" means American Express. A "D" means Discover. The values
	// under these headers indicate the requirements for the card brands. A "R" means Required, a
	// "C" means Conditional (if it was included in the Auth, it should be included in the Capture), a
	// "O" means Optional.
	// If a data element is optional and data is not sent, the field must be space or zero filled.

	// V M A D Length Format Reference Contents
	TransactionType               ChipCardTransactionType // O R R R 2 A/N 4.318 Transaction Type
	TerminalTransactionDate       string                  // O R R R 6 NUM 4.302 Terminal Transaction Date (YYMMDD)
	TerminalVerificationResults   string                  // O R R R 10 A/N 4.303 Terminal Verification Results
	TerminalCurrencyCode          CurrencyCodeType        // O R R R 3 NUM 4.298 Terminal Currency Code
	ApplicationTransactionCounter string                  // O R R R 4 A/N 4.31 Application Transaction Counter
	ApplicationInterchangeProfile string                  // O R R R 4 A/N 4.30 Application Interchange Profile
	ApplicationCryptogram         string                  // O R R R 16 A/N 4.29 Application Cryptogram
	UnpredictableNumber           string                  // O R R C 8 A/N 4.329 Unpredictable Number
	IssuerApplicationData         string                  // O C R C 64 A/N 4.151 Issuer Application Data
	CryptogramInformationData     string                  // O R R O 2 A/N 4.82 Cryptogram Information Data
	TerminalCapabilityProfile     string                  // O O O R 6 A/N 4.296 Terminal Capability Profile
	CardSequenceNumber            string                  // O C R R 3 NUM 4.53 Card Sequence Number
	IssuerAuthenticationData      string                  // O O O C 32 A/N 4.153 Issuer Authentication Data
	CVMResults                    string                  // O R O O 6 A/N 4.89 CVM Results
	IssuerScriptResults           string                  // O O O C 50 A/N 4.154 Issuer Script Results
	FormFactorIdentifier          string                  // O O O O 8 A/N 4.128 Form Factor Identifier

	//(OPTIONAL DATA GROUP 1)
	CryptogramAmount uint64 //12 N 4.81 Cryptogram Amount
}

func NewBatchChipCardAddendumRecord(g1request *G1AuthorizationMessageRequest,
	g2request *G2AuthorizationMessageRequest,
	g1response *G1AuthorizationMessageResponse,
	g3v055request *G3v055IntegratedChipCardEmvTlvRequest,
	g3v055response *G3v055IntegratedChipCardEmvTlvResponse) *BatchChipCardAddendumRecord {

	d := BatchChipCardAddendumRecord{}

	d.RecordFormat = Settlement
	d.ApplicationType = SingleBatchPerConnection
	d.MessageDelimiter = bytes.MessageDelimiter
	d.RoutingId = 'Z'
	d.RecordType = NewOptionalDataGroupMap(ChipCardAddendumRecord)

	return &d
}

func (d *BatchChipCardAddendumRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {
	buffer.AppendFormat("%c", d.RecordFormat)     //1 1 A/N Record Formad K 4.242
	buffer.AppendFormat("%d", d.ApplicationType)  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", d.MessageDelimiter) //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", d.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", d.RecordType)       //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%-2s", d.TransactionType)               // O R R R 2 A/N 4.318 Transaction Type
	buffer.AppendFormat("%-6s", d.TerminalTransactionDate)       // O R R R 6 NUM 4.302 Terminal Transaction Date (YYMMDD)
	buffer.AppendFormat("%-10s", d.TerminalVerificationResults)  // O R R R 10 A/N 4.303 Terminal Verification Results
	buffer.AppendFormat("%03d", d.TerminalCurrencyCode)          // O R R R 3 NUM 4.298 Terminal Currency Code
	buffer.AppendFormat("%-4s", d.ApplicationTransactionCounter) // O R R R 4 A/N 4.31 Application Transaction Counter
	buffer.AppendFormat("%-4s", d.ApplicationInterchangeProfile) // O R R R 4 A/N 4.30 Application Interchange Profile
	buffer.AppendFormat("%-16s", d.ApplicationCryptogram)        // O R R R 16 A/N 4.29 Application Cryptogram
	buffer.AppendFormat("%-8s", d.UnpredictableNumber)           // O R R C 8 A/N 4.329 Unpredictable Number
	buffer.AppendFormat("%-64s", d.IssuerApplicationData)        // O C R C 64 A/N 4.151 Issuer Application Data
	buffer.AppendFormat("%-2s", d.CryptogramInformationData)     // O R R O 2 A/N 4.82 Cryptogram Information Data
	buffer.AppendFormat("%-6s", d.TerminalCapabilityProfile)     // O O O R 6 A/N 4.296 Terminal Capability Profile
	buffer.AppendFormat("%-3s", d.CardSequenceNumber)            // O C R R 3 NUM 4.53 Card Sequence Number
	buffer.AppendFormat("%-32s", d.IssuerAuthenticationData)     // O O O C 32 A/N 4.153 Issuer Authentication Data
	buffer.AppendFormat("%-6s", d.CVMResults)                    // O R O O 6 A/N 4.89 CVM Results
	buffer.AppendFormat("%-50s", d.IssuerScriptResults)          // O O O C 50 A/N 4.154 Issuer Script Results
	buffer.AppendFormat("%-8s", d.FormFactorIdentifier)          // O O O O 8 A/N 4.128 Form Factor Identifier

	// NOTE: ODG 1 should only be sent when the ARQC is sent in the Cryptogram field. If
	// a TC cryptogram is sent in the Chip Card Addendum record this data
	// group should not be sent .
	if true == d.RecordType.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%012d", &d.CryptogramAmount) //12 N 4.81 Cryptogram Amount
	}
}

func (d *BatchChipCardAddendumRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	d.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (d *BatchChipCardAddendumRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(256)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", d.RecordFormat)       //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", d.ApplicationType) //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", d.MessageDelimiter)       //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RoutingId", d.RoutingId)             // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", d.RecordType)                 //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%-32s%-2s (%[2]s)\n", "TransactionType", string(d.TransactionType), d.TransactionType) // O R R R 2 A/N 4.318 Transaction Type
	buffer.AppendFormat("%-32s%-6s (YYMMDD)\n", "TerminalTransactionDate", d.TerminalTransactionDate)           // O R R R 6 NUM 4.302 Terminal Transaction Date (YYMMDD)
	buffer.AppendFormat("%-32s%-10s\n", "TerminalVerificationResults", d.TerminalVerificationResults)           // O R R R 10 A/N 4.303 Terminal Verification Results
	buffer.AppendFormat("%-32s%03d (%[2]s)\n", "TerminalCurrencyCode", d.TerminalCurrencyCode)                  // O R R R 3 NUM 4.298 Terminal Currency Code
	buffer.AppendFormat("%-32s%-4s\n", "ApplicationTransactionCounter", d.ApplicationTransactionCounter)        // O R R R 4 A/N 4.31 Application Transaction Counter
	buffer.AppendFormat("%-32s%-4s\n", "ApplicationInterchangeProfile", d.ApplicationInterchangeProfile)        // O R R R 4 A/N 4.30 Application Interchange Profile
	buffer.AppendFormat("%-32s%-16s\n", "ApplicationCryptogram", d.ApplicationCryptogram)                       // O R R R 16 A/N 4.29 Application Cryptogram
	buffer.AppendFormat("%-32s%-8s\n", "UnpredictableNumber", d.UnpredictableNumber)                            // O R R C 8 A/N 4.329 Unpredictable Number
	buffer.AppendFormat("%-32s%-64s\n", "IssuerApplicationData", d.IssuerApplicationData)                       // O C R C 64 A/N 4.151 Issuer Application Data
	buffer.AppendFormat("%-32s%-2s\n", "CryptogramInformationData", d.CryptogramInformationData)                // O R R O 2 A/N 4.82 Cryptogram Information Data
	buffer.AppendFormat("%-32s%-6s\n", "TerminalCapabilityProfile", d.TerminalCapabilityProfile)                // O O O R 6 A/N 4.296 Terminal Capability Profile
	buffer.AppendFormat("%-32s%-3s\n", "CardSequenceNumber", d.CardSequenceNumber)                              // O C R R 3 NUM 4.53 Card Sequence Number
	buffer.AppendFormat("%-32s%-32s\n", "IssuerAuthenticationData", d.IssuerAuthenticationData)                 // O O O C 32 A/N 4.153 Issuer Authentication Data
	buffer.AppendFormat("%-32s%-6s\n", "CVMResults", d.CVMResults)                                              // O R O O 6 A/N 4.89 CVM Results
	buffer.AppendFormat("%-32s%-50s\n", "IssuerScriptResults", d.IssuerScriptResults)                           // O O O C 50 A/N 4.154 Issuer Script Results
	buffer.AppendFormat("%-32s%-8s\n", "FormFactorIdentifier", d.FormFactorIdentifier)                          // O O O O 8 A/N 4.128 Form Factor Identifier

	// NOTE: ODG 1 should only be sent when the ARQC is sent in the Cryptogram field. If
	// a TC cryptogram is sent in the Chip Card Addendum record this data
	// group should not be sent .
	if true == d.RecordType.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%-32s%012d", "CryptogramAmount", &d.CryptogramAmount) //12 N 4.81 Cryptogram Amount
	}

	return buffer.String()
}

func ParseBatchChipCardAddendumRecord(s string) (*BatchChipCardAddendumRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!
	d := BatchChipCardAddendumRecord{}

	i := 0
	if 232 > len(s) {
		panic("minimum EIS1081 Batch Chip Card Record length is 232")
	}

	//i++ //advance past the STX if not already removed

	fmt.Sscanf(s[0:4], "%c%1d%c%c",
		&d.RecordFormat,     //1 1 A/N Record Format K 4.242
		&d.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&d.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&d.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	d.RecordType = ParseOptionalDataGroups(s[4:9]) //5-9 5 A/N Record Type H@@@@ 4.243

	d.TransactionType = ChipCardTransactionType(s[9:11])  // O R R R 2 A/N 4.318 Transaction Type
	d.TerminalTransactionDate = s[11:17]                  // O R R R 6 NUM 4.302 Terminal Transaction Date (YYMMDD)
	d.TerminalVerificationResults = s[17:27]              // O R R R 10 A/N 4.303 Terminal Verification Results
	fmt.Sscanf(s[27:30], "%03d", &d.TerminalCurrencyCode) // O R R R 3 NUM 4.298 Terminal Currency Code
	d.ApplicationTransactionCounter = s[30:34]            // O R R R 4 A/N 4.31 Application Transaction Counter
	d.ApplicationInterchangeProfile = s[34:38]            // O R R R 4 A/N 4.30 Application Interchange Profile
	d.ApplicationCryptogram = s[38:54]                    // O R R R 16 A/N 4.29 Application Cryptogram
	d.UnpredictableNumber = s[54:62]                      // O R R C 8 A/N 4.329 Unpredictable Number
	d.IssuerApplicationData = s[62:126]                   // O C R C 64 A/N 4.151 Issuer Application Data
	d.CryptogramInformationData = s[126:128]              // O R R O 2 A/N 4.82 Cryptogram Information Data
	d.TerminalCapabilityProfile = s[128:134]              // O O O R 6 A/N 4.296 Terminal Capability Profile
	d.CardSequenceNumber = s[134:137]                     // O C R R 3 NUM 4.53 Card Sequence Number
	d.IssuerAuthenticationData = s[137:169]               // O O O C 32 A/N 4.153 Issuer Authentication Data
	d.CVMResults = s[169:175]                             // O R O O 6 A/N 4.89 CVM Results
	d.IssuerScriptResults = s[175:225]                    // O O O C 50 A/N 4.154 Issuer Script Results
	d.FormFactorIdentifier = s[225:233]                   // O O O O 8 A/N 4.128 Form Factor Identifier

	i = 233

	// NOTE: ODG 1 should only be sent when the ARQC is sent in the Cryptogram field. If
	// a TC cryptogram is sent in the Chip Card Addendum record this data
	// group should not be sent .
	if true == d.RecordType.GetGroupBit(GroupMapBitType(1)) {
		fmt.Sscanf(s[233:245], "%012d", &d.CryptogramAmount) //12 N 4.81 Cryptogram Amount
		i = 245
	}

	return &d, i
}
