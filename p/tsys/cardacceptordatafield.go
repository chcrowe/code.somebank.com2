package tsys

import (
	"fmt"
	"strconv"
	"strings"

	"code.somebank.com/p/bytes"
)

type CardAcceptorDataField struct {
	MerchantName string //NNNN... 1-25 A/N left-justified and space-filled
	//name provided must correspond to the name printed on the customer receipt

	// For preferred customer/passenger transport transactions, characters 1-12 of name field should
	// contain the shortened name, left-justified and space-filled to 12. Characters 13-25 of this field
	// should contain the ticket number, left-justified and space-filled to 25

	MerchantCity       string //LL.....LL 26-38 A/N
	MerchantState      string //SS 39-40 A/N chars must be upper case
	TicketNumber       string //Characters 13-25 of merchant name field for preferred customer/passenger transport transactions
	CustSvcPhoneNumber uint64 // 13 A/N 999-9999999 (dash is required)
}

func NewCardAcceptorDataField(name string, city string, state string, ticket string, phone uint64) *CardAcceptorDataField {
	return &CardAcceptorDataField{name, city, state, ticket, phone}
}

func (c *CardAcceptorDataField) Copy(cadf *CardAcceptorDataField) {
	c.MerchantName = cadf.MerchantName
	c.MerchantCity = cadf.MerchantCity
	c.MerchantState = cadf.MerchantState
	c.TicketNumber = cadf.TicketNumber
	c.CustSvcPhoneNumber = cadf.CustSvcPhoneNumber
}

func (c *CardAcceptorDataField) String(transcode TransactionCodeType) string {
	switch transcode {
	case PurchaseCardNotPresent, PurchaseCardNotPresentRepeat, CreditAccountFundingOrPayment, CreditAccountFundingOrPaymentRepeat, CardholderFundsTransferCardNotPresent, CardholderFundsTransferCardNotPresentRepeat:
		ph := fmt.Sprintf("%d", c.CustSvcPhoneNumber)
		return fmt.Sprintf("%-25s%03s-%-9s%s", c.MerchantName, ph[0:3], ph[3], c.MerchantState[0:2])
	default:
		return fmt.Sprintf("%-25s%-13s%s", c.MerchantName, c.MerchantCity, c.MerchantState[0:2])
	}
	return ""
}

func (c *CardAcceptorDataField) VerboseString(transcode TransactionCodeType) string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	switch transcode {
	case PurchaseCardNotPresent, PurchaseCardNotPresentRepeat, CreditAccountFundingOrPayment, CreditAccountFundingOrPaymentRepeat, CardholderFundsTransferCardNotPresent, CardholderFundsTransferCardNotPresentRepeat:
		ph := fmt.Sprintf("%d", c.CustSvcPhoneNumber)
		ph2 := fmt.Sprintf("%-3s-%-9s", ph[0:3], ph[3])

		buffer.AppendFormat("%8s%-32s%-25s\n", " ", "MerchantName", c.MerchantName)
		buffer.AppendFormat("%8s%-32s%-13s\n", " ", "CustomerServicePhoneNumber", ph2)
		buffer.AppendFormat("%8s%-32s%s\n", " ", "MerchantState", c.MerchantState)

	default:
		buffer.AppendFormat("%8s%-32s%-25s\n", " ", "MerchantName", c.MerchantName)
		buffer.AppendFormat("%8s%-32s%-13s\n", " ", "MerchantCity", c.MerchantCity)
		buffer.AppendFormat("%8s%-32s%s\n", " ", "MerchantState", c.MerchantState)
	}

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}

func ParseCardAcceptorDataField(transcode TransactionCodeType, s string) *CardAcceptorDataField {
	c := CardAcceptorDataField{"", "", "", "", 0}

	if 40 > len(s) {
		return &c
		//panic("CardAcceptorDataField string must be at least 40 characters in length")
	}

	c.MerchantName = strings.Trim(s[0:25], " ")
	switch transcode {
	case PurchaseCardNotPresent, PurchaseCardNotPresentRepeat, CreditAccountFundingOrPayment, CreditAccountFundingOrPaymentRepeat, CardholderFundsTransferCardNotPresent, CardholderFundsTransferCardNotPresentRepeat:
		c.CustSvcPhoneNumber, _ = strconv.ParseUint(strings.Replace(strings.Trim(s[25:38], " "), "-", "", -1), 10, 64)

	default:
		c.MerchantCity = strings.Trim(s[25:38], " ")

	}

	c.MerchantState = strings.Trim(s[38:], " ")

	return &c
}
