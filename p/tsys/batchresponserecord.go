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

type BatchResponseRecord struct {
	RecordFormat     RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType  ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId        rune                     // 4 1 A/N X.25 Routing ID Z 4.341
	RecordType       *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243

	BatchRecordCount  uint                  //9 NUM 4.42 Batch Record Count
	BatchNetDeposit   uint64                //16 NUM 4.40 Batch Net Deposit
	BatchResponseCode BatchResponseCodeType //2 A/N 4.43 Batch Response Code - GB
	//2 NUM 4.125 Filler
	BatchNumber uint //3 NUM 4.41 Batch Number

	//GB
	BatchResponseText string //9 A/N 4.44 Batch Response Text
	//16 A/N 4.124 Filler

	//RB
	ErrorType                 string //1 A/N 4.116 Error Type
	ErrorRecordSequenceNumber uint   //4 NUM 4.114 Error Record Sequence Number
	ErrorRecordType           string //1 A/N 4.115 Error Record Type
	ErrorDataFieldNumber      uint   //2 NUM 4.113 Error Data Field Number
	ErrorData                 string //32 A/N 4.112 Error Data

	//QD
	BatchTransmissionDate uint // 4 A/N 4.45 Batch Transmission Date - MMDD
	// 21 A/N 4.124 Filler
}

func NewBatchResponseRecord(recordcount uint,
	netdeposit uint64,
	responsecode BatchResponseCodeType,
	batchnum uint,
	responsetext string) *BatchResponseRecord {

	b := BatchResponseRecord{}

	b.RecordFormat = Settlement
	b.ApplicationType = SingleBatchPerConnection
	b.MessageDelimiter = bytes.MessageDelimiter
	b.RoutingId = 'Z'
	b.RecordType = NewOptionalDataGroupMap(TrailerResponseRecord)

	b.BatchRecordCount = recordcount
	b.BatchNetDeposit = netdeposit
	b.BatchResponseCode = responsecode
	b.BatchNumber = batchnum
	b.BatchResponseText = responsetext

	return &b
}

func (b *BatchResponseRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", b.RecordFormat)     //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%d", b.ApplicationType)  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", b.MessageDelimiter) //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", b.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%s", b.RecordType)       //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%09d", b.BatchRecordCount)          // 9 NUM 4.42 Batch Record Count
	buffer.AppendFormat("%016d", b.BatchNetDeposit)          // 16 NUM 4.40 Batch Net Deposit
	buffer.AppendFormat("%-2s", string(b.BatchResponseCode)) //2 A/N 4.43 Batch Response Code - GB
	buffer.AppendFormat("%02d", 0)                           //2 NUM 4.125 Filler
	buffer.AppendFormat("%03d", b.BatchNumber)               // 3 NUM 4.41 Batch Number (001-999)

	switch b.BatchResponseCode {
	case GoodBatch:
		buffer.AppendFormat("%-9s", b.BatchResponseText) //9 A/N 4.44 Batch Response Text
		buffer.AppendFormat("%-16s", "")                 //16 A/N 4.124 Filler

	case RejectedBatch:
		buffer.AppendFormat("%-1s", b.ErrorType)                 //1 A/N 4.116 Error Type
		buffer.AppendFormat("%04d", b.ErrorRecordSequenceNumber) //4 NUM 4.114 Error Record Sequence Number
		buffer.AppendFormat("%-1s", b.ErrorRecordType)           //1 A/N 4.115 Error Record Type
		buffer.AppendFormat("%02d", b.ErrorDataFieldNumber)      //2 NUM 4.113 Error Data Field Number
		buffer.AppendFormat("%-32s", b.ErrorData)                //32 A/N 4.112 Error Data

	case DuplicateBatch:
		buffer.AppendFormat("%04d", b.BatchTransmissionDate) // 4 A/N 4.45 Batch Transmission Date - MMDD
		buffer.AppendFormat("%-21s", "")                     // 21 A/N 4.124 Filler
	}
}

func (b *BatchResponseRecord) String() string {
	buffer := bytes.NewSafeBuffer(128)
	b.CopyToBuffer(buffer)
	return buffer.String()
}

func (b *BatchResponseRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", b.RecordFormat)       //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", b.ApplicationType) //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", b.MessageDelimiter)       //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", b.RoutingId)            // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", b.RecordType)                 //5-9 5 A/N Record Type H@@@@ 4.243

	buffer.AppendFormat("%-32s%09d\n", "BatchRecordCount", b.BatchRecordCount)                                     // 9 NUM 4.42 Batch Record Count
	buffer.AppendFormat("%-32s%016d\n", "BatchNetDeposit", b.BatchNetDeposit)                                      // 16 NUM 4.40 Batch Net Deposit
	buffer.AppendFormat("%-32s%-2s (%s)\n", "BatchResponseCode", string(b.BatchResponseCode), b.BatchResponseCode) //2 A/N 4.43 Batch Response Code - GB
	buffer.AppendFormat("%-32s%02d\n", "Filler", 0)                                                                //2 NUM 4.125 Filler
	buffer.AppendFormat("%-32s%03d\n", "BatchNumber", b.BatchNumber)                                               // 3 NUM 4.41 Batch Number (001-999)

	switch b.BatchResponseCode {
	case GoodBatch:
		buffer.AppendFormat("%-32s%-9s\n", "BatchResponseText", b.BatchResponseText) //9 A/N 4.44 Batch Response Text
		buffer.AppendFormat("%-32s%-16s\n", "Filler", "")                            //16 A/N 4.124 Filler

	case RejectedBatch:
		buffer.AppendFormat("%-32s%-1s", "ErrorType", b.ErrorType)                                   //1 A/N 4.116 Error Type
		buffer.AppendFormat("%-32s%04d\n", "ErrorRecordSequenceNumber", b.ErrorRecordSequenceNumber) //4 NUM 4.114 Error Record Sequence Number
		buffer.AppendFormat("%-32s%-1s\n", "ErrorRecordType", b.ErrorRecordType)                     //1 A/N 4.115 Error Record Type
		buffer.AppendFormat("%-32s%02d\n", "ErrorDataFieldNumber", b.ErrorDataFieldNumber)           //2 NUM 4.113 Error Data Field Number
		buffer.AppendFormat("%-32s%-32s\n", "ErrorData", b.ErrorData)                                //32 A/N 4.112 Error Data

	case DuplicateBatch:
		buffer.AppendFormat("%-32s%04d\n", "BatchTransmissionDate", b.BatchTransmissionDate) // 4 A/N 4.45 Batch Transmission Date - MMDD
		buffer.AppendFormat("%-32s%-21s\n", "Filler", "")                                    // 21 A/N 4.124 Filler
	}

	return buffer.String()
}

func ParseBatchResponseRecord(s string) (*BatchResponseRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	b := BatchResponseRecord{}

	if 73 > len(s) {
		panic("minimum EIS1081 Batch Trailer Record length is 73")
	}

	fmt.Sscanf(s[0:4], "%c%1d%c%c",
		&b.RecordFormat,     //1 1 A/N Record Format K 4.242
		&b.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&b.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&b.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	b.RecordType = ParseOptionalDataGroups(s[4:9]) //5-9 5 A/N Record Type H@@@@ 4.243

	fmt.Sscanf(s[9:18], "%09d", &b.BatchRecordCount)      // 9 NUM 4.42 Batch Record Count
	fmt.Sscanf(s[18:34], "%016d", &b.BatchNetDeposit)     // 16 NUM 4.40 Batch Net Deposit
	b.BatchResponseCode = BatchResponseCodeType(s[34:36]) //2 A/N 4.43 Batch Response Code - GB
	//2 NUM 4.125 Filler
	fmt.Sscanf(s[38:41], "%03d", &b.BatchNumber) // 3 NUM 4.41 Batch Number (001-999)

	i := 41

	switch b.BatchResponseCode {
	case GoodBatch:
		b.BatchResponseText = s[41:50] //9 A/N 4.44 Batch Response Text
		//16 A/N 4.124 Filler
		i = 66

	case RejectedBatch:
		b.ErrorType = s[41:42]                                     //1 A/N 4.116 Error Type
		fmt.Sscanf(s[42:46], "%04d", &b.ErrorRecordSequenceNumber) //4 NUM 4.114 Error Record Sequence Number
		b.ErrorRecordType = s[46:47]                               //1 A/N 4.115 Error Record Type
		fmt.Sscanf(s[47:49], "%02d", &b.ErrorDataFieldNumber)      //2 NUM 4.113 Error Data Field Number
		b.ErrorData = s[49:81]                                     //32 A/N 4.112 Error Data
		i = 81

	case DuplicateBatch:
		fmt.Sscanf(s[41:45], "%04d", &b.BatchTransmissionDate) // 4 A/N 4.45 Batch Transmission Date - MMDD
		// 21 A/N 4.124 Filler
		i = 66
	}

	return &b, i
}
