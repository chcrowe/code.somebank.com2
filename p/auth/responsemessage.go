package auth

import (
	"code.somebank.com/p/bytes"
	"code.somebank.com/p/list"

	//"code.somebank.com/p/tsys"
	"encoding/xml"
	//"fmt"
	"strings"
)

type ResponseMessage struct {
	VerifiedAmount string `xml:"VerifiedAmount"`
	ErrorCode      string `xml:"ErrorCode"`
	ErrorText      string `xml:"ErrorText"`

	ApprovalCode              string `xml:"ApprovalCode"`
	ResponseCode              string `xml:"ResponseCode"`
	ResponseText              string `xml:"ResponseText"`
	CVVResultCode             string `xml:"CVVResultCode"`
	AVSResultCode             string `xml:"AVSResultCode"`
	TransactionSequenceNumber string `xml:"TransactionSequenceNumber"`
	TransactionID             string `xml:"TransactionID"`
	LocalTransactionDate      string `xml:"LocalTransactionDate"`
	LocalTransactionTime      string `xml:"LocalTransactionTime"`
	RetrievalReferenceNumber  string `xml:"RetrievalReferenceNumber"`

	MerchantNumber   string `xml:"MerchantNumber"`
	TerminalIDNumber string `xml:"TerminalIDNumber"`

	CardType       string `xml:"CardType"`
	NameOnCard     string `xml:"NameOnCard"`
	AccountNumber  string `xml:"AccountNumber"`
	ExpirationDate string `xml:"ExpirationDate"`

	PurchaseIdentifier1           string `xml:"PurchaseIdentifier1"`
	PurchaseIdentifier2           string `xml:"PurchaseIdentifier2"`
	OperatorID                    string `xml:"OperatorID"`
	CustomerServicePhoneNumber    string `xml:"CustomerServicePhoneNumber"`
	VoiceAuthorizationPhoneNumber string `xml:"VoiceAuthorizationPhoneNumber"`

	TaxAmount           string `xml:"TaxAmount"`
	TaxPercentage       string `xml:"TaxPercentage"`
	SurchargeAmount     string `xml:"SurchargeAmount"`
	SurchargePercentage string `xml:"SurchargePercentage"`

	TransactionStatus string `xml:"TransactionStatus"`
	MessageFromHost   string `xml:"MessageFromHost"`
	TransactionType   string `xml:"TransactionType"`
}

// func UnmarshallFrom1080Request(rm *ResponseMessage, g1 *tsys.G1AuthorizationMessageRequest, g2 *tsys.G2AuthorizationMessageRequest, g3array []string) {
// 	rm.MerchantNumber = fmt.Sprintf("%d", g1.MerchantNumber)
// 	rm.TerminalIDNumber = fmt.Sprintf("%d", g1.TerminalNumber)

// }

// func UnmarshallFrom1080Response(rm *ResponseMessage, g1 *tsys.G1AuthorizationMessageResponse, g2 *tsys.G2AuthorizationMessageResponse, g3array []string) {
// 	rm.ApprovalCode = g1.ApprovalCode
// 	rm.ResponseCode = string(g1.ResponseCode)
// 	rm.ResponseText = g1.AuthorizationResponseText
// 	rm.TransactionSequenceNumber = fmt.Sprintf("%d", g1.TransactionSequenceNumber)
// 	rm.TransactionID = g1.TransactionIdentifier
// 	rm.LocalTransactionDate = fmt.Sprintf("%02d%02d%02d", g1.LocalTxnDateTime.Month(), g1.LocalTxnDateTime.Day(), g1.LocalTxnDateTime.Year())
// 	rm.LocalTransactionTime = fmt.Sprintf("%02d%02d%02d", g1.LocalTxnDateTime.Hour(), g1.LocalTxnDateTime.Minute(), g1.LocalTxnDateTime.Second())
// 	rm.AVSResultCode = string(g1.AVSResultCode)
// 	rm.RetrievalReferenceNumber = g1.RetrievalReferenceNumber
// }

func (r *ResponseMessage) Xml() string {
	b, _ := xml.Marshal(r)
	return string(b)
}

func (r *ResponseMessage) Values() string {

	m := list.NameValuePairMap{}

	m.Set("VerifiedAmount", r.VerifiedAmount)
	m.Set("ErrorCode", r.ErrorCode)
	m.Set("ErrorText", r.ErrorText)

	m.Set("ApprovalCode", r.ApprovalCode)
	m.Set("ResponseCode", r.ResponseCode)
	m.Set("ResponseText", r.ResponseText)
	m.Set("CVVResultCode", r.CVVResultCode)
	m.Set("AVSResultCode", r.AVSResultCode)
	m.Set("TransactionSequenceNumber", r.TransactionSequenceNumber)
	m.Set("TransactionID", r.TransactionID)
	m.Set("LocalTransactionDate", r.LocalTransactionDate)
	m.Set("LocalTransactionTime", r.LocalTransactionTime)
	m.Set("RetrievalReferenceNumber", r.RetrievalReferenceNumber)

	m.Set("MerchantNumber", r.MerchantNumber)
	m.Set("TerminalIDNumber", r.TerminalIDNumber)

	m.Set("CardType", r.CardType)
	m.Set("NameOnCard", r.NameOnCard)
	m.Set("AccountNumber", r.AccountNumber)
	m.Set("ExpirationDate", r.ExpirationDate)

	m.Set("PurchaseIdentifier1", r.PurchaseIdentifier1)
	m.Set("PurchaseIdentifier2", r.PurchaseIdentifier2)
	m.Set("OperatorID", r.OperatorID)
	m.Set("CustomerServicePhoneNumber", r.CustomerServicePhoneNumber)
	m.Set("VoiceAuthorizationPhoneNumber", r.VoiceAuthorizationPhoneNumber)

	m.Set("TaxAmount", r.TaxAmount)
	m.Set("TaxPercentage", r.TaxPercentage)
	m.Set("SurchargeAmount", r.SurchargeAmount)
	m.Set("SurchargePercentage", r.SurchargePercentage)

	m.Set("TransactionStatus", r.TransactionStatus)
	m.Set("MessageFromHost", r.MessageFromHost)
	m.Set("TransactionType", r.TransactionType)

	return m.String()
}

func (rsp *ResponseMessage) CreatePOSCP4XmlMiniResponse(action string, req *RequestMessage) string {
	sb := bytes.NewSafeBuffer(1024)
	sb.AppendFormat("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	sb.AppendFormat("<responsemessage>\n")

	switch action {
	case "GET_CARD_INFO":
		sb.AppendFormat("<td>%s</td>\n", req.TrackData)
		sb.AppendFormat("<noc>%s</noc>\n", req.NameOnCard)
		sb.AppendFormat("<an>%s</an>\n", req.AccountNumber)
		sb.AppendFormat("<ct>%s</ct>\n", req.CardType)
		sb.AppendFormat("<ed>%s</ed>\n", req.ExpirationDate)
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)
		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)

	case "GET_HOSTMESSAGES":
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)

	case "GET_MERCHANT_INFO":

		// sb.AppendFormat("<mn>%s</mn>\n", MerchantName)
		// sb.AppendFormat("<mnb>%s</mnb>\n", rsp.MerchantNumber)
		// sb.AppendFormat("<murl>%s</murl>\n", MerchantURL)
		// sb.AppendFormat("<mem>%s</mem>\n", MerchantEMail)
		// sb.AppendFormat("<ms1>%s</ms1>\n", MerchantStreet1)
		// sb.AppendFormat("<ms2>%s</ms2>\n", MerchantStreet2)
		// sb.AppendFormat("<mcty>%s</mcty>\n", MerchantCity)
		// sb.AppendFormat("<ms>%s</ms>\n", MerchantState)
		// sb.AppendFormat("<mz>%s</mz>\n", MerchantZip)
		// sb.AppendFormat("<mph>%s</mph>\n", MerchantPhone)
		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)
		sb.AppendFormat("<rl1>%s</rl1>\n", "")
		sb.AppendFormat("<rl2>%s</rl2>\n", "")
		sb.AppendFormat("<rl3>%s</rl3>\n", "")
		// sb.AppendFormat("<tc>%s</tc>\n", TerminalConfigurations)

	case "AUTHORIZE_ONLINE_GIFT", "REFUND_ONLINE_GIFT", "RELOAD_GIFT", "ISSUE_GIFT", "RELOAD_REVERSAL_GIFT", "ISSUE_REVERSAL_GIFT":

		sb.AppendFormat("<va>%s</va>\n", rsp.VerifiedAmount)
		sb.AppendFormat("<ac>%s</ac>\n", rsp.ApprovalCode)
		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<pid1>%s</pid1>\n", rsp.PurchaseIdentifier1)
		sb.AppendFormat("<pid2>%s</pid2>\n", rsp.PurchaseIdentifier2)
		sb.AppendFormat("<oid>%s</oid>\n", rsp.OperatorID)
		sb.AppendFormat("<tsn>%s</tsn>\n", rsp.TransactionSequenceNumber)
		sb.AppendFormat("<tid>%s</tid>\n", rsp.TransactionID)
		sb.AppendFormat("<ltd>%s</ltd>\n", rsp.LocalTransactionDate)
		sb.AppendFormat("<ltt>%s</ltt>\n", rsp.LocalTransactionTime)
		sb.AppendFormat("<cspn>%s</cspn>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<mn>%s</mn>\n", rsp.MerchantNumber)
		sb.AppendFormat("<ts>%s</ts>\n", strings.ToUpper(rsp.TransactionStatus))
		sb.AppendFormat("<noc>%s</noc>\n", rsp.NameOnCard)
		sb.AppendFormat("<an>%s</an>\n", rsp.AccountNumber)
		sb.AppendFormat("<ct>%s</ct>\n", rsp.CardType)
		sb.AppendFormat("<ed>%s</ed>\n", rsp.ExpirationDate)
		// sb.AppendFormat("<cb>%s</cb>\n", rsp.CurrentBalance)
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)

	case "GET_BALANCE_GIFT":

		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<cspn>%s</cspn>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<mn>%s</mn>\n", rsp.MerchantNumber)
		sb.AppendFormat("<noc>%s</noc>\n", rsp.NameOnCard)
		sb.AppendFormat("<an>%s</an>\n", rsp.AccountNumber)
		sb.AppendFormat("<ct>%s</ct>\n", rsp.CardType)
		sb.AppendFormat("<ed>%s</ed>\n", rsp.ExpirationDate)
		// sb.AppendFormat("<cb>%s</cb>\n", rsp.CurrentBalance)
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)

	default:

		sb.AppendFormat("<va>%s</va>\n", rsp.VerifiedAmount)
		sb.AppendFormat("<ac>%s</ac>\n", rsp.ApprovalCode)
		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<cvvrc>%s</cvvrc>\n", rsp.CVVResultCode)
		sb.AppendFormat("<arc>%s</arc>\n", rsp.AVSResultCode)
		sb.AppendFormat("<pid1>%s</pid1>\n", rsp.PurchaseIdentifier1)
		sb.AppendFormat("<pid2>%s</pid2>\n", rsp.PurchaseIdentifier2)
		sb.AppendFormat("<oid>%s</oid>\n", rsp.OperatorID)
		sb.AppendFormat("<tsn>%s</tsn>\n", rsp.TransactionSequenceNumber)
		sb.AppendFormat("<tid>%s</tid>\n", rsp.TransactionID)
		sb.AppendFormat("<ltd>%s</ltd>\n", rsp.LocalTransactionDate)
		sb.AppendFormat("<ltt>%s</ltt>\n", rsp.LocalTransactionTime)
		sb.AppendFormat("<cspn>%s</cspn>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<vapn>%s</vapn>\n", rsp.VoiceAuthorizationPhoneNumber)
		sb.AppendFormat("<mn>%s</mn>\n", rsp.MerchantNumber)
		sb.AppendFormat("<tidn>%s</tidn>\n", rsp.TerminalIDNumber)
		sb.AppendFormat("<rrn>%s</rrn>\n", rsp.RetrievalReferenceNumber)
		sb.AppendFormat("<ct>%s</ct>\n", rsp.CardType)

		if 0 < len(rsp.AccountNumber) {
			sb.AppendFormat("<noc>%s</noc>\n", rsp.NameOnCard)
			sb.AppendFormat("<an>%s</an>\n", rsp.AccountNumber)
			sb.AppendFormat("<ed>%s</ed>\n", rsp.ExpirationDate)
		}

	case "AUTHORIZE_ONLINE_FS", "REFUND_ONLINE_FS", "AUTHORIZE_VOUCHER_FS", "GET_BALANCE_FS", "AUTHORIZE_ONLINE_CB", "WITHDRAW_ONLINE_CB", "GET_BALANCE_CB":
		// sb.AppendFormat("<lbcb>%s</lbcb>\n", rsp.LedgerBalanceCashBenefits)
		// sb.AppendFormat("<lbfs>%s</lbfs>\n", rsp.LedgerBalanceFoodStamps)

		sb.AppendFormat("<ts>%s</ts>\n", strings.ToUpper(rsp.TransactionStatus))
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)

		// sb.AppendFormat("<tt>%s</tt>\n", GetTransactionType(action))

	}

	sb.AppendFormat("</responsemessage>")
	return sb.String()
}

func (rsp *ResponseMessage) CreatePOSCP4XmlResponse(action string, req *RequestMessage) string {
	sb := bytes.NewSafeBuffer(1024)
	sb.AppendFormat("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	sb.AppendFormat("<responsemessage>\n")

	switch action {
	case "GET_CARD_INFO":
		sb.AppendFormat("<TrackData>%s</TrackData>\n", req.TrackData)
		sb.AppendFormat("<NameOnCard>%s</NameOnCard>\n", req.NameOnCard)
		sb.AppendFormat("<AccountNumber>%s</AccountNumber>\n", req.AccountNumber)
		sb.AppendFormat("<CardType>%s</CardType>\n", req.CardType)
		sb.AppendFormat("<ExpirationDate>%s</ExpirationDate>\n", req.ExpirationDate)
		sb.AppendFormat("<MessageFromHost>%s</MessageFromHost>\n", rsp.MessageFromHost)
		sb.AppendFormat("<ResponseCode>%s</ResponseCode>\n", rsp.ResponseCode)
		sb.AppendFormat("<ResponseText>%s</ResponseText>\n", rsp.ResponseText)

	case "GET_HOSTMESSAGES":

		sb.AppendFormat("<MessageFromHost>%s</MessageFromHost>\n", rsp.MessageFromHost)
	case "GET_MERCHANT_INFO":

		// sb.AppendFormat("<MerchantName>%s</MerchantName>\n", MerchantName)
		// sb.AppendFormat("<MerchantNumber>%s</MerchantNumber>\n", rsp.MerchantNumber)
		// sb.AppendFormat("<MerchantURL>%s</MerchantURL>\n", MerchantURL)
		// sb.AppendFormat("<MerchantEMail>%s</MerchantEMail>\n", MerchantEMail)
		// sb.AppendFormat("<MerchantStreet1>%s</MerchantStreet1>\n", MerchantStreet1)
		// sb.AppendFormat("<MerchantStreet2>%s</MerchantStreet2>\n", MerchantStreet2)
		// sb.AppendFormat("<MerchantCity>%s</MerchantCity>\n", MerchantCity)
		// sb.AppendFormat("<MerchantState>%s</MerchantState>\n", MerchantState)
		// sb.AppendFormat("<MerchantZip>%s</MerchantZip>\n", MerchantZip)
		// sb.AppendFormat("<MerchantPhone>%s</MerchantPhone>\n", MerchantPhone)
		sb.AppendFormat("<ResponseCode>%s</ResponseCode>\n", rsp.ResponseCode)
		sb.AppendFormat("<ResponseText>%s</ResponseText>\n", rsp.ResponseText)
		sb.AppendFormat("<rl1>%s</rl1>\n", "")
		sb.AppendFormat("<rl2>%s</rl2>\n", "")
		sb.AppendFormat("<rl3>%s</rl3>\n", "")
		// sb.AppendFormat("<tc>%s</tc>\n", TerminalConfigurations)

	case "AUTHORIZE_ONLINE_GIFT", "REFUND_ONLINE_GIFT", "RELOAD_GIFT", "ISSUE_GIFT", "RELOAD_REVERSAL_GIFT", "ISSUE_REVERSAL_GIFT":

		sb.AppendFormat("<VerifiedAmount>%s</VerifiedAmount>\n", rsp.VerifiedAmount)
		sb.AppendFormat("<ApprovalCode>%s</ApprovalCode>\n", rsp.ApprovalCode)
		sb.AppendFormat("<ResponseCode>%s</ResponseCode>\n", rsp.ResponseCode)
		sb.AppendFormat("<ResponseText>%s</ResponseText>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<PurchaseIdentifier1>%s</PurchaseIdentifier1>\n", rsp.PurchaseIdentifier1)
		sb.AppendFormat("<PurchaseIdentifier2>%s</PurchaseIdentifier2>\n", rsp.PurchaseIdentifier2)
		sb.AppendFormat("<OperatorID>%s</OperatorID>\n", rsp.OperatorID)
		sb.AppendFormat("<TransactionSequenceNumber>%s</TransactionSequenceNumber>\n", rsp.TransactionSequenceNumber)
		sb.AppendFormat("<TransactionID>%s</TransactionID>\n", rsp.TransactionID)
		sb.AppendFormat("<LocalTransactionDate>%s</LocalTransactionDate>\n", rsp.LocalTransactionDate)
		sb.AppendFormat("<LocalTransactionTime>%s</LocalTransactionTime>\n", rsp.LocalTransactionTime)
		sb.AppendFormat("<CustomerServicePhoneNumber>%s</CustomerServicePhoneNumber>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<MerchantNumber>%s</MerchantNumber>\n", rsp.MerchantNumber)
		sb.AppendFormat("<TransactionStatus>%s</TransactionStatus>\n", strings.ToUpper(rsp.TransactionStatus))
		sb.AppendFormat("<NameOnCard>%s</NameOnCard>\n", rsp.NameOnCard)
		sb.AppendFormat("<AccountNumber>%s</AccountNumber>\n", rsp.AccountNumber)
		sb.AppendFormat("<CardType>%s</CardType>\n", rsp.CardType)
		sb.AppendFormat("<ExpirationDate>%s</ExpirationDate>\n", rsp.ExpirationDate)
		// sb.AppendFormat("<CurrentBalance>%s</CurrentBalance>\n", rsp.CurrentBalance)
		sb.AppendFormat("<MessageFromHost>%s</MessageFromHost>\n", rsp.MessageFromHost)

	case "GET_BALANCE_GIFT":

		sb.AppendFormat("<ResponseCode>%s</ResponseCode>\n", rsp.ResponseCode)
		sb.AppendFormat("<ResponseText>%s</ResponseText>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<CustomerServicePhoneNumber>%s</CustomerServicePhoneNumber>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<MerchantNumber>%s</MerchantNumber>\n", rsp.MerchantNumber)
		sb.AppendFormat("<NameOnCard>%s</NameOnCard>\n", rsp.NameOnCard)
		sb.AppendFormat("<AccountNumber>%s</AccountNumber>\n", rsp.AccountNumber)
		sb.AppendFormat("<CardType>%s</CardType>\n", rsp.CardType)
		sb.AppendFormat("<ExpirationDate>%s</ExpirationDate>\n", rsp.ExpirationDate)
		// sb.AppendFormat("<CurrentBalance>%s</CurrentBalance>\n", rsp.CurrentBalance)
		sb.AppendFormat("<MessageFromHost>%s</MessageFromHost>\n", rsp.MessageFromHost)

	default:

		sb.AppendFormat("<VerifiedAmount>%s</VerifiedAmount>\n", rsp.VerifiedAmount)
		sb.AppendFormat("<ApprovalCode>%s</ApprovalCode>\n", rsp.ApprovalCode)
		sb.AppendFormat("<ResponseCode>%s</ResponseCode>\n", rsp.ResponseCode)
		sb.AppendFormat("<ResponseText>%s</ResponseText>\n", rsp.ResponseText)
		sb.AppendFormat("<ec>%s</ec>\n", "")
		sb.AppendFormat("<et>%s</et>\n", "")
		sb.AppendFormat("<CVVResultCode>%s</CVVResultCode>\n", rsp.CVVResultCode)
		sb.AppendFormat("<AVSResultCode>%s</AVSResultCode>\n", rsp.AVSResultCode)
		sb.AppendFormat("<PurchaseIdentifier1>%s</PurchaseIdentifier1>\n", rsp.PurchaseIdentifier1)
		sb.AppendFormat("<PurchaseIdentifier2>%s</PurchaseIdentifier2>\n", rsp.PurchaseIdentifier2)
		sb.AppendFormat("<OperatorID>%s</OperatorID>\n", rsp.OperatorID)
		sb.AppendFormat("<TransactionSequenceNumber>%s</TransactionSequenceNumber>\n", rsp.TransactionSequenceNumber)
		sb.AppendFormat("<TransactionID>%s</TransactionID>\n", rsp.TransactionID)
		sb.AppendFormat("<LocalTransactionDate>%s</LocalTransactionDate>\n", rsp.LocalTransactionDate)
		sb.AppendFormat("<LocalTransactionTime>%s</LocalTransactionTime>\n", rsp.LocalTransactionTime)
		sb.AppendFormat("<CustomerServicePhoneNumber>%s</CustomerServicePhoneNumber>\n", rsp.CustomerServicePhoneNumber)
		sb.AppendFormat("<VoiceAuthorizationPhoneNumber>%s</VoiceAuthorizationPhoneNumber>\n", rsp.VoiceAuthorizationPhoneNumber)
		sb.AppendFormat("<MerchantNumber>%s</MerchantNumber>\n", rsp.MerchantNumber)
		sb.AppendFormat("<TerminalIDNumber>%s</TerminalIDNumber>\n", rsp.TerminalIDNumber)
		sb.AppendFormat("<RetrievalReferenceNumber>%s</RetrievalReferenceNumber>\n", rsp.RetrievalReferenceNumber)
		sb.AppendFormat("<CardType>%s</CardType>\n", rsp.CardType)

		if 0 < len(rsp.AccountNumber) {
			sb.AppendFormat("<NameOnCard>%s</NameOnCard>\n", rsp.NameOnCard)
			sb.AppendFormat("<AccountNumber>%s</AccountNumber>\n", rsp.AccountNumber)
			sb.AppendFormat("<ExpirationDate>%s</ExpirationDate>\n", rsp.ExpirationDate)
		}

	case "AUTHORIZE_ONLINE_FS", "REFUND_ONLINE_FS", "AUTHORIZE_VOUCHER_FS", "GET_BALANCE_FS", "AUTHORIZE_ONLINE_CB", "WITHDRAW_ONLINE_CB", "GET_BALANCE_CB":
		// sb.AppendFormat("<lbcb>%s</lbcb>\n", rsp.LedgerBalanceCashBenefits)
		// sb.AppendFormat("<lbfs>%s</lbfs>\n", rsp.LedgerBalanceFoodStamps)

		sb.AppendFormat("<TransactionStatus>%s</TransactionStatus>\n", strings.ToUpper(rsp.TransactionStatus))
		sb.AppendFormat("<MessageFromHost>%s</MessageFromHost>\n", rsp.MessageFromHost)

		// sb.AppendFormat("<tt>%s</tt>\n", GetTransactionType(action))

	}

	sb.AppendFormat("</responsemessage>")
	return sb.String()
}

func (rsp *ResponseMessage) CreatePOSCP308XmlResponse(action string, req *RequestMessage) string {
	sb := bytes.NewSafeBuffer(1024)
	sb.AppendFormat("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	sb.AppendFormat("<responsemessage>\n")

	switch action {

	case "GET_CARD_INFO":
		sb.AppendFormat("<noc>%s</noc>\n", rsp.NameOnCard)
		sb.AppendFormat("<an>%s</an>\n", rsp.AccountNumber)
		sb.AppendFormat("<ct>%s</ct>\n", rsp.CardType)
		sb.AppendFormat("<ed>%s</ed>\n", rsp.ExpirationDate)
		sb.AppendFormat("<mfh>%s</mfh\n", rsp.MessageFromHost)

	default:
		sb.AppendFormat("<va>%s</va>\n", rsp.VerifiedAmount)
		sb.AppendFormat("<ac>%s</ac>\n", rsp.ApprovalCode)
		sb.AppendFormat("<rc>%s</rc>\n", rsp.ResponseCode)
		sb.AppendFormat("<rt>%s</rt>\n", rsp.ResponseText)
		sb.AppendFormat("<cvvrc>%s</cvvrc>\n", rsp.CVVResultCode)
		sb.AppendFormat("<arc>%s</arc>\n", rsp.AVSResultCode)
		sb.AppendFormat("<pid1>%s</pid1>\n", rsp.PurchaseIdentifier1)
		sb.AppendFormat("<pid2>%s</pid2>\n", rsp.PurchaseIdentifier2)
		sb.AppendFormat("<tsn>%s</tsn>\n", rsp.TransactionSequenceNumber)
		sb.AppendFormat("<ltd>%s</ltd>\n", rsp.LocalTransactionDate)
		sb.AppendFormat("<ltt>%s</ltt>\n", rsp.LocalTransactionTime)
		sb.AppendFormat("<vapn>%s</vapn>\n", rsp.VoiceAuthorizationPhoneNumber)
		sb.AppendFormat("<mfh>%s</mfh>\n", rsp.MessageFromHost)

	}

	sb.AppendFormat("</responsemessage>")
	return sb.String()
}

func (rsp *ResponseMessage) CreatePOSCP3NameValueResponse(action string) string {

	sb := bytes.NewSafeBuffer(1024)

	sb.AppendFormat("[VA]=[%s]\n", rsp.VerifiedAmount)
	sb.AppendFormat("[AC]=[%s]\n", rsp.ApprovalCode)
	sb.AppendFormat("[RC]=[%s]\n", rsp.ResponseCode)
	sb.AppendFormat("[RT]=[%s]\n", rsp.ResponseText)
	sb.AppendFormat("[ARC]=[%s]\n", rsp.AVSResultCode)
	sb.AppendFormat("[PID1]=[%s]\n", rsp.PurchaseIdentifier1)
	sb.AppendFormat("[PID2]=[%s]\n", rsp.PurchaseIdentifier2)
	sb.AppendFormat("[TSN]=[%s]\n", rsp.TransactionSequenceNumber)
	sb.AppendFormat("[LTD]=[%s]\n", rsp.LocalTransactionDate)
	sb.AppendFormat("[LTT]=[%s]\n", rsp.LocalTransactionTime)
	sb.AppendFormat("[EC]=[%s]\n", rsp.ErrorCode)
	sb.AppendFormat("[ET]=[%s]\n", rsp.ErrorText)

	return sb.String()
}
