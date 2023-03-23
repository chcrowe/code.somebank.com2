package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"code.somebank.com/p/auth"
	"code.somebank.com/p/tsys"
)

func main() {
	DoNvpRequest()
	Do1080Request()
	DoXmlRequest()
}

func DoNvpRequest() {
	request := auth.RequestMessage{Action: "AUTHORIZE_ONLINE", ClientID: "9999", MasterKey: "999999", TrackData: "9999999999991234=101010100000000000", TransactionAmount: "10.00"}
	request.LocalTransactionDate = time.Now().String()
	resp, e := http.Post("http://127.0.0.1:8080/POSCP3/main.asp", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(request.Values())))
	if nil != e {
		panic(e)
	}

	body, e2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("NVP recv: %s\n\n", body)
	if nil != e2 {
		panic(e2)
	}
}

func DoXmlRequest() {
	request := auth.RequestMessage{Action: "AUTHORIZE_ONLINE", ClientID: "9999", MasterKey: "999999", TrackData: "4123456789012349=171210100000000000", TransactionAmount: "10.00"}
	request.LocalTransactionDate = time.Now().String()
	resp, e := http.Post("http://127.0.0.1:8080", "application/xml", bytes.NewBuffer([]byte(request.Xml())))
	if nil != e {
		panic(e)
	}

	body, e2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("XML recv: %s\n\n", body)
	if nil != e2 {
		panic(e2)
	}
}

func Do1080Request() {
	s := CreateVisaRetailAuthorization()
	resp, e := http.Post("http://127.0.0.1:8080", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(s)))
	if nil != e {
		panic(e)
	}

	body, e2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("1080 recv: %s\n\n", body)
	if nil != e2 {
		panic(e2)
	}
}

func CreateVisaRetailAuthorization() string {

	G1 := tsys.NewG1AuthorizationMessageRequest(999995, tsys.Purchase, 1, 8800, 0)
	G1.SetTerminalIdentification(999999999911, 11, 9911, tsys.ReservedForThirdPartyDevelopers, tsys.Retail, tsys.MiscellaneousAndSpecialtyRetailStores)
	G1.SetTerminalLocale(tsys.UsdUnitedStatesUsDollar, tsys.UnitedStates, tsys.English, "22030", tsys.NewTimeZoneDifferential(tsys.HoursBehindGMTWithDaylightSavings, 5)) //cc := ParseCreditCard("%B4123456789012349^MAVERICK INTERNATIONAL VSA^17121010000000000000000?;4123456789012349=171210100000000000?
	customeraccountdata := tsys.NewCustomerDataField(tsys.FullMagneticStripeReadAndTransmitTrack2, "", "9999999999991234=101010100000000000")
	cardholderidentification := tsys.NewCardholderIdentificationDataField(tsys.CardholderSignatureTerminalHasPinPad, "", "22031", "", "")
	G1.SetCardholderData(customeraccountdata, cardholderidentification)
	G1.CardAcceptorData = tsys.NewCardAcceptorDataField("IKEA STORE", "FAIRFAX", "VA", "", 8664445555)

	additionalAmounts := [...]tsys.AdditionalAmountInfo{*(tsys.NewAdditionalAmountInfo(tsys.CreditCardAccount, tsys.DepositAccountAvailableBalance, tsys.UsdUnitedStatesUsDollar, tsys.PositiveBalance, 500)),
		*(tsys.NewAdditionalAmountInfo(tsys.CreditCardAccount, tsys.DepositAccountAvailableBalance, tsys.UsdUnitedStatesUsDollar, tsys.PositiveBalance, 700)),
		*(tsys.NewAdditionalAmountInfo(tsys.CreditCardAccount, tsys.DepositAccountAvailableBalance, tsys.UsdUnitedStatesUsDollar, tsys.PositiveBalance, 900)),
		*(tsys.NewAdditionalAmountInfo(tsys.CreditCardAccount, tsys.DepositAccountAvailableBalance, tsys.UsdUnitedStatesUsDollar, tsys.PositiveBalance, 1100))}

	g3records := []string{
		tsys.NewG3v001CommercialCardRequest().String(),
		//NewG3v007CardVerificationCodeRequest(Visa, "123").String(),
		tsys.NewG3v020DeveloperInformationRequest().String(),
		tsys.NewG3v022AdditionalAmountsRequest(&additionalAmounts).String(),
		tsys.NewG3v021MerchantVerificationValueRequest("1EB2C1772D").String(),
		tsys.NewG3v025TransactionFeeAmountRequest(tsys.DebitTransactionFee, 150).String(),
		tsys.NewG3v027PosDataRequest(tsys.TerminalInputICCReaderAndContactlessTerminal,
			tsys.CardholderAuthenticationPINEntryCapability,
			tsys.CardCaptureCapability,
			tsys.EnvironmentOnCardAcceptorPremisesAttendedTerminal,
			tsys.CardholderPresent,
			tsys.CardPresent,
			tsys.CardInputOnlineChip,
			tsys.CardholderPINAuthenticated,
			tsys.AuthenticationEntityAuthorizingAgentOnlinePIN,
			tsys.CardOutputCapabilityICC,
			tsys.TerminalOutputCapabilityICC,
			tsys.PINCaptureCapability06CharactersMaximum).String(),
		tsys.NewG3v032CurrencyConversionDataRequest().String(),
		tsys.NewG3v034CardProductCodeRequest().String(),
		tsys.NewG3v040VisaISAChargeIndicatorRequest().String(),
		tsys.NewG3v046CardTypeGroupRequest().String(),
		tsys.NewG3v061SpendQualifiedIndicatorRequest().String()}

	//NewG3v048AmexCardholderVerificationResultsRequest().String()
	//NewG3v050AssociationTimestampRequest().String() //mastercard
	//NewG3v051MasterCardEventMonitoringRealTimeScoringRequest(1).String() //mastercard

	//NewG3v058AlternateAccountIDRequest().String() //mastercard combine w G3v059
	//NewG3v059MasterCardPayPassMappingServiceRequest().String()

	authrequest := tsys.NewAuthorizationMessageRequest("", G1, nil, g3records)

	return authrequest.Buffer.String()
}
