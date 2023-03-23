package tsys

import (
	"fmt"
	"strconv"
	"strings"

	"code.somebank.com/p/bytes"
)

type G2AuthorizationMessageRequest struct {
	ReversalAndCancelData *ReversalAndCancelDataII //0, 7 A/N

	// (SharingGroup) This one to 30-character field contains a listing of direct debit and EBT networks that a POS
	//device can access. This field should be configured as a parameter
	SharingGroup string //NetworkIdentificationType.SpaceOrEmpty//1-30 A/N
	// (SharingGroup) a value of “V” or “5” in this field returned in the response indicates
	// that the transaction was processed as a credit card purchase transaction (Transaction Code 54)
	// even though it was originally submitted as a debit transaction. When these values (“V” or “5”)
	// are returned, the transaction must be processed and settled as a credit card (non-debit)
	// transaction. Please note that this does not apply to “R3” type transactions

	MerchantABANumber         uint64                     //0, 9 NUM
	MerchantSettlementAgentNo string                     //0, 4 A/N
	AgentBankNumber           uint64                     //6 NUM
	AgentChainNumber          uint64                     //6 NUM
	BatchNumber               uint                       //3 NUM
	ReimbursementAttribute    ReimbursementAttributeType //1 A/N
	OriginalPurchaseData      string                     //0, 12 A/N YYMMDDhhmmss
}

func NewG2AuthorizationMessageRequest(reversalData *ReversalAndCancelDataII,
	sharegroup string,
	abanumber uint64,
	settlementagent string,
	agentbank uint64,
	agentchain uint64,
	batchnum uint,
	reimburseattr ReimbursementAttributeType,
	purchasedata string) *G2AuthorizationMessageRequest {

	t := G2AuthorizationMessageRequest{}

	// ABA: 990000408	(US to Vital)
	// Settlement Agent: V040
	// Sharing Group:	VNGK7F3EHYL8MIQZW

	t.ReversalAndCancelData = reversalData

	if 0 < len(sharegroup) {
		t.SharingGroup = sharegroup
	} else {
		t.SharingGroup = "VNGK7F3EHYL8MIQZW"
	}

	t.MerchantABANumber = abanumber

	if 0 < len(settlementagent) {
		t.MerchantSettlementAgentNo = settlementagent
	} else {
		t.MerchantSettlementAgentNo = "V040"
	}

	t.AgentBankNumber = agentbank
	t.AgentChainNumber = agentchain
	t.BatchNumber = batchnum
	t.ReimbursementAttribute = reimburseattr
	t.OriginalPurchaseData = purchasedata

	return &t
}

func (t *G2AuthorizationMessageRequest) String() string {

	s := fmt.Sprintf("%c%s%c%09d%s%c%06d%06d%03d%c%s%c",
		bytes.FileSeparator,         //1 A/N
		t.SharingGroup,              //1-30 A/N
		bytes.FileSeparator,         //1 A/N
		t.MerchantABANumber,         //0, 9 NUM
		t.MerchantSettlementAgentNo, //0, 4 A/N
		bytes.FileSeparator,         //1 A/N
		t.AgentBankNumber,           //6 NUM
		t.AgentChainNumber,          //6 NUM
		t.BatchNumber,               //3 NUM
		t.ReimbursementAttribute,    //1 A/N
		t.OriginalPurchaseData,      //don't send for POS Check
		bytes.FileSeparator)         //1 A/N

	if 0 < t.ReversalAndCancelData.SystemTraceAuditNumber { //Don't send for POS Check
		return t.ReversalAndCancelData.String() + s //0, 7 A/N
	}

	return s
}

func (t *G2AuthorizationMessageRequest) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)

	var sharinggroups string = ""
	var netidtype NetworkIdentificationType
	for i := 0; i < len(t.SharingGroup); i++ {
		netidtype = NetworkIdentificationType(t.SharingGroup[i])
		if 0 == i {
			sharinggroups = fmt.Sprintf("%c=%[1]s", netidtype)
		} else {
			sharinggroups = sharinggroups + ", " + fmt.Sprintf("%c=%[1]s", netidtype)
		}
	}

	buffer.AppendFormat("%-32s%s %s\n", "ReversalAndCancelDataII", t.ReversalAndCancelData, t.ReversalAndCancelData.VerboseString()) //1-30 A/N

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)
	buffer.AppendFormat("%-32s%s (%s)\n", "SharingGroup", t.SharingGroup, sharinggroups) //1-30 A/N
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)
	buffer.AppendFormat("%-32s%09d\n", "MerchantABANumber", t.MerchantABANumber)               //0, 9 NUM
	buffer.AppendFormat("%-32s%s\n", "MerchantSettlementAgentNo", t.MerchantSettlementAgentNo) //0, 4 A/N
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)
	buffer.AppendFormat("%-32s%06d\n", "AgentBankNumber", t.AgentBankNumber)                     //6 NUM
	buffer.AppendFormat("%-32s%06d\n", "AgentChainNumber", t.AgentChainNumber)                   //6 NUM
	buffer.AppendFormat("%-32s%03d\n", "BatchNumber", t.BatchNumber)                             //3 NUM
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "ReimbursementAttribute", t.ReimbursementAttribute) //1 A/N
	buffer.AppendFormat("%-32s%s\n", "OriginalPurchaseData", t.OriginalPurchaseData)             //0, 12 A/N YYMMDDhhmmss
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	return buffer.String()
}

func ParseG2AuthorizationMessageRequest(s string) (*G2AuthorizationMessageRequest, int) {

	t := G2AuthorizationMessageRequest{}
	i := 0

	fields := strings.Split(s, string(bytes.FileSeparator))

	for j := 0; j < 4 && j < len(fields); j++ {
		fieldlen := len(fields[j])
		switch j {
		case 0:
			if 7 == fieldlen {
				t.ReversalAndCancelData = ParseReversalAndCancelDataII(fields[j])
			}
		case 1:
			t.SharingGroup = fields[j]
		case 2:
			if 9 <= fieldlen {
				t.MerchantABANumber, _ = strconv.ParseUint(fields[j][0:9], 10, 64)
			}
			if 13 <= fieldlen {
				t.MerchantSettlementAgentNo = fields[j][9:13]
			}
		case 3:
			if 16 <= fieldlen {
				fmt.Sscanf(fields[j], "%06d%06d%03d%c", &t.AgentBankNumber, &t.AgentChainNumber, &t.BatchNumber, &t.ReimbursementAttribute)
			}
			if 28 <= fieldlen {
				t.OriginalPurchaseData = fields[j][16:28]
			}
		}
		i += fieldlen + 1 //in addition to field length, add a "1" to account for each field separator
	}

	return &t, i
}
