package tsys

import (
	"strings"

	"code.somebank.com/p/list"
)

type TerminalConfig struct {
	AcquirerBIN    uint64           //6 NUM 4.2 (999995)
	MerchantNumber uint64           //12 NUM 4.59 (999999999911)
	StoreNumber    uint             //4 NUM 4.81 (1515)
	TerminalNumber uint             //4 NUM 4.84 (5999)
	DeviceCode     DeviceCodeType   //1 A/N 4.34 (ReservedForThirdPartyDevelopers)
	IndustryCode   IndustryCodeType //1 A/N 4.46 (UnknownOrUnsure)
	CurrencyCode   CurrencyCodeType //3 NUM 4.30 (UsdUnitedStatesUsDollar)
	CountryCode    CountryCodeType  //3 NUM 4.28 (UnitedStates)
	CityCode       string           //9 A/N 4.25 City Code must be one of the following three formats...
	// U.S. ZIP Code 5 character numeric, left-justified, space-filled.
	// U.S. ZIP Code + Extension 9 character numeric.
	// Canadian Postal Code 6 character "AnAnAn" format, left-justified, spacefilled.

	LanguageIndicator    LanguageIndicatorType //2 NUM 4.51 (English)
	TimeZoneDiff         *TimeZoneDifferential //3 NUM 4.86 (705 EASTERN)
	MerchantCategoryCode MerchantCategoryType  //4 NUM 4.57

	RequestedACI              RequestAciType //1 A/N 4.69 (DeviceIsCpsMeritCapableOrCreditOrOffline)
	TransactionSequenceNumber uint           //4 NUM 4.92

	AgentBankNumber              uint   //6 NUM
	AgentChainNumber             uint   //6 NUM
	BatchNumber                  uint   //3 NUM 4.18
	MerchantLocalTelephoneNo     uint64 //11 A/N 4.119
	CardholderServiceTelephoneNo uint64 //11 A/N 4.34

	MerchantCity  string //LL.....LL 26-38 A/N
	MerchantName  string //NNNN... 1-25 A/N left-justified and space-filled
	MerchantState string //SS 39-40 A/N chars must be upper case

	PosDataCodeProfile string
}

func UnmarshallFromIniFileMap(m *list.IniFileMap, t *TerminalConfig) {

	t.AcquirerBIN = m.GetUint64("Terminal.AcquirerBIN")
	t.MerchantNumber = m.GetUint64("Terminal.MerchantNumber")
	t.StoreNumber = m.GetUint("Terminal.StoreNumber")
	t.TerminalNumber = m.GetUint("Terminal.TerminalNumber")
	t.DeviceCode = DeviceCodeType(m.GetUint64("Terminal.DeviceCode"))
	t.IndustryCode = IndustryCodeType(m.GetUint64("Terminal.IndustryCode"))
	t.CurrencyCode = CurrencyCodeType(m.GetUint64("Terminal.CurrencyCode"))
	t.CountryCode = CountryCodeType(m.GetUint64("Terminal.CountryCode"))
	t.CityCode = m.Get("Terminal.CityCode")
	t.LanguageIndicator = LanguageIndicatorType(m.GetUint("Terminal.LanguageIndicator"))
	t.TimeZoneDiff = ParseTimeZoneDifferential(m.Get("Terminal.TimeZoneDiff"))
	t.MerchantCategoryCode = MerchantCategoryType(m.GetUint("Terminal.MerchantCategoryCode"))
	t.TransactionSequenceNumber = m.GetUint("Terminal.TransactionSequenceNumber")
	t.AgentBankNumber = m.GetUint("Terminal.AgentBankNumber")
	t.AgentChainNumber = m.GetUint("Terminal.AgentChainNumber")
	t.BatchNumber = m.GetUint("Terminal.BatchNumber")
	t.MerchantLocalTelephoneNo = m.GetUint64("Terminal.MerchantLocalTelephoneNo")
	t.CardholderServiceTelephoneNo = m.GetUint64("Terminal.CardholderServiceTelephoneNo")

	t.MerchantCity = strings.ToUpper(m.Get("Terminal.MerchantCity"))
	t.MerchantName = strings.ToUpper(m.Get("Terminal.MerchantName"))
	t.MerchantState = strings.ToUpper(m.Get("Terminal.MerchantState"))

	aci := m.Get("Terminal.RequestedACI")
	if 0 < len(aci) {
		t.RequestedACI = RequestAciType(aci[0])
	}

	t.PosDataCodeProfile = strings.ToUpper(m.Get("Terminal.PosDataCodeProfile"))
}
