package tsys

import (
	"code.somebank.com/p/auth"
	"code.somebank.com/p/card"
	"code.somebank.com/p/card/msr"
	estrings "code.somebank.com/p/strings"
)

func CreateGroup3Elements(t *TerminalConfig, r *auth.RequestMessage, issuer card.CardIssuerType) []string {
	g3records := make([]string, 3)

	switch issuer {
	case card.Amex, card.Visa:
		g3records = append(g3records, NewG3v001CommercialCardRequest().String())
		break
	}

	if 0 < len(r.CVV) {
		g3records = append(g3records, NewG3v007CardVerificationCodeRequest(EnhancedResponseRequested, r.CVV).String())
	}

	if card.Visa == issuer && 0 < len(r.XID) && 0 < len(r.CAVV) {
		g3records = append(g3records, NewG3v017Visa3dSecurEcomVerificationRequest(r.XID, r.CAVV).String())
	}

	if card.Mastercard == issuer && 0 < len(r.UCAF) {
		g3records = append(g3records, NewG3v019MastercardUniversalCardholderAuthenticationRequest(UcafDataWasPopulated, r.UCAF).String())
	}

	g3records = append(g3records, NewG3v020DeveloperInformationRequest().String())

	if card.Visa == issuer && 0 < len(r.MerchantVerificationValue) {
		g3records = append(g3records, NewG3v021MerchantVerificationValueRequest(r.MerchantVerificationValue).String())
	}

	gratuityamount := estrings.StringDecimalAmountToUint64(r.GratuityAmount)
	roomrateamount := estrings.StringDecimalAmountToUint64(r.RoomRateAmount)
	settlementamount := estrings.StringDecimalAmountToUint64(r.SettlementAmount)
	taxamount := estrings.StringDecimalAmountToUint64(r.TaxAmount)
	if 0 < gratuityamount || 0 < roomrateamount || 0 < settlementamount || 0 < taxamount {

		additionalAmounts := [...]AdditionalAmountInfo{
			*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UnitedStatesUsDollar, PositiveBalance, gratuityamount)),
			*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UnitedStatesUsDollar, PositiveBalance, roomrateamount)),
			*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UnitedStatesUsDollar, PositiveBalance, settlementamount)),
			*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UnitedStatesUsDollar, PositiveBalance, taxamount))}

		g3records = append(g3records, NewG3v022AdditionalAmountsRequest(&additionalAmounts).String())
	}

	g3records = append(g3records, NewG3v025TransactionFeeAmountRequest(DebitTransactionFee, 150).String())

	if 0 < len(t.PosDataCodeProfile) {
		g3records = append(g3records, ParseG3v027PosDataRequest(t.PosDataCodeProfile).String())
	} else {
		g3records = append(g3records, NewG3v027PosDataRequest(TerminalInputICCReaderAndContactlessTerminal,
			CardholderAuthenticationPINEntryCapability,
			CardCaptureCapability,
			EnvironmentOnCardAcceptorPremisesAttendedTerminal,
			CardholderPresent,
			CardPresent,
			CardInputOnlineChip,
			CardholderPINAuthenticated,
			AuthenticationEntityAuthorizingAgentOnlinePIN,
			CardOutputCapabilityICC,
			TerminalOutputCapabilityICC,
			PINCaptureCapability06CharactersMaximum).String())
	}

	g3records = append(g3records, NewG3v032CurrencyConversionDataRequest().String())
	g3records = append(g3records, NewG3v034CardProductCodeRequest().String())

	if card.Visa == issuer {
		g3records = append(g3records, NewG3v040VisaISAChargeIndicatorRequest().String())
	}

	g3records = append(g3records, NewG3v046CardTypeGroupRequest().String())

	if card.Amex == issuer {
		g3records = append(g3records, NewG3v048AmexCardholderVerificationResultsRequest().String())
	}

	if card.Mastercard == issuer {
		g3records = append(g3records, NewG3v050AssociationTimestampRequest().String()) //mastercard
		g3records = append(g3records, NewG3v051MasterCardEventMonitoringRealTimeScoringRequest(1).String())
		g3records = append(g3records, NewG3v058AlternateAccountIDRequest().String()) //mastercard combine w G3v059
		g3records = append(g3records, NewG3v059MasterCardPayPassMappingServiceRequest().String())
	}

	g3records = append(g3records, NewG3v061SpendQualifiedIndicatorRequest().String())

	return g3records
}

func CreateCardholderData(t *TerminalConfig, r *auth.RequestMessage) (*CustomerDataField, *CardholderIdentificationDataField, card.CardIssuerType) {
	var customeraccountdata *CustomerDataField
	var cardholderidentification *CardholderIdentificationDataField
	var issuer card.CardIssuerType

	if 0 < len(r.TrackData) {
		msr1, msr2 := msr.ParseMagneticStripe(r.TrackData)
		if nil != msr2 && 0 < len(msr2.TrackData) {
			customeraccountdata = NewCustomerDataField(FullMagneticStripeReadAndTransmitTrack2, "", msr2.TrackData)
			issuer = card.GetCardIssuerType(msr2.AccountNumber)
		} else if nil != msr1 && 0 < len(msr1.TrackData) {
			customeraccountdata = NewCustomerDataField(FullMagneticStripeReadAndTransmitTrack1, "", msr1.TrackData)
			issuer = card.GetCardIssuerType(msr1.AccountNumber)
		}

		// SpaceOrEmptyAccountDataSource                           AccountDataSourceType = ' '
		// BarCodeRead                                             AccountDataSourceType = 'A'
		// MicrDataWasAcquiredByOcrReader                          AccountDataSourceType = 'B'
		// CheckImagingDevicePosCheckService                       AccountDataSourceType = 'C'
		// FullMagneticStripeReadAndTransmitTrack2                 AccountDataSourceType = 'D'
		// ChipCardReadDataCCPS                                    AccountDataSourceType = 'G'
		// FullMagneticStripeReadAndTransmitTrack1                 AccountDataSourceType = 'H'
		// ManuallyKeyedChipCardReadCapableTerminal                AccountDataSourceType = 'P'
		// ProximityPaymentUsingTrackDataRules                     AccountDataSourceType = 'Q'
		// ProximityPaymentUsingEmvRules                           AccountDataSourceType = 'R'
		// FullMagneticStripeReadChipCardCapableTerminal           AccountDataSourceType = 'S'
		// ManuallyKeyedTrack2Capable                              AccountDataSourceType = 'T'
		// ManuallyKeyedTrack1Capable                              AccountDataSourceType = 'X'
		// ManuallyKeyedTerminalHasNoCardReadingCapability         AccountDataSourceType = '@'
		// ChipCardTransProcessedAsMagStripe_NoEmvAppOnTerminal    AccountDataSourceType = 'W'
		// ChipCardTransProcessedAsMagStripe_CardOrTerminalFailure AccountDataSourceType = 'Z'
		cardholderidentification = NewCardholderIdentificationDataField(CardholderSignatureTerminalHasPinPad, r.AVSStreet1, r.AVSZip, r.EncryptedPINData, r.PINPadSequenceNumber)

	} else if 0 < len(r.AccountNumber) {
		var ads AccountDataSourceType = ManuallyKeyedTrack2Capable
		if 0 < len(r.AccountDataSource) {
			ads = AccountDataSourceType(r.AccountDataSource[0])
		}

		var month, year uint = 0, 0
		switch len(r.ExpirationDate) {
		case 4:
			month = estrings.StringToUint(r.ExpirationDate[0:2])
			year = estrings.StringToUint(r.ExpirationDate[2:]) + 2000
		case 6:
			month = estrings.StringToUint(r.ExpirationDate[0:2])
			year = estrings.StringToUint(r.ExpirationDate[2:])
		}

		customeraccountdata = NewManuallyEnteredCustomerDataField(ads, r.AccountNumber, month, year, r.CVV)
		cardholderidentification = NewCardholderIdentificationDataField(CardholderSignatureTerminalHasPinPad, r.AVSStreet1, r.AVSZip, r.EncryptedPINData, r.PINPadSequenceNumber)
		issuer = card.GetCardIssuerType(r.AccountNumber)
	}

	// PersonalIdentificationNumber32CharStaticKeyNonUsa          CardholderIdCodeType = 'A'
	// PinAtAutomatedDispensingMachine32CharStaticKey             CardholderIdCodeType = 'B'
	// SelfServiceLimitedAmountTerminalNoIdMethodAvailable        CardholderIdCodeType = 'C'
	// SelfServiceTerminalNoIdMethodAvailable                     CardholderIdCodeType = 'D'
	// CustomerActivatedAutomatedFuelDispenserNoIdMethodAvailable CardholderIdCodeType = 'E'
	// PinAuthenticationByIccChipCard                             CardholderIdCodeType = 'F'
	// ClearTextPin                                               CardholderIdCodeType = 'G'
	// PinAtAutomatedDispensingMachine32CharDukpt                 CardholderIdCodeType = 'J'
	// PersonalIdentificationNumber32CharDukpt                    CardholderIdCodeType = 'K'
	// CardPresentUnableToReadMagStripeSendingAvsData             CardholderIdCodeType = 'M'
	// CardNotPresentIncludesEcomAndFullAvsData                   CardholderIdCodeType = 'N'
	// NoCvmRequiredEMVContactOrContactless                       CardholderIdCodeType = 'P'
	// PersonalIdentificationNumber32CharStaticKey                CardholderIdCodeType = 'S'
	// CardholderSignatureTerminalHasPinPad                       CardholderIdCodeType = 'Z'
	// CardholderSignatureNoPinPadAvailable                       CardholderIdCodeType = '@'

	return customeraccountdata, cardholderidentification, issuer
}

func CreateAuthorization(t *TerminalConfig, r *auth.RequestMessage) *AuthorizationMessageRequest {

	amount := estrings.StringDecimalAmountToUint64(r.TransactionAmount)
	cashbackamount := estrings.StringDecimalAmountToUint64(r.CashbackAmount)

	var transactCode transactionCodeType
	switch r.Action {
	case "AUTHORIZE_ONLINE":
		transactCode = Purchase
	case "AUTHORIZE_ONLINE":
		transactCode = Purchase

		// Purchase                                          TransactionCodeType = "54"
		// PurchaseRepeat                                    TransactionCodeType = "64"
		// PurchaseCardNotPresent                            TransactionCodeType = "56" // Used primarily in Direct Marketing, Mail, and Telephone Order Environments
		// PurchaseCardNotPresentRepeat                      TransactionCodeType = "66"
		// CardAuthentication                                TransactionCodeType = "58" // To determine if an account is open for use or to verify cardholder information. Transaction amounts must be zero filled. AVS and CVV2 may be attempted. (Visa, MasterCard, American Express, Discover, PayPal only)
		// CardAuthenticationRepeat                          TransactionCodeType = "68"
		// OnlineAuthorizationReversal                       TransactionCodeType = "59" // Used to reverse a completed credit authorization prior to batch settlement.
		// OnlineAuthorizationReversalRepeat                 TransactionCodeType = "69"
		// StoreAndForwardAuthorizationReversal              TransactionCodeType = "5A" // Same as code “59,” except the reversal is being submitted after batch settlement occurred
		// StoreAndForwardAuthorizationReversalRepeat        TransactionCodeType = "6A"
		// BillPayTransaction                                TransactionCodeType = "5B" // By Credit Card only
		// BillPayTransactionRepeat                          TransactionCodeType = "6B"
		// CreditAdvice                                      TransactionCodeType = "5C" // Advice message for an AFD final sale amount (Visa and MasterCard only)
		// CreditAccountFundingOrPayment                     TransactionCodeType = "5G" // Credit card used for payment of account to account transactions
		// CreditAccountFundingOrPaymentRepeat               TransactionCodeType = "6G" // Credit Account Funding Transaction (Visa), Credit Payment Transaction (MC)
		// CardNotPresentCreditAccountFundingOrPayment       TransactionCodeType = "5H" // Card-not-present credit payment of account to account transactions
		// CardNotPresentCreditAccountFundingOrPaymentRepeat TransactionCodeType = "6H"
		// CardPresentCreditCardholderFundsTransfer          TransactionCodeType = "5J" // Card-present credit payment of cardholder funds transfer
		// CardPresentCreditCardholderFundsTransferRepeat    TransactionCodeType = "6J"
		// CardholderFundsTransferCardNotPresent             TransactionCodeType = "5K"
		// CardholderFundsTransferCardNotPresentRepeat       TransactionCodeType = "6K"

		// FoodStampsReturnEbt                               TransactionCodeType = "92"
		// DirectDebitPurchase                               TransactionCodeType = "93"
		// DirectDebitPurchaseReturn                         TransactionCodeType = "94"
		// CashBenefitsCashWithdrawalEbt                     TransactionCodeType = "96"
		// FoodStampPurchaseEbt                              TransactionCodeType = "98"
		// DirectDebitBalanceInquiry                         TransactionCodeType = "9A"

		// DebitBillPaymentTransaction        TransactionCodeType = "9B"
		// PinlessDebitBillPaymentTransaction TransactionCodeType = "9C" // Debit transactions without a PIN. The Network ID of 0000 can not be used.

	}

	G1 := NewG1AuthorizationMessageRequest(t.AcquirerBIN, Purchase, t.TransactionSequenceNumber, amount, cashbackamount)
	G1.SetTerminalIdentification(t.MerchantNumber, t.StoreNumber, t.TerminalNumber, t.DeviceCode, t.IndustryCode, t.MerchantCategoryCode)
	G1.SetTerminalLocale(t.CurrencyCode, t.CountryCode, t.LanguageIndicator, t.CityCode, t.TimeZoneDiff) //cc := ParseCreditCard("%B4123456789012349^MAVERICK INTERNATIONAL VSA^17121010000000000000000?;4123456789012349=171210100000000000?

	customeraccountdata, cardholderidentification, issuer := CreateCardholderData(t, r)

	G1.SetCardholderData(customeraccountdata, cardholderidentification)
	G1.CardAcceptorData = NewCardAcceptorDataField(t.MerchantName, t.MerchantCity, t.MerchantState, "", t.CardholderServiceTelephoneNo)

	g3records := CreateGroup3Elements(t, r, issuer)

	return NewAuthorizationMessageRequest("", G1, nil, g3records)
}

func CreateReversal(t *TerminalConfig, r *auth.RequestMessage) *AuthorizationMessageRequest {

	amount := estrings.StringDecimalAmountToUint64(r.TransactionAmount)
	cashbackamount := estrings.StringDecimalAmountToUint64(r.CashbackAmount)

	G1 := NewG1AuthorizationMessageRequest(t.AcquirerBIN, DirectDebitPurchase, t.TransactionSequenceNumber, amount, cashbackamount)
	G1.SetTerminalIdentification(t.MerchantNumber, t.StoreNumber, t.TerminalNumber, t.DeviceCode, t.IndustryCode, t.MerchantCategoryCode)
	G1.SetTerminalLocale(t.CurrencyCode, t.CountryCode, t.LanguageIndicator, t.CityCode, t.TimeZoneDiff) //cc := ParseCreditCard("%B4123456789012349^MAVERICK INTERNATIONAL VSA^17121010000000000000000?;4123456789012349=171210100000000000?

	customeraccountdata, cardholderidentification, issuer := CreateCardholderData(t, r)

	G1.SetCardholderData(customeraccountdata, cardholderidentification)
	G1.CardAcceptorData = NewCardAcceptorDataField(t.MerchantName, t.MerchantCity, t.MerchantState, "", t.CardholderServiceTelephoneNo)

	g3records := CreateGroup3Elements(t, r, issuer)

	return NewAuthorizationMessageRequest("", G1, nil, g3records)
}
