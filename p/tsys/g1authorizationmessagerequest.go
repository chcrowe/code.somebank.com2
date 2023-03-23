package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"

	//"regexp"
	"strconv"
	"strings"
	"time"
)

type G1AuthorizationMessageRequest struct {
	RecordFormat     RecordFormatType         //1 A/N 4.67 (CreditCardAuthorizationRequest)
	ApplicationType  ApplicationIndicatorType //1 A/N 4.9 (MultipleAuthorizationsPerConnectionFullDuplexInterleaved)
	MessageDelimiter bytes.AsciiCharacterType //1 A/N 4.62 (MessageDelimiter)

	//BEGIN - Merchant Persistent Data
	AcquirerBIN    uint64           //6 NUM 4.2 (999995)
	MerchantNumber uint64           //12 NUM 4.59 (999999999911)
	StoreNumber    uint             //4 NUM 4.81 (1515)
	TerminalNumber uint             //4 NUM 4.84 (5999)
	DeviceCode     DeviceCodeType   //1 A/N 4.34 (ReservedForThirdPartyDevelopers)
	IndustryCode   IndustryCodeType //1 A/N 4.46 (UnknownOrUnsure)
	CurrencyCode   CurrencyCodeType //3 NUM 4.30 (UsdUnitedStatesUsDollar)
	CountryCode    CountryCodeType  //3 NUM 4.28 (UnitedStates)
	CityCode       string           //9 A/N 4.25 City Code must be one of the following three formats...
	// 									U.S. ZIP Code 5 character numeric, left-justified, space-filled.
	// 									U.S. ZIP Code + Extension 9 character numeric.
	// 									Canadian Postal Code 6 character "AnAnAn" format, left-justified, spacefilled.

	LanguageIndicator    LanguageIndicatorType //2 NUM 4.51 (English)
	TimeZoneDiff         *TimeZoneDifferential //3 NUM 4.86 (705 EASTERN)
	MerchantCategoryCode MerchantCategoryType  //4 NUM 4.57
	//END - Merchant Persistent Data

	//BEGIN - Dynamic Transaction Data
	RequestedACI                 RequestAciType        //1 A/N 4.69 (DeviceIsCpsMeritCapableOrCreditOrOffline)
	TransactionSequenceNumber    uint                  //4 NUM 4.92
	TransactionCode              TransactionCodeType   //2 A/N 4.89 (Purchase)
	CardholderIdentificationCode CardholderIdCodeType  //1 A/N 4.23
	AccountDataSource            AccountDataSourceType //1 A/N 4.1

	//all request messages have same exact Group I structure up to this point...

	CustomerData                 *CustomerDataField                 //1 - 79 A/N 4.31
	CardholderIdentificationData *CardholderIdentificationDataField //0, 128 A/N 4.43
	ReceivingInstitutionId       ReceivingInstitutionIdType         //0, 6 NUM 4.66 used for Check Authorizations and Private Label
	TransactionAmount            uint64                             //1 - 12 NUM 4.88
	// 									In industries where cash back is permitted on purchase transactions, SecondaryAmount could contain a
	// 									cash back amount.
	// 									For Authorization Reversal transactions, this field must contain the final settlement amount for
	//  								a partial reversal, and is not used for a full reversal.
	SecondaryAmount uint64 //0, 12 NUM 4.76

	MarketSpecificData           *MarketSpecificDataField //0, 4 A/N 4.54
	CardAcceptorData             *CardAcceptorDataField   //0, 40 A/N 4.17
	ReversalOrIncrementalTransID string                   //15 A/N 4.75 see Transaction ID in G1 Auth Response

	// This 15-character field can contain a Transaction Identifier (Visa, American Express, PayPal or
	// Discover) or Reference Number (MasterCard) (see Table 3.6 for record format and version
	// number). The POS device does not attempt to interpret the meaning of any data appearing in
	// this field. Data returned in this field is recorded and submitted as part of the data capture
	// settlement format.
	// For incremental authorization requests, the Transaction ID must be the same Transaction ID
	// returned in the original authorization response.
	// If “MAV” is in positions 5-7 of this field, the transaction should not be submitted for capture.

	ReversalAndCancelData *ReversalAndCancelDataI //0, 30 A/N 4.73

	//END - Dynamic Transaction Data
}

func NewG1AuthorizationMessageRequest(acquirerid uint64, transcode TransactionCodeType, sequenceno uint, amount uint64, cashback uint64) *G1AuthorizationMessageRequest {
	t := G1AuthorizationMessageRequest{}

	if true == IsDebitCardTransaction(transcode) {
		t.RecordFormat = DebitEbtRequest
	} else {
		t.RecordFormat = CreditCardAuthorizationRequest //1 A/N 4.67 Record Format (eg CreditCardAuthorizationRequest)
	}

	t.ApplicationType = MultipleAuthorizationsPerConnectionFullDuplexInterleaved //1 A/N 4.9 eg SingleAuthorizationPerConnection, MultipleAuthorizationsPerConnectionFullDuplexInterleaved
	t.MessageDelimiter = bytes.MessageDelimiter                                  //1 A/N 4.62
	t.AcquirerBIN = acquirerid                                                   //6 NUM 4.2 eg. 999995
	t.TransactionCode = transcode                                                //2 A/N 4.89 eg Purchase
	t.TransactionSequenceNumber = sequenceno                                     //4 NUM 4.92
	t.TransactionAmount = amount
	t.SecondaryAmount = cashback

	t.MarketSpecificData = NewMarketSpecificDataField(AutoRentalOrNonParticipatingProperty, OtherIndustries, 0)
	t.ReversalAndCancelData = NewReversalAndCancelDataI("", time.Now(), 0)

	return &t
}

func (t *G1AuthorizationMessageRequest) SetTerminalIdentification(merchant uint64, store uint, terminal uint, device DeviceCodeType, industry IndustryCodeType, category MerchantCategoryType) {

	t.MerchantNumber = merchant                               //12 NUM 4.59 eg 999999999911
	t.StoreNumber = store                                     //4 NUM 4.81 eg 11
	t.TerminalNumber = terminal                               //4 NUM 4.84 eg 5999
	t.DeviceCode = device                                     //1 A/N 4.34 eg ReservedForThirdPartyDevelopers
	t.MerchantCategoryCode = category                         //4 NUM 4.57 eg SpaceOrEmptyMerchantCategory
	t.IndustryCode = industry                                 //1 A/N 4.46 eg AllTypesExceptHotelAndAuto
	t.RequestedACI = DeviceIsCpsMeritCapableOrCreditOrOffline //1 A/N 4.69 eg DeviceIsCpsMeritCapableOrCreditOrOffline
}

func (t *G1AuthorizationMessageRequest) SetTerminalLocale(currency CurrencyCodeType, country CountryCodeType, language LanguageIndicatorType, zipcode string, timezone *TimeZoneDifferential) {

	t.CurrencyCode = currency      //3 NUM 4.30 eg UsdUnitedStatesUsDollar
	t.CountryCode = country        //3 NUM 4.28 eg UnitedStates
	t.CityCode = zipcode           //9 A/N 4.25
	t.LanguageIndicator = language //2 NUM 4.51 eg English
	t.TimeZoneDiff = timezone      //3 NUM 4.86 eg 705
}

func (t *G1AuthorizationMessageRequest) SetCardholderData(customeraccountdata *CustomerDataField, cardholderidentification *CardholderIdentificationDataField) {
	t.CardholderIdentificationCode = cardholderidentification.CardholderIdCode //1 A/N 4.23 eg SpaceOrEmptyCardholderId
	t.AccountDataSource = customeraccountdata.AccountDataSource                //1 A/N 4.1 eg SpaceOrEmptyAccountDataSource

	t.CustomerData = customeraccountdata
	t.CardholderIdentificationData = cardholderidentification
}

func (t *G1AuthorizationMessageRequest) SetMarketAndCardAcceptorData(marketdata *MarketSpecificDataField, cardacceptor *CardAcceptorDataField) {
	t.MarketSpecificData = marketdata
	t.CardAcceptorData = cardacceptor
}

func (t *G1AuthorizationMessageRequest) CopyToBuffer(buffer *bytes.SafeBuffer) {

	buffer.AppendFormat("%c", t.RecordFormat)                 //1 A/N 4.67
	buffer.AppendFormat("%d", t.ApplicationType)              //1 A/N 4.9
	buffer.AppendFormat("%c", t.MessageDelimiter)             //1 A/N 4.62
	buffer.AppendFormat("%06d", t.AcquirerBIN)                //6 NUM 4.2
	buffer.AppendFormat("%012d", t.MerchantNumber)            //12 NUM 4.59
	buffer.AppendFormat("%04d", t.StoreNumber)                //4 NUM 4.81
	buffer.AppendFormat("%04d", t.TerminalNumber)             //4 NUM 4.84
	buffer.AppendFormat("%c", t.DeviceCode)                   //1 A/N 4.34
	buffer.AppendFormat("%c", t.IndustryCode)                 //1 A/N 4.46
	buffer.AppendFormat("%03d", t.CurrencyCode)               //3 NUM 4.30
	buffer.AppendFormat("%03d", t.CountryCode)                //3 NUM 4.28
	buffer.AppendFormat("%-9s", t.CityCode)                   //9 A/N 4.25
	buffer.AppendFormat("%02d", t.LanguageIndicator)          //2 NUM 4.51
	buffer.AppendFormat("%03s", t.TimeZoneDiff)               //3 NUM 4.86
	buffer.AppendFormat("%04d", t.MerchantCategoryCode)       //4 NUM 4.57
	buffer.AppendFormat("%c", t.RequestedACI)                 //1 A/N 4.69
	buffer.AppendFormat("%04d", t.TransactionSequenceNumber)  //4 NUM 4.92
	buffer.AppendFormat("%2s", string(t.TransactionCode))     //2 A/N 4.89
	buffer.AppendFormat("%c", t.CardholderIdentificationCode) //1 A/N 4.23
	buffer.AppendFormat("%c", t.AccountDataSource)            //1 A/N 4.1

	// //all request messages have same exact Group I structure up to this point...(first 64 bytes only)

	buffer.AppendString(t.CustomerData.String()) //1 - 79 A/N 4.31
	buffer.AppendByte(byte(bytes.FileSeparator))

	buffer.AppendString(t.CardholderIdentificationData.String()) //0, 128 A/N 4.24
	buffer.AppendByte(byte(bytes.FileSeparator))

	if SpaceOrEmptyReceivingInst != t.ReceivingInstitutionId {
		buffer.AppendFormat("%06d", t.ReceivingInstitutionId) //0, 6 NUM 4.66
	}
	buffer.AppendByte(byte(bytes.FileSeparator))

	buffer.AppendFormat("%012d", t.TransactionAmount) //1 - 12 NUM 4.88
	buffer.AppendByte(byte(bytes.FileSeparator))

	if DirectDebitPurchase == t.TransactionCode {
		buffer.AppendFormat("%012d", t.SecondaryAmount) //0, 12 NUM 4.76
	}
	buffer.AppendByte(byte(bytes.FileSeparator))

	switch t.IndustryCode {
	case AutoRental, HotelAndLodging:
		buffer.AppendFormat("%s", t.MarketSpecificData.String()) //0, 4 A/N 4.54
	}
	buffer.AppendByte(byte(bytes.FileSeparator))

	buffer.AppendFormat("%s", t.CardAcceptorData.String(t.TransactionCode)) //0, 40 A/N 4.17
	buffer.AppendByte(byte(bytes.FileSeparator))

	if 0 < len(t.ReversalOrIncrementalTransID) {
		buffer.AppendFormat("%s", t.ReversalOrIncrementalTransID) //15 A/N 4.75
	}
	buffer.AppendByte(byte(bytes.FileSeparator))

	if true == IsReturnOrReversal(t.TransactionCode) {
		buffer.AppendFormat("%s", t.ReversalAndCancelData.String()) //0, 30 A/N 4.73
	}
	buffer.AppendByte(byte(bytes.FileSeparator))
}

func (t *G1AuthorizationMessageRequest) String() string {
	buffer := bytes.NewSafeBuffer(512)
	t.CopyToBuffer(buffer)
	return buffer.String()
}

func (t *G1AuthorizationMessageRequest) VerboseString() string {

	buffer := bytes.NewSafeBuffer(512)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RecordFormat", t.RecordFormat)                                      //1 A/N 4.67
	buffer.AppendFormat("%-32s%d (%[2]s)\n", "ApplicationType", t.ApplicationType)                                //1 A/N 4.9
	buffer.AppendFormat("%-32s%c\n", "MessageDelimiter", t.MessageDelimiter)                                      //1 A/N 4.62
	buffer.AppendFormat("%-32s%06d\n", "AcquirerBIN", t.AcquirerBIN)                                              //6 NUM 4.2
	buffer.AppendFormat("%-32s%012d\n", "MerchantNumber", t.MerchantNumber)                                       //12 NUM 4.59
	buffer.AppendFormat("%-32s%04d\n", "StoreNumber", t.StoreNumber)                                              //4 NUM 4.81
	buffer.AppendFormat("%-32s%04d\n", "TerminalNumber", t.TerminalNumber)                                        //4 NUM 4.84
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "DeviceCode", t.DeviceCode, t.DeviceCode)                            //1 A/N 4.34
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "IndustryCode", t.IndustryCode)                                      //1 A/N 4.46
	buffer.AppendFormat("%-32s%03d (%[2]s)\n", "CurrencyCode", t.CurrencyCode)                                    //3 NUM 4.30
	buffer.AppendFormat("%-32s%03d (%[2]s)\n", "CountryCode", t.CountryCode)                                      //3 NUM 4.28
	buffer.AppendFormat("%-32s%-9s\n", "CityCode", t.CityCode)                                                    //9 A/N 4.25
	buffer.AppendFormat("%-32s%02d (%[2]s)\n", "LanguageIndicator", t.LanguageIndicator)                          //2 NUM 4.51
	buffer.AppendFormat("%-32s%03s %s\n", "TimeZoneDifferential", t.TimeZoneDiff, t.TimeZoneDiff.VerboseString()) //3 NUM 4.86
	buffer.AppendFormat("%-32s%04d (%[2]s)\n", "MerchantCategoryCode", t.MerchantCategoryCode)                    //4 NUM 4.57
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "RequestedACI", t.RequestedACI)                                      //1 A/N 4.69
	buffer.AppendFormat("%-32s%04d\n", "TransactionSequenceNumber", t.TransactionSequenceNumber)                  //4 NUM 4.92
	buffer.AppendFormat("%-32s%2s (%[3]s)\n", "TransactionCode", string(t.TransactionCode), t.TransactionCode)    //2 A/N 4.89
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "CardholderIdentificationCode", t.CardholderIdentificationCode)      //1 A/N 4.23
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "AccountDataSource", t.AccountDataSource)                            //1 A/N 4.1

	// //all request messages have same exact Group I structure up to this point...(first 64 bytes only)

	buffer.AppendFormat("%-32s%s %s\n", "CustomerData", t.CustomerData.String(), t.CustomerData.VerboseString()) //1 - 79 A/N 4.31
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	buffer.AppendFormat("%-32s%s %s\n", "CardholderIdData", t.CardholderIdentificationData.String(), t.CardholderIdentificationData.VerboseString()) //0, 128 A/N 4.24
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	if SpaceOrEmptyReceivingInst != t.ReceivingInstitutionId {
		buffer.AppendFormat("%-32s%06d (%[2]s)\n", "ReceivingInstitutionId", t.ReceivingInstitutionId) //0, 6 NUM 4.66
	} else {
		buffer.AppendFormat("%-32s (%s)\n", "ReceivingInstitutionId", t.ReceivingInstitutionId) //0, 6 NUM 4.66
	}
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	buffer.AppendFormat("%-32s%012d\n", "TransactionAmount", t.TransactionAmount) //1 - 12 NUM 4.88
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	if DirectDebitPurchase == t.TransactionCode {
		buffer.AppendFormat("%-32s%012d\n", "SecondaryAmount", t.SecondaryAmount) //0, 12 NUM 4.76
	}
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	switch t.IndustryCode {
	case AutoRental, HotelAndLodging:
		buffer.AppendFormat("%-32s%s %s\n", "MarketSpecificData", t.MarketSpecificData.String(), t.MarketSpecificData.VerboseString()) //0, 4 A/N 4.54
		// default:
		// 	buffer.AppendFormat("%-32sN/A (Hotel and AutoRental Only)\n", "MarketSpecificData") //0, 4 A/N 4.54
	}
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	buffer.AppendFormat("%-32s%s %s\n", "CardAcceptorData", t.CardAcceptorData.String(t.TransactionCode), t.CardAcceptorData.VerboseString(t.TransactionCode)) //0, 40 A/N 4.17
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	if 0 < len(t.ReversalOrIncrementalTransID) {
		buffer.AppendFormat("%-32s%s\n", "ReversalOrIncrementalTransID", t.ReversalOrIncrementalTransID) //15 A/N 4.75
	}
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	if true == IsReturnOrReversal(t.TransactionCode) {
		buffer.AppendFormat("%-32s%s %s\n", "ReversalAndCancelDataI", t.ReversalAndCancelData.String(), t.ReversalAndCancelData.VerboseString()) //0, 30 A/N 4.73
	}
	buffer.AppendFormat("%-32s%c (%[2]s)\n", "FS", bytes.FileSeparator)

	return buffer.String()
}

func ParseG1AuthorizationMessageRequest(s string) (*G1AuthorizationMessageRequest, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!

	t := G1AuthorizationMessageRequest{}

	i := 0
	if 67 > len(s) {
		panic("minimum EIS1080 Group 1 request message length is 67")
	}
	//i++ //advance past the STX

	fmt.Sscanf(s[i:i+37], "%c%1d%c%06d%012d%04d%04d%c%c%03d%03d",
		&t.RecordFormat,     //1 A/N 4.67
		&t.ApplicationType,  //1 A/N 4.9
		&t.MessageDelimiter, //1 A/N 4.62
		&t.AcquirerBIN,      //6 NUM 4.2
		&t.MerchantNumber,   //12 NUM 4.59
		&t.StoreNumber,      //4 NUM 4.81
		&t.TerminalNumber,   //4 NUM 4.84
		&t.DeviceCode,       //1 A/N 4.34
		&t.IndustryCode,     //1 A/N 4.46
		&t.CurrencyCode,     //3 NUM 4.30
		&t.CountryCode)      //3 NUM 4.28

	i += 37

	t.CityCode = s[i : i+9] //9 A/N 4.25
	i += 9

	var tmzonediff uint
	fmt.Sscanf(s[i:i+14], "%02d%03d%04d%c%04d",
		&t.LanguageIndicator,         //2 NUM 4.51
		&tmzonediff,                  //3 NUM 4.86
		&t.MerchantCategoryCode,      //4 NUM 4.57
		&t.RequestedACI,              //1 A/N 4.69
		&t.TransactionSequenceNumber) //4 NUM 4.92

	t.TimeZoneDiff = NewTimeZoneDifferentialFromUint(tmzonediff)
	i += 14

	t.TransactionCode = TransactionCodeType(s[i : i+2])           //2 A/N 4.89
	t.CardholderIdentificationCode = CardholderIdCodeType(s[i+2]) //1 A/N 4.23
	t.AccountDataSource = AccountDataSourceType(s[i+3])           //1 A/N 4.1
	i += 4

	// // //all request messages have same exact Group I structure up to this point...(first 64 bytes only)
	//the customer data field potentially has embedded field separators (FS); this must be parsed BEFORE attempting to split the remainder of group1 by FS
	customerdata, charsread := ParseCustomerDataField(t.AccountDataSource, s[i:])
	t.CustomerData = customerdata

	i += charsread

	/*
	   The remainder of Group I should have this structure (8 field-<FS> separated elements)
	   0	1 ASCII 4.80 Field Separator
	   1	0, 128 A/N 4.43 Cardholder Identification Data
	   2	1 ASCII 4.80 Field Separator
	   3	0, 6 NUM 4.127 Receiving Institution ID
	   4	1 ASCII 4.80 Field Separator
	   5	1 - 12 NUM 4.161 Transaction Amount
	   6	1 ASCII 4.80 Field Separator
	   7	0-12 NUM 4.138 Secondary Amount
	   8	1 ASCII 4.80 Field Separator
	   9	0, 4 A/N 4.105 Market Specific Data
	   0	1 ASCII 4.80 Field Separator
	   1	0, 40 A/N 4.33 Card Acceptor Data
	   2	1 ASCII 4.80 Field Separator
	   3	15 A/N 4.136 Reversal and Incremental Transaction ID
	   4	1 ASCII 4.80 Field Separator
	   5	0, 30 A/N 4.134 Reversal and Cancel Data I
	   6	1 ASCII 4.80 Field Separator
	*/

	i += 1 // advance past last FS
	gp1fields := strings.Split(s[i:], string(bytes.FileSeparator))
	gp1fieldslen := len(gp1fields)

	const maxremaininggroupfields int = 8

	for j := 0; j < maxremaininggroupfields && j < gp1fieldslen; j++ {
		switch j {
		case 0:
			t.CardholderIdentificationData, _ = ParseCardholderIdentificationDataField(t.CardholderIdentificationCode, gp1fields[j])
		case 1:
			if 6 == len(gp1fields[j]) {
				fmt.Sscanf(gp1fields[j], "%06d", &t.ReceivingInstitutionId)
			}

		case 2:
			t.TransactionAmount, _ = strconv.ParseUint(gp1fields[j], 10, 64)
		case 3:
			if DirectDebitPurchase == t.TransactionCode && 0 < len(gp1fields[j]) {
				t.SecondaryAmount, _ = strconv.ParseUint(gp1fields[j], 10, 64)
			}
		case 4:
			t.MarketSpecificData = ParseMarketSpecificDataField(gp1fields[j])
		case 5:
			t.CardAcceptorData = ParseCardAcceptorDataField(t.TransactionCode, gp1fields[j])
		case 6:
			t.ReversalOrIncrementalTransID = gp1fields[j]
		case 7:
			t.ReversalAndCancelData = ParseReversalAndCancelDataI(gp1fields[j])
		}

		i += len(gp1fields[j]) + 1
	}

	return &t, i
}

func IsOfflineTransaction(code AuthorizationSourceType) bool {
	return OffLineApprovalAuthorizationCodeManuallyKeyed == code || CreditRefundOrNonauthorizedTransactions == code || OffLineApprovalPosDeviceGenerated == code
}

func IsReturnOrReversal(code TransactionCodeType) bool {
	switch code {
	case OnlineAuthorizationReversal, OnlineAuthorizationReversalRepeat, StoreAndForwardAuthorizationReversal, StoreAndForwardAuthorizationReversalRepeat,
		BalanceInquiryReversal, BalanceInquiryReversalRepeat, PosCheckReversalConversionWithGuarantee, PosCheckReversalConversionWithVerification, PosCheckReversalConversionOnly,
		FoodStampsReturnEbt, DirectDebitPurchaseReturn, DebitAccountFundingReturn, DebitFundsTransferReturn, AutomaticReversalDirectDebitPurchase,
		AutomaticReversalDirectDebitPurchaseReturn, AutomaticReversalInterlinkDirectDebitCancel, ATMCashDisbursementReversal, ATMDepositReversal,
		ATMCardholderAccountTransferReversal, GiftCardReturnRefund, GiftCardMerchantInitiatedReversal, Q2PrepaidCardActivationReversal, Q4PrepaidCardLoadReversal,
		PrepaidCardActivationReversal, PrepaidCardLoadReversal, CreditReturn, AccountFundingCreditReturn, CardholderFundsTransferCreditReturn:

		return true
	default:
		return false
	}
}

func IsCreditCardTransaction(code TransactionCodeType) bool {
	switch code {
	case Purchase, PurchaseRepeat, CashAdvance, CashAdvanceRepeat, PurchaseCardNotPresent, PurchaseCardNotPresentRepeat,
		QuasiCash, QuasiCashRepeat, CardAuthentication, CardAuthenticationRepeat, OnlineAuthorizationReversal, OnlineAuthorizationReversalRepeat, StoreAndForwardAuthorizationReversal, StoreAndForwardAuthorizationReversalRepeat, BillPayTransaction, BillPayTransactionRepeat, CreditAdvice, CreditAccountFundingOrPayment, CreditAccountFundingOrPaymentRepeat, CardNotPresentCreditAccountFundingOrPayment, CardNotPresentCreditAccountFundingOrPaymentRepeat, CardPresentCreditCardholderFundsTransfer,
		CardPresentCreditCardholderFundsTransferRepeat, CardholderFundsTransferCardNotPresent, CardholderFundsTransferCardNotPresentRepeat, CreditStoredValueBalanceInquiry, CreditStoredValueBalanceInquiryRepeat, HealthcareEligibilityInquiry, HealthcareEligibilityInquiryRepeat, BalanceInquiryReversal, BalanceInquiryReversalRepeat, ProductEligibilityInquiry:
		return true
	default:
		return false
	}
}

func IsGiftCardTransaction(code TransactionCodeType) bool {
	switch code {
	case GiftCardCloseCard, GiftCardBalanceInquiry, GiftCardPurchaseRedemption, GiftCardReturnRefund, GiftCardAddValueLoadCard,
		GiftCardDecreaseValueUnloadCard, GiftCardStandAloneTip, GiftCardIssueGiftCard, GiftCardIssueVirtualGiftCard,
		GiftCardMerchantInitiatedCancel, GiftCardMerchantInitiatedReversal, GiftCardCashBack:
		return true
	default:
		return false
	}
}

func IsDebitCardTransaction(code TransactionCodeType) bool {
	switch code {
	case FoodStampsReturnEbt, DirectDebitPurchase, DirectDebitPurchaseReturn, CashBenefitsCashWithdrawalEbt, FoodStampPurchaseEbt,
		DirectDebitBalanceInquiry, DebitBillPaymentTransaction, PinlessDebitBillPaymentTransaction, FoodStampsElectronicVoucherEbt,
		EbtCashBenefitsPurchaseOrPurchaseWithCashBack, DebitAccountFundingPurchase, DebitAccountFundingReturn, DebitCardholderFundsTransfer,
		DebitFundsTransferReturn, EbtFoodStampBalanceInquiry, EbtCashBenefitsBalanceInquiry, ChipCardTransactionAdviceRecordLimitedAvailability,
		AutomaticReversalDirectDebitPurchase, AutomaticReversalDirectDebitPurchaseReturn, AutomaticReversalInterlinkDirectDebitCancel:
		return true
	default:
		return false
	}
}

func IsCheckTransaction(code TransactionCodeType) bool {
	switch code {
	case CheckGuarantee, PosCheckConversionWithGuarantee, PosCheckConversionWithVerification, PosCheckConversionOnly, PosCheckReversalConversionWithGuarantee,
		PosCheckReversalConversionWithVerification, PosCheckReversalConversionOnly:
		return true
	default:
		return false
	}
}

func (t *G1AuthorizationMessageRequest) IsEcomOrMotoTransaction() bool {
	return (ManuallyKeyedTerminalHasNoCardReadingCapability == t.AccountDataSource) && (PurchaseCardNotPresent == t.TransactionCode || PurchaseCardNotPresentRepeat == t.TransactionCode)
}
