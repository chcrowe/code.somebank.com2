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

type BatchDetailRecord struct {
	RecordFormat     RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType  ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId        rune                     //4 1 A/N X.25 Routing ID Z 4.341
	RecordType       *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243

	TransactionCode              TransactionCodeType     // 10-11 2 A/N Transaction Code 54 - Purchase 56 - Card not present CR - Credit/Return 4.312
	CardholderIdentificationCode CardholderIdCodeType    // 12 1 A/N Cardholder Identification Code @ - Signature M - Bad Mag read 4.57
	AccountDataSource            AccountDataSourceType   // 13 1 A/N Account Data Source Code @ - No Cardreader D - Track 2 Read H - Track 1 Read T - Keyed/Track 2 Capable X - Keyed/Track 1 Capable 4.1
	CardholderAccountNumber      string                  // 14-35 22 A/N Cardholder Account Number Left-justified/space-filled 4.55
	RequestedACI                 RequestAciType          // 36 1 A/N Requested ACI 4.251
	ReturnedACI                  ReturnAciType           // 37 1 A/N Returned ACI 4.261
	AuthorizationSourceCode      AuthorizationSourceType // 38 1 A/N Authorization Source Code 4.36
	TransactionSequenceNumber    uint                    // 39-42 4 NUM Transaction Sequence Number Right-justified/zero-filled 4.315
	ResponseCode                 ResponseCodeType        // 43-44 2 A/N Response Code 4.253
	ApprovalCode                 string                  // 45-50 6 A/N Approval Code Left-justified/space filled 4.33
	LocalTransactionDate         uint64                  // 51-54 4 NUM Local Transaction Date MMDD 4.175
	LocalTransactionTime         uint64                  // 55-60 6 NUM Local Transaction Time HHMMSS 4.176
	AVSResultCode                AvsResultType           // 61 1 A/N AVS Result Code 0 - No AVS 4.21
	TransactionIdentifier        string                  // 62-76 15 A/N Transaction Identifier Left-justified/space-filled 4.313
	ValidationCode               string                  // 77-80 4 A/N Validation Code 4.332
	VoidIndicator                VoidIndicatorType       // 81 1 A/N Void Indicator <SPACE> - Not Voided 4.339
	// 2 sub-fields 82-83 2 NUM Transaction Status Code 00 4.317
	ReversalTransmissionStatus ReversalTransmissionStatusType //position 1 REVERSAL TRANSMITTED 0 - No Authorization Reversal Transmitted 1 - An Authorization Reversal Was Transmitted
	ProtocolCompletionStatus   ProtocolCompletionStatusType   //position 2 PROTOCOL COMPLETION STATUS 0 - Host Acknowledgment <ACK> Received 1 - Host Acknowledgment NOT Received
	ReimbursementAttribute     ReimbursementAttributeType     // 84 1 A/N Reimbursement Attribute 4.245
	SettlementAmount           uint64                         // 85-96 12 NUM Settlement Amount Right-justified/zero-filled 4.269
	AuthorizedAmount           uint64                         // 97-108 12 NUM Authorized Amount Right-justified/zero-filled 4.37

	OptionalData *BatchDetailOptionalDataRecord
}

func NewBatchDetailRecord(g1request *G1AuthorizationMessageRequest,
	g2request *G2AuthorizationMessageRequest,
	g1response *G1AuthorizationMessageResponse,
	settlementAmount uint64) *BatchDetailRecord {

	d := BatchDetailRecord{}

	d.RecordFormat = Settlement
	d.ApplicationType = SingleBatchPerConnection
	d.MessageDelimiter = bytes.MessageDelimiter
	d.RoutingId = 'Z'
	d.RecordType = NewOptionalDataGroupMap(DetailRecord)

	d.TransactionCode = g1request.TransactionCode
	d.CardholderIdentificationCode = g1request.CardholderIdentificationCode
	d.AccountDataSource = g1request.AccountDataSource
	//d.CardholderAccountNumber = g1request.CustomerData
	d.RequestedACI = g1request.RequestedACI
	d.ReturnedACI = g1response.ReturnedACI
	d.AuthorizationSourceCode = g1response.AuthorizationSourceCode
	d.TransactionSequenceNumber = g1request.TransactionSequenceNumber
	d.ResponseCode = g1response.ResponseCode
	d.ApprovalCode = g1response.ApprovalCode

	localdt := fmt.Sprintf("%02d%02d", g1response.LocalTxnDateTime.Month(), g1response.LocalTxnDateTime.Day())
	localtm := fmt.Sprintf("%02d%02d%02d", g1response.LocalTxnDateTime.Hour(), g1response.LocalTxnDateTime.Minute(), g1response.LocalTxnDateTime.Second())

	fmt.Sscanf(localdt, "%04d", &d.LocalTransactionDate)
	fmt.Sscanf(localtm, "%06d", &d.LocalTransactionTime)

	d.AVSResultCode = g1response.AVSResultCode
	d.TransactionIdentifier = g1response.TransactionIdentifier
	d.ValidationCode = g1response.ValidationCode
	d.VoidIndicator = NotVoided

	d.ReversalTransmissionStatus = AuthorizationReversalNotTransmitted
	d.ProtocolCompletionStatus = HostAcknowledgementReceived
	d.ReimbursementAttribute = EbtNonDebitOrNonInterlinkDebit

	d.SettlementAmount = settlementAmount
	d.AuthorizedAmount = g1request.TransactionAmount

	d.OptionalData = NewBatchDetailOptionalDataRecord(d.RecordType,
		g1request, g2request, g1response, settlementAmount)

	return &d
}

func (d *BatchDetailRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {
	buffer.AppendFormat("%c", d.RecordFormat)                 //1 1 A/N Record Formad K 4.242
	buffer.AppendFormat("%d", d.ApplicationType)              //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", d.MessageDelimiter)             //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", d.RoutingId)                    // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", d.RecordType)                   //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%2s", string(d.TransactionCode))     // 10-11 2 A/N Transaction Code 54 - Purchase 56 - Card not present CR - Credit/Return 4.312
	buffer.AppendFormat("%c", d.CardholderIdentificationCode) // 12 1 A/N Cardholder Identification Code @ - Signature M - Bad Mag read 4.57
	buffer.AppendFormat("%c", d.AccountDataSource)            // 13 1 A/N Account Data Source Code @ - No Cardreader D - Track 2 Read H - Track 1 Read T - Keyed/Track 2 Capable X - Keyed/Track 1 Capable 4.1
	buffer.AppendFormat("%-22s", d.CardholderAccountNumber)   // 14-35 22 A/N Cardholder Account Number Left-justified/space-filled 4.55
	buffer.AppendFormat("%c", d.RequestedACI)                 // 36 1 A/N Requested ACI 4.251
	buffer.AppendFormat("%c", d.ReturnedACI)                  // 37 1 A/N Returned ACI 4.261
	buffer.AppendFormat("%c", d.AuthorizationSourceCode)      // 38 1 A/N Authorization Source Code 4.36
	buffer.AppendFormat("%04d", d.TransactionSequenceNumber)  // 39-42 4 NUM Transaction Sequence Number Right-justified/zero-filled 4.315
	buffer.AppendFormat("%2s", string(d.ResponseCode))        // 43-44 2 A/N Response Code 4.253
	buffer.AppendFormat("%-6s", d.ApprovalCode)               // 45-50 6 A/N Approval Code Left-justified/space filled 4.33
	buffer.AppendFormat("%04d", d.LocalTransactionDate)       // 51-54 4 NUM Local Transaction Date MMDD 4.175
	buffer.AppendFormat("%06d", d.LocalTransactionTime)       // 55-60 6 NUM Local Transaction Time HHMMSS 4.176
	buffer.AppendFormat("%c", d.AVSResultCode)                // 61 1 A/N AVS Result Code 0 - No AVS 4.21
	buffer.AppendFormat("%-15s", d.TransactionIdentifier)     // 62-76 15 A/N Transaction Identifier Left-justified/space-filled 4.313
	buffer.AppendFormat("%-4s", d.ValidationCode)             // 77-80 4 A/N Validation Code 4.332
	buffer.AppendFormat("%c", d.VoidIndicator)                // 81 1 A/N Void Indicator <SPACE> - Not Voided 4.339
	// 2 sub-fields 82-83 2 NUM Transaction Status Code 00 4.317
	buffer.AppendFormat("%d", d.ReversalTransmissionStatus) //position 1 REVERSAL TRANSMITTED 0 - No Authorization Reversal Transmitted 1 - An Authorization Reversal Was Transmitted
	buffer.AppendFormat("%d", d.ProtocolCompletionStatus)   //position 2 PROTOCOL COMPLETION STATUS 0 - Host Acknowledgment <ACK> Received 1 - Host Acknowledgment NOT Received
	buffer.AppendFormat("%c", d.ReimbursementAttribute)     // 84 1 A/N Reimbursement Attribute 4.245
	buffer.AppendFormat("%012d", d.SettlementAmount)        // 85-96 12 NUM Settlement Amount Right-justified/zero-filled 4.269
	buffer.AppendFormat("%012d", d.AuthorizedAmount)        // 97-108 12 NUM Authorized Amount Right-justified/zero-filled 4.37
}

func (d *BatchDetailRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	d.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (d *BatchDetailRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(256)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", d.RecordFormat)                                 //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", d.ApplicationType)                           //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", d.MessageDelimiter)                                 //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RoutingId", d.RoutingId)                                       // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", d.RecordType)                                           //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%-32s%2s\n", "TransactionCode", string(d.TransactionCode))                          // 10-11 2 A/N Transaction Code 54 - Purchase 56 - Card not present CR - Credit/Return 4.312
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "CardholderIdentificationCode", d.CardholderIdentificationCode) // 12 1 A/N Cardholder Identification Code @ - Signature M - Bad Mag read 4.57
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AccountDataSource", d.AccountDataSource)                       // 13 1 A/N Account Data Source Code @ - No Cardreader D - Track 2 Read H - Track 1 Read T - Keyed/Track 2 Capable X - Keyed/Track 1 Capable 4.1
	buffer.AppendFormat("%-32s%-22s\n", "CardholderAccountNumber", d.CardholderAccountNumber)                // 14-35 22 A/N Cardholder Account Number Left-justified/space-filled 4.55
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RequestedACI", d.RequestedACI)                                 // 36 1 A/N Requested ACI 4.251
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "ReturnedACI", d.ReturnedACI)                                   // 37 1 A/N Returned ACI 4.261
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AuthorizationSourceCode", d.AuthorizationSourceCode)           // 38 1 A/N Authorization Source Code 4.36
	buffer.AppendFormat("%-32s%04d\n", "TransactionSequenceNumber", d.TransactionSequenceNumber)             // 39-42 4 NUM Transaction Sequence Number Right-justified/zero-filled 4.315
	buffer.AppendFormat("%-32s%2s (%[2]s)\n", "ResponseCode", string(d.ResponseCode))                        // 43-44 2 A/N Response Code 4.253
	buffer.AppendFormat("%-32s%-6s\n", "ApprovalCode", d.ApprovalCode)                                       // 45-50 6 A/N Approval Code Left-justified/space filled 4.33
	buffer.AppendFormat("%-32s%04d (MMDD)\n", "LocalTransactionDate", d.LocalTransactionDate)                // 51-54 4 NUM Local Transaction Date MMDD 4.175
	buffer.AppendFormat("%-32s%06d (HHMMSS)\n", "LocalTransactionTime", d.LocalTransactionTime)              // 55-60 6 NUM Local Transaction Time HHMMSS 4.176
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AVSResultCode", d.AVSResultCode)                               // 61 1 A/N AVS Result Code 0 - No AVS 4.21
	buffer.AppendFormat("%-32s%-15s\n", "TransactionIdentifier", d.TransactionIdentifier)                    // 62-76 15 A/N Transaction Identifier Left-justified/space-filled 4.313
	buffer.AppendFormat("%-32s%-4s\n", "ValidationCode", d.ValidationCode)                                   // 77-80 4 A/N Validation Code 4.332
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "VoidIndicator", d.VoidIndicator)                               // 81 1 A/N Void Indicator <SPACE> - Not Voided 4.339
	// 2 sub-fields 82-83 2 NUM Transaction Status Code 00 4.317
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ReversalTransmissionStatus", d.ReversalTransmissionStatus) //position 1 REVERSAL TRANSMITTED 0 - No Authorization Reversal Transmitted 1 - An Authorization Reversal Was Transmitted
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ProtocolCompletionStatus", d.ProtocolCompletionStatus)     //position 2 PROTOCOL COMPLETION STATUS 0 - Host Acknowledgment <ACK> Received 1 - Host Acknowledgment NOT Received
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "ReimbursementAttribute", d.ReimbursementAttribute)         // 84 1 A/N Reimbursement Attribute 4.245
	buffer.AppendFormat("%-32s%012d\n", "SettlementAmount", d.SettlementAmount)                          // 85-96 12 NUM Settlement Amount Right-justified/zero-filled 4.269
	buffer.AppendFormat("%-32s%012d\n", "AuthorizedAmount", d.AuthorizedAmount)                          // 97-108 12 NUM Authorized Amount Right-justified/zero-filled 4.37

	return buffer.String()
}

func ParseBatchDetailRecord(s string) (*BatchDetailRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!
	d := BatchDetailRecord{}

	i := 0
	if 108 > len(s) {
		panic("minimum EIS1081 Batch Detail Record length is 108")
	}

	//i++ //advance past the STX if not already removed

	fmt.Sscanf(s[0:4], "%c%1d%c%c",
		&d.RecordFormat,     //1 1 A/N Record Format K 4.242
		&d.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&d.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&d.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	d.RecordType = ParseOptionalDataGroups(s[4:9]) //5-9 5 A/N Record Type H@@@@ 4.243

	d.TransactionCode = TransactionCodeType(s[9:11])             // 10-11 2 A/N Transaction Code 54 - Purchase 56 - Card not present CR - Credit/Return 4.312
	d.CardholderIdentificationCode = CardholderIdCodeType(s[11]) // 12 1 A/N Cardholder Identification Code @ - Signature M - Bad Mag read 4.57
	d.AccountDataSource = AccountDataSourceType(s[12])           // 13 1 A/N Account Data Source Code @ - No Cardreader D - Track 2 Read H - Track 1 Read T - Keyed/Track 2 Capable X - Keyed/Track 1 Capable 4.1
	d.CardholderAccountNumber = s[13:35]                         // 14-35 22 A/N Cardholder Account Number Left-justified/space-filled 4.55
	d.RequestedACI = RequestAciType(s[35])                       // 36 1 A/N Requested ACI 4.251
	d.ReturnedACI = ReturnAciType(s[36])                         // 37 1 A/N Returned ACI 4.261
	d.AuthorizationSourceCode = AuthorizationSourceType(s[37])   // 38 1 A/N Authorization Source Code 4.36
	fmt.Sscanf(s[38:42], "%04d", &d.TransactionSequenceNumber)   // 39-42 4 NUM Transaction Sequence Number Right-justified/zero-filled 4.315
	d.ResponseCode = ResponseCodeType(s[42:44])                  // 43-44 2 A/N Response Code 4.253
	d.ApprovalCode = s[44:50]                                    // 45-50 6 A/N Approval Code Left-justified/space filled 4.33
	fmt.Sscanf(s[50:54], "%04d", &d.LocalTransactionDate)        // 51-54 4 NUM Local Transaction Date MMDD 4.175
	fmt.Sscanf(s[54:60], "%06d", &d.LocalTransactionTime)        // 55-60 6 NUM Local Transaction Time HHMMSS 4.176
	d.AVSResultCode = AvsResultType(s[60])                       // 61 1 A/N AVS Result Code 0 - No AVS 4.21
	d.TransactionIdentifier = s[61:76]                           // 62-76 15 A/N Transaction Identifier Left-justified/space-filled 4.313
	d.ValidationCode = s[76:80]                                  // 77-80 4 A/N Validation Code 4.332
	d.VoidIndicator = VoidIndicatorType(s[80])                   // 81 1 A/N Void Indicator <SPACE> - Not Voided 4.339
	// 2 sub-fields 82-83 2 NUM Transaction Status Code 00 4.317
	fmt.Sscanf(s[81:82], "%01d", &d.ReversalTransmissionStatus)  //position 1 REVERSAL TRANSMITTED 0 - No Authorization Reversal Transmitted 1 - An Authorization Reversal Was Transmitted
	fmt.Sscanf(s[82:83], "%01d", &d.ProtocolCompletionStatus)    //position 2 PROTOCOL COMPLETION STATUS 0 - Host Acknowledgment <ACK> Received 1 - Host Acknowledgment NOT Received
	d.ReimbursementAttribute = ReimbursementAttributeType(s[83]) // 84 1 A/N Reimbursement Attribute 4.245
	fmt.Sscanf(s[84:96], "%012d", &d.SettlementAmount)           // 85-96 12 NUM Settlement Amount Right-justified/zero-filled 4.269
	fmt.Sscanf(s[96:108], "%012d", &d.AuthorizedAmount)          // 97-108 12 NUM Authorized Amount Right-justified/zero-filled 4.37

	i = 108

	return &d, i
}

// Table 3.2 Transaction header record - optional data groups
// 0 Header Record
// 1 Passenger Travel
// 2 Hotel, Auto Rental, and Card not Present
// 4 Gen2 Authentication Genkey (Only for Gen2 Authentication participants)
// 5 Developer ID
// 6 Gateway ID

// Table 3.6 Transaction detail record - optional data groups
// 1 Cash Back
// 	12 NUM 4.63 Cash Back Amount
// 2 Restaurant
// 	12 NUM 4.137 Gratuity Amount
// 3 Direct Marketing
// 	12 NUM 4.306 Total Authorized Amount
// 	1 A/N 4.236 Purchase Identifier Format Code
// 	25 A/N 4.235 Purchase Identifier

// 4 Auto Rental
// 	12 NUM 4.306 Total Authorized Amount
// 	1 A/N 4.236 Purchase Identifier Format Code
// 	25 A/N 4.235 Purchase Identifier
// 	1 A/N 4.181 Market Specific Data Identifier
// 	1 A/N 4.215 No Show Indicator
// 	6 A/N 4.121.1 Extra Charges
// 	6 NUM 4.246 Rental Date (YYMMDD)
// 5 Hotel
// 	12 NUM 4.306 Total Authorized Amount
// 	1 A/N 4.236 Purchase Identifier Format Code
// 	25 A/N 4.235 Purchase Identifier
// 	1 A/N 4.181 Market Specific Data Identifier
// 	1 A/N 4.215 4.278 No Show or Special Program Ind.
// 	6 A/N 4.121 Extra Charges
// 	6 NUM 4.72 Check-in Date (YYMMDD)
// 6 Hotel - Not Front Desk (Non-Lodging Purchase)
// 	4 NUM 4.185 Merchant Category Code Override
// 7 Hotel - American Express
// 	1 A/N 4.71 Charge Type
// 	6 NUM 4.73 Checkout Date (YYMMDD)
// 	2 NUM 4.281 Stay Duration
// 	12 NUM 4.180 Lodging Room Rate
// 8 Non-T&E Commercial Card Level II (Visa and MasterCard Only)
// 	1 A/N 4.221 Optional Amount Identifier
// 	12 NUM 4.220 Optional Amount
// 	16 A/N 4.237 Purchase Order Number
// 9 Passenger Transport (Visa Only)
// 	1 A/N 4.254 Restricted Ticket Indicator
// 	2 NUM 4.207 Multiple Clearing Sequence Number
// 	2 NUM 4.206 Multiple Clearing Sequence Count
// 	13 A/N 4.304 Ticket Number
// 	20 A/N 4.228 Passenger Name
// 	6 NUM 4.93 Departure Date (MMDDYY)
// 	3 A/N 4.224 Origination City/Airport Code
// 	7 A/N 4.323 Trip Leg 1 Information
// 	7 A/N 4.323 Trip Leg 2 Information
// 	7 A/N 4.323 Trip Leg 3 Information
// 	7 A/N 4.323 Trip Leg 4 Information
// 10 Direct Debit
// 	12 A/N 4.259 Retrieval Reference Number
// 	6 NUM 4.288 System Trace Audit Number
// 	1 A/N 4.213 Network Identification Code
// 	4 NUM 4.270 Settlement Date (MMDD)
// 	11 Multi-Terminal Stores
// 	4 NUM 4.301 Terminal Number Override
// 12 Direct Marketing - Installment or Recurring Payments
// 	2 NUM 4.207 Multiple Clearing Sequence Number
// 	2 NUM 4.206 Multiple Clearing Sequence Count
// 	1 A/N 4.205 MOTO/e-Commerce Indicator
// 13 Auto Rental
// 	20 A/N 4.250 Renter Name
// 	18 A/N 4.247 Rental Return City
// 	3 A/N 4.249 Rental Return State/Country
// 	10 A/N 4.248 Rental Return Location ID
// 14 Service Development Indicator
// 	1 NUM 4.267 Service Development Indicator
// 16 Existing Debt Indicator (Visa Only)
// 	1 NUM 4.119 Existing Debt Indicator
// 17 Universal Cardholder Authentication
// 	1 A/N 4.324 UCAF Collection Indicator
// 18 e-Commerce Goods Indicator
// 	2 A/N 4.108 e-Commerce Goods Indicator
// 19 CCS Private Label Restricted for CCS use only (Mandatory for Private Label transactions)
// 	5 NUM 4.80 Credit Plan Number (Not Null)
// 	4 A/N 4.233 Product Code 1
// 	4 A/N 4.233 Product Code 2
// 	4 A/N 4.233 Product Code 3
// 	4 A/N 4.233 Product Code 4
// 	4 A/N 4.233 Product Code 5
// 	3 NUM 4.223 Org. ID (Not Null)
// 	9 NUM 4.283 Store Code (Not Null)
// 	6 NUM 4.241 Receiving Institution ID (“911111”)
// 20 Merchant Verification Value
// 	10 A/N 4.195 Merchant Verification Value
// 21 American Express Corporate Purchasing Card
// 	9 A/N 4.287 Supplier Reference Number
// 	17 A/N 4.59 Cardholder Reference Number
// 	6 A/N 4.277 Shipped To ZIP Code
// 	6 N 4.262 Sales Tax
// 	40 A/N 4.65 Charge Descriptor 1
// 22 U.S. Only Non-T&E Commercial Card Level II (Visa and MasterCard)
// 	1 A/N 4.221 Optional Amount Identifier
// 	12 NUM 4.220 Optional Amount
// 	17 A/N 4.237 Purchase Order No. - Customer Ref. ID
// 23 Group Map Extension (Group 25 - 48) Reserved
// 	4 A/N 4.243 Record Type @@@@
// 24 Reserved (POS-partner) Use Only

// EXTENDED BIT MAP SECTION

// 25 Commercial Card / Level II (Non-T&E and Fleet Non-fuel)
// 	12 NUM 4.306 Total Authorized Amount
// 	1 A/N 4.236 Purchase Identifier Format Code
// 	25 A/N 4.235 Purchase Identifier
// 	1 NUM 4.174 Local Tax Included Flag
// 	12 NUM 4.173 Local Tax
// 	1 NUM 4.209 National Tax Included Flag
// 	12 NUM 4.208 National Tax Amount
// 	17 A/N 4.237 Purchase Order No. / Cust. Ref. ID
// 26 Visa Card Enhanced Data (Non-T&E and Fleet Non-fuel)
// 	20 A/N 4.194 Merchant VAT Registration Number
// 	13 A/N 4.87 Customer VAT Registration Number
// 	4 A/N 4.285 Summary Commodity Code
// 	12 NUM 4.101 Discount Amount
// 	12 NUM 4.130 Freight Amount
// 	12 NUM 4.107 Duty Amount
// 	10 A/N 4.98 Destination Postal / ZIP Code
// 	10 A/N 4.271 Ship from Postal / ZIP Code
// 	3 A/N 4.97 Destination Country Code
// 	15 A/N 4.325 Unique VAT Invoice Reference Number
// 	6 NUM 4.222 Order Date
// 	12 NUM 4.333 VAT / Tax Amount (Freight/Shipping)
// 	4 NUM 4.334 VAT / Tax Rate (Freight/Shipping)
// 	3 NUM 4.167 Line Item Count
// 27 MasterCard Enhanced Data (Non-T&E and Fleet Non-fuel)
// 	12 NUM 4.130 Freight Amount
// 	12 NUM 4.107 Duty Amount
// 	10 A/N 4.98 Destination Postal / Zip Code
// 	10 A/N 4.271 Ship From Postal / Zip Code
// 	3 A/N 4.97 Destination Country Code
// 	1 A/N 4.25 Alternate Tax Amount Indicator
// 	9 NUM 4.24 Alternate Tax Amount
// 	3 NUM 4.167 Line Item Count
// 28 Visa Card Enhanced Data (Fleet)
// 	4 NUM 4.238 Purchase Time
// 	1 A/N 4.168 Service Type
// 	2 NUM 4.133 Fuel Type
// 	12 NUM 4.203 Motor Fuel Unit Price
// 	1 A/N 4.326 Motor Fuel Unit of Measure
// 	12 NUM 4.240 Motor Fuel Quantity
// 	12 NUM 4.138 Gross Fuel Price
// 	12 NUM 4.210 Net Fuel Price
// 	12 NUM 4.212 Net Non-fuel Price
// 	7 A/N 4.217 Odometer Reading
// 	17 A/N 4.237 Cust. Code / Cust. Ref. Identifier (CRI)
// 	1 NUM 4.239 Type of Purchase
// 	12 NUM 4.279 State Fuel Tax Amount
// 	12 NUM 4.172 Local Fuel Tax Amount
// 	4 NUM 4.334 VAT Tax Rate
// 	2 A/N 4.233 Product Code 1
// 	2 A/N 4.233 Product Code 2
// 	2 A/N 4.233 Product Code 3
// 	2 A/N 4.233 Product Code 4
// 	2 A/N 4.233 Product Code 5
// 	2 A/N 4.233 Product Code 6
// 	2 A/N 4.233 Product Code 7
// 	2 A/N 4.233 Product Code 8
// 	12 NUM 4.173 Local Sales Tax
// 	1 A/N 4.174 Local Tax Included
// 	9 NUM 4.139 Gross Non-fuel price
// 	1 A/N 4.198 Misc. Fuel Tax Exemption Status
// 	1 A/N 4.199 Misc. Non-Fuel Tax Exemption Status
// 	3 NUM 4.167 Line Item Count (Must be 000)
// 29 MasterCard Enhanced Data (Fleet - Fuel Only)
// 	4 NUM 4.218 Oil Company Brand Names
// 	30 A/N 4.193 Merchant Street Address
// 	9 A/N 4.190 Merchant Postal/ZIP Code and Ext.
// 	4 NUM 4.238 Purchase Time
// 	1 A/N 4.268 Motor Fuel Service Type
// 	3 NUM 4.233 Motor Fuel Product Code
// 	5 NUM 4.203 Motor Fuel Unit Price
// 	1 A/N 4.327 Motor Fuel Unit of Measure
// 	6 NUM 4.240 Motor Fuel Quantity
// 	9 NUM 4.202 Motor Fuel Sale Amount
// 	7 A/N 4.217 Odometer Reading
// 	17 A/N 4.335 Vehicle Number
// 	17 A/N 4.106 Driver Number / ID No.
// 	1 NUM 4.234 Product Type Code
// 	5 NUM 4.101 Coupon/Discount Amount
// 	6 A/N 4.291 Tax Amount 1
// 	6 A/N 4.292 Tax Amount 2
// 	3 NUM 4.167 Line Item Count
// 30 American Express CPC Charge Descriptors
// 	40 A/N 4.66 Charge Descriptor 2
// 	40 A/N 4.67 Charge Descriptor 3
// 	40 A/N 4.68 Charge Descriptor 4
// 31 Market Specific Data Indicator
// 	1 A/N 4.181 Market Specific Data Indicator
// 32 POS Data Code
// 	12 A/N 4.231 POS Data Code
// 33 Retail Transaction Advice Addenda
// 	40 A/N 4.255 Retail Department Name
// 34 Lodging Transaction Advice Addenda
// 	26 A/N 4.179 Lodging Renter Name
// 35 Transaction Advice Line Item Counts
// 	3 NUM 4.290 TAD Line Item Count
// 	3 NUM 4.289 TAA Line Item Count
// 36 Location Detail Transaction Advice Addenda
// 	38 A/N 4.177 Location Detail Name
// 	38 A/N 4.177 Location Detail Address
// 	21 A/N 4.177 Location Detail City
// 	3 A/N 4.177 Location Detail Region Code
// 	3 A/N 4.177 Country Code
// 	15 A/N 4.177 Location Detail Postal Code
// 37 MasterCard Miscellaneous Fields
// 	3 NUM 4.338 Version Number
// 	3 A/N 4.265 Service Code
// 38 Visa Miscellaneous Fields
// 	3 NUM 4.338 Version Number
// 	2 A/N 4.51 Card product code
// 39 Amex CAPN Corporate Purchasing Solution Extension
// 	38 A/N 4.252 Requester Name
// 	12 Num 4.311 Total Tax Amount
// 	3 Num 4.294 Tax Type Code
// 40 Discover/ PayPal Miscellaneous Fields
// 	3 Num 4.338 Version Number
// 	13 A/N 4.103 Discover/ PayPal POS Data Code
// 	1 A/N 4.227 Partial Shipment Indicator
// 41 Detail Extension
// 	4 Num 4.141 Group Length
// 	Variable A/N 4.140 Group Fields

// 	ODG 41 Tag Description
// 		TAG Max size Format Reference Content/Description
// 		OAI 1 A/N 4.221 Optional Amount Identifier
// 		OA 12 NUM 4.220 Optional Amount
// 		PON 25 A/N 4.237 Purchase Order Number
// 		PTI 3 NUM 4.229 Payment Transaction Indicator
// 		IIA 1 A/N 4.182 MC IIAS Indicator
// 		SID 20 A/N 4.191 Merchant Seller ID
// 		CDV 82 A/N 4.88 Custom Identifier Detail
// 		ATS 10 NUM 4.35 Association Timestamp
// 		TTT 2 NUM 4.319 Transit Transaction Type Indicator
// 		TMI 2 NUM 4.320 Transportation Mode Indicator
// 		TFA 9 A/N 4.313 Transaction Fee Amount with Indicator
// 		MDO 1 A/N 4.104 Domain Server
// 		MDE 2 A/N 4.201 Mobile Device Type
// 		DTC 1 A/N 4.149 Invoice Level Discount Treatment Code
// 		TTC 1 A/N 4.293 Tax Treatment Code
// 		SQI 1 A/N 4.279 Spend Qualified Indicator
// 		TAL 2 A/N 4.306 Token Assurance Level
// 		TRI 11 A/N 4.307 Token Requestor ID
// 		ARS 1 A/N 4.2 Account Range Status
// 		WID 3 A/N 4.183 MasterCard Wallet Identifier
// 		ISO 11 NUM 4.151 ISO ID
// 		PFI 11 NUM 4.229 Payment Facilitator ID NOTE: The presence of this tag requires use of the SMI TAG as well as Usage 2 of ODG 36 “Location Detail Transaction Advice Addenda.”
// 		SMI 15 A/N 4.285 Sub-Merchant ID NOTE: The presence of this tag requires use of the PFI TAG as well as Usage 2 of ODG 36 “Location Detail Transaction Advice Addenda.”

// 42 Amex Auto Rental
// 	14 A/N 4.27.1 Amex Auto Rental Agreement Number
// 	38 A/N 4.27.2 Amex Auto Rental Pickup Location
// 	18 A/N 4.27.3 Amex Auto Rental Pickup City Name
// 	3 A/N 4.27.4 Amex Auto Rental Pickup Region Code
// 	3 A/N 4.27.5 Amex Auto Rental Pickup Country Code
// 	8 NUM 4.27.6 Amex Auto Rental Pickup Date
// 	6 NUM 4.27.7 Amex Auto Rental Pickup Time
// 	18 A/N 4.27.8 Amex Auto Rental Return City Name
// 	3 A/N 4.27.9 Amex Auto Rental Return Region Code
// 	3 A/N 4.27.10 Amex Auto Rental Return Country Code
// 	8 NUM 4.27.11 Amex Auto Rental Return Date
// 	6 NUM 4.27.12 Amex Auto Rental Return Time
// 	26 A/N 4.27.13 Amex Auto Rental Renter Name
// 	4 A/N 4.27.14 Amex Auto Rental Vehicle Class ID
// 	5 NUM 4.27.15 Amex Auto Rental Distance
// 	1 A/N 4.27.16 Amex Auto Rental Distance Unit of Measurement
// 	1 A/N 4.27.17 Amex Auto Rental Audit Adjustment Indicator
// 	12 NUM 4.27.18 Amex Auto Rental Audit Adjustment Amount

// 43 Amex Insurance
// 	23 A/N 4.28.1 Policy Number
// 	8 NUM 4.28.2 Coverage Start Date
// 	8 NUM 4.28.3 Coverage End Date
// 	7 A/N 4.28.4 Policy Premium Frequency
// 	23 A/N 4.28.5 Additional Insurance Policy Number
// 	25 A/N 4.28.6 Type of Policy
// 	30 A/N 4.28.7 Name of Insured

// 45 Transaction Security Indicator
// 	1 A/N 4.315 Transaction Security Indicator

// 46 Merchant Soft Descriptor
// 	25 A/N 4.192 Merchant Soft Descriptor
