package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
	//"strings"
	//"time"
)

type G2AuthorizationMessageResponse struct {
	AcquirerBIN            uint                      //6 A/N
	HostMessageIdentifier  string                    //8 A/N
	SystemTraceAuditNumber uint                      //6 NUM
	NetworkIDCode          NetworkIdentificationType //1 A/N
	//SettlementDate time.Time //4 NUM MMDD
	SettlementMonth uint
	SettlementDay   uint
	//<FS>
}

func (t *G2AuthorizationMessageResponse) String() string {
	return fmt.Sprintf("%06d%8s%06d%c%02d%02d%c",
		t.AcquirerBIN,            //6 NUM
		t.HostMessageIdentifier,  //8 A/N
		t.SystemTraceAuditNumber, //6 NUM
		t.NetworkIDCode,          //1 A/N
		t.SettlementMonth,        //2 NUM
		t.SettlementDay,          //2 NUM
		bytes.FileSeparator)
}

func ParseG2AuthorizationMessageResponse(s string) (*G2AuthorizationMessageResponse, int) {
	t := G2AuthorizationMessageResponse{}

	if len(s) >= 25 {
		fmt.Sscanf(s, "%06d", &t.AcquirerBIN) //6 NUM
		t.HostMessageIdentifier = s[6:14]     //8 A/N

		fmt.Sscanf(s[14:], "%06d%c%02d%02d",
			&t.SystemTraceAuditNumber, //6 NUM
			&t.NetworkIDCode,          //1 A/N
			&t.SettlementMonth,        //2 NUM
			&t.SettlementDay)          //2 NUM

	}
	i := 26

	return &t, i
}
