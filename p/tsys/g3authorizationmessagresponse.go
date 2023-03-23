package tsys

import (
	"fmt"
	"strconv"
	"strings"

	"code.somebank.com/p/bytes"
)

// based upon EIS 1080 VERSION 12.0 FEBRUARY 11, 2014
type CommercialCardResponseType rune

const (
	BusinessCard                    CommercialCardResponseType = 'B'
	VisaCommerceReserved            CommercialCardResponseType = 'D'
	CorporateCard                   CommercialCardResponseType = 'R'
	PurchasingCard                  CommercialCardResponseType = 'S'
	NonCommercialCard               CommercialCardResponseType = '0'
	InvalidRequestIndicatorReceived CommercialCardResponseType = ' '
)

func (x CommercialCardResponseType) String() string {
	switch x {
	case BusinessCard:
		return "BusinessCard"
	case VisaCommerceReserved:
		return "VisaCommerceReserved"
	case CorporateCard:
		return "CorporateCard"
	case PurchasingCard:
		return "PurchasingCard"
	case NonCommercialCard:
		return "NonCommercialCard"
	case InvalidRequestIndicatorReceived:
		return "InvalidRequestIndicatorReceived"
	}
	return ""
}

// ----------------------------------------------------
type G3v001CommercialCardResponse struct {
	VersionNumber                   G3VersionType              // 3 NUM 4.85 Group III Version Number 001
	CommercialCardResponseIndicator CommercialCardResponseType // 0-1 A/N 4.52 Commercial Card Response Indicator B,R,L,S,D,0
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v001CommercialCardResponse(indicator CommercialCardResponseType) *G3v001CommercialCardResponse {
	return &G3v001CommercialCardResponse{G3v001CommercialCard, indicator}
}

func (g *G3v001CommercialCardResponse) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.CommercialCardResponseIndicator)
}

func ParseG3v001CommercialCardResponse(s string) *G3v001CommercialCardResponse {
	g := G3v001CommercialCardResponse{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.CommercialCardResponseIndicator)
	return &g
}

// ----------------------------------------------------
type G3v007CardVerificationCodeResponse struct {
	VersionNumber              G3VersionType      // 3 NUM 4.85 Group III Version Number 007
	VerificationCodeResultCode Cvv2ResultCodeType // 0-1 A/N 4.176 Verification Code Result Code
}

func NewG3v007CardVerificationCodeResponse(code Cvv2ResultCodeType) *G3v007CardVerificationCodeResponse {
	return &G3v007CardVerificationCodeResponse{G3v007CardVerificationCode, code}
}

func (g *G3v007CardVerificationCodeResponse) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.VerificationCodeResultCode)
}

func ParseG3v007CardVerificationCodeResponse(s string) *G3v007CardVerificationCodeResponse {
	g := G3v007CardVerificationCodeResponse{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.VerificationCodeResultCode)
	return &g
}

// ----------------------------------------------------
type G3v008FleetFuelingCardResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 008
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v008FleetFuelingCardResponse() *G3v008FleetFuelingCardResponse {
	return &G3v008FleetFuelingCardResponse{G3v008FleetFuelingCard}
}

func (g *G3v008FleetFuelingCardResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v008FleetFuelingCardResponse(s string) *G3v008FleetFuelingCardResponse {
	g := G3v008FleetFuelingCardResponse{VersionNumber: G3v008FleetFuelingCard}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v009SeteCommerceResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 009
	//1 ASCII 4.86 Group Separator <GS>
}

func NewG3v009SeteCommerceResponse() *G3v009SeteCommerceResponse {
	return &G3v009SeteCommerceResponse{G3v009SeteCommerce}
}

func (g *G3v009SeteCommerceResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v009SeteCommerceResponse(s string) *G3v009SeteCommerceResponse {
	g := G3v009SeteCommerceResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
// reserved for future use
type G3v010CcpsResponse struct {
	VersionNumber                   G3VersionType //3 NUM 4.85 Group III Version Number 010
	AuthorizationResponseCryptogram string        // 16 A/N 4.27 Authorization Response Cryptogram
	// 1 ASCII 4.80 Field Separator <FS>
	IssuerScript string // 0,512 A/N 4.96 Issuer Script
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v010CcpsResponse() *G3v010CcpsResponse {
	return &G3v010CcpsResponse{VersionNumber: G3v010Ccps}
}

func (g *G3v010CcpsResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v010CcpsResponse(s string) *G3v010CcpsResponse {
	g := G3v010CcpsResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	g.AuthorizationResponseCryptogram = fields[0]
	g.IssuerScript = fields[1]

	return &g
}

// ----------------------------------------------------
type G3v011ChipConditionResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 011
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v011ChipConditionResponse() *G3v011ChipConditionResponse {
	return &G3v011ChipConditionResponse{G3v011ChipCondition}
}

func (g *G3v011ChipConditionResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v011ChipConditionResponse(s string) *G3v011ChipConditionResponse {
	g := G3v011ChipConditionResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

//----------------------------------------------------
// type G3v012CommercialCardLargeTicketResponse struct {
// 	VersionNumber G3VersionType //3 NUM
// }

// func NewG3v012CommercialCardLargeTicketResponse() *G3v012CommercialCardLargeTicketResponse {
// 	return &G3v012CommercialCardLargeTicketResponse{G3v012CommercialCardLargeTicket, ""}
// }

// func (g *G3v012CommercialCardLargeTicketResponse) String() string {
// 	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.xxx)
// }

// ----------------------------------------------------
type G3v013ElectronicBenefitsTransferResponse struct {
	VersionNumber    G3VersionType // 3 NUM 4.85 Group III Version Number 013
	AvailableBalance uint64        // 0,12 NUM 4.30 Available Balance
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v013ElectronicBenefitsTransferResponse(balance uint64) *G3v013ElectronicBenefitsTransferResponse {
	return &G3v013ElectronicBenefitsTransferResponse{G3v013ElectronicBenefitsTransfer, balance}
}

func (g *G3v013ElectronicBenefitsTransferResponse) String() string {
	return fmt.Sprintf("%03d%012d", g.VersionNumber, g.AvailableBalance)
}

func ParseG3v013ElectronicBenefitsTransferResponse(s string) *G3v013ElectronicBenefitsTransferResponse {
	g := G3v013ElectronicBenefitsTransferResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	if 3 < len(s) {
		g.AvailableBalance, _ = strconv.ParseUint(s[3:], 10, 64)
	}

	return &g
}

// ----------------------------------------------------
type G3v014MotoEcommerceResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 014
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v014MotoEcommerceResponse() *G3v014MotoEcommerceResponse {
	return &G3v014MotoEcommerceResponse{G3v014MotoEcommerce}
}

func (g *G3v014MotoEcommerceResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v014MotoEcommerceResponse(s string) *G3v014MotoEcommerceResponse {
	g := G3v014MotoEcommerceResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v015ServiceDevelopmentIndicatorResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 015
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v015ServiceDevelopmentIndicatorResponse() *G3v015ServiceDevelopmentIndicatorResponse {
	return &G3v015ServiceDevelopmentIndicatorResponse{G3v015ServiceDevelopmentIndicator}
}

func (g *G3v015ServiceDevelopmentIndicatorResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v015ServiceDevelopmentIndicatorResponse(s string) *G3v015ServiceDevelopmentIndicatorResponse {
	g := G3v015ServiceDevelopmentIndicatorResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v017Visa3dSecurEcomVerificationResponse struct {
	VersionNumber   G3VersionType      // 3 NUM 4.85 Group III Version Number 017
	CAVVResultsCode CavvResultCodeType // 0 or 1 A/N 4.47 CAVV Results Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v017Visa3dSecurEcomVerificationResponse(code CavvResultCodeType) *G3v017Visa3dSecurEcomVerificationResponse {
	return &G3v017Visa3dSecurEcomVerificationResponse{G3v017Visa3dSecurEcomVerification, code}
}

func (g *G3v017Visa3dSecurEcomVerificationResponse) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.CAVVResultsCode)
}

func ParseG3v017Visa3dSecurEcomVerificationResponse(s string) *G3v017Visa3dSecurEcomVerificationResponse {
	g := G3v017Visa3dSecurEcomVerificationResponse{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.CAVVResultsCode)
	return &g
}

// ----------------------------------------------------
type G3v018ExistingDebtIndicatorResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 018
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v018ExistingDebtIndicatorResponse() *G3v018ExistingDebtIndicatorResponse {
	return &G3v018ExistingDebtIndicatorResponse{G3v018ExistingDebtIndicator}
}

func (g *G3v018ExistingDebtIndicatorResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v018ExistingDebtIndicatorResponse(s string) *G3v018ExistingDebtIndicatorResponse {
	g := G3v018ExistingDebtIndicatorResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v019MastercardUniversalCardholderAuthenticationResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 019
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v019MastercardUniversalCardholderAuthenticationResponse() *G3v019MastercardUniversalCardholderAuthenticationResponse {
	return &G3v019MastercardUniversalCardholderAuthenticationResponse{G3v019MastercardUniversalCardholderAuthentication}
}

func (g *G3v019MastercardUniversalCardholderAuthenticationResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v019MastercardUniversalCardholderAuthenticationResponse(s string) *G3v019MastercardUniversalCardholderAuthenticationResponse {
	g := G3v019MastercardUniversalCardholderAuthenticationResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v020DeveloperInformationResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 020
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v020DeveloperInformationResponse() *G3v020DeveloperInformationResponse {
	return &G3v020DeveloperInformationResponse{G3v020DeveloperInformation}
}

func (g *G3v020DeveloperInformationResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v020DeveloperInformationResponse(s string) *G3v020DeveloperInformationResponse {
	g := G3v020DeveloperInformationResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v021MerchantVerificationValueResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 021
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v021MerchantVerificationValueResponse() *G3v021MerchantVerificationValueResponse {
	return &G3v021MerchantVerificationValueResponse{G3v021MerchantVerificationValue}
}

func (g *G3v021MerchantVerificationValueResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v021MerchantVerificationValueResponse(s string) *G3v021MerchantVerificationValueResponse {
	g := G3v021MerchantVerificationValueResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
// Group 3 Version 22 for Additional Amounts can support on the request either no additional
// amounts or four field separated additional amounts. If the transaction does not require
// additional amounts in the request but requires additional amounts in the response, the POS
// device should send only the 022 version number. If the transaction does require additional
// amounts in the request, the POS device should send the 022 version number plus the field
// separated additional amounts. The POS device should send the field separators for all four
// additional amounts even if some of the amount data fields are not used. All five subfields for
// an amount must be present for the amount to be valid.
type G3v022AdditionalAmountsResponse struct {
	VersionNumber     G3VersionType //3 NUM
	AdditionalAmounts [4]AdditionalAmountInfo
}

func NewG3v022AdditionalAmountsResponse(amounts []AdditionalAmountInfo) *G3v022AdditionalAmountsResponse {
	g := G3v022AdditionalAmountsResponse{VersionNumber: G3v022AdditionalAmounts}

	for n := 0; n < 4 && n < len(amounts); n++ {
		g.AdditionalAmounts[n].Copy(&amounts[n])
	}

	return &g
}

func (g *G3v022AdditionalAmountsResponse) String() string {
	return fmt.Sprintf("%03d%s%s%s%s", g.VersionNumber, g.AdditionalAmounts[0].String(), g.AdditionalAmounts[1].String(),
		g.AdditionalAmounts[2].String(), g.AdditionalAmounts[3].String())
}

func ParseG3v022AdditionalAmountsResponse(s string) *G3v022AdditionalAmountsResponse {
	g := G3v022AdditionalAmountsResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	n := 0
	for i := 3; i < (ADDITIONALAMOUNTINFOLEN * 4); i += ADDITIONALAMOUNTINFOLEN {
		amt := ParseAdditionalAmountInfo(s[i : i+ADDITIONALAMOUNTINFOLEN])
		g.AdditionalAmounts[n].Copy(amt)
		n++
	}

	return &g
}

// ----------------------------------------------------
// Response message - Visa and MasterCard healthcare data (version 023)
type G3v023VisaMastercardHealthcareResponse struct {
	VersionNumber  G3VersionType // 3 NUM 4.85 Group III Version Number 023
	HealthcareData string        // III 0-199* ANS 4.87 Healthcare, MasterCard or Discover Member Defined Data
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v023VisaMastercardHealthcareResponse(health string) *G3v023VisaMastercardHealthcareResponse {
	return &G3v023VisaMastercardHealthcareResponse{G3v023VisaMastercardHealthcare, health}
}

func (g *G3v023VisaMastercardHealthcareResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.HealthcareData, bytes.FileSeparator)
}

func ParseG3v023VisaMastercardHealthcareResponse(s string) *G3v023VisaMastercardHealthcareResponse {
	g := G3v023VisaMastercardHealthcareResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 203 {
		g.HealthcareData = s[3:i]
	}

	return &g
}

type MerchantAdviceCodeType uint

const (
	AdviceCodeNewAccountInformation                    = 01
	AdviceCodeTryAgainLater                            = 02
	AdviceCodeDoNotRetryForRecurringPaymentTransaction = 03
	AdviceCodeRecurringPaymentCancellation             = 21
)

// ----------------------------------------------------
// Response message - Merchant Advice Code (MAC) (version 024)
type G3v024MerchantAdviceCodeResponse struct {
	VersionNumber      G3VersionType          //3 NUM
	MerchantAdviceCode MerchantAdviceCodeType // 2 A/N 4.109 Merchant Advice Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v024MerchantAdviceCodeResponse(advice MerchantAdviceCodeType) *G3v024MerchantAdviceCodeResponse {
	return &G3v024MerchantAdviceCodeResponse{G3v024MerchantAdviceCode, advice}
}

func (g *G3v024MerchantAdviceCodeResponse) String() string {
	return fmt.Sprintf("%03d%02d%c", g.VersionNumber, g.MerchantAdviceCode, bytes.FileSeparator)
}

func ParseG3v024MerchantAdviceCodeResponse(s string) *G3v024MerchantAdviceCodeResponse {
	g := G3v024MerchantAdviceCodeResponse{}
	fmt.Sscanf(s, "%03d%02d", &g.VersionNumber, &g.MerchantAdviceCode)
	return &g
}

// ----------------------------------------------------
// Response message - Transaction fee amount (version 025)
type G3v025TransactionFeeAmountResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 025
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v025TransactionFeeAmountResponse() *G3v025TransactionFeeAmountResponse {
	return &G3v025TransactionFeeAmountResponse{G3v025TransactionFeeAmount}
}

func (g *G3v025TransactionFeeAmountResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v025TransactionFeeAmountResponse(s string) *G3v025TransactionFeeAmountResponse {
	g := G3v025TransactionFeeAmountResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v026ProductParticipationGroupResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 026
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v026ProductParticipationGroupResponse() *G3v026ProductParticipationGroupResponse {
	return &G3v026ProductParticipationGroupResponse{G3v026ProductParticipationGroup}
}

func (g *G3v026ProductParticipationGroupResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v026ProductParticipationGroupResponse(s string) *G3v026ProductParticipationGroupResponse {
	g := G3v026ProductParticipationGroupResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type G3v027PosDataResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 027
}

func NewG3v027PosDataResponse() *G3v027PosDataResponse {
	return &G3v027PosDataResponse{G3v027PosData}
}

func (g *G3v027PosDataResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v027PosDataResponse(s string) *G3v027PosDataResponse {
	g := G3v027PosDataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v028AmericanExpressAdditionalDataResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 028
}

func NewG3v028AmericanExpressAdditionalDataResponse() *G3v028AmericanExpressAdditionalDataResponse {
	return &G3v028AmericanExpressAdditionalDataResponse{G3v028AmericanExpressAdditionalData}
}

func (g *G3v028AmericanExpressAdditionalDataResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v028AmericanExpressAdditionalDataResponse(s string) *G3v028AmericanExpressAdditionalDataResponse {
	g := G3v028AmericanExpressAdditionalDataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v029ExtendedAVSdataResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 029
}

func NewG3v029ExtendedAVSdataResponse() *G3v029ExtendedAVSdataResponse {
	return &G3v029ExtendedAVSdataResponse{G3v029ExtendedAVSdata}
}

func (g *G3v029ExtendedAVSdataResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v029ExtendedAVSdataResponse(s string) *G3v029ExtendedAVSdataResponse {
	g := G3v029ExtendedAVSdataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v030AmericanExpressMerchantNameLocationResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 030
}

func NewG3v030AmericanExpressMerchantNameLocationResponse() *G3v030AmericanExpressMerchantNameLocationResponse {
	return &G3v030AmericanExpressMerchantNameLocationResponse{VersionNumber: G3v030AmericanExpressMerchantNameLocation}
}

func (g *G3v030AmericanExpressMerchantNameLocationResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v030AmericanExpressMerchantNameLocationResponse(s string) *G3v030AmericanExpressMerchantNameLocationResponse {
	g := G3v030AmericanExpressMerchantNameLocationResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v031AgentIdentificationServiceResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 031
}

func NewG3v031AgentIdentificationServiceResponse() *G3v031AgentIdentificationServiceResponse {
	return &G3v031AgentIdentificationServiceResponse{G3v031AgentIdentificationService}
}

func (g *G3v031AgentIdentificationServiceResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v031AgentIdentificationServiceResponse(s string) *G3v031AgentIdentificationServiceResponse {
	g := G3v031AgentIdentificationServiceResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v032CurrencyConversionDataResponse struct {
	VersionNumber G3VersionType //3 NUM

	CardholderBillingAmount uint64 // 0-12 NUM 4.38 Cardholder Billing Amount
	// 1 ASCII 4.80 Field Separator <FS>
	SettlementConversionRateDecimals uint
	SettlementConversionRate         uint64 // 0-8 NUM 4.142 Settlement Conversion Rate
	// 1 ASCII 4.80 Field Separator <FS>
	CardholderBillingConversionRateDecimals uint
	CardholderBillingConversionRate         uint64 // 0-8 NUM 4.39 Cardholder Billing Conversion Rate
	// 1 ASCII 4.80 Field Separator <FS>
	ConversionDateMonth uint // 0-4 NUM 4.53 Conversion Date
	ConversionDateDay   uint
	// 1 ASCII 4.80 Field Separator <FS>
	AcquirerTransactionCurrencyCode CurrencyCodeType // 0-3 NUM 4.5 Acquirer Transaction Currency Code
	// 1 ASCII 4.80 Field Separator <FS>
	SettlementCurrencyCode CurrencyCodeType // 0-3 NUM 4.143 Settlement Currency Code
	// 1 ASCII 4.80 Field Separator <FS>
	CardholderBillingCurrencyCode CurrencyCodeType // 0-3 NUM 4.40 Cardholder Billing Currency Code
	// 1 ASCII 4.80 Field Separator <FS>
	ActualTransactionAmount uint64 // 0-12 NUM 4.8 Actual Amount, Transaction
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-12 NUM Reserved
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-12 NUM Reserved
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v032CurrencyConversionDataResponse(cardholderamount uint64, settledecimals uint, settlerate uint64, billingdecimals uint, billingrate uint64,
	convmonth uint, convday uint, acquirercurrency CurrencyCodeType, settlecurrency CurrencyCodeType, cardholdercurrency CurrencyCodeType, actualamount uint64) *G3v032CurrencyConversionDataResponse {
	return &G3v032CurrencyConversionDataResponse{G3v032CurrencyConversionData, cardholderamount, settledecimals, settlerate, billingdecimals, billingrate,
		convmonth, convday, acquirercurrency, settlecurrency, cardholdercurrency, actualamount}
}

func (g *G3v032CurrencyConversionDataResponse) String() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%03d", g.VersionNumber)

	buffer.AppendFormat("%d", g.CardholderBillingAmount) // 0-12 NUM 4.38 Cardholder Billing Amount
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%d", g.SettlementConversionRateDecimals)
	buffer.AppendFormat("%d", g.SettlementConversionRate) // 0-8 NUM 4.142 Settlement Conversion Rate
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%d", g.CardholderBillingConversionRateDecimals)
	buffer.AppendFormat("%d", g.CardholderBillingConversionRate) // 0-8 NUM 4.39 Cardholder Billing Conversion Rate
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%02d", g.ConversionDateMonth) // 0-4 NUM 4.53 Conversion Date
	buffer.AppendFormat("%02d", g.ConversionDateDay)
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%03d", g.AcquirerTransactionCurrencyCode) // 0-3 NUM 4.5 Acquirer Transaction Currency Code
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%03d", g.SettlementCurrencyCode) // 0-3 NUM 4.143 Settlement Currency Code
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%03d", g.CardholderBillingCurrencyCode) // 0-3 NUM 4.40 Cardholder Billing Currency Code
	buffer.AppendByte(byte(bytes.FileSeparator))
	buffer.AppendFormat("%d", g.ActualTransactionAmount) // 0-12 NUM 4.8 Actual Amount, Transaction
	buffer.AppendByte(byte(bytes.FileSeparator))
	// 0-12 NUM Reserved
	buffer.AppendByte(byte(bytes.FileSeparator))
	// 0-12 NUM Reserved
	buffer.AppendByte(byte(bytes.FileSeparator))

	return buffer.String()
}

func ParseG3v032CurrencyConversionDataResponse(s string) *G3v032CurrencyConversionDataResponse {
	g := G3v032CurrencyConversionDataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.CardholderBillingAmount, _ = strconv.ParseUint(fields[j], 10, 64)
			case 1:
				fmt.Sscanf(fields[j], "%d", &g.SettlementConversionRateDecimals)
				g.SettlementConversionRate, _ = strconv.ParseUint(fields[j][1:], 10, 64)
			case 2:
				fmt.Sscanf(fields[j], "%d", &g.CardholderBillingConversionRateDecimals)
				g.CardholderBillingConversionRate, _ = strconv.ParseUint(fields[j][1:], 10, 64)
			case 3:
				fmt.Sscanf(fields[j], "%02d%02d", &g.ConversionDateMonth, &g.ConversionDateDay)
			case 4:
				fmt.Sscanf(fields[j], "%03d", &g.AcquirerTransactionCurrencyCode)
			case 5:
				fmt.Sscanf(fields[j], "%03d", &g.SettlementCurrencyCode)
			case 6:
				fmt.Sscanf(fields[j], "%03d", &g.CardholderBillingCurrencyCode)
			case 7:
				g.ActualTransactionAmount, _ = strconv.ParseUint(fields[j], 10, 64)
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v033ReversalRequestCodeResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 033
}

func NewG3v033ReversalRequestCodeResponse() *G3v033ReversalRequestCodeResponse {
	return &G3v033ReversalRequestCodeResponse{G3v033ReversalRequestCode}
}

func (g *G3v033ReversalRequestCodeResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v033ReversalRequestCodeResponse(s string) *G3v033ReversalRequestCodeResponse {
	g := G3v033ReversalRequestCodeResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type CardProductCodeType string

const (
	CardProductVisaTraditional        CardProductCodeType = "A^"
	CardProductAmericanExpressCard    CardProductCodeType = "AX"
	CardProductVisaTraditionalRewards CardProductCodeType = "B^"
	CardProductVisaSignature          CardProductCodeType = "C^"
	CardProductVisaSignaturePreferred CardProductCodeType = "D^"
	CardProductDiscoverCard           CardProductCodeType = "DI"
	CardProductDinersCard             CardProductCodeType = "DN"
	CardProductProprietaryATM         CardProductCodeType = "E^"
	CardProductVisaClassic            CardProductCodeType = "F^"
	CardProductVisaBusiness           CardProductCodeType = "G^"
	CardProductVisaSignatureBusiness  CardProductCodeType = "G1"
	// G2 Reserved
	CardProductVisaBusinessEnhancedOrPlatinum CardProductCodeType = "G3"
	CardProductVisaInfinitePrivilegeBusiness  CardProductCodeType = "G4"
	// H^ Reserved
	CardProductVisaInfinite          CardProductCodeType = "I^"
	CardProductVisaInfinitePrivilege CardProductCodeType = "I1"
	CardProductUltraHighNetWorth     CardProductCodeType = "I2"
	// J^ Reserved
	// J1 Reserved
	// J2 Reserved
	CardProductVisaHealthcare CardProductCodeType = "J3"
	// J4 Reserved

	CardProductJCBCard                    CardProductCodeType = "JC"
	CardProductVisaCorporateTAndE         CardProductCodeType = "K^"
	CardProductVisaGSCCorporateTAndE      CardProductCodeType = "K1"
	CardProductElectron                   CardProductCodeType = "L^"
	CardProductMasterCard                 CardProductCodeType = "M^"
	CardProductVisaPlatinum               CardProductCodeType = "N^"
	CardProductVisaRewards                CardProductCodeType = "N1"
	CardProductVisaSelect                 CardProductCodeType = "N2"
	CardProductVisaGold                   CardProductCodeType = "P^"
	CardProductPrivateLabel               CardProductCodeType = "Q^"
	CardProductReserved                   CardProductCodeType = "Q1"
	CardProductPrivateLabelBasic          CardProductCodeType = "Q2"
	CardProductPrivateLabelStandard       CardProductCodeType = "Q3"
	CardProductPrivateLabelEnhanced       CardProductCodeType = "Q4"
	CardProductPrivateLabelSpecialized    CardProductCodeType = "Q5"
	CardProductPrivateLabelPremium        CardProductCodeType = "Q6"
	CardProductProprietary                CardProductCodeType = "R^"
	CardProductVisaPurchasing             CardProductCodeType = "S^"
	CardProductVisaPurchasingwithFleet    CardProductCodeType = "S1"
	CardProductVisaGSAPurchasing          CardProductCodeType = "S2"
	CardProductVisaGSAPurchasingwithFleet CardProductCodeType = "S3"
	CardProductCommercialLoan             CardProductCodeType = "S4"
	CardProductCommercialTransportEBT     CardProductCodeType = "S5"

	CardProductBusinessLoan CardProductCodeType = "S6"
	// S7 Reserved
	// T^ Reserved
	CardProductVisaTravelMoney CardProductCodeType = "U^"
	CardProductVPay            CardProductCodeType = "V^"

// V1 Reserved
// W^ Reserved
// X^ Reserved
// Y^ Reserved
// Z^ Reserved
)

func (x CardProductCodeType) String() string {
	switch x {
	case CardProductVisaTraditional:
		return "CardProductVisaTraditional"
	case CardProductAmericanExpressCard:
		return "CardProductAmericanExpressCard"
	case CardProductVisaTraditionalRewards:
		return "CardProductVisaTraditionalRewards"
	case CardProductVisaSignature:
		return "CardProductVisaSignature"
	case CardProductVisaSignaturePreferred:
		return "CardProductVisaSignaturePreferred"
	case CardProductDiscoverCard:
		return "CardProductDiscoverCard"
	case CardProductDinersCard:
		return "CardProductDinersCard"
	case CardProductProprietaryATM:
		return "CardProductProprietaryATM"
	case CardProductVisaClassic:
		return "CardProductVisaClassic"
	case CardProductVisaBusiness:
		return "CardProductVisaBusiness"
	case CardProductVisaSignatureBusiness:
		return "CardProductVisaSignatureBusiness"
	// G2 Reserved
	case CardProductVisaBusinessEnhancedOrPlatinum:
		return "CardProductVisaBusinessEnhancedOrPlatinum"
	case CardProductVisaInfinitePrivilegeBusiness:
		return "CardProductVisaInfinitePrivilegeBusiness"
	// H^ Reserved
	case CardProductVisaInfinite:
		return "CardProductVisaInfinite"
	case CardProductVisaInfinitePrivilege:
		return "CardProductVisaInfinitePrivilege"
	case CardProductUltraHighNetWorth:
		return "CardProductUltraHighNetWorth"
	// J^ Reserved
	// J1 Reserved
	// J2 Reserved
	case CardProductVisaHealthcare:
		return "CardProductVisaHealthcare"
	// J4 Reserved

	case CardProductJCBCard:
		return "CardProductJCBCard"
	case CardProductVisaCorporateTAndE:
		return "CardProductVisaCorporateTAndE"
	case CardProductVisaGSCCorporateTAndE:
		return "CardProductVisaGSCCorporateTAndE"
	case CardProductElectron:
		return "CardProductElectron"
	case CardProductMasterCard:
		return "CardProductMasterCard"
	case CardProductVisaPlatinum:
		return "CardProductVisaPlatinum"
	case CardProductVisaRewards:
		return "CardProductVisaRewards"
	case CardProductVisaSelect:
		return "CardProductVisaSelect"
	case CardProductVisaGold:
		return "CardProductVisaGold"
	case CardProductPrivateLabel:
		return "CardProductPrivateLabel"
	case CardProductReserved:
		return "CardProductReserved"
	case CardProductPrivateLabelBasic:
		return "CardProductPrivateLabelBasic"
	case CardProductPrivateLabelStandard:
		return "CardProductPrivateLabelStandard"
	case CardProductPrivateLabelEnhanced:
		return "CardProductPrivateLabelEnhanced"
	case CardProductPrivateLabelSpecialized:
		return "CardProductPrivateLabelSpecialized"
	case CardProductPrivateLabelPremium:
		return "CardProductPrivateLabelPremium"
	case CardProductProprietary:
		return "CardProductProprietary"
	case CardProductVisaPurchasing:
		return "CardProductVisaPurchasing"
	case CardProductVisaPurchasingwithFleet:
		return "CardProductVisaPurchasingwithFleet"
	case CardProductVisaGSAPurchasing:
		return "CardProductVisaGSAPurchasing"
	case CardProductVisaGSAPurchasingwithFleet:
		return "CardProductVisaGSAPurchasingwithFleet"
	case CardProductCommercialLoan:
		return "CardProductCommercialLoan"
	case CardProductCommercialTransportEBT:
		return "CardProductCommercialTransportEBT"

	case CardProductBusinessLoan:
		return "CardProductBusinessLoan"
	// S7 Reserved
	// T^ Reserved
	case CardProductVisaTravelMoney:
		return "CardProductVisaTravelMoney"
	case CardProductVPay:
		return "CardProductVPay"
	}
	return ""
}

// ----------------------------------------------------
type G3v034CardProductCodeResponse struct {
	VersionNumber   G3VersionType       //3 NUM
	CardProductCode CardProductCodeType // 0 or 2 A/N 4.34 Card Product Code
	// 1 ASCII 4.80 Field Separator <FS>
}

func NewG3v034CardProductCodeResponse(card CardProductCodeType) *G3v034CardProductCodeResponse {
	return &G3v034CardProductCodeResponse{G3v034CardProductCode, card}
}

func (g *G3v034CardProductCodeResponse) String() string {
	return fmt.Sprintf("%03d%2s", g.VersionNumber, string(g.CardProductCode))
}

func ParseG3v034CardProductCodeResponse(s string) *G3v034CardProductCodeResponse {

	g := G3v034CardProductCodeResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	if uint8(bytes.FileSeparator) != s[3] && 5 < len(s) {
		g.CardProductCode = CardProductCodeType(strings.Replace(s[3:5], " ", "^", -1))
	}
	return &g
}

// ----------------------------------------------------
type G3v035PromotionalCodeResponse struct {
	VersionNumber  G3VersionType // 3 NUM 4.85 Group III Version Number 035
	PromotionCodes string        // 0-50 A/N 4.125 Promotion Codes Card specific format
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v035PromotionalCodeResponse(promocodes string) *G3v035PromotionalCodeResponse {
	return &G3v035PromotionalCodeResponse{G3v035PromotionalCode, promocodes}
}

func (g *G3v035PromotionalCodeResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.PromotionCodes, bytes.FileSeparator)
}

func ParseG3v035PromotionalCodeResponse(s string) *G3v035PromotionalCodeResponse {
	g := G3v035PromotionalCodeResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 53 {
		g.PromotionCodes = s[3:i]
	}
	return &g
}

// ----------------------------------------------------
type G3v036PaymentTransactionIdentifierResponse struct {
	VersionNumber                G3VersionType                    // 3 NUM 4.85 Group III Version Number 036
	PaymentTransactionIdentifier PaymentTransactionIdentifierType // 3 A/N 4.122 Payment Transaction Identifier
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v036PaymentTransactionIdentifierResponse(transactionid PaymentTransactionIdentifierType) *G3v036PaymentTransactionIdentifierResponse {
	return &G3v036PaymentTransactionIdentifierResponse{G3v036PaymentTransactionIdentifier, transactionid}
}

func (g *G3v036PaymentTransactionIdentifierResponse) String() string {
	return fmt.Sprintf("%03d%-3s", g.VersionNumber, g.PaymentTransactionIdentifier)
}

func ParseG3v036PaymentTransactionIdentifierResponse(s string) *G3v036PaymentTransactionIdentifierResponse {
	g := G3v036PaymentTransactionIdentifierResponse{}

	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.PaymentTransactionIdentifier = PaymentTransactionIdentifierType(s[3:6])

	return &g
}

// Version 037 is only valid for MasterCard
// ----------------------------------------------------
type G3v037RealTimeSubstantiationResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 037
}

func NewG3v037RealTimeSubstantiationResponse() *G3v037RealTimeSubstantiationResponse {
	return &G3v037RealTimeSubstantiationResponse{G3v037RealTimeSubstantiation}
}

func (g *G3v037RealTimeSubstantiationResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v037RealTimeSubstantiationResponse(s string) *G3v037RealTimeSubstantiationResponse {
	g := G3v037RealTimeSubstantiationResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// This field contains the digital value of the magnetic signature from the card if it is captured
// when the card is swiped. This data will not be returned in the response, and it must not be
// stored after authorization. This must not be submitted in contactless or chip transactions. (see
// Table 3.72 for record format and version number).
// ----------------------------------------------------
type G3v038ElectroMagneticSignatureResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 038
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v038ElectroMagneticSignatureResponse() *G3v038ElectroMagneticSignatureResponse {
	return &G3v038ElectroMagneticSignatureResponse{G3v038ElectroMagneticSignature}
}

func (g *G3v038ElectroMagneticSignatureResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v038ElectroMagneticSignatureResponse(s string) *G3v038ElectroMagneticSignatureResponse {
	g := G3v038ElectroMagneticSignatureResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// This field is used by MasterCard in advice requests and reversals. A value of P indicates that a
// PIN was used, and an S indicates a signature or other method. For auto fuel dispensing advice
// messages, the value must be S. (See Table 3.74 for record format and version number).
// ----------------------------------------------------
type G3v039CardholderVerificationMethodResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 039
}

func NewG3v039CardholderVerificationMethodResponse() *G3v039CardholderVerificationMethodResponse {
	return &G3v039CardholderVerificationMethodResponse{G3v039CardholderVerificationMethod}
}

func (g *G3v039CardholderVerificationMethodResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v039CardholderVerificationMethodResponse(s string) *G3v039CardholderVerificationMethodResponse {
	g := G3v039CardholderVerificationMethodResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type ISAChargeIndicatorType rune

const (
	SingleCurrencyISA ISAChargeIndicatorType = 'C'
	MulticurrencyISA  ISAChargeIndicatorType = 'S'
)

func (x ISAChargeIndicatorType) String() string {
	switch x {
	case SingleCurrencyISA:
		return "SingleCurrencyISA"
	case MulticurrencyISA:
		return "MulticurrencyISA"
	}
	return ""
}

// ----------------------------------------------------
// The Visa International Service Assessment (ISA) charge is payable by the U.S. acquirers. The
// new Acquirer ISA applies to single currency and multicurrency transactions that are submitted
// by acquirers in the U.S. region. The merchant must be in the U.S. and the issuer country is non
// US.
type G3v040VisaISAChargeIndicatorResponse struct {
	VersionNumber      G3VersionType          // 3 NUM 4.85 Group III Version Number 040
	ISAChargeIndicator ISAChargeIndicatorType // 0-1 ASCII 4.93 ISA Charge Indicator
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v040VisaISAChargeIndicatorResponse(isa ISAChargeIndicatorType) *G3v040VisaISAChargeIndicatorResponse {
	return &G3v040VisaISAChargeIndicatorResponse{G3v040VisaISAChargeIndicator, isa}
}

func (g *G3v040VisaISAChargeIndicatorResponse) String() string {
	return fmt.Sprintf("%03d%c%c", g.VersionNumber, g.ISAChargeIndicator, bytes.FileSeparator)
}

func ParseG3v040VisaISAChargeIndicatorResponse(s string) *G3v040VisaISAChargeIndicatorResponse {
	g := G3v040VisaISAChargeIndicatorResponse{}
	if uint8(bytes.FileSeparator) == s[3] {
		fmt.Sscanf(s, "%03d", &g.VersionNumber)
	} else {
		fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.ISAChargeIndicator)
	}

	return &g
}

// ----------------------------------------------------
type G3v041NtiaUpcSkuDataResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 041
}

func NewG3v041NtiaUpcSkuDataResponse() *G3v041NtiaUpcSkuDataResponse {
	return &G3v041NtiaUpcSkuDataResponse{G3v041NtiaUpcSkuData}
}

func (g *G3v041NtiaUpcSkuDataResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v041NtiaUpcSkuDataResponse(s string) *G3v041NtiaUpcSkuDataResponse {
	g := G3v041NtiaUpcSkuDataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v042VisaContactlessResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 042
}

func NewG3v042VisaContactlessResponse() *G3v042VisaContactlessResponse {
	return &G3v042VisaContactlessResponse{G3v042VisaContactless}
}

func (g *G3v042VisaContactlessResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v042VisaContactlessResponse(s string) *G3v042VisaContactlessResponse {
	g := G3v042VisaContactlessResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v043NetworkIDResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 043
}

func NewG3v043NetworkIDResponse() *G3v043NetworkIDResponse {
	return &G3v043NetworkIDResponse{G3v043NetworkID}
}

func (g *G3v043NetworkIDResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v043NetworkIDResponse(s string) *G3v043NetworkIDResponse {
	g := G3v043NetworkIDResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v044AutomatedTellerMachineResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 044
}

func NewG3v044AutomatedTellerMachineResponse() *G3v044AutomatedTellerMachineResponse {
	return &G3v044AutomatedTellerMachineResponse{G3v044AutomatedTellerMachine}
}

func (g *G3v044AutomatedTellerMachineResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v044AutomatedTellerMachineResponse(s string) *G3v044AutomatedTellerMachineResponse {
	g := G3v044AutomatedTellerMachineResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v045IntegratedChipCardResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 045
	IssuerScript  string        // 0-510 AN 4.96 Issuer Script Can repeat up to 10 times
	// 1 ASCII 4.80 Field Separator <FS>
	AuthorizationResponseCryptogram string // 0-510 AN 4.27 Authorization Response Cryptogram (ARPC)
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v045IntegratedChipCardResponse(script string, cryptogram string) *G3v045IntegratedChipCardResponse {
	return &G3v045IntegratedChipCardResponse{G3v045IntegratedChipCard, script, cryptogram}
}

func (g *G3v045IntegratedChipCardResponse) String() string {
	return fmt.Sprintf("%03d%s%c%s%c", g.VersionNumber,
		g.IssuerScript, bytes.FileSeparator,
		g.AuthorizationResponseCryptogram, bytes.FileSeparator)
}

func ParseG3v045IntegratedChipCardResponse(s string) *G3v045IntegratedChipCardResponse {
	g := G3v045IntegratedChipCardResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.IssuerScript = fields[j]
			case 1:
				g.AuthorizationResponseCryptogram = fields[j]
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v046CardTypeGroupResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 046
	CardType      string        // 1-20 ANS 4.36 Card Type If POS sends G3V46 in the request the host will respond with card type.
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v046CardTypeGroupResponse(card string) *G3v046CardTypeGroupResponse {
	return &G3v046CardTypeGroupResponse{G3v046CardTypeGroup, card}
}

func (g *G3v046CardTypeGroupResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.CardType, bytes.FileSeparator)
}

func ParseG3v046CardTypeGroupResponse(s string) *G3v046CardTypeGroupResponse {
	g := G3v046CardTypeGroupResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	g.CardType = fields[0]
	return &g
}

//----------------------------------------------------
// type G3v047TSYSInternalUseOnlyResponse struct {
// 	VersionNumber G3VersionType //3 NUM
// }

// func NewG3v047TSYSInternalUseOnlyResponse() *G3v047TSYSInternalUseOnlyResponse {
// 	return &G3v047TSYSInternalUseOnlyResponse{G3v047TSYSInternalUseOnly, ""}
// }

// func (g *G3v047TSYSInternalUseOnlyResponse) String() string {
// 	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.xxx)
// }

type CardholderVerificationResultsCodeType rune

const (
	FirstNameMatchesLastNameDoesNotMatch    CardholderVerificationResultsCodeType = 'F'
	FirstNameDoesNotMatchLastNameMatches    CardholderVerificationResultsCodeType = 'L'
	FirstNameAndLastNameMatch               CardholderVerificationResultsCodeType = 'M'
	CardholderVerificationNoMatch           CardholderVerificationResultsCodeType = 'N'
	CardholderVerificationRetry             CardholderVerificationResultsCodeType = 'R'
	CardholderVerificationServiceNotAllowed CardholderVerificationResultsCodeType = 'S'
	CardholderVerificationResultNotSent     CardholderVerificationResultsCodeType = ' '
	RetrySystemUnavailableOrDataUnchecked   CardholderVerificationResultsCodeType = 'U'
	NoDataFromIssuerAuthorizationSystem     CardholderVerificationResultsCodeType = 'W'
	CardholderDataMatches                   CardholderVerificationResultsCodeType = 'Y'
)

func (x CardholderVerificationResultsCodeType) String() string {
	switch x {
	case FirstNameMatchesLastNameDoesNotMatch:
		return "FirstNameMatchesLastNameDoesNotMatch"
	case FirstNameDoesNotMatchLastNameMatches:
		return "FirstNameDoesNotMatchLastNameMatches"
	case FirstNameAndLastNameMatch:
		return "FirstNameAndLastNameMatch"
	case CardholderVerificationNoMatch:
		return "CardholderVerificationNoMatch"
	case CardholderVerificationRetry:
		return "CardholderVerificationRetry"
	case CardholderVerificationServiceNotAllowed:
		return "CardholderVerificationServiceNotAllowed"
	case CardholderVerificationResultNotSent:
		return "CardholderVerificationResultNotSent"
	case RetrySystemUnavailableOrDataUnchecked:
		return "RetrySystemUnavailableOrDataUnchecked"
	case NoDataFromIssuerAuthorizationSystem:
		return "NoDataFromIssuerAuthorizationSystem"
	case CardholderDataMatches:
		return "CardholderDataMatches"
	}
	return ""
}

// ----------------------------------------------------
type G3v048AmexCardholderVerificationResultsResponse struct {
	VersionNumber                  G3VersionType // 3 NUM 4.85 Group III Version Number 048
	VerificationResults            string        // 9 AN 4.18 Amex Cardholder Verification Results or Discover Cardholder Name Verification
	BillingZIPResultCode           CardholderVerificationResultsCodeType
	BillingStreetMatchResultCode   CardholderVerificationResultsCodeType
	BillingNameMatchResultCode     CardholderVerificationResultsCodeType
	TelephoneNumberMatchResultCode CardholderVerificationResultsCodeType
	EmailAddressMatchResultCode    CardholderVerificationResultsCodeType
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v048AmexCardholderVerificationResultsResponse(zip CardholderVerificationResultsCodeType, street CardholderVerificationResultsCodeType, name CardholderVerificationResultsCodeType, phone CardholderVerificationResultsCodeType, email CardholderVerificationResultsCodeType) *G3v048AmexCardholderVerificationResultsResponse {
	return &G3v048AmexCardholderVerificationResultsResponse{G3v048AmexCardholderVerificationResults, "", zip, street, name, phone, email}
}

func (g *G3v048AmexCardholderVerificationResultsResponse) String() string {
	return fmt.Sprintf("%03d%c%c%c%c%c    ", g.VersionNumber, g.BillingZIPResultCode, g.BillingStreetMatchResultCode, g.BillingNameMatchResultCode, g.TelephoneNumberMatchResultCode, g.EmailAddressMatchResultCode)
}

func ParseG3v048AmexCardholderVerificationResultsResponse(s string) *G3v048AmexCardholderVerificationResultsResponse {
	g := G3v048AmexCardholderVerificationResultsResponse{}
	fmt.Sscanf(s, "%03d%c%c%c%c%c", &g.VersionNumber, &g.BillingZIPResultCode, &g.BillingStreetMatchResultCode, &g.BillingNameMatchResultCode, &g.TelephoneNumberMatchResultCode, &g.EmailAddressMatchResultCode)
	g.VerificationResults = s[3:12]
	return &g
}

// ----------------------------------------------------
type G3v049Gen2TerminalAuthenticationResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 049
	GenKey        string        // 0 or 24 ASCII 4.84.4 GenKey ASCII
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v049Gen2TerminalAuthenticationResponse(key string) *G3v049Gen2TerminalAuthenticationResponse {
	return &G3v049Gen2TerminalAuthenticationResponse{G3v049Gen2TerminalAuthentication, key}
}

func (g *G3v049Gen2TerminalAuthenticationResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.GenKey, bytes.FileSeparator)
}

func ParseG3v049Gen2TerminalAuthenticationResponse(s string) *G3v049Gen2TerminalAuthenticationResponse {
	g := G3v049Gen2TerminalAuthenticationResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 27 {
		g.GenKey = s[3:i]
	}
	return &g
}

// ----------------------------------------------------
// The Association timestamp version supports a specific date/time and other data elements
// provided by the association for message matching. Currently only valid on MasterCard AFD
// credit advice messages
type G3v050AssociationTimestampResponse struct {
	VersionNumber        G3VersionType // 3 NUM 4.85 Group III Version Number 050
	AssociationTimestamp string        // 0, 10 NUM 4.25 Association Timestamp mmdddhhmmss
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v050AssociationTimestampResponse(tmstamp string) *G3v050AssociationTimestampResponse {
	return &G3v050AssociationTimestampResponse{G3v050AssociationTimestamp, tmstamp}
}

func (g *G3v050AssociationTimestampResponse) String() string {
	return fmt.Sprintf("%03d%-10s%c", g.VersionNumber, g.AssociationTimestamp, bytes.FileSeparator)
}

func ParseG3v050AssociationTimestampResponse(s string) *G3v050AssociationTimestampResponse {
	g := G3v050AssociationTimestampResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	i := strings.Index(s[3:], string(bytes.FileSeparator))
	if 3 < i && i < 13 {
		g.AssociationTimestamp = s[3:13]
	}
	return &g
}

type EMSScoringResponseIndicatorType uint

const EMSScoringWasPerformed EMSScoringResponseIndicatorType = 90

func (x EMSScoringResponseIndicatorType) String() string {
	switch x {
	case EMSScoringWasPerformed:
		return "EMSScoringWasPerformed"
	}
	return ""
}

type EMSResultsCodeType rune

const (
	EMSScoringSuccessful                        EMSResultsCodeType = 'C'
	InvalidTransactionNotQualifiedForEMSScoring EMSResultsCodeType = 'I'
	ScoringWasNotSuccessful                     EMSResultsCodeType = 'U'
)

func (x EMSResultsCodeType) String() string {
	switch x {
	case EMSScoringSuccessful:
		return "EMSScoringSuccessful"
	case InvalidTransactionNotQualifiedForEMSScoring:
		return "InvalidTransactionNotQualifiedForEMSScoring"
	case ScoringWasNotSuccessful:
		return "ScoringWasNotSuccessful"
	}
	return ""
}

type EMSAdditionalInformationType rune

const (
	EMSScoringValueNotPresent EMSAdditionalInformationType = ' '
	NotQualifiedForEMSScoring EMSAdditionalInformationType = 'N'
)

func (x EMSAdditionalInformationType) String() string {
	switch x {
	case EMSScoringValueNotPresent:
		return "EMSScoringValueNotPresent"
	case NotQualifiedForEMSScoring:
		return "NotQualifiedForEMSScoring"
	}
	return ""
}

//----------------------------------------------------
// The Event Monitoring Real-time Scoring Service (EMS) is used by participating Card Not
// Present (CNP) merchants to request a predictive risk score that may assist in determining if a
// CNP transaction is fraudulent. This EMS Service is available to MasterCard transactions only.

type G3v051MasterCardEventMonitoringRealTimeScoringResponse struct {
	VersionNumber               G3VersionType                   // 3 NUM 4.85 Group III Version Number 051
	EMSScoringResponseIndicator EMSScoringResponseIndicatorType // 2 NUM 4.73 EMS Scoring Response Indicator
	// 	90 - EMS scoring was performed
	// 1 ASCII 4.80 Field Separator <FS>
	EMSResultsCode EMSResultsCodeType // 1 ASCII 4.70 EMS Results Code
	// 	Valid values:
	// 		C - EMS scoring successful
	// 		I - Invalid, transaction does not qualify for EMS scoring
	// 		U - Scoring was not successful
	// 1 ASCII 4.80 Field Separator <FS>
	EMSAdditionalInformation EMSAdditionalInformationType // 1 ASCII 4.69 EMS Additional Information
	// 		N - Not qualified for EMS scoring (blank) - No value present
	// 1 ASCII 4.80 Field Separator <FS>
	EMSRiskScore uint // 0, 3 NUM 4.71 EMS Risk Score
	// 		Possible values = 000 - 999
	// 		May not be present
	// 1 ASCII 4.80 Field Separator <FS>
	RiskScoreReasonCode string // 0, 2 ANS 4.72 Risk score reason code; May not be present; Values and meaning are available to participating acquirers from MasterCard only.
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v051MasterCardEventMonitoringRealTimeScoringResponse(responseind EMSScoringResponseIndicatorType, resultcode EMSResultsCodeType, additional EMSAdditionalInformationType, score uint, reason string) *G3v051MasterCardEventMonitoringRealTimeScoringResponse {
	return &G3v051MasterCardEventMonitoringRealTimeScoringResponse{G3v051MasterCardEventMonitoringRealTimeScoring, responseind, resultcode, additional, score, reason}
}

func (g *G3v051MasterCardEventMonitoringRealTimeScoringResponse) String() string {
	return fmt.Sprintf("%03d%02d%c%c%c%c%c%03d%c%2s%c", g.VersionNumber, g.EMSScoringResponseIndicator, bytes.FileSeparator,
		g.EMSResultsCode, bytes.FileSeparator,
		g.EMSAdditionalInformation, bytes.FileSeparator,
		g.EMSRiskScore, bytes.FileSeparator,
		g.RiskScoreReasonCode, bytes.FileSeparator)
}

func ParseG3v051MasterCardEventMonitoringRealTimeScoringResponse(s string) *G3v051MasterCardEventMonitoringRealTimeScoringResponse {
	g := G3v051MasterCardEventMonitoringRealTimeScoringResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				fmt.Sscanf(fields[j], "%02d", &g.EMSScoringResponseIndicator)
			case 1:
				g.EMSResultsCode = EMSResultsCodeType(fields[j][0])
			case 2:
				g.EMSAdditionalInformation = EMSAdditionalInformationType(fields[j][0])
			case 3:
				fmt.Sscanf(fields[j], "%03d", &g.EMSRiskScore)
			case 4:
				g.RiskScoreReasonCode = fields[j]
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v052VoltageEncryptionTransmissionBlockResponse struct {
	VersionNumber G3VersionType // 5 NUM 4.85 Group III Version Number 052
}

func NewG3v052VoltageEncryptionTransmissionBlockResponse() *G3v052VoltageEncryptionTransmissionBlockResponse {
	return &G3v052VoltageEncryptionTransmissionBlockResponse{G3v052VoltageEncryptionTransmissionBlock}
}

func (g *G3v052VoltageEncryptionTransmissionBlockResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v052VoltageEncryptionTransmissionBlockResponse(s string) *G3v052VoltageEncryptionTransmissionBlockResponse {
	g := G3v052VoltageEncryptionTransmissionBlockResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type TokenStatusIndicatorType uint

const (
	TokenRetrievalWasSuccessfulAndIsIncluded TokenStatusIndicatorType = 0
	TerminalNotConfiguredToReceiveAToken     TokenStatusIndicatorType = 1
	TerminalWasNotAuthenticatedForToken      TokenStatusIndicatorType = 2
	TokenSystemFailure                       TokenStatusIndicatorType = 3
)

func (x TokenStatusIndicatorType) String() string {
	switch x {
	case TokenRetrievalWasSuccessfulAndIsIncluded:
		return "TokenRetrievalWasSuccessfulAndIsIncluded"
	case TerminalNotConfiguredToReceiveAToken:
		return "TerminalNotConfiguredToReceiveAToken"
	case TerminalWasNotAuthenticatedForToken:
		return "TerminalWasNotAuthenticatedForToken"
	case TokenSystemFailure:
		return "TokenSystemFailure"
	}
	return ""
}

//----------------------------------------------------

// Presence of Group 3 Version 53 with any Transaction Code (4.163) indicates the POS Device
// is requesting a token. The token is delivered in the response

// NOTE Group 3 Version 049 - Gen2 Terminal Authentication is required in order to use Tokens.

type G3v053TokenResponse struct {
	VersionNumber G3VersionType            //3 NUM 4.85 Group III Version Number 053
	TokenStatus   TokenStatusIndicatorType // 2 AN 4.159 Token Status Status of token retrieval
	// 1 ASCII 4.80 Field Separator <FS>
	Token string // 0, 13-19 A/N 4.157 Token Format preserved token with last 4 digits preserved.
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v053TokenResponse(status TokenStatusIndicatorType, pantoken string) *G3v053TokenResponse {
	return &G3v053TokenResponse{G3v053Token, status, pantoken}
}

func (g *G3v053TokenResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v053TokenResponse(s string) *G3v053TokenResponse {
	g := G3v053TokenResponse{}
	fmt.Sscanf(s, "%03d%02d", &g.VersionNumber, &g.TokenStatus)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	g.Token = fields[1]

	return &g
}

// ----------------------------------------------------
type G3v054TransitProgramResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 054
}

func NewG3v054TransitProgramResponse() *G3v054TransitProgramResponse {
	return &G3v054TransitProgramResponse{G3v054TransitProgram}
}

func (g *G3v054TransitProgramResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v054TransitProgramResponse(s string) *G3v054TransitProgramResponse {
	g := G3v054TransitProgramResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v055IntegratedChipCardEmvTlvResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 055

	// This variable length field is composed solely of hexadecimal characters (0-9, A-F or a-f). Each
	// pair of characters represents one byte of information. The string of characters represents a
	// series of TLV data that represent the information passed between the card and the terminal.
	// Each datum has a one or two byte (two or four characters) tag, a one byte (two characters)
	// length, and a one or more byte (two or more characters) value or payload. The length byte
	// always represents the number of bytes following the length byte in the TLV datum.
	EmvTlvDatum string // 6-255 Hex 4.157 Refer to Appendix B TLV data, 2 characters per byte

	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v055IntegratedChipCardEmvTlvResponse(tlv string) *G3v055IntegratedChipCardEmvTlvResponse {
	return &G3v055IntegratedChipCardEmvTlvResponse{G3v055IntegratedChipCardEmvTlv, tlv}
}

func (g *G3v055IntegratedChipCardEmvTlvResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.EmvTlvDatum, bytes.FileSeparator)
}

func ParseG3v055IntegratedChipCardEmvTlvResponse(s string) *G3v055IntegratedChipCardEmvTlvResponse {
	g := G3v055IntegratedChipCardEmvTlvResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 258 {
		g.EmvTlvDatum = s[3:i]
	}

	return &g
}

// ----------------------------------------------------
type G3v056MessageReasonCodeResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 056
}

func NewG3v056MessageReasonCodeResponse() *G3v056MessageReasonCodeResponse {
	return &G3v056MessageReasonCodeResponse{G3v056MessageReasonCode}
}

func (g *G3v056MessageReasonCodeResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v056MessageReasonCodeResponse(s string) *G3v056MessageReasonCodeResponse {
	g := G3v056MessageReasonCodeResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v057DiscoverOrPayPalAdditionalDataResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 057
	// This field contains additional data from a Discover/PayPal AFD authorization response
	// message that must be used in a Discover/PayPal AFD completion advice message.
	AdditionalResponseData string // 0-25 A/N/S 4.10 Additional Response Data Additional data for Discover/ PayPal AFD messages
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v057DiscoverOrPayPalAdditionalDataResponse(data string) *G3v057DiscoverOrPayPalAdditionalDataResponse {
	return &G3v057DiscoverOrPayPalAdditionalDataResponse{G3v057DiscoverOrPayPalAdditionalData, data}
}

func (g *G3v057DiscoverOrPayPalAdditionalDataResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.AdditionalResponseData, bytes.FileSeparator)
}

func ParseG3v057DiscoverOrPayPalAdditionalDataResponse(s string) *G3v057DiscoverOrPayPalAdditionalDataResponse {
	g := G3v057DiscoverOrPayPalAdditionalDataResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 28 {
		g.AdditionalResponseData = s[3:i]
	}

	return &g
}

// It is recommended that Group 3, version 58 be accompanied by Group 3, version 59 for
// MasterCard transactions
// ----------------------------------------------------
type G3v058AlternateAccountIDResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 058
	// For MasterCard PayPass transit transactions, this field will contain the Primary Account
	// Number (PAN).
	AlternateAccountID string // 1-28 NUM 4.16 Alternate Account ID 1
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>

}

func NewG3v058AlternateAccountIDResponse(accountid string) *G3v058AlternateAccountIDResponse {
	return &G3v058AlternateAccountIDResponse{G3v058AlternateAccountID, accountid}
}

func (g *G3v058AlternateAccountIDResponse) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.AlternateAccountID, bytes.FileSeparator)
}

func ParseG3v058AlternateAccountIDResponse(s string) *G3v058AlternateAccountIDResponse {
	g := G3v058AlternateAccountIDResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 28 {
		g.AlternateAccountID = s[3:i]
	}

	return &g
}

// It is recommended that Group 3, version 59 be accompanied by Group 3, version 58 for
// MasterCard transactions.

// ----------------------------------------------------
type G3v059MasterCardPayPassMappingServiceResponse struct {
	VersionNumber      G3VersionType // 3 NUM 4.85 Group III Version Number 059
	MappedPANIndicator byte          // 1 A/N 4.103 Mapped PAN Indicator
	// 1 ASCII 4.80 Field Separator <FS>
	MappedCardExpirationDateMonth uint // 4 NUM 4.106 Mapped Card Expiration Date
	MappedCardExpirationDateYear  uint
	// 1 ASCII 4.80 Field Separator <FS>
	MappedProductCode string // 3 A/N 4.104 Mapped Product Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v059MasterCardPayPassMappingServiceResponse(panind byte, expmonth uint, expyear uint, product string) *G3v059MasterCardPayPassMappingServiceResponse {
	return &G3v059MasterCardPayPassMappingServiceResponse{G3v059MasterCardPayPassMappingService, panind, expmonth, expyear, product}
}

func (g *G3v059MasterCardPayPassMappingServiceResponse) String() string {
	return fmt.Sprintf("%03d%c%c%02d%02d%c%s%c", g.VersionNumber, g.MappedPANIndicator, bytes.FileSeparator,
		g.MappedCardExpirationDateMonth, g.MappedCardExpirationDateYear, bytes.FileSeparator,
		g.MappedProductCode, bytes.FileSeparator)
}

func ParseG3v059MasterCardPayPassMappingServiceResponse(s string) *G3v059MasterCardPayPassMappingServiceResponse {
	g := G3v059MasterCardPayPassMappingServiceResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.MappedPANIndicator = fields[j][0]
			case 1:
				fmt.Sscanf(fields[j], "%02d%02d", &g.MappedCardExpirationDateMonth, &g.MappedCardExpirationDateYear)
			case 2:
				g.MappedProductCode = fields[j]
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// PayPass Mobile (G3v060) - This group should be sent on all MasterCard PayPass transactions.
// ----------------------------------------------------
type G3v060PayPassMobileResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 060
}

func NewG3v060PayPassMobileResponse() *G3v060PayPassMobileResponse {
	return &G3v060PayPassMobileResponse{G3v060PayPassMobile}
}

func (g *G3v060PayPassMobileResponse) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v060PayPassMobileResponse(s string) *G3v060PayPassMobileResponse {
	g := G3v060PayPassMobileResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v061SpendQualifiedIndicatorResponse struct {
	//This group should be sent on all Visa transactions
	VersionNumber           G3VersionType // 3 NUM 4.85 Group III Version Number 061
	SpendQualifiedIndicator byte          // 0, 1 A/N 4.146 Spend Qualified Indicator
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v061SpendQualifiedIndicatorResponse(spendind byte) *G3v061SpendQualifiedIndicatorResponse {
	return &G3v061SpendQualifiedIndicatorResponse{G3v061SpendQualifiedIndicator, spendind}
}

func (g *G3v061SpendQualifiedIndicatorResponse) String() string {
	return fmt.Sprintf("%03d%c%c", g.VersionNumber, g.SpendQualifiedIndicator, bytes.FileSeparator)
}

func ParseG3v061SpendQualifiedIndicatorResponse(s string) *G3v061SpendQualifiedIndicatorResponse {
	g := G3v061SpendQualifiedIndicatorResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	if 0 < len(fields[0]) {
		g.SpendQualifiedIndicator = fields[0][0]
	}

	return &g
}

// ----------------------------------------------------
type G3v200GiftCardResponse struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 200

	AvailableBalance uint64 // 0 or 12 NUM 4.30 Available Balance
	// 1 ASCII 4.80 Field Separator <FS>
	ActualAmount uint64 // 0 or 12 NUM 4.20 Amount Actually Used From Card
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v200GiftCardResponse(balance uint64, amount uint64) *G3v200GiftCardResponse {
	return &G3v200GiftCardResponse{G3v200GiftCard, balance, amount}
}

func (g *G3v200GiftCardResponse) String() string {
	return fmt.Sprintf("%03d%012d%c%012d%c", g.VersionNumber, g.AvailableBalance, bytes.FileSeparator, g.ActualAmount, bytes.FileSeparator)
}

func ParseG3v200GiftCardResponse(s string) *G3v200GiftCardResponse {
	g := G3v200GiftCardResponse{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.AvailableBalance, _ = strconv.ParseUint(fields[j], 10, 64)
			case 1:
				g.ActualAmount, _ = strconv.ParseUint(fields[j], 10, 64)
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

//----------------------------------------------------
