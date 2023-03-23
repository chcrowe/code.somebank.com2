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

type BatchTrailerRecord struct {
	RecordFormat     RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType  ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId        rune                     // 4 1 A/N X.25 Routing ID Z 4.341
	RecordType       *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243

	BatchTransmissionDate uint   // 4 NUM 4.45 Batch Transmission Date (MMDD)
	BatchNumber           uint   // 3 NUM 4.41 Batch Number (001-999)
	BatchRecordCount      uint   // 9 NUM 4.42 Batch Record Count
	BatchHashingTotal     uint64 // 16 NUM 4.39 Batch Hashing Total
	CashBackTotal         uint64 // 16 NUM 4.64 Cash Back Total
	BatchNetDeposit       uint64 // 16 NUM 4.40 Batch Net Deposit
}

func NewBatchTrailerRecord(batchdt uint,
	batchno uint,
	recordCount uint,
	grossdeposit uint64,
	cashback uint64,
	netdeposit uint64) *BatchTrailerRecord {

	h := BatchTrailerRecord{}

	h.RecordFormat = Settlement
	h.ApplicationType = SingleBatchPerConnection
	h.MessageDelimiter = bytes.MessageDelimiter
	h.RoutingId = 'Z'
	h.RecordType = NewOptionalDataGroupMap(TrailerRecord)

	return &h
}

func (h *BatchTrailerRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", h.RecordFormat)     //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%d", h.ApplicationType)  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", h.MessageDelimiter) //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", h.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", h.RecordType)       //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%04d", h.BatchTransmissionDate) // 4 NUM 4.45 Batch Transmission Date (MMDD)
	buffer.AppendFormat("%03d", h.BatchNumber)           // 3 NUM 4.41 Batch Number (001-999)
	buffer.AppendFormat("%09d", h.BatchRecordCount)      // 9 NUM 4.42 Batch Record Count
	buffer.AppendFormat("%016d", h.BatchHashingTotal)    // 16 NUM 4.39 Batch Hashing Total
	buffer.AppendFormat("%016d", h.CashBackTotal)        // 16 NUM 4.64 Cash Back Total
	buffer.AppendFormat("%016d", h.BatchNetDeposit)      // 16 NUM 4.40 Batch Net Deposit
}

func (h *BatchTrailerRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	h.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (h *BatchTrailerRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", h.RecordFormat)       //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", h.ApplicationType) //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", h.MessageDelimiter)       //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", h.RoutingId)            // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", h.RecordType)                 //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%-32s%04d\n", "BatchTransmissionDate", h.BatchTransmissionDate) // 4 NUM 4.45 Batch Transmission Date (MMDD)
	buffer.AppendFormat("%-32s%03d\n", "BatchNumber", h.BatchNumber)                     // 3 NUM 4.41 Batch Number (001-999)
	buffer.AppendFormat("%-32s%09d\n", "BatchRecordCount", h.BatchRecordCount)           // 9 NUM 4.42 Batch Record Count
	buffer.AppendFormat("%-32s%016d\n", "BatchHashingTotal", h.BatchHashingTotal)        // 16 NUM 4.39 Batch Hashing Total
	buffer.AppendFormat("%-32s%016d\n", "CashBackTotal", h.CashBackTotal)                // 16 NUM 4.64 Cash Back Total
	buffer.AppendFormat("%-32s%016d\n", "BatchNetDeposit", h.BatchNetDeposit)            // 16 NUM 4.40 Batch Net Deposit

	return buffer.String()
}

func ParseBatchTrailerRecord(s string) (*BatchTrailerRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	h := BatchTrailerRecord{}

	if 73 > len(s) {
		panic("minimum EIS1081 Batch Trailer Record length is 73")
	}

	fmt.Sscanf(s[0:4], "%c%1d%c%c",
		&h.RecordFormat,     //1 1 A/N Record Format K 4.242
		&h.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&h.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&h.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	h.RecordType = ParseOptionalDataGroups(s[4:9]) //5-9 5 A/N Record Type H@@@@ 4.243

	fmt.Sscanf(s[9:73], "%04d%03d%09d%016d%016d%016d",
		&h.BatchTransmissionDate, // 4 NUM 4.45 Batch Transmission Date (MMDD)
		&h.BatchNumber,           // 3 NUM 4.41 Batch Number (001-999)
		&h.BatchRecordCount,      // 9 NUM 4.42 Batch Record Count
		&h.BatchHashingTotal,     // 16 NUM 4.39 Batch Hashing Total
		&h.CashBackTotal,         // 16 NUM 4.64 Cash Back Total
		&h.BatchNetDeposit)       // 16 NUM 4.40 Batch Net Deposit

	i := 73

	return &h, i
}
