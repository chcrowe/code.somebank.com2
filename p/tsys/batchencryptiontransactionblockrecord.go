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

type BatchEncryptionTransactionBlockRecord struct {
	RecordFormat               RecordFormatType         //1 1 A/N Record Format K 4.242
	ApplicationType            ApplicationIndicatorType //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	MessageDelimiter           bytes.AsciiCharacterType //3 1 A/N Message Delimiter . 4.196
	RoutingId                  rune                     // 4 1 A/N X.25 Routing ID Z 4.341
	RecordType                 *OptionalDataGroupMap    //5-9 5 A/N Record Type H@@@@ 4.243
	EncryptionType             rune                     //10 1 A/N Encryption Type V 4.111
	EncryptionDataSize         int                      //11-15 5 NUM Encryption Data Size 264 4.110
	EncryptedTransmissionBlock string                   //- VAR A/N Encrypted Transmission Block 4.109
}

func NewBatchEncryptionTransactionBlockRecord(encryptedBlock string) *BatchEncryptionTransactionBlockRecord {

	h := BatchEncryptionTransactionBlockRecord{}

	h.RecordFormat = Settlement
	h.ApplicationType = SingleBatchPerConnection
	h.MessageDelimiter = bytes.MessageDelimiter
	h.RoutingId = 'Z'
	h.RecordType = NewOptionalDataGroupMap(EncryptedTransmissionBlockRecord)
	h.EncryptionType = 'V' //Voltage Encryption
	h.EncryptionDataSize = 264
	h.EncryptedTransmissionBlock = encryptedBlock

	return &h
}

func (h *BatchEncryptionTransactionBlockRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", h.RecordFormat)               //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%d", h.ApplicationType)            //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%c", h.MessageDelimiter)           //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%c", h.RoutingId)                  // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%5s", h.RecordType)                //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%c", h.EncryptionType)             //10 1 A/N Encryption Type V 4.111
	buffer.AppendFormat("%05d", h.EncryptionDataSize)       //11-15 5 NUM Encryption Data Size 264 4.110
	buffer.AppendFormat("%c", h.EncryptedTransmissionBlock) //- VAR A/N Encrypted Transmission Block 4.109
}

func (h *BatchEncryptionTransactionBlockRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	h.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (h *BatchEncryptionTransactionBlockRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", h.RecordFormat)                     //1 1 A/N Record Format K 4.242
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", h.ApplicationType)               //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", h.MessageDelimiter)                     //3 1 A/N Message Delimiter . 4.196
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", h.RoutingId)                          // 4 1 A/N X.25 Routing ID Z 4.341
	buffer.AppendFormat("%-32s%-5s\n", "RecordType", h.RecordType)                               //5-9 5 A/N Record Type H@@@@ 4.243
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "EncryptionType", h.EncryptionType)                 //10 1 A/N Encryption Type V 4.111
	buffer.AppendFormat("%-32s%05d (%[2]s)\n", "EncryptionDataSize", h.EncryptionDataSize)       //11-15 5 NUM Encryption Data Size 264 4.110
	buffer.AppendFormat("%-32s%s\n", "EncryptedTransmissionBlock", h.EncryptedTransmissionBlock) //- VAR A/N Encrypted Transmission Block 4.109

	return buffer.String()
}

func ParseBatchEncryptionTransactionBlockRecord(s string) (*BatchEncryptionTransactionBlockRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	h := BatchEncryptionTransactionBlockRecord{}

	i := 0
	if 15 > len(s) {
		panic("minimum EIS1081 Batch Parameter Record length is 65")
	}
	//i++ //advance past the STX if not already removed

	fmt.Sscanf(s[i:i+4], "%c%1d%c%c",
		&h.RecordFormat,     //1 1 A/N Record Format K 4.242
		&h.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
		&h.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
		&h.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341
	i += 4

	h.RecordType = ParseOptionalDataGroups(s[i : i+5]) //5-9 5 A/N Record Type H@@@@ 4.243
	i += 5

	fmt.Sscanf(s[i:i+6], "%c%05d",
		&h.EncryptionType,     //10 1 A/N Encryption Type V 4.111
		&h.EncryptionDataSize) //11-15 5 NUM Encryption Data Size 264 4.110
	i += 6

	h.EncryptedTransmissionBlock = s[i : i+h.EncryptionDataSize]
	i += h.EncryptionDataSize

	return &h, i
}
