package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

const ADDITIONALAMOUNTINFOLEN = 25

type AdditionalAmountInfo struct {
	AccountType AdditionalAmountAccountType //2 NUM 4.9.1 First Additional Amount Account Type
	// 1 ASCII 4.80 Field Separator <FS>
	AmountType AdditionalAmountType // 2 NUM 4.9.2 First Additional Amount Amount Type
	// 1 ASCII 4.80 Field Separator <FS>
	CurrencyCode CurrencyCodeType // 3 NUM 4.9.3 First Additional Amount Currency Code
	// 1 ASCII 4.80 Field Separator <FS>
	Sign AdditionalAmountSignType // 1 ALPHA 4.9.4 First Additional Amount Sign
	// 1 ASCII 4.80 Field Separator <FS>
	Amount uint64 // 12 N 4.9.5 First Additional Amount
	// 1 ASCII 4.80 Field Separator <FS>
}

// AdditionalAmountAccountType -> AccountNotSpecified; //2 NUM
// AdditionalAmountType -> DepositAccountAvailableBalance; //2 NUM
// CurrencyCodeType -> UsdUnitedStatesUsDollar; //3 NUM
// AdditionalAmountSignType -> PositiveBalance; //1 A/N

func NewAdditionalAmountInfo(typeaccount AdditionalAmountAccountType, typeamount AdditionalAmountType, currency CurrencyCodeType, amountsign AdditionalAmountSignType, txnamount uint64) *AdditionalAmountInfo {
	return &AdditionalAmountInfo{typeaccount, typeamount, currency, amountsign, txnamount}
}

func (a *AdditionalAmountInfo) Copy(addamt *AdditionalAmountInfo) {
	a.AccountType = addamt.AccountType
	a.AmountType = addamt.AmountType
	a.CurrencyCode = addamt.CurrencyCode
	a.Sign = addamt.Sign
	a.Amount = addamt.Amount
}

func (a *AdditionalAmountInfo) String() string {
	return fmt.Sprintf("%02d%c%2s%c%03d%c%c%c%012d%c", uint(a.AccountType), bytes.FileSeparator,
		string(a.AmountType), bytes.FileSeparator,
		uint(a.CurrencyCode), bytes.FileSeparator,
		a.Sign, bytes.FileSeparator,
		a.Amount, bytes.FileSeparator)
}

func ParseAdditionalAmountInfo(s string) *AdditionalAmountInfo {
	a := AdditionalAmountInfo{AccountNotSpecified, DepositAccountAvailableBalance, UnitedStatesUsDollar, PositiveBalance, 0}

	if ADDITIONALAMOUNTINFOLEN > len(s) {
		panic("AdditionalAmountInfo string must be at least 25 characters long")
	}

	var fs1, fs2, fs3, fs4, fs5 byte = 0, 0, 0, 0, 0

	fmt.Sscanf(s, "%02d%c%02d%c%03d%c%c%c%012d%c",
		&a.AccountType,
		&fs1,
		&a.AmountType,
		&fs2,
		&a.CurrencyCode,
		&fs3,
		&a.Sign,
		&fs4,
		&a.Amount,
		&fs5)

	return &a
}
