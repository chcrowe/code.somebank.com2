package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

type DebitEbtConfirmationRecord struct {
	RecordFormat           RecordFormatType         //1 A/N 4.67
	ApplicationType        ApplicationIndicatorType //1 A/N 4.9
	Delimiter              bytes.AsciiCharacterType //1 A/N 4.62
	AcquirerBIN            uint                     //6 NUM 4.2
	HostMessageIdentifier  string                   //8 A/N
	SystemTraceAuditNumber uint                     //6 NUM
}

func NewDebitEbtConfirmationRecord(Bin uint, HostMessageId string, SystemTraceNumber uint) *DebitEbtConfirmationRecord {
	return &DebitEbtConfirmationRecord{DebitEbtConfirmation, SingleAuthorizationPerConnection, bytes.MessageDelimiter, Bin, HostMessageId, SystemTraceNumber}
}

func (d *DebitEbtConfirmationRecord) String() string {
	return fmt.Sprintf("%c%c%c%06d%8s%06d", d.RecordFormat, d.ApplicationType, d.Delimiter,
		d.AcquirerBIN, d.HostMessageIdentifier, d.SystemTraceAuditNumber)
}

func ParseDebitEbtConfirmationRecord(s string) *DebitEbtConfirmationRecord {
	d := DebitEbtConfirmationRecord{DebitEbtConfirmation, SingleAuthorizationPerConnection, bytes.MessageDelimiter, 0, "", 0}

	if 23 > len(s) {
		panic("DebitEbtConfirmationRecord string must be at least 23 characters in length")
	}

	fmt.Sscanf(s, "%c%c%c%06d%8s%06d", &d.RecordFormat, &d.ApplicationType, &d.Delimiter, &d.AcquirerBIN, &d.HostMessageIdentifier, &d.SystemTraceAuditNumber)

	return &d
}
