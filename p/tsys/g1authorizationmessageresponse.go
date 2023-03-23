package tsys

import (
	"fmt"
	"strings"
	"time"

	"code.somebank.com/p/bytes"
)

const REPONSEGRP1MINLEN = 67

type G1AuthorizationMessageResponse struct {
	RecordFormat              RecordFormatType         //1 A/N 4.67
	ApplicationType           ApplicationIndicatorType //1 A/N 4.9
	MessageDelimiter          bytes.AsciiCharacterType //1 A/N 4.62
	ReturnedACI               ReturnAciType            //1 A/N
	StoreNumber               uint                     //4 NUM
	TerminalNumber            uint                     //4 NUM
	AuthorizationSourceCode   AuthorizationSourceType  //1 A/N
	TransactionSequenceNumber uint                     //4 NUM
	ResponseCode              ResponseCodeType         //2 A/N
	ApprovalCode              string                   //6 A/N
	LocalTxnDateTime          time.Time
	AuthorizationResponseText string                 //16 A/N
	AVSResultCode             AvsResultType          //1 A/N
	RetrievalReferenceNumber  string                 //12 A/N
	MarketSpecificDataId      MarketSpecificDataType //1 A/N
	TransactionIdentifier     string                 //0, 15 A/N
	// This 15-character field can contain a Transaction Identifier (Visa, American Express, PayPal or
	// Discover) or Reference Number (MasterCard) (see Table 3.6 for record format and version
	// number). The POS device does not attempt to interpret the meaning of any data appearing in
	// this field. Data returned in this field is recorded and submitted as part of the data capture
	// settlement format.
	// For incremental authorization requests, the Transaction ID must be the same Transaction ID
	// returned in the original authorization response.
	// If “MAV” is in positions 5-7 of this field, the transaction should not be submitted for capture.

	ValidationCode string //0, 4 A/N
}

func NewG1AuthorizationMessageResponse(record RecordFormatType, apptype ApplicationIndicatorType, store uint, terminal uint, authsource AuthorizationSourceType) *G1AuthorizationMessageResponse {

	t := G1AuthorizationMessageResponse{RecordFormat: record, ApplicationType: apptype, StoreNumber: store, TerminalNumber: terminal, AuthorizationSourceCode: authsource}

	t.LocalTxnDateTime = time.Now()

	return &t
}

func (t *G1AuthorizationMessageResponse) UpdateResults(responseTime time.Time, responsecd ResponseCodeType, responseText string, approval string) {
	t.LocalTxnDateTime = responseTime
	t.ResponseCode = responsecd
	t.AuthorizationResponseText = responseText
	t.ApprovalCode = approval
}

func (t *G1AuthorizationMessageResponse) getTransactionIdentifier() string {
	if 0 < len(t.TransactionIdentifier) {
		return fmt.Sprintf("%15s", t.TransactionIdentifier)
	}
	return ""
}
func (t *G1AuthorizationMessageResponse) getValidationCode() string {
	if 0 < len(t.ValidationCode) {
		return fmt.Sprintf("%4s", t.ValidationCode)
	}
	return ""
}

func (t *G1AuthorizationMessageResponse) String() string {

	return fmt.Sprintf("%c%c%c%c%04d%04d%c%04d%2s%6s%02d%02d%02d%02d%02d%02d%16s%c%12s%c%s%c%s%c",
		t.RecordFormat,              //1 A/N 4.67
		t.ApplicationType,           //1 A/N 4.9
		t.MessageDelimiter,          //1 A/N 4.62
		t.ReturnedACI,               //1 A/N 4.62
		t.StoreNumber,               //4 NUM 4.81
		t.TerminalNumber,            //4 NUM 4.84
		t.AuthorizationSourceCode,   //1 A/N 4.34
		t.TransactionSequenceNumber, //4 NUM 4.92
		string(t.ResponseCode),      //2 A/N
		t.ApprovalCode,              //6 A/N

		t.LocalTxnDateTime.Month(), t.LocalTxnDateTime.Day(), t.LocalTxnDateTime.Year(), //MMDDYY"
		t.LocalTxnDateTime.Hour(), t.LocalTxnDateTime.Minute(), t.LocalTxnDateTime.Second(), //HHMMSS

		t.AuthorizationResponseText,  //16 A/N
		t.AVSResultCode,              //1 A/N
		t.RetrievalReferenceNumber,   //12 A/N
		t.MarketSpecificDataId,       //1 A/N
		t.getTransactionIdentifier(), //0,15 A/N
		bytes.FileSeparator,          //1 ASCII
		t.getValidationCode(),        //0,4 A/N
		bytes.FileSeparator)          //1 ASCII
}

func (t *G1AuthorizationMessageResponse) VerboseString() string {

	buffer := bytes.NewSafeBuffer(512)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", t.RecordFormat)                                                                                                 //1 A/N 4.67
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", t.ApplicationType)                                                                                           //1 A/N 4.9
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", t.MessageDelimiter)                                                                                                 //1 A/N 4.62
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "ReturnedACI", t.ReturnedACI)                                                                                                   //1 A/N 4.69
	buffer.AppendFormat("%-32s%04d\n", "StoreNumber", t.StoreNumber)                                                                                                         //4 NUM 4.81
	buffer.AppendFormat("%-32s%04d\n", "TerminalNumber", t.TerminalNumber)                                                                                                   //4 NUM 4.84
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AuthorizationSourceCode", t.AuthorizationSourceCode)                                                                           //1 A/N 4.1
	buffer.AppendFormat("%-32s%04d\n", "TransactionSequenceNumber", t.TransactionSequenceNumber)                                                                             //4 NUM 4.92
	buffer.AppendFormat("%-32s%s (%s)\n", "ResponseCode", string(t.ResponseCode), t.ResponseCode)                                                                            // 2 A/N
	buffer.AppendFormat("%-32s%s\n", "ApprovalCode", t.ApprovalCode)                                                                                                         // 6 A/N
	buffer.AppendFormat("%-32s%02d%02d%02d %s\n", "Local Transaction Date", t.LocalTxnDateTime.Month(), t.LocalTxnDateTime.Day(), t.LocalTxnDateTime.Year(), "(MMDDYY)")     //MMDDYY"
	buffer.AppendFormat("%-32s%02d%02d%02d %s\n", "Local Transaction Time", t.LocalTxnDateTime.Hour(), t.LocalTxnDateTime.Minute(), t.LocalTxnDateTime.Second(), "(HHMMSS)") //HHMMSS
	buffer.AppendFormat("%-32s%s\n", "AuthorizationResponseText", t.AuthorizationResponseText)                                                                               // 6 A/N
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AVSResultCode", t.AVSResultCode)                                                                                               //1 A/N 4.1
	buffer.AppendFormat("%-32s%s\n", "RetrievalReferenceNumber", t.RetrievalReferenceNumber)                                                                                 // 6 A/N
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "MarketSpecificDataId", t.MarketSpecificDataId)                                                                                 //1 A/N 4.1
	buffer.AppendFormat("%-32s%s\n", "TransactionIdentifier", t.TransactionIdentifier)                                                                                       // 6 A/N
	buffer.AppendFormat("%-32s%s\n", "ValidationCode", t.ValidationCode)                                                                                                     // 6 A/N

	return buffer.String()
}

func ParseG1AuthorizationMessageResponse(s string) (*G1AuthorizationMessageResponse, int) {
	t := G1AuthorizationMessageResponse{}

	// advance buffer scan position past any StartOfText or other control character artifacts left from serial communications
	// by locating first occurence of the STX
	i := 0
	if 67 > len(s) {
		panic("minimum EIS1080 Group 1 response message length is 67")
	}

	var mon, day, year, hour, min, sec int

	fmt.Sscanf(s[i:i+17], "%c%1d%c%c%04d%04d%c%04d",
		&t.RecordFormat,              //1 A/N 4.67
		&t.ApplicationType,           //1 A/N 4.9
		&t.MessageDelimiter,          //1 A/N 4.62
		&t.ReturnedACI,               //1 A/N 4.62
		&t.StoreNumber,               //4 NUM 4.81
		&t.TerminalNumber,            //4 NUM 4.84
		&t.AuthorizationSourceCode,   //1 A/N 4.34
		&t.TransactionSequenceNumber) //4 NUM 4.92

	//fmt.Printf("DEBUG --> %s\n", s)

	t.ResponseCode = ResponseCodeType(s[i+17 : i+19]) //2 A/N
	t.ApprovalCode = s[i+19 : i+25]                   //6 A/N

	fmt.Sscanf(s[i+25:i+37], "%02d%02d%02d%02d%02d%02d",
		&mon, &day, &year, //MMDDYY
		&hour, &min, &sec) //HHMMSS

	t.AuthorizationResponseText = s[i+37 : i+53]             //16 A/N
	t.AVSResultCode = AvsResultType(s[i+53])                 //1 A/N
	t.RetrievalReferenceNumber = s[i+54 : i+66]              //12 A/N
	t.MarketSpecificDataId = MarketSpecificDataType(s[i+66]) //1 A/N

	t.LocalTxnDateTime = time.Date(year, time.Month(mon), day, hour, min, sec, 0, time.Local)

	i = 67
	fields := strings.Split(s[i:], string(bytes.FileSeparator))
	for j := 0; j < 2 && j < len(fields); j++ {
		switch j {
		case 0:
			if 15 == len(fields[j]) {
				t.TransactionIdentifier = fields[j] //0,15 A/N
			}
		case 1:
			if 4 == len(fields[j]) {
				t.ValidationCode = fields[j] //0,4 A/N
			}
		}
		i += len(fields[j]) + 1
	}

	//t.ApprovalCode = strings.Replace(strings.Replace(t.ApprovalCode, "VITAL", "SOMEBANK", -1), "TAS", "SOMEBANK", -1)

	//t.AuthorizationResponseText = strings.Replace(strings.Replace(t.AuthorizationResponseText, "VITAL", "SOMEBANK", -1), "TAS", "SOMEBANK", -1)

	return &t, i
}

func IsApproved(code ResponseCodeType) bool {
	return "85" == code || "00" == code
}
