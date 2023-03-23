package tsys

import (
	"fmt"
	"strconv"
	"time"

	"code.somebank.com/p/bytes"
)

type ReversalAndCancelDataI struct {
	ApprovalCode         string    //6 A/N returned in the original authorization response record
	LocalTransactionDate time.Time //6 NUM MMDDYY returned in the original authorization response record
	//LocalTransactionTime Time  //6 NUM HHMMSS returned in the original authorization response record
	RetrievalReferenceNumber uint64 //12 NUM returned in the original authorization response record
}

func NewReversalAndCancelDataI(approval string, tm time.Time, retrievalno uint64) *ReversalAndCancelDataI {
	return &ReversalAndCancelDataI{approval, tm, retrievalno}
}

func (rd *ReversalAndCancelDataI) String() string {
	//if 0 >= len(rd.ApprovalCode) {
	//	return ""
	//}

	return fmt.Sprintf("%6s%02d%02d%02d%02d%02d%02d%012d", rd.ApprovalCode, rd.LocalTransactionDate.Month(), rd.LocalTransactionDate.Day(),
		rd.LocalTransactionDate.Year(), rd.LocalTransactionDate.Hour(), rd.LocalTransactionDate.Minute(), rd.LocalTransactionDate.Second(),
		rd.RetrievalReferenceNumber)
}

func (rd *ReversalAndCancelDataI) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	buffer.AppendFormat("%8s%-32s\"%s\"\n", " ", "ApprovalCode", rd.ApprovalCode)
	buffer.AppendFormat("%8s%-32s%02d%02d%02d (MMDDYY)\n", " ", "LocalTransactionDate", rd.LocalTransactionDate.Month(), rd.LocalTransactionDate.Day(), rd.LocalTransactionDate.Year())
	buffer.AppendFormat("%8s%-32s%02d%02d%02d (HHMMSS)\n", " ", "LocalTransactionTime", rd.LocalTransactionDate.Hour(), rd.LocalTransactionDate.Minute(), rd.LocalTransactionDate.Second())
	buffer.AppendFormat("%8s%-32s%012d\n", " ", "RetrievalReferenceNumber", rd.RetrievalReferenceNumber)

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}

func (rd *ReversalAndCancelDataI) Copy(reversal *ReversalAndCancelDataI) {
	rd.ApprovalCode = reversal.ApprovalCode
	rd.LocalTransactionDate = reversal.LocalTransactionDate
	rd.RetrievalReferenceNumber = reversal.RetrievalReferenceNumber
}

func ParseReversalAndCancelDataI(s string) *ReversalAndCancelDataI {

	rd := ReversalAndCancelDataI{"", time.Now(), 0}

	if 18 > len(s) {
		return &rd
		//panic("ReversalAndCancelDataI must be at least 18 characters in length")
	}

	rd.ApprovalCode = s[0:6]

	var mon, day, year, hour, min, sec int
	fmt.Sscanf(s[6:18], "%02d%02d%02d%02d%02d%02d", &mon, &day, &year, &hour, &min, &sec)

	rd.LocalTransactionDate = time.Date(year, time.Month(mon), day, hour, min, sec, 0, time.Local)

	if 30 == len(s) {
		rd.RetrievalReferenceNumber, _ = strconv.ParseUint(s[18:30], 10, 64)
	}

	return &rd
}

type ReversalAndCancelDataII struct {
	SystemTraceAuditNumber    uint64                    //6 NUM
	NetworkIdentificationCode NetworkIdentificationType //1 A/N 4.135
}

func NewReversalAndCancelDataII(tracenumber uint64, netidcode NetworkIdentificationType) *ReversalAndCancelDataII {
	return &ReversalAndCancelDataII{tracenumber, netidcode}
}

func (rd *ReversalAndCancelDataII) Copy(reversal *ReversalAndCancelDataII) {
	if nil != reversal {
		rd.SystemTraceAuditNumber = reversal.SystemTraceAuditNumber
		rd.NetworkIdentificationCode = reversal.NetworkIdentificationCode
	}
}

func (rd *ReversalAndCancelDataII) String() string {
	if 0 == rd.SystemTraceAuditNumber {
		return ""
	}
	return fmt.Sprintf("%06d%c", rd.SystemTraceAuditNumber, rd.NetworkIdentificationCode)
}

func (rd *ReversalAndCancelDataII) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	buffer.AppendFormat("%8s%-32s%06d\n", " ", "SystemTraceAuditNumber", rd.SystemTraceAuditNumber)
	buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "NetworkIdentificationCode", rd.NetworkIdentificationCode)

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}

func ParseReversalAndCancelDataII(s string) *ReversalAndCancelDataII {

	rd := ReversalAndCancelDataII{0, SpaceOrEmptyNetworkId}

	if 7 > len(s) {
		return &rd
		//panic("ReversalAndCancelDataII must be at least 7 characters in length")
	}

	fmt.Sscanf(s[0:6], "%06d%c", &rd.SystemTraceAuditNumber, &rd.NetworkIdentificationCode)

	return &rd
}
