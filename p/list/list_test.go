package list

import (
	"testing"
)

func TestParseIniFileMap(t *testing.T) {

	s := `[Contact]
	adminemail=abc@abc.com
	email=abc@abc.com


	[CreditCardOptions]
	AllowDuplicates=false


	[Deposit]
	AccountNumber=00099999999
	BankName=BOA
	RoutingNumber=099000000


	[Limits]
	AATAllowance=50
	ADDAAllowance=50
	ADDCAllowance=50
	ATPDAllowance=50
	AvgAverageTicket_90=0
	AvgDailyDepositAmount_90=0
	AvgDailyDepositCount_90=0
	AvgTransactionsPerDeposit_90=0
	MaxAuthorizationAmount=10000
	MaxCashbackAmount=200
	MaxCreditAmount=5000
	MaxGratuityAmount=1000
	MaxSettlementAmount=10000


	[Organization]
	ClientId=9999
	CompanyName=XYZ Corp
	DBAName=XYZ Corp
	MasterKey=9999999


	[Pmoney]
	BusinessId=99999


	[Products]
	ACH=true
	CreditCard=true
	DebitCard=true
	EFT=true
	GiftCard=true


	[Rates]
	DR_MC_Spread_CE=0
	DR_Original_C=0
	DR_VISA_Spread_CE=0
	DiscountRate_C=0
	DiscountRate_E=0
	DiscountRate_P=0
	PI_Original_C=0
	PI_Spread_CE=0
	PartnerPoints=0
	PerItemFee_C=0
	PerItemFee_E=0
	PerItemFee_P=0
	SetupFee=0
	YearlyFee=0


	[Tax]
	LocalTaxRate=0.03


	[Terminal]
	AcquirerBIN=999995
	AgentBankNumber=599999
	AgentChainNumber=599999
	BatchNumber=889
	CityCode=22033
	CountryCode=840
	CurrencyCode=840
	DeviceCode=ReservedForThirdPartyDevelopers
	IndustryCode=Retail
	LanguageIndicator=English
	MerchantCategoryCode=5999
	MerchantCity=FAIRFAX
	CardholderServiceTelephoneNo=8664321100
	MerchantLocalTelephoneNo=8008347790
	MerchantName=IKEA Store
	MerchantNumber=999999999911
	MerchantState=VA
	PosDataCodeProfile=027H11101C13336
	RequestedACI=Y
	StoreNumber=1515
	TerminalNumber=5999
	TimeZoneDiff=705
	TransactionSequenceNumber=1108
	`
	cfg := ParseIniFileMap(s)

	t.Logf("%s\n\n", cfg.String())

	t.Logf("%s\n\n", cfg.Serialize())

}

// func TestIniFileMap() {

// 	cfg := list.LoadIniFileMap("merchant.ini")

// 	fmt.Printf("%s\n\n", cfg.String())

// 	fmt.Printf("%s\n\n", cfg.Serialize())

// 	cfg.Save("merchantcopy.ini")
// }

// var sampleinifile string = `
// [Organization]
// ClientId=9999
// MasterKey=999999
// CompanyName=XYZ Corp
// DBAName=XYZ Corp

// [Terminal]
// AcquirerBIN=999995
// MerchantNumber=999999999911
// StoreNumber=1515
// TerminalNumber=5999
// DeviceCode=ReservedForThirdPartyDevelopers
// IndustryCode=Retail
// CurrencyCode=840
// CountryCode=840
// CityCode=22033
// LanguageIndicator=English
// TimeZoneDiff=705
// MerchantCategoryCode=5999
// RequestedACI=Y
// TransactionSequenceNumber=1108
// AgentBankNumber=599999
// AgentChainNumber=599999
// BatchNumber=889
// MerchantLocalTelephoneNo=8008347790
// CardholderServiceTelephoneNo=8664321100

// [CreditCardOptions]
// AllowDuplicates=false

// [Rates]
// PerItemFee_E=0
// DiscountRate_E=0
// PerItemFee_C=0
// DiscountRate_C=0
// PerItemFee_P=0
// DiscountRate_P=0
// DR_Original_C=0
// PI_Original_C=0
// DR_VISA_Spread_CE=0
// DR_MC_Spread_CE=0
// PI_Spread_CE=0
// SetupFee=0
// YearlyFee=0
// PartnerPoints=0

// [Tax]
// LocalTaxRate=0.03

// [Limits]
// MaxAuthorizationAmount=10000
// MaxCreditAmount=5000
// MaxGratuityAmount=1000
// MaxCashbackAmount=200
// MaxSettlementAmount=10000
// AvgDailyDepositAmount_90=0
// AvgAverageTicket_90=0
// AvgTransactionsPerDeposit_90=0
// AvgDailyDepositCount_90=0
// ADDAAllowance=50
// AATAllowance=50
// ATPDAllowance=50
// ADDCAllowance=50

// [Contact]
// adminemail=abc@abc.com
// email=abc@abc.com

// [Products]
// CreditCard=true
// DebitCard=true
// EFT=true
// ACH=true
// GiftCard=true

// [Deposit]
// BankName=BOA
// RoutingNumber=999999999
// AccountNumber=00099999999

// [Pmoney]
// BusinessId=99999
// `
