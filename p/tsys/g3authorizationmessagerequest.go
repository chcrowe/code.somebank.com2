package tsys

import (
	"fmt"
	"strconv"
	"strings"

	"code.somebank.com/p/bytes"
)

const TSYS_DEVELOPER_ID string = "000136"
const TSYS_DEVELOPER_VERSION string = "B248"

//based upon EIS 1080 VERSION 12.0 FEBRUARY 11, 2014

// ----------------------------------------------------
type G3v001CommercialCardRequest struct {
	VersionNumber                  G3VersionType //3 NUM
	CommercialCardRequestIndicator string        // "!010" //4 A/N 4.51
}

func NewG3v001CommercialCardRequest() *G3v001CommercialCardRequest {
	return &G3v001CommercialCardRequest{G3v001CommercialCard, "!010"}
}

func (g *G3v001CommercialCardRequest) String() string {
	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.CommercialCardRequestIndicator)
}

func ParseG3v001CommercialCardRequest(s string) *G3v001CommercialCardRequest {
	g := G3v001CommercialCardRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.CommercialCardRequestIndicator = s[3:7]
	return &g
}

// ----------------------------------------------------
type G3v007CardVerificationCodeRequest struct {
	VersionNumber    G3VersionType         //3 NUM
	VerificationCode *VerificationCodeInfo //6 A/N 4.98
}

func NewG3v007CardVerificationCodeRequest(desiredresponse VerificationCodeDesiredResponseType, code string) *G3v007CardVerificationCodeRequest {
	vc := NewVerificationCodeInfo(desiredresponse, code)
	return &G3v007CardVerificationCodeRequest{G3v007CardVerificationCode, vc}
}

func (g *G3v007CardVerificationCodeRequest) String() string {
	return fmt.Sprintf("%03d%s", g.VersionNumber, g.VerificationCode.String())
}

func ParseG3v007CardVerificationCodeRequest(s string) *G3v007CardVerificationCodeRequest {
	v := ParseVerificationCodeInfo(s[3:])
	g := G3v007CardVerificationCodeRequest{G3v007CardVerificationCode, v}
	return &g
}

type G3v008FleetFuelingCardIdPromptType uint

const (
	G3v008IdNumberAndOdometerReading G3v008FleetFuelingCardIdPromptType = 1
	G3v008VehicleNumber              G3v008FleetFuelingCardIdPromptType = 2
	G3v008DriverIdAndOdometerReading G3v008FleetFuelingCardIdPromptType = 3
	G3v008OdometerReading            G3v008FleetFuelingCardIdPromptType = 4
	G3v008NoPrompting                G3v008FleetFuelingCardIdPromptType = 5
	G3v008IdNumber                   G3v008FleetFuelingCardIdPromptType = 6
)

func (x G3v008FleetFuelingCardIdPromptType) String() string {
	switch x {
	case G3v008IdNumberAndOdometerReading:
		return "G3v008IdNumberAndOdometerReading"
	case G3v008VehicleNumber:
		return "G3v008VehicleNumber"
	case G3v008DriverIdAndOdometerReading:
		return "G3v008DriverIdAndOdometerReading"
	case G3v008OdometerReading:
		return "G3v008OdometerReading"
	case G3v008NoPrompting:
		return "G3v008NoPrompting"
	case G3v008IdNumber:
		return "G3v008IdNumber"
	}
	return ""
}

// ----------------------------------------------------
type G3v008FleetFuelingCardRequest struct {
	VersionNumber   G3VersionType // 3 NUM 4.85 Group III Version Number 008
	PromptIndicator G3v008FleetFuelingCardIdPromptType
	// This 17-character field contains either a Driver ID, Vehicle ID, or other Identification number
	// to be used in the authorization of a Visa Fleet Fueling Card. Visa cards issued in the range
	// 448460 - 448699 contain instructions for customized prompts in the last position of the
	// magnetic stripe (before the end sentinel character)
	IdentificationNumber string // 0, 17 A/N 4.89 Identification Number
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v008FleetFuelingCardRequest(prompt G3v008FleetFuelingCardIdPromptType, id string) *G3v008FleetFuelingCardRequest {
	return &G3v008FleetFuelingCardRequest{G3v008FleetFuelingCard, prompt, id}
}

func (g *G3v008FleetFuelingCardRequest) String() string {
	if G3v008NoPrompting == g.PromptIndicator {
		return fmt.Sprintf("%03d%c", g.VersionNumber, bytes.FileSeparator)
	}
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.IdentificationNumber, bytes.FileSeparator)
}

func ParseG3v008FleetFuelingCardRequest(s string) *G3v008FleetFuelingCardRequest {
	g := G3v008FleetFuelingCardRequest{VersionNumber: G3v008FleetFuelingCard}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 4 < i && i < 21 {
		g.IdentificationNumber = s[3:i]
	}

	return &g
}

// ----------------------------------------------------
type G3v009SeteCommerceRequest struct {
	VersionNumber                     G3VersionType //3 NUM 4.85 Group III Version Number 009
	CardholderCertificateSerialNumber string        //0, 32 A/N 4.41 Cardholder Certificate Serial Number
	//1 ASCII 4.80 Field Separator <FS>
	MerchantCertificateSerialNumber string //1-32 A/N 4.108 Merchant Certificate Serial Number
	//1 ASCII 4.80 Field Separator <FS>
	XID       string //40 A/N 4.180 XID
	Transtain string //40 A/N 4.170 Transtain
	//1 ASCII 4.86 Group Separator <GS>
}

func NewG3v009SeteCommerceRequest(cardholdercert string, merchantcert string, id string, trans string) *G3v009SeteCommerceRequest {
	return &G3v009SeteCommerceRequest{G3v009SeteCommerce, cardholdercert, merchantcert, id, trans}
}

func (g *G3v009SeteCommerceRequest) String() string {
	return fmt.Sprintf("%03d%s%c%s%c%-40s%-40s", g.VersionNumber, g.CardholderCertificateSerialNumber, bytes.FileSeparator, g.MerchantCertificateSerialNumber, bytes.FileSeparator, g.XID, g.Transtain)
}

func ParseG3v009SeteCommerceRequest(s string) *G3v009SeteCommerceRequest {
	g := G3v009SeteCommerceRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	if 3 < len(fields[0]) {
		g.CardholderCertificateSerialNumber = fields[0]
	}
	g.MerchantCertificateSerialNumber = fields[1]
	if 80 == len(fields[2]) {
		g.XID = fields[2][0:40]
		g.Transtain = fields[2][40:80]
	}

	return &g
}

// ----------------------------------------------------
// reserved for future use
type G3v010CcpsRequest struct {
	VersionNumber               G3VersionType //3 NUM 4.85 Group III Version Number 010
	CardSequenceNumber          uint          // 3 NUM 4.35 Card Sequence Number
	TerminalCapabilityProfile   string        // 6 A/N 4.149 Terminal Capability Profile
	TerminalVerificationResults string        // 10 A/N 4.155 Terminal Verification Results
	UnpredictableNumber         string        // 8 A/N 4.173 Unpredictable Number
	InterfaceDeviceSerialNumber uint64        // 0, 8 NUM 4.92 Interface Device Serial Number
	// 1 ASCII 4.80 Field Separator <FS>
	DerivationKeyIndex      string // 2 A/N 4.63 Derivation Key Index
	CryptogramVersionNumber string // 2 A/N 4.60 Cryptogram Version Number
	CardVerificationResults string // 8 A/N 4.37 Card Verification Results
	IssuerDiscretionaryData string // 0, 16 A/N 4.95 Issuer Discretionary Data
	// 1 ASCII 4.80 Field Separator <FS>
	AuthorizationRequestCryptogram string // 16 A/N 4.25 Authorization Request Cryptogram
	ApplicationTransactionCounter  string // 4 A/N 4.22 Application Transaction Counter
	ApplicationInterchangeProfile  string // 4 A/N 4.21 Application Interchange Profile
	//TransactionDate                string // 6 NUM 4.164 Transaction Date YYMMDD

	TransactionDateYear  uint
	TransactionDateMonth uint
	TransactionDateDay   uint
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v010CcpsRequest() *G3v010CcpsRequest {
	return &G3v010CcpsRequest{VersionNumber: G3v010Ccps}
}

func (g *G3v010CcpsRequest) String() string {

	buffer := bytes.NewSafeBuffer(128)

	buffer.AppendFormat("%03d", g.VersionNumber)
	buffer.AppendFormat("%03d", g.CardSequenceNumber)
	buffer.AppendFormat("%-6s", g.TerminalCapabilityProfile)    // 6 A/N 4.149 Terminal Capability Profile
	buffer.AppendFormat("%-10s", g.TerminalVerificationResults) // 10 A/N 4.155 Terminal Verification Results
	buffer.AppendFormat("%-8s", g.UnpredictableNumber)          // 8 A/N 4.173 Unpredictable Number
	buffer.AppendFormat("%08d", g.InterfaceDeviceSerialNumber)  // 0, 8 NUM 4.92 Interface Device Serial Number
	buffer.AppendFormat("%c", bytes.FileSeparator)
	buffer.AppendFormat("%-2s", g.DerivationKeyIndex)       // 2 A/N 4.63 Derivation Key Index
	buffer.AppendFormat("%-2s", g.CryptogramVersionNumber)  // 2 A/N 4.60 Cryptogram Version Number
	buffer.AppendFormat("%-8s", g.CardVerificationResults)  // 8 A/N 4.37 Card Verification Results
	buffer.AppendFormat("%-16s", g.IssuerDiscretionaryData) // 0, 16 A/N 4.95 Issuer Discretionary Data
	buffer.AppendFormat("%c", bytes.FileSeparator)
	buffer.AppendFormat("%-16s", g.AuthorizationRequestCryptogram)                                           // 16 A/N 4.25 Authorization Request Cryptogram
	buffer.AppendFormat("%-4s", g.ApplicationTransactionCounter)                                             // 4 A/N 4.22 Application Transaction Counter
	buffer.AppendFormat("%-4s", g.ApplicationInterchangeProfile)                                             // 4 A/N 4.21 Application Interchange Profile
	buffer.AppendFormat("%02d%02d%02d", g.TransactionDateYear, g.TransactionDateMonth, g.TransactionDateDay) // 6 NUM 4.164 Transaction Date YYMMDD

	return buffer.String()
}

func ParseG3v010CcpsRequest(s string) *G3v010CcpsRequest {
	g := G3v010CcpsRequest{}
	fmt.Sscanf(s, "%03d%03d", &g.VersionNumber, &g.CardSequenceNumber)

	fields := strings.Split(s[6:], string(bytes.FileSeparator))
	if 23 < len(fields[0]) {
		g.TerminalCapabilityProfile = fields[0][0:6]
		g.TerminalVerificationResults = fields[0][6:16]
		g.UnpredictableNumber = fields[0][16:24]
		if 31 < len(fields[0]) {
			fmt.Sscanf(fields[0][24:32], "%08d", &g.InterfaceDeviceSerialNumber)
		}
	}

	if 11 < len(fields[1]) {
		g.DerivationKeyIndex = fields[1][0:2]
		g.CryptogramVersionNumber = fields[1][2:4]
		g.CardVerificationResults = fields[1][4:12]
		if 27 < len(fields[1]) {
			g.IssuerDiscretionaryData = fields[1][12:28]
		}
	}

	if 29 < len(fields[2]) {
		g.AuthorizationRequestCryptogram = fields[2][0:16]
		g.ApplicationTransactionCounter = fields[2][16:20]
		g.ApplicationInterchangeProfile = fields[2][20:24]
		fmt.Sscanf(fields[2][24:30], "%02d%02d%02d", &g.TransactionDateYear, &g.TransactionDateMonth, &g.TransactionDateDay)
	}

	return &g
}

// ----------------------------------------------------
type G3v011ChipConditionRequest struct {
	VersionNumber     G3VersionType         // 3 NUM 4.85 Group III Version Number 011
	ChipConditionCode ChipConditionCodeType // 1 A/N 4.49 Chip Condition Code
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v011ChipConditionRequest(chipcondition ChipConditionCodeType) *G3v011ChipConditionRequest {
	return &G3v011ChipConditionRequest{G3v011ChipCondition, chipcondition}
}

func (g *G3v011ChipConditionRequest) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.ChipConditionCode)
}

func ParseG3v011ChipConditionRequest(s string) *G3v011ChipConditionRequest {
	g := G3v011ChipConditionRequest{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.ChipConditionCode)
	return &g
}

//----------------------------------------------------
// type G3v012CommercialCardLargeTicketRequest struct {
// 	VersionNumber G3VersionType //3 NUM
// }

// func NewG3v012CommercialCardLargeTicketRequest() *G3v012CommercialCardLargeTicketRequest {
// 	return &G3v012CommercialCardLargeTicketRequest{G3v012CommercialCardLargeTicket, ""}
// }

// func (g *G3v012CommercialCardLargeTicketRequest) String() string {
// 	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.xxx)
// }

// ----------------------------------------------------
type G3v013ElectronicBenefitsTransferRequest struct {
	VersionNumber     G3VersionType // 3 NUM 4.85 Group III Version Number 013
	FoodAndConsumerId string        // 0, 7 A/N 4.79 FCS ID Food and Consumer Identifier; FCS ID identifies the Merchant as being certified and approved to accept Food Stamps
	// 1 ASCII 4.80 Field Separator <FS>
	ElectronicVoucherSerialNumber string // 0, 15 A/N 4.68 Electronic Voucher Serial Number
	// 1 ASCII 4.80 Field Separator <FS>
	VoucherApprovalCode string // 0, 6 A/N 4.179 Voucher Approval Code
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v013ElectronicBenefitsTransferRequest(fcsid string, voucherno string, voucherapproval string) *G3v013ElectronicBenefitsTransferRequest {
	return &G3v013ElectronicBenefitsTransferRequest{G3v013ElectronicBenefitsTransfer, fcsid, voucherno, voucherapproval}
}

func (g *G3v013ElectronicBenefitsTransferRequest) String() string {
	return fmt.Sprintf("%03d%s%c%s%c%s", g.VersionNumber, g.FoodAndConsumerId, bytes.FileSeparator, g.ElectronicVoucherSerialNumber, bytes.FileSeparator, g.VoucherApprovalCode)
}

func ParseG3v013ElectronicBenefitsTransferRequest(s string) *G3v013ElectronicBenefitsTransferRequest {
	g := G3v013ElectronicBenefitsTransferRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	g.FoodAndConsumerId = fields[0]
	g.ElectronicVoucherSerialNumber = fields[1]
	g.VoucherApprovalCode = fields[2]

	return &g
}

// ----------------------------------------------------
type G3v014MotoEcommerceRequest struct {
	VersionNumber          G3VersionType         // 3 NUM 4.85 Group III Version Number 014
	MOTOeCommerceIndicator MotoEcomIndicatorType // 1 A/N 4.116 MOTO/e-Commerce Indicator
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v014MotoEcommerceRequest(motoecomindicator MotoEcomIndicatorType) *G3v014MotoEcommerceRequest {
	return &G3v014MotoEcommerceRequest{G3v014MotoEcommerce, motoecomindicator}
}

func (g *G3v014MotoEcommerceRequest) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.MOTOeCommerceIndicator)
}

func ParseG3v014MotoEcommerceRequest(s string) *G3v014MotoEcommerceRequest {
	g := G3v014MotoEcommerceRequest{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.MOTOeCommerceIndicator)
	return &g
}

// ----------------------------------------------------
type G3v015ServiceDevelopmentIndicatorRequest struct {
	VersionNumber               G3VersionType          // 3 NUM 4.85 Group III Version Number 015
	ServiceDevelopmentIndicator ServiceDevelopmentType // 1 NUM 4.140 Service Development Indicator 6,5,7
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v015ServiceDevelopmentIndicatorRequest(serviceindicator ServiceDevelopmentType) *G3v015ServiceDevelopmentIndicatorRequest {
	return &G3v015ServiceDevelopmentIndicatorRequest{G3v015ServiceDevelopmentIndicator, serviceindicator}
}

func (g *G3v015ServiceDevelopmentIndicatorRequest) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.ServiceDevelopmentIndicator)
}

func ParseG3v015ServiceDevelopmentIndicatorRequest(s string) *G3v015ServiceDevelopmentIndicatorRequest {
	g := G3v015ServiceDevelopmentIndicatorRequest{}
	fmt.Sscanf(s, "%03d%d", &g.VersionNumber, &g.ServiceDevelopmentIndicator)
	return &g
}

//----------------------------------------------------
// type G3v016PosCheckServiceAuthorizationRequest struct {
// 	VersionNumber G3VersionType //3 NUM
// }

// func NewG3v016PosCheckServiceAuthorizationRequest() *G3v016PosCheckServiceAuthorizationRequest {
// 	return &G3v016PosCheckServiceAuthorizationRequest{G3v016PosCheckServiceAuthorization, ""}
// }

// func (g *G3v016PosCheckServiceAuthorizationRequest) String() string {
// 	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.xxx)
// }

// ----------------------------------------------------
type G3v017Visa3dSecurEcomVerificationRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 017
	XID           string        // 0 or 40 A/N 4.180 XID
	CAVV          string        // 40 A/N 4.46 CAVV
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v017Visa3dSecurEcomVerificationRequest(id string, authvalue string) *G3v017Visa3dSecurEcomVerificationRequest {
	return &G3v017Visa3dSecurEcomVerificationRequest{G3v017Visa3dSecurEcomVerification, id, authvalue}
}

func (g *G3v017Visa3dSecurEcomVerificationRequest) String() string {
	return fmt.Sprintf("%03d%s%-40s", g.VersionNumber, g.XID, g.CAVV)
}

func ParseG3v017Visa3dSecurEcomVerificationRequest(s string) *G3v017Visa3dSecurEcomVerificationRequest {
	g := G3v017Visa3dSecurEcomVerificationRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	if 83 <= len(s) {
		g.XID = s[3:43]
		g.CAVV = s[43:83]
	} else if 43 <= len(s) {
		g.CAVV = s[3:43]
	}

	return &g
}

// ----------------------------------------------------
type G3v018ExistingDebtIndicatorRequest struct {
	VersionNumber         G3VersionType             // 3 NUM 4.85 Group III Version Number 018
	ExistingDebtIndicator ExistingDebtIndicatorType // 1 NUM 4.69 Existing Debt Indicator 9
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v018ExistingDebtIndicatorRequest(debt ExistingDebtIndicatorType) *G3v018ExistingDebtIndicatorRequest {
	return &G3v018ExistingDebtIndicatorRequest{G3v018ExistingDebtIndicator, debt}
}

func (g *G3v018ExistingDebtIndicatorRequest) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.ExistingDebtIndicator)
}

func ParseG3v018ExistingDebtIndicatorRequest(s string) *G3v018ExistingDebtIndicatorRequest {
	g := G3v018ExistingDebtIndicatorRequest{}
	fmt.Sscanf(s, "%03d%1d", &g.VersionNumber, &g.ExistingDebtIndicator)
	return &g
}

// ----------------------------------------------------
type G3v019MastercardUniversalCardholderAuthenticationRequest struct {
	VersionNumber           G3VersionType      // 3 NUM 4.85 Group III Version Number 019
	UcafCollectionIndicator UcafCollectionType // 1 NUM 4.172 UCAF Collection Indicator
	UcafAuthenticationData  string             // 0-32 A/N (and Special) 4.171 UCAF Authentication Data
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v019MastercardUniversalCardholderAuthenticationRequest(ucafindicator UcafCollectionType, ucafdata string) *G3v019MastercardUniversalCardholderAuthenticationRequest {
	return &G3v019MastercardUniversalCardholderAuthenticationRequest{G3v019MastercardUniversalCardholderAuthentication, ucafindicator, ucafdata}
}

func (g *G3v019MastercardUniversalCardholderAuthenticationRequest) String() string {
	return fmt.Sprintf("%03d%c%s%c", g.VersionNumber, g.UcafCollectionIndicator, g.UcafAuthenticationData, bytes.FileSeparator)
}

func ParseG3v019MastercardUniversalCardholderAuthenticationRequest(s string) *G3v019MastercardUniversalCardholderAuthenticationRequest {
	g := G3v019MastercardUniversalCardholderAuthenticationRequest{}
	fmt.Sscanf(s, "%03d%1d", &g.VersionNumber, &g.UcafCollectionIndicator)
	i := strings.Index(s[4:], string(bytes.FileSeparator))
	if 4 < i && i < 36 {
		g.UcafAuthenticationData = s[4:i]
	}
	return &g
}

// ----------------------------------------------------
type G3v020DeveloperInformationRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 020
	DeveloperID   string        // 6 A/N 4.64 Developer ID (TSYS Acquiring Solutions assigned)
	VersionID     string        // 4 A/N 4.177 Version ID (TSYS Acquiring Solutions assigned)
	// 1 ASCII 4.80 Field Separator <FS>
	GatewayID string // 0, 10 A/N Gateway ID This data is optional
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v020DeveloperInformationRequest() *G3v020DeveloperInformationRequest {
	return &G3v020DeveloperInformationRequest{G3v020DeveloperInformation, TSYS_DEVELOPER_ID, TSYS_DEVELOPER_VERSION, ""}
}

func (g *G3v020DeveloperInformationRequest) String() string {
	return fmt.Sprintf("%03d%-6s%-4s%c%s%c", g.VersionNumber, g.DeveloperID, g.VersionID, bytes.FileSeparator, g.GatewayID, bytes.FileSeparator)
}

func ParseG3v020DeveloperInformationRequest(s string) *G3v020DeveloperInformationRequest {
	g := G3v020DeveloperInformationRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.DeveloperID = s[3:9]
	g.VersionID = s[9:13]

	if byte(bytes.FileSeparator) != s[14] {
		i := strings.Index(s[14:], string(bytes.FileSeparator))
		if 10 == i {
			g.GatewayID = s[14:24]
		}
	}

	return &g
}

// ----------------------------------------------------
type G3v021MerchantVerificationValueRequest struct {
	VersionNumber             G3VersionType // 3 NUM 4.85 Group III Version Number 021
	MerchantVerificationValue string        // 10 A/N 4.112 Merchant Verification Value 0-9, A-F only
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v021MerchantVerificationValueRequest(verificationValue string) *G3v021MerchantVerificationValueRequest {
	return &G3v021MerchantVerificationValueRequest{G3v021MerchantVerificationValue, verificationValue}
}

func (g *G3v021MerchantVerificationValueRequest) String() string {
	return fmt.Sprintf("%03d%-10s", g.VersionNumber, g.MerchantVerificationValue)
}

func ParseG3v021MerchantVerificationValueRequest(s string) *G3v021MerchantVerificationValueRequest {
	g := G3v021MerchantVerificationValueRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.MerchantVerificationValue = s[3:13]
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
type G3v022AdditionalAmountsRequest struct {
	VersionNumber     G3VersionType //3 NUM
	AdditionalAmounts [4]AdditionalAmountInfo
}

func NewG3v022AdditionalAmountsRequest(amounts *[4]AdditionalAmountInfo) *G3v022AdditionalAmountsRequest {
	g := G3v022AdditionalAmountsRequest{VersionNumber: G3v022AdditionalAmounts}

	for n := 0; n < 4 && n < len(amounts); n++ {
		g.AdditionalAmounts[n].Copy(&amounts[n])
	}

	return &g
}

func (g *G3v022AdditionalAmountsRequest) String() string {
	return fmt.Sprintf("%03d%s%s%s%s", g.VersionNumber, g.AdditionalAmounts[0].String(), g.AdditionalAmounts[1].String(),
		g.AdditionalAmounts[2].String(), g.AdditionalAmounts[3].String())
}

func ParseG3v022AdditionalAmountsRequest(s string) *G3v022AdditionalAmountsRequest {
	g := G3v022AdditionalAmountsRequest{}
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
// Request message - Visa and MasterCard healthcare data (version 023)
type G3v023VisaMastercardHealthcareRequest struct {
	VersionNumber  G3VersionType // 3 NUM 4.85 Group III Version Number 023
	HealthcareData string        // III 0-199* ANS 4.87 Healthcare, MasterCard or Discover Member Defined Data
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v023VisaMastercardHealthcareRequest(health string) *G3v023VisaMastercardHealthcareRequest {
	return &G3v023VisaMastercardHealthcareRequest{G3v023VisaMastercardHealthcare, health}
}

func (g *G3v023VisaMastercardHealthcareRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.HealthcareData, bytes.FileSeparator)
}

func ParseG3v023VisaMastercardHealthcareRequest(s string) *G3v023VisaMastercardHealthcareRequest {
	g := G3v023VisaMastercardHealthcareRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 203 {
		g.HealthcareData = s[3:i]
	}

	return &g
}

// ----------------------------------------------------
// Request message - Merchant Advice Code (MAC) (version 024)
type G3v024MerchantAdviceCodeRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v024MerchantAdviceCodeRequest() *G3v024MerchantAdviceCodeRequest {
	return &G3v024MerchantAdviceCodeRequest{G3v024MerchantAdviceCode}
}

func (g *G3v024MerchantAdviceCodeRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v024MerchantAdviceCodeRequest(s string) *G3v024MerchantAdviceCodeRequest {
	g := G3v024MerchantAdviceCodeRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type TransactionFeeAmountType rune

const (
	CreditTransactionFee TransactionFeeAmountType = 'C'
	DebitTransactionFee  TransactionFeeAmountType = 'D'
)

func (x TransactionFeeAmountType) String() string {
	switch x {
	case CreditTransactionFee:
		return "CreditTransactionFee"
	case DebitTransactionFee:
		return "DebitTransactionFee"
	}
	return ""
}

// ----------------------------------------------------
// Request message - Transaction fee amount (version 025)
type G3v025TransactionFeeAmountRequest struct {
	VersionNumber        G3VersionType // 3 NUM 4.85 Group III Version Number 025
	FeeAmountIndicator   TransactionFeeAmountType
	TransactionFeeAmount uint64 // 0 or 9 A/N 4.165 Transaction Fee Amount (includes fee amount indicator C/D above)
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v025TransactionFeeAmountRequest(amountind TransactionFeeAmountType, fee uint64) *G3v025TransactionFeeAmountRequest {
	return &G3v025TransactionFeeAmountRequest{G3v025TransactionFeeAmount, amountind, fee}
}

func (g *G3v025TransactionFeeAmountRequest) String() string {
	// The format of the nine characters is “annnnnnnn” where “a” is either “D” for debit or “C” for
	// credit and where “nnnnnnnn” is the numeric fee amount with the decimal implied
	// “D00000150” is a $1.50 transaction fee amount debited to the cardholder's account.
	return fmt.Sprintf("%03d%c%08d%c", g.VersionNumber, g.FeeAmountIndicator, g.TransactionFeeAmount, bytes.FileSeparator)
}

func ParseG3v025TransactionFeeAmountRequest(s string) *G3v025TransactionFeeAmountRequest {
	g := G3v025TransactionFeeAmountRequest{}
	fmt.Sscanf(s, "%03d%c%08d", &g.VersionNumber, &g.FeeAmountIndicator, &g.TransactionFeeAmount)

	return &g
}

type ProductParticipationGroupType byte

const (
	// Values for merchants accepting Visa, MasterCard, Discover, PayPal and American Express cards are for merchandise sales only:
	MerchandiseSalesCanBePartiallyApproved ProductParticipationGroupType = 'A' // Cash Over can be partially approved (Not applicable to merchandise-only merchants).

	// Additional values for Discover merchants providing Cash Over with merchandise sales:
	// A MerchandiseAndCashOver Can Be Partially Approved
	MerchandiseOnlyCanBePartiallyApproved               ProductParticipationGroupType = 'B' // Cash Over must be fully approved or declined
	MerchandiseMustBeFullyApprovedOrDeclined            ProductParticipationGroupType = 'C' // Cash Over can be partially approved (only if merchandise fully approved)
	MerchandiseAndCashOverMustBeFullyApprovedOrDeclined ProductParticipationGroupType = 'D' // Cash Over must be fully approved (if merchandise is approved) or declined
)

// ----------------------------------------------------
type G3v026ProductParticipationGroupRequest struct {
	VersionNumber             G3VersionType // 3 NUM 4.85 Group III Version Number 026
	ProductParticipationGroup string        // 1 to 10 ALPHA 4.124 Product Participation Group
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v026ProductParticipationGroupRequest(participationgroup string) *G3v026ProductParticipationGroupRequest {
	return &G3v026ProductParticipationGroupRequest{G3v026ProductParticipationGroup, participationgroup}
}

func (g *G3v026ProductParticipationGroupRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.ProductParticipationGroup, bytes.FileSeparator)
}

func ParseG3v026ProductParticipationGroupRequest(s string) *G3v026ProductParticipationGroupRequest {
	g := G3v026ProductParticipationGroupRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 13 {
		g.ProductParticipationGroup = s[3:i]
	}
	return &g
}

type TerminalInputCapabilityType rune

const (
	TerminalInputUnspecifiedTerminal                              TerminalInputCapabilityType = '0'
	TerminalInputManualNoTerminal                                 TerminalInputCapabilityType = '1'
	TerminalInputMagneticStripeReaderTerminal                     TerminalInputCapabilityType = '2'
	TerminalInputBarCodePaymentCodeTerminal                       TerminalInputCapabilityType = '3'
	TerminalInputOpticalCharacterReaderTerminal                   TerminalInputCapabilityType = '4' //[MC] only
	TerminalInputIntegratedCircuitCardTerminal                    TerminalInputCapabilityType = '5'
	TerminalInputKeyEntryOnlyTerminal                             TerminalInputCapabilityType = '6'
	TerminalInputPANAutoEntryViaContactlessMagneticStripeTerminal TerminalInputCapabilityType = 'A'
	TerminalInputMagneticStripeReaderAndKeyEntryTerminal          TerminalInputCapabilityType = 'B'
	TerminalInputMagneticStripeReaderICCAndKeyEntryTerminal       TerminalInputCapabilityType = 'C'
	TerminalInputMagneticStripeReaderAndICCTerminal               TerminalInputCapabilityType = 'D'
	TerminalInputICCAndKeyEntryTerminal                           TerminalInputCapabilityType = 'E'
	TerminalInputICCReaderAndContactlessTerminal                  TerminalInputCapabilityType = 'H'
	TerminalInputPANAutoEntryViaContactlessChipTerminal           TerminalInputCapabilityType = 'M'
	TerminalInputOtherTerminal                                    TerminalInputCapabilityType = 'V' //[MC] only
	TerminalInputMagneticStripeSignatureTerminal                  TerminalInputCapabilityType = 'X' //[AX] only
)

func (x TerminalInputCapabilityType) String() string {
	switch x {
	case TerminalInputUnspecifiedTerminal:
		return "TerminalInputUnspecifiedTerminal"
	case TerminalInputManualNoTerminal:
		return "TerminalInputManualNoTerminal"
	case TerminalInputMagneticStripeReaderTerminal:
		return "TerminalInputMagneticStripeReaderTerminal"
	case TerminalInputBarCodePaymentCodeTerminal:
		return "TerminalInputBarCodePaymentCodeTerminal"
	case TerminalInputOpticalCharacterReaderTerminal:
		return "TerminalInputOpticalCharacterReaderTerminal"
	case TerminalInputIntegratedCircuitCardTerminal:
		return "TerminalInputIntegratedCircuitCardTerminal"
	case TerminalInputKeyEntryOnlyTerminal:
		return "TerminalInputKeyEntryOnlyTerminal"
	case TerminalInputPANAutoEntryViaContactlessMagneticStripeTerminal:
		return "TerminalInputPANAutoEntryViaContactlessMagneticStripeTerminal"
	case TerminalInputMagneticStripeReaderAndKeyEntryTerminal:
		return "TerminalInputMagneticStripeReaderAndKeyEntryTerminal"
	case TerminalInputMagneticStripeReaderICCAndKeyEntryTerminal:
		return "TerminalInputMagneticStripeReaderICCAndKeyEntryTerminal"
	case TerminalInputMagneticStripeReaderAndICCTerminal:
		return "TerminalInputMagneticStripeReaderAndICCTerminal"
	case TerminalInputICCAndKeyEntryTerminal:
		return "TerminalInputICCAndKeyEntryTerminal"
	case TerminalInputICCReaderAndContactlessTerminal:
		return "TerminalInputICCReaderAndContactlessTerminal"
	case TerminalInputPANAutoEntryViaContactlessChipTerminal:
		return "TerminalInputPANAutoEntryViaContactlessChipTerminal"
	case TerminalInputOtherTerminal:
		return "TerminalInputOtherTerminal"
	case TerminalInputMagneticStripeSignatureTerminal:
		return "TerminalInputMagneticStripeSignatureTerminal"
	}
	return ""
}

type CardholderAuthenticationCapabilityType uint

const (
	CardholderAuthenticationNoElectronicAuthenticationCapability          CardholderAuthenticationCapabilityType = 0
	CardholderAuthenticationPINEntryCapability                            CardholderAuthenticationCapabilityType = 1
	CardholderAuthenticationElectronicSignatureAnalysisCapability         CardholderAuthenticationCapabilityType = 2
	CardholderAuthenticationElectronicAuthenticationCapabilityInoperative CardholderAuthenticationCapabilityType = 5
	CardholderAuthenticationOther                                         CardholderAuthenticationCapabilityType = 6
	CardholderAuthenticationUnspecified                                   CardholderAuthenticationCapabilityType = 9
)

func (x CardholderAuthenticationCapabilityType) String() string {
	switch x {
	case CardholderAuthenticationNoElectronicAuthenticationCapability:
		return "CardholderAuthenticationNoElectronicAuthenticationCapability"
	case CardholderAuthenticationPINEntryCapability:
		return "CardholderAuthenticationPINEntryCapability"
	case CardholderAuthenticationElectronicSignatureAnalysisCapability:
		return "CardholderAuthenticationElectronicSignatureAnalysisCapability"
	case CardholderAuthenticationElectronicAuthenticationCapabilityInoperative:
		return "CardholderAuthenticationElectronicAuthenticationCapabilityInoperative"
	case CardholderAuthenticationOther:
		return "CardholderAuthenticationOther"
	case CardholderAuthenticationUnspecified:
		return "CardholderAuthenticationUnspecified"
	}
	return ""
}

type CardCaptureCapabilityType uint

const (
	CardCaptureNoCaptureCapability   CardCaptureCapabilityType = 0
	CardCaptureCapability            CardCaptureCapabilityType = 1
	CardCaptureCapabilityUnspecified CardCaptureCapabilityType = 9
)

func (x CardCaptureCapabilityType) String() string {
	switch x {
	case CardCaptureNoCaptureCapability:
		return "CardCaptureNoCaptureCapability"
	case CardCaptureCapability:
		return "CardCaptureCapability"
	case CardCaptureCapabilityUnspecified:
		return "CardCaptureCapabilityUnspecified"
	}
	return ""
}

type TerminalOperatingEnvironmentType rune

const (
	EnvironmentNoTerminalUsed                            TerminalOperatingEnvironmentType = '0'
	EnvironmentOnCardAcceptorPremisesAttendedTerminal    TerminalOperatingEnvironmentType = '1'
	EnvironmentOnCardAcceptorPremisesUnattendedTerminal  TerminalOperatingEnvironmentType = '2'
	EnvironmentOffCardAcceptorPremisesAttendedTerminal   TerminalOperatingEnvironmentType = '3'
	EnvironmentOffCardAcceptorPremisesUnattendedTerminal TerminalOperatingEnvironmentType = '4'
	EnvironmentOnCardholderPremisesUnattendedTerminal    TerminalOperatingEnvironmentType = '5'
	EnvironmentOffCardholderPremisesUnattendedTerminal   TerminalOperatingEnvironmentType = '6'
	EnvironmentUnspecifiedDataNotAvailableTerminal       TerminalOperatingEnvironmentType = '9'
	EnvironmentMobileOffPremisesTerminal                 TerminalOperatingEnvironmentType = 'M' //[V] only
	EnvironmentMobileOnPremisesTerminal                  TerminalOperatingEnvironmentType = 'P' //[V] only
	EnvironmentElectronicDeliveryOfProductTerminal       TerminalOperatingEnvironmentType = 'S' //[AX] only
	EnvironmentPhysicalDeliveryOfProductTerminal         TerminalOperatingEnvironmentType = 'T' //[AX] only
)

func (x TerminalOperatingEnvironmentType) String() string {
	switch x {
	case EnvironmentNoTerminalUsed:
		return "EnvironmentNoTerminalUsed"
	case EnvironmentOnCardAcceptorPremisesAttendedTerminal:
		return "EnvironmentOnCardAcceptorPremisesAttendedTerminal"
	case EnvironmentOnCardAcceptorPremisesUnattendedTerminal:
		return "EnvironmentOnCardAcceptorPremisesUnattendedTerminal"
	case EnvironmentOffCardAcceptorPremisesAttendedTerminal:
		return "EnvironmentOffCardAcceptorPremisesAttendedTerminal"
	case EnvironmentOffCardAcceptorPremisesUnattendedTerminal:
		return "EnvironmentOffCardAcceptorPremisesUnattendedTerminal"
	case EnvironmentOnCardholderPremisesUnattendedTerminal:
		return "EnvironmentOnCardholderPremisesUnattendedTerminal"
	case EnvironmentOffCardholderPremisesUnattendedTerminal:
		return "EnvironmentOffCardholderPremisesUnattendedTerminal"
	case EnvironmentUnspecifiedDataNotAvailableTerminal:
		return "EnvironmentUnspecifiedDataNotAvailableTerminal"
	case EnvironmentMobileOffPremisesTerminal:
		return "EnvironmentMobileOffPremisesTerminal"
	case EnvironmentMobileOnPremisesTerminal:
		return "EnvironmentMobileOnPremisesTerminal"
	case EnvironmentElectronicDeliveryOfProductTerminal:
		return "EnvironmentElectronicDeliveryOfProductTerminal"
	case EnvironmentPhysicalDeliveryOfProductTerminal:
		return "EnvironmentPhysicalDeliveryOfProductTerminal"
	}
	return ""
}

type CardholderPresentDataType uint

const (
	CardholderPresent                        CardholderPresentDataType = 0
	CardholderNotPresentUnspecifiedReason    CardholderPresentDataType = 1
	CardholderNotPresentMailTransaction      CardholderPresentDataType = 2
	CardholderNotPresentPhoneTransaction     CardholderPresentDataType = 3
	CardholderNotPresentRecurringTransaction CardholderPresentDataType = 4
	CardholderNotPresentElectronicCommerce   CardholderPresentDataType = 5
	CardholderNotPresentRecurrentBilling     CardholderPresentDataType = 6
	CardholderPresentDataUnspecified         CardholderPresentDataType = 7
)

func (x CardholderPresentDataType) String() string {
	switch x {
	case CardholderPresent:
		return "CardholderPresent"
	case CardholderNotPresentUnspecifiedReason:
		return "CardholderNotPresentUnspecifiedReason"
	case CardholderNotPresentMailTransaction:
		return "CardholderNotPresentMailTransaction"
	case CardholderNotPresentPhoneTransaction:
		return "CardholderNotPresentPhoneTransaction"
	case CardholderNotPresentRecurringTransaction:
		return "CardholderNotPresentRecurringTransaction"
	case CardholderNotPresentElectronicCommerce:
		return "CardholderNotPresentElectronicCommerce"
	case CardholderNotPresentRecurrentBilling:
		return "CardholderNotPresentRecurrentBilling"
	case CardholderPresentDataUnspecified:
		return "CardholderPresentDataUnspecified"
	}
	return ""
}

type CardPresentDataType rune

const (
	CardNotPresent                 CardPresentDataType = '0'
	CardPresent                    CardPresentDataType = '1'
	CardPresentDataUnspecified     CardPresentDataType = '9'
	CardPresentDataTransponder     CardPresentDataType = 'W' //[AX] only
	CardPresentDataContactlessChip CardPresentDataType = 'X'
)

func (x CardPresentDataType) String() string {
	switch x {
	case CardNotPresent:
		return "CardNotPresent"
	case CardPresent:
		return "CardPresent"
	case CardPresentDataUnspecified:
		return "CardPresentDataUnspecified"
	case CardPresentDataTransponder:
		return "CardPresentDataTransponder"
	case CardPresentDataContactlessChip:
		return "CardPresentDataContactlessChip"
	}
	return ""
}

type CardInputModeType rune

const (
	CardInputUnspecified                                                            CardInputModeType = '0'
	CardInputManualInputNoTerminal                                                  CardInputModeType = '1'
	CardInputMagneticStripeReader                                                   CardInputModeType = '2'
	CardInputBarCodePaymentCode                                                     CardInputModeType = '3'
	CardInputKeyEntered                                                             CardInputModeType = '6'
	CardInputPANAutoEntryViaContactlessMagneticStripe                               CardInputModeType = 'A'
	CardInputMSRInputTrackDataCapturedAndPassedUnaltered                            CardInputModeType = 'B'
	CardInputOnlineChip                                                             CardInputModeType = 'C'
	CardInputOfflineChip                                                            CardInputModeType = 'F'
	CardInputPANAutoEntryViaContactlessChipCard                                     CardInputModeType = 'M' //(EMV Mode)
	CardInputTrackDataReadAndSentUnAlteredChipCapableTerminalChipDataCouldNotBeRead CardInputModeType = 'N'
	CardInputEmptyCandidateListFallback                                             CardInputModeType = 'P'
	CardInputEcomNoSecurityChannelEncryptedOrSetWithoutCardholderCertificate        CardInputModeType = 'S' //[MC]
	CardInputManuallyEnteredWithKeyedCID                                            CardInputModeType = 'V' //[AX JCB Canada]
	CardInputSwipedTransactionWithKeyedCID                                          CardInputModeType = 'W' //[AX JCBCanada]
	CardInputMagneticStripeSignature                                                CardInputModeType = 'X' //[AX JCB Canada]
	CardInputMagneticStripeSignatureWithKeyedCID                                    CardInputModeType = 'Y' //[AX JCB Canada]
)

func (x CardInputModeType) String() string {
	switch x {
	case CardInputUnspecified:
		return "CardInputUnspecified"
	case CardInputManualInputNoTerminal:
		return "CardInputManualInputNoTerminal"
	case CardInputMagneticStripeReader:
		return "CardInputMagneticStripeReader"
	case CardInputBarCodePaymentCode:
		return "CardInputBarCodePaymentCode"
	case CardInputKeyEntered:
		return "CardInputKeyEntered"
	case CardInputPANAutoEntryViaContactlessMagneticStripe:
		return "CardInputPANAutoEntryViaContactlessMagneticStripe"
	case CardInputMSRInputTrackDataCapturedAndPassedUnaltered:
		return "CardInputMSRInputTrackDataCapturedAndPassedUnaltered"
	case CardInputOnlineChip:
		return "CardInputOnlineChip"
	case CardInputOfflineChip:
		return "CardInputOfflineChip"
	case CardInputPANAutoEntryViaContactlessChipCard:
		return "CardInputPANAutoEntryViaContactlessChipCard"
	case CardInputTrackDataReadAndSentUnAlteredChipCapableTerminalChipDataCouldNotBeRead:
		return "CardInputTrackDataReadAndSentUnAlteredChipCapableTerminalChipDataCouldNotBeRead"
	case CardInputEmptyCandidateListFallback:
		return "CardInputEmptyCandidateListFallback"
	case CardInputEcomNoSecurityChannelEncryptedOrSetWithoutCardholderCertificate:
		return "CardInputEcomNoSecurityChannelEncryptedOrSetWithoutCardholderCertificate"
	case CardInputManuallyEnteredWithKeyedCID:
		return "CardInputManuallyEnteredWithKeyedCID"
	case CardInputSwipedTransactionWithKeyedCID:
		return "CardInputSwipedTransactionWithKeyedCID"
	case CardInputMagneticStripeSignature:
		return "CardInputMagneticStripeSignature"
	case CardInputMagneticStripeSignatureWithKeyedCID:
		return "CardInputMagneticStripeSignatureWithKeyedCID"
	}
	return ""
}

type CardholderAuthenticationMethodType rune

const (
	CardholderNotAuthenticated                          CardholderAuthenticationMethodType = '0'
	CardholderPINAuthenticated                          CardholderAuthenticationMethodType = '1'
	CardholderElectronicSignatureAnalysisAuthenticated  CardholderAuthenticationMethodType = '2'
	CardholderManualSignatureVerification               CardholderAuthenticationMethodType = '5'
	CardholderOtherManualVerification                   CardholderAuthenticationMethodType = '6' //(such as a driver’s license number)
	CardholderUnspecifiedAuthentication                 CardholderAuthenticationMethodType = '9'
	CardholderOtherSystematicVerification               CardholderAuthenticationMethodType = 'S'
	CardholderElectronicTicketEnvironmentAuthentication CardholderAuthenticationMethodType = 'T' //[AX] Amex only
)

func (x CardholderAuthenticationMethodType) String() string {
	switch x {
	case CardholderNotAuthenticated:
		return "CardholderNotAuthenticated"
	case CardholderPINAuthenticated:
		return "CardholderPINAuthenticated"
	case CardholderElectronicSignatureAnalysisAuthenticated:
		return "CardholderElectronicSignatureAnalysisAuthenticated"
	case CardholderManualSignatureVerification:
		return "CardholderManualSignatureVerification"
	case CardholderOtherManualVerification:
		return "CardholderOtherManualVerification"
	case CardholderUnspecifiedAuthentication:
		return "CardholderUnspecifiedAuthentication"
	case CardholderOtherSystematicVerification:
		return "CardholderOtherSystematicVerification"
	case CardholderElectronicTicketEnvironmentAuthentication:
		return "CardholderElectronicTicketEnvironmentAuthentication"
	}
	return ""
}

type CardholderAuthenticationEntityType uint

const (
	AuthenticationEntityNotAuthenticated              CardholderAuthenticationEntityType = 0
	AuthenticationEntityICCOfflinePIN                 CardholderAuthenticationEntityType = 1
	AuthenticationEntityCardAcceptanceDevice          CardholderAuthenticationEntityType = 2 //(CAD)
	AuthenticationEntityAuthorizingAgentOnlinePIN     CardholderAuthenticationEntityType = 3
	AuthenticationEntityMerchantCardAcceptorSignature CardholderAuthenticationEntityType = 4
	AuthenticationEntityOther                         CardholderAuthenticationEntityType = 5
	AuthenticationEntityUnspecified                   CardholderAuthenticationEntityType = 9
)

func (x CardholderAuthenticationEntityType) String() string {
	switch x {
	case AuthenticationEntityNotAuthenticated:
		return "AuthenticationEntityNotAuthenticated"
	case AuthenticationEntityICCOfflinePIN:
		return "AuthenticationEntityICCOfflinePIN"
	case AuthenticationEntityCardAcceptanceDevice:
		return "AuthenticationEntityCardAcceptanceDevice"
	case AuthenticationEntityAuthorizingAgentOnlinePIN:
		return "AuthenticationEntityAuthorizingAgentOnlinePIN"
	case AuthenticationEntityMerchantCardAcceptorSignature:
		return "AuthenticationEntityMerchantCardAcceptorSignature"
	case AuthenticationEntityOther:
		return "AuthenticationEntityOther"
	case AuthenticationEntityUnspecified:
		return "AuthenticationEntityUnspecified"
	}
	return ""
}

type CardOutputCapabilityType rune

const (
	CardOutputCapabilityUnspecified         CardOutputCapabilityType = '0'
	CardOutputCapabilityNone                CardOutputCapabilityType = '1'
	CardOutputCapabilityMagneticStripeWrite CardOutputCapabilityType = '2'
	CardOutputCapabilityICC                 CardOutputCapabilityType = '3'
	CardOutputCapabilityOther               CardOutputCapabilityType = 'S'
)

func (x CardOutputCapabilityType) String() string {
	switch x {
	case CardOutputCapabilityUnspecified:
		return "CardOutputCapabilityUnspecified"
	case CardOutputCapabilityNone:
		return "CardOutputCapabilityNone"
	case CardOutputCapabilityMagneticStripeWrite:
		return "CardOutputCapabilityMagneticStripeWrite"
	case CardOutputCapabilityICC:
		return "CardOutputCapabilityICC"
	case CardOutputCapabilityOther:
		return "CardOutputCapabilityOther"
	}
	return ""
}

type TerminalOutputCapabilityType uint

const (
	TerminalOutputCapabilityUnspecified         TerminalOutputCapabilityType = 0
	TerminalOutputCapabilityNone                TerminalOutputCapabilityType = 1
	TerminalOutputCapabilityMagneticStripeWrite TerminalOutputCapabilityType = 2
	TerminalOutputCapabilityICC                 TerminalOutputCapabilityType = 3
	TerminalOutputCapabilityOther               TerminalOutputCapabilityType = 4
)

func (x TerminalOutputCapabilityType) String() string {
	switch x {
	case TerminalOutputCapabilityUnspecified:
		return "TerminalOutputCapabilityUnspecified"
	case TerminalOutputCapabilityNone:
		return "TerminalOutputCapabilityNone"
	case TerminalOutputCapabilityMagneticStripeWrite:
		return "TerminalOutputCapabilityMagneticStripeWrite"
	case TerminalOutputCapabilityICC:
		return "TerminalOutputCapabilityICC"
	case TerminalOutputCapabilityOther:
		return "TerminalOutputCapabilityOther"
	}
	return ""
}

type PINCaptureCapabilityType rune

const (
	NoPINCaptureCapability                  PINCaptureCapabilityType = '0'
	PINCaptureCapabilityUnspecified         PINCaptureCapabilityType = '1'
	PINCaptureCapability02Reserved          PINCaptureCapabilityType = '2'
	PINCaptureCapability03Reserved          PINCaptureCapabilityType = '3'
	PINCaptureCapability04CharactersMaximum PINCaptureCapabilityType = '4'
	PINCaptureCapability05CharactersMaximum PINCaptureCapabilityType = '5'
	PINCaptureCapability06CharactersMaximum PINCaptureCapabilityType = '6'
	PINCaptureCapability07CharactersMaximum PINCaptureCapabilityType = '7'
	PINCaptureCapability08CharactersMaximum PINCaptureCapabilityType = '8'
	PINCaptureCapability09CharactersMaximum PINCaptureCapabilityType = '9'
	PINCaptureCapability10CharactersMaximum PINCaptureCapabilityType = 'A'
	PINCaptureCapability11CharactersMaximum PINCaptureCapabilityType = 'B'
	PINCaptureCapability12CharactersMaximum PINCaptureCapabilityType = 'C'
)

func (x PINCaptureCapabilityType) String() string {
	switch x {
	case NoPINCaptureCapability:
		return "NoPINCaptureCapability"
	case PINCaptureCapabilityUnspecified:
		return "PINCaptureCapabilityUnspecified"
	case PINCaptureCapability02Reserved:
		return "PINCaptureCapability02Reserved"
	case PINCaptureCapability03Reserved:
		return "PINCaptureCapability03Reserved"
	case PINCaptureCapability04CharactersMaximum:
		return "PINCaptureCapability04CharactersMaximum"
	case PINCaptureCapability05CharactersMaximum:
		return "PINCaptureCapability05CharactersMaximum"
	case PINCaptureCapability06CharactersMaximum:
		return "PINCaptureCapability06CharactersMaximum"
	case PINCaptureCapability07CharactersMaximum:
		return "PINCaptureCapability07CharactersMaximum"
	case PINCaptureCapability08CharactersMaximum:
		return "PINCaptureCapability08CharactersMaximum"
	case PINCaptureCapability09CharactersMaximum:
		return "PINCaptureCapability09CharactersMaximum"
	case PINCaptureCapability10CharactersMaximum:
		return "PINCaptureCapability10CharactersMaximum"
	case PINCaptureCapability11CharactersMaximum:
		return "PINCaptureCapability11CharactersMaximum"
	case PINCaptureCapability12CharactersMaximum:
		return "PINCaptureCapability12CharactersMaximum"
	}
	return ""
}

//----------------------------------------------------
// The POS data code is a fixed string of 12 characters S1 to S12 that indicate the condition of
// the POS device at the time of transaction (see Table 3.50 for record format and version
// number).
// The information in the POS data code takes precedence over the values in the Account Data
// Source (ADSC) and the Cardholder Identification Code (CID) fields. Information in the POS
// data code should accurately indicate the condition of the POS device at the time of transaction.
// Even though the POS data code has precedence over the ADSC and CID fields, the ADSC
// and CID fields still need to be populated to match the condition of the transaction as closely
// as possible.
// Group 3 Version 27 must be sent on all transactions from any device that is capable of reading
// an ICC Chip Card or an NFC Chip Card.

type G3v027PosDataRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 027
	//12 A/N 4.123 POS Data Code
	//Subfield 1: Terminal data - card data input capability
	TerminalInputCapability TerminalInputCapabilityType //1 A/N
	//Subfield 2: terminal data - cardholder authentication capability
	CardholderAuthenticationCapability CardholderAuthenticationCapabilityType //1 NUM
	//Subfield 3: terminal data - card capture capability
	CardCaptureCapability CardCaptureCapabilityType //1 NUM
	//Subfield 4: terminal operating environment
	TerminalOperatingEnvironment TerminalOperatingEnvironmentType //1 A/N
	//Subfield 5: cardholder present data
	CardholderPresentData CardholderPresentDataType //1 NUM
	//Subfield 6: card present data
	CardPresentData CardPresentDataType //1 A/N
	//Subfield 7: card data - input mode
	CardInputMode CardInputModeType //1 A/N
	//Subfield 8: cardholder authentication method
	CardholderAuthenticationMethod CardholderAuthenticationMethodType //1 A/N
	//Subfield 9: cardholder authentication entity
	CardholderAuthenticationEntity CardholderAuthenticationEntityType //1 NUM
	//Subfield 10: card data output capability
	CardOutputCapability CardOutputCapabilityType //1 A/N
	//Subfield 11: terminal data output capability
	TerminalOutputCapability TerminalOutputCapabilityType //1 NUM
	//Subfield 12: PIN capture capability
	PINCaptureCapability PINCaptureCapabilityType //1 A/N
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v027PosDataRequest(termcap TerminalInputCapabilityType,
	authcap CardholderAuthenticationCapabilityType,
	cardcapture CardCaptureCapabilityType,
	operenvironment TerminalOperatingEnvironmentType,
	cardholderpresent CardholderPresentDataType,
	cardpresent CardPresentDataType,
	inputmode CardInputModeType,
	authmethod CardholderAuthenticationMethodType,
	authentity CardholderAuthenticationEntityType,
	outputcaps CardOutputCapabilityType,
	termoutput TerminalOutputCapabilityType,
	pincapture PINCaptureCapabilityType) *G3v027PosDataRequest {
	return &G3v027PosDataRequest{G3v027PosData, termcap, authcap, cardcapture, operenvironment, cardholderpresent, cardpresent, inputmode, authmethod, authentity, outputcaps, termoutput, pincapture}
}

func (g *G3v027PosDataRequest) String() string {
	return fmt.Sprintf("%03d%c%d%d%c%d%c%c%c%d%c%d%c", g.VersionNumber, g.TerminalInputCapability, g.CardholderAuthenticationCapability, g.CardCaptureCapability,
		g.TerminalOperatingEnvironment, g.CardholderPresentData, g.CardPresentData, g.CardInputMode, g.CardholderAuthenticationMethod,
		g.CardholderAuthenticationEntity, g.CardOutputCapability, g.TerminalOutputCapability, g.PINCaptureCapability)
}

func ParseG3v027PosDataRequest(s string) *G3v027PosDataRequest {
	g := G3v027PosDataRequest{}
	fmt.Sscanf(s, "%03d%c%1d%1d%c%1d%c%c%c%1d%c%1d%c", &g.VersionNumber, &g.TerminalInputCapability, &g.CardholderAuthenticationCapability, &g.CardCaptureCapability,
		&g.TerminalOperatingEnvironment, &g.CardholderPresentData, &g.CardPresentData, &g.CardInputMode, &g.CardholderAuthenticationMethod,
		&g.CardholderAuthenticationEntity, &g.CardOutputCapability, &g.TerminalOutputCapability, &g.PINCaptureCapability)

	return &g
}

func (g *G3v027PosDataRequest) VerboseString() string {

	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.VersionNumber", g.VersionNumber)                                           //1 A/N
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.TerminalInputCapability", g.TerminalInputCapability)                       //1 A/N
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.CardholderAuthenticationCapability", g.CardholderAuthenticationCapability) //1 NUM
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.CardCaptureCapability", g.CardCaptureCapability)                           //1 NUM
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.TerminalOperatingEnvironment", g.TerminalOperatingEnvironment)             //1 A/N
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.CardholderPresentData", g.CardholderPresentData)                           //1 NUM
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.CardPresentData", g.CardPresentData)                                       //1 A/N
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.CardInputMode", g.CardInputMode)                                           //1 A/N
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.CardholderAuthenticationMethod", g.CardholderAuthenticationMethod)         //1 A/N
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.CardholderAuthenticationEntity", g.CardholderAuthenticationEntity)         //1 NUM
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.CardOutputCapability", g.CardOutputCapability)                             //1 A/N
	buffer.AppendFormat("%-42s %[2]d (%[2]s)\n", "G3v027.TerminalOutputCapability", g.TerminalOutputCapability)                     //1 NUM
	buffer.AppendFormat("%-42s %[2]c (%[2]s)\n", "G3v027.PINCaptureCapability", g.PINCaptureCapability)                             //1 A/N

	return buffer.String()
}

//----------------------------------------------------
// !!! incomplete

// This field is optionally used specifically for American Express transactions, to hold additional
// data. Only one of these formats can be used at a time.
// • The ITD format has a minimum length of 74 bytes and a maximum of 265.
// • The APD format has a minimum of 151 bytes and a maximum of 290.
// • The Card Present - Goods Sold format has a specific length of 16 bytes.

type G3v028AmericanExpressAdditionalDataRequest struct {
	VersionNumber      G3VersionType // 3 NUM 4.85 Group III Version Number 028
	AmexAdditionalData string        // 16-287 ANS 4.17 Amex Additional Data
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v028AmericanExpressAdditionalDataRequest(data string) *G3v028AmericanExpressAdditionalDataRequest {
	return &G3v028AmericanExpressAdditionalDataRequest{G3v028AmericanExpressAdditionalData, data}
}

func (g *G3v028AmericanExpressAdditionalDataRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.AmexAdditionalData, bytes.FileSeparator)
}

func ParseG3v028AmericanExpressAdditionalDataRequest(s string) *G3v028AmericanExpressAdditionalDataRequest {
	g := G3v028AmericanExpressAdditionalDataRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 290 {
		g.AmexAdditionalData = s[3:i]
	}

	return &g
}

type ContactAddressType struct {
	PostalCode string // 0-9 A/N 4.78.9 Ship-to Postal Code
	// 1 ASCII 4.80 Field Separator <FS>
	Address string // 0-50 A/N 4.78.10 Ship-to Address
	// 1 ASCII 4.80 Field Separator <FS>
	City string // 0 ALPHA 4.78.11 Ship-to City (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	StateProvince string // 0 A/N 4.78.12 Ship-to State/Province (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	CountryCode CountryCodeType // 0 or 3 NUM 4.78.13 Ship-to Country Code
	// 1 ASCII 4.80 Field Separator <FS>
	FirstName string // 0-15 ALPHA 4.78.14 Ship-to First Name
	// 1 ASCII 4.80 Field Separator <FS>
	LastName string // 0-30 ALPHA 4.78.15 Ship-to Last Name
	// 1 ASCII 4.80 Field Separator <FS>
	PhoneNumber uint64 // 0-10 NUM 4.78.16 Ship-to Phone Number
	// 1 ASCII 4.80 Field Separator <FS>
}

func NewContactAddressType(postal string, street string, city string, state string, country CountryCodeType, first string, last string, phone uint64) *ContactAddressType {
	return &ContactAddressType{postal, street, city, state, country, first, last, phone}
}

func (a *ContactAddressType) String() string {
	return fmt.Sprintf("%s%c%s%c%s%c%s%c%03d%c%s%c%s%c%d%c", a.PostalCode, bytes.FileSeparator,
		a.Address, bytes.FileSeparator,
		a.City, bytes.FileSeparator,
		a.StateProvince, bytes.FileSeparator,
		a.CountryCode, bytes.FileSeparator,
		a.FirstName, bytes.FileSeparator,
		a.LastName, bytes.FileSeparator,
		a.PhoneNumber, bytes.FileSeparator)
}

func ParseContactAddressType(s string) (*ContactAddressType, int) {
	a := ContactAddressType{}

	i := 0
	fields := strings.Split(s, string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				a.PostalCode = fields[j]
			case 1:
				a.Address = fields[j]
			case 2:
				a.City = fields[j]
			case 3:
				a.StateProvince = fields[j]
			case 4:
				fmt.Sscanf(fields[j], "%03d", &a.CountryCode)
			case 5:
				a.FirstName = fields[j]
			case 6:
				a.LastName = fields[j]
			case 7:
				a.PhoneNumber, _ = strconv.ParseUint(fields[j], 10, 64)
			}
		}
		i += len(fields[j]) + 1
	}

	return &a, i
}

//----------------------------------------------------
// For American Express transactions, Group 3 Version 29 should be sent on all American
// Express transactions that contain Address Verification Data. For American Express, extended
// AVS Data in G3V29 takes priority over AVS Data sent in the Cardholder Identification Data
// field. If Extended AVS Data is sent in G3V29, AVS Data should not be sent in the Cardholder
// Identification Data field. If Extended AVS Data is sent in G3V29 and AVS Data is sent in the
// Cardholder Identification Data field, the AVS Data in the Cardholder Identification Data field
// will not be used. If this Group is used, there must at least be data in the Cardholder Billing
// Postal Code field or the transaction will be rejected.
// For Discover transactions, this structure is used for the Enhanced Address Verification
// Service, for Cardholder First and Last Names only. The postal code and first five characters of
// the street address should still be sent in Cardholder identification data as usual. For Discover
// transactions, the Cardholder Billing Postal Code is not required.

// NOTE The POS Device should send an empty Group 3 Version 48 whenever Group 3 Version
// 29 data is sent. This is to receive the response

type G3v029ExtendedAVSdataRequest struct {
	VersionNumber     G3VersionType // 3 NUM 4.85 Group III Version Number 029
	CardholderAddress *ContactAddressType
	ShipToAddress     *ContactAddressType
	// 5-9 A/N 4.78.1 Cardholder Billing Postal Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-20 A/N 4.78.2 Cardholder Billing (Street) Address
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 ALPHA 4.78.3 Cardholder Billing City (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 A/N 4.78.4 Cardholder Billing State/Province (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 NUM 4.78.5 Cardholder Billing Country Code (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-35 ALPHA 4.78.6 Cardholder Billing First Name (Amex) Cardholder First Name (Discover) Value is truncated at 15 characters for Amex transactions
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-35 ALPHA 4.78.7 Cardholder Billing Last Name (Amex) Cardholder Last Name (Discover) Value is truncated at 30 characters for Amex transactions
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-10 NUM 4.78.8 Cardholder Billing Phone Number
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-9 A/N 4.78.9 Ship-to Postal Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-50 A/N 4.78.10 Ship-to Address
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 ALPHA 4.78.11 Ship-to City (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 A/N 4.78.12 Ship-to State/Province (future use)
	// 1 ASCII 4.80 Field Separator <FS>
	// 0 or 3 NUM 4.78.13 Ship-to Country Code
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-15 ALPHA 4.78.14 Ship-to First Name
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-30 ALPHA 4.78.15 Ship-to Last Name
	// 1 ASCII 4.80 Field Separator <FS>
	// 0-10 NUM 4.78.16 Ship-to Phone Number
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>

}

func NewG3v029ExtendedAVSdataRequest(cardholder *ContactAddressType, shipping *ContactAddressType) *G3v029ExtendedAVSdataRequest {
	return &G3v029ExtendedAVSdataRequest{G3v029ExtendedAVSdata, cardholder, shipping}
}

func (g *G3v029ExtendedAVSdataRequest) String() string {
	return fmt.Sprintf("%03d%s%s", g.VersionNumber, g.CardholderAddress.String(), g.ShipToAddress.String())
}

func ParseG3v029ExtendedAVSdataRequest(s string) *G3v029ExtendedAVSdataRequest {
	g := G3v029ExtendedAVSdataRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 3
	g.CardholderAddress, i = ParseContactAddressType(s[3:])
	g.ShipToAddress, i = ParseContactAddressType(s[i:])

	return &g
}

// ----------------------------------------------------
type G3v030AmericanExpressMerchantNameLocationRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 030
	//AmexMerchantName string        // 15-99 ANS 4.19 Amex Merchant
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>

	//AmexMerchantName (Location Data) has the following format
	//Oil Company CATs the name/location data will have the format:
	//S#ssssssssss\\\ppppp~~~~~\\ ---> \\\ is 3 delimiters, \\ is 2 delimiters
	//StationLocationCode string //ssssssssss is a variable length (12 bytes max), merchant assigned Station Location Code. 1-12 A/N
	/*PostalCode string */ //ppppp~~~~~ is the postal code, left justified space filled to 10 characters

	//For Aggregators, the name/location data/Telephone Number/Email Address will have the format:
	//S#vvvvvvvvvv\1234~abcdef\ccccccccc\ppppp~~~~~\bbbbbbbbbb\abc@123
	VendorOrStationLocationCode string //vvvvvvvvvv is a variable length (16 bytes max), merchant assigned Seller/Vendor Code.
	VendorStreetAddress         string //123~abcdef is the variable length (20 bytes max) seller/vendor street address.
	VendorCity                  string //ccccccccc is the variable length (13 bytes max) seller/vendor city.
	PostalCode                  string //ppppp~~~~~ is the postal code, left justified space filled to 10 characters.
	PhoneNumber                 uint64 //bbbbbbbbbb is the seller's business phone number
	Email                       string //abc@123 is the seller's email address (19 positions maximum)

}

func NewG3v030AmericanExpressMerchantNameLocationCATRequest(stationcode string, postal string) *G3v030AmericanExpressMerchantNameLocationRequest {
	return &G3v030AmericanExpressMerchantNameLocationRequest{VersionNumber: G3v030AmericanExpressMerchantNameLocation, VendorOrStationLocationCode: stationcode, PostalCode: postal}
}

func NewG3v030AmericanExpressMerchantNameLocationAggregatorRequest(vendorcode string, street string, city string, postal string, phone uint64, email string) *G3v030AmericanExpressMerchantNameLocationRequest {
	return &G3v030AmericanExpressMerchantNameLocationRequest{VersionNumber: G3v030AmericanExpressMerchantNameLocation,
		VendorOrStationLocationCode: vendorcode, VendorStreetAddress: street, VendorCity: city, PostalCode: postal,
		PhoneNumber: phone, Email: email}
}

func (g *G3v030AmericanExpressMerchantNameLocationRequest) String() string {

	// For Oil Company CATs the name/location data will have the format:
	// S#ssssssssss\\\ppppp~~~~~\\
	// For Aggregators, the name/location data/Telephone Number/Email Address will have the format:
	// S#vvvvvvvvvv\1234~abcdef\ccccccccc\ppppp~~~~~\bbbbbbbbbb\abc@123
	emailadd := g.Email
	if 19 < len(emailadd) {
		emailadd = emailadd[0:19]
	}

	return fmt.Sprintf("%03dS#%s\\%s\\%s\\%s\\%d\\%s%c", g.VersionNumber, g.VendorOrStationLocationCode, g.VendorStreetAddress, g.VendorCity, g.PostalCode, g.PhoneNumber, emailadd, bytes.FileSeparator)
}

func ParseG3v030AmericanExpressMerchantNameLocationRequest(s string) *G3v030AmericanExpressMerchantNameLocationRequest {
	g := G3v030AmericanExpressMerchantNameLocationRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	if 'S' != s[3] && '#' != s[4] {
		panic("Amex merchant name/location data should begin with a S# character sequence")
	}

	i := 0
	fields := strings.Split(s[5:], "\\")
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.VendorOrStationLocationCode = fields[j]
			case 1:
				g.VendorStreetAddress = fields[j]
			case 2:
				g.VendorCity = fields[j]
			case 3:
				g.PostalCode = fields[j]
			case 4:
				g.PhoneNumber, _ = strconv.ParseUint(fields[j], 10, 64)
			case 5:
				g.Email = fields[j]
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
// supports Visa’s Agent Identification service Service program.
type G3v031AgentIdentificationServiceRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 031
	AgentId       string        // 5 A/N 4.14 Agent Identification
	AgentIdResult string        // 12 A/N 4.15 Agent Identification Result
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v031AgentIdentificationServiceRequest(id string, result string) *G3v031AgentIdentificationServiceRequest {
	return &G3v031AgentIdentificationServiceRequest{G3v031AgentIdentificationService, id, result}
}

func (g *G3v031AgentIdentificationServiceRequest) String() string {
	return fmt.Sprintf("%03d%-5s%-12s", g.VersionNumber, g.AgentId, g.AgentIdResult)
}

func ParseG3v031AgentIdentificationServiceRequest(s string) *G3v031AgentIdentificationServiceRequest {
	g := G3v031AgentIdentificationServiceRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.AgentId = s[3:8]
	g.AgentIdResult = s[8:20]
	return &g
}

// ----------------------------------------------------
type G3v032CurrencyConversionDataRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v032CurrencyConversionDataRequest() *G3v032CurrencyConversionDataRequest {
	return &G3v032CurrencyConversionDataRequest{G3v032CurrencyConversionData}
}

func (g *G3v032CurrencyConversionDataRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v032CurrencyConversionDataRequest(s string) *G3v032CurrencyConversionDataRequest {
	g := G3v032CurrencyConversionDataRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type ReversalOrAdjustmentCodeType uint

const (
	CustomerDirectedReversal   ReversalOrAdjustmentCodeType = 17
	PartialReversal            ReversalOrAdjustmentCodeType = 32
	CardNotPresentSuspectFraud ReversalOrAdjustmentCodeType = 34
)

// ----------------------------------------------------
type G3v033ReversalRequestCodeRequest struct {
	VersionNumber            G3VersionType                // 3 NUM 4.85 Group III Version Number 033
	ReversalOrAdjustmentCode ReversalOrAdjustmentCodeType // 2 NUM 4.137 Reversal Request Code/Adjustment Response Code
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v033ReversalRequestCodeRequest(responsecode ReversalOrAdjustmentCodeType) *G3v033ReversalRequestCodeRequest {
	return &G3v033ReversalRequestCodeRequest{G3v033ReversalRequestCode, responsecode}
}

func (g *G3v033ReversalRequestCodeRequest) String() string {
	return fmt.Sprintf("%03d%02d", g.VersionNumber, g.ReversalOrAdjustmentCode)
}

func ParseG3v033ReversalRequestCodeRequest(s string) *G3v033ReversalRequestCodeRequest {
	g := G3v033ReversalRequestCodeRequest{}
	fmt.Sscanf(s, "%03d%02d", &g.VersionNumber, &g.ReversalOrAdjustmentCode)
	return &g
}

// ----------------------------------------------------
type G3v034CardProductCodeRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v034CardProductCodeRequest() *G3v034CardProductCodeRequest {
	return &G3v034CardProductCodeRequest{G3v034CardProductCode}
}

func (g *G3v034CardProductCodeRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v034CardProductCodeRequest(s string) *G3v034CardProductCodeRequest {
	g := G3v034CardProductCodeRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

//----------------------------------------------------
// The Discover/PayPal Network may establish promotional relationships with merchants. This
// field indicates the promotion to be used by the merchant for rewarding the cardholder at the
// point of service. The value sent in the authorization request will be echoed back in the
// authorization response

type G3v035PromotionalCodeRequest struct {
	VersionNumber  G3VersionType // 3 NUM 4.85 Group III Version Number 035
	PromotionCodes string        // 0-50 A/N 4.125 Promotion Codes Card specific format
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v035PromotionalCodeRequest(promocodes string) *G3v035PromotionalCodeRequest {
	return &G3v035PromotionalCodeRequest{G3v035PromotionalCode, promocodes}
}

func (g *G3v035PromotionalCodeRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.PromotionCodes, bytes.FileSeparator)
}

func ParseG3v035PromotionalCodeRequest(s string) *G3v035PromotionalCodeRequest {
	g := G3v035PromotionalCodeRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 53 {
		g.PromotionCodes = s[3:i]
	}
	return &g
}

type PaymentTransactionIdentifierType string

const (
	PersonToPerson                          PaymentTransactionIdentifierType = "C01"
	MasterCardRebate                        PaymentTransactionIdentifierType = "C02"
	RePowerLoadValue                        PaymentTransactionIdentifierType = "C03"
	GamingRepay                             PaymentTransactionIdentifierType = "C04"
	OtherPaymentTransaction                 PaymentTransactionIdentifierType = "C05"
	PaymentCreditCardBalanceWithCashOrCheck PaymentTransactionIdentifierType = "C06"
)

func (x PaymentTransactionIdentifierType) String() string {
	switch x {
	case PersonToPerson:
		return "PersonToPerson"
	case MasterCardRebate:
		return "MasterCardRebate"
	case RePowerLoadValue:
		return "RePowerLoadValue"
	case GamingRepay:
		return "GamingRepay"
	case OtherPaymentTransaction:
		return "OtherPaymentTransaction"
	case PaymentCreditCardBalanceWithCashOrCheck:
		return "PaymentCreditCardBalanceWithCashOrCheck"
	}
	return ""
}

// ----------------------------------------------------
type G3v036PaymentTransactionIdentifierRequest struct {
	VersionNumber                G3VersionType                    // 3 NUM 4.85 Group III Version Number 036
	PaymentTransactionIdentifier PaymentTransactionIdentifierType // 3 A/N 4.122 Payment Transaction Identifier
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v036PaymentTransactionIdentifierRequest(transactionid PaymentTransactionIdentifierType) *G3v036PaymentTransactionIdentifierRequest {
	return &G3v036PaymentTransactionIdentifierRequest{G3v036PaymentTransactionIdentifier, transactionid}
}

func (g *G3v036PaymentTransactionIdentifierRequest) String() string {
	return fmt.Sprintf("%03d%-3s", g.VersionNumber, string(g.PaymentTransactionIdentifier))
}

func ParseG3v036PaymentTransactionIdentifierRequest(s string) *G3v036PaymentTransactionIdentifierRequest {
	g := G3v036PaymentTransactionIdentifierRequest{}

	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.PaymentTransactionIdentifier = PaymentTransactionIdentifierType(s[3:6])

	return &g
}

type RealTimeSubstantiationCodeType uint

const (
	NotVerifiedWithIIAS         = 0
	VerifiedWithIIAS            = 1
	ExemptFromIIAS90PercentRule = 2
)

// Version 037 is only valid for MasterCard
//----------------------------------------------------
// One character field that identifies if the merchant verified the purchased items against an
// Inventory Information Approval System (IIAS)

type G3v037RealTimeSubstantiationRequest struct {
	VersionNumber          G3VersionType                  // 3 NUM 4.85 Group III Version Number 037
	RealTimeSubstantiation RealTimeSubstantiationCodeType // 1 NUM 4.126 Real Time Substantiation
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v037RealTimeSubstantiationRequest(substantiation RealTimeSubstantiationCodeType) *G3v037RealTimeSubstantiationRequest {
	return &G3v037RealTimeSubstantiationRequest{G3v037RealTimeSubstantiation, substantiation}
}

func (g *G3v037RealTimeSubstantiationRequest) String() string {
	return fmt.Sprintf("%03d%d", g.VersionNumber, g.RealTimeSubstantiation)
}

func ParseG3v037RealTimeSubstantiationRequest(s string) *G3v037RealTimeSubstantiationRequest {
	g := G3v037RealTimeSubstantiationRequest{}
	fmt.Sscanf(s, "%03d%1d", &g.VersionNumber, &g.RealTimeSubstantiation)
	return &g
}

// This field contains the digital value of the magnetic signature from the card if it is captured
// when the card is swiped. This data will not be returned in the response, and it must not be
// stored after authorization. This must not be submitted in contactless or chip transactions. (see
// Table 3.72 for record format and version number).
// ----------------------------------------------------
type G3v038ElectroMagneticSignatureRequest struct {
	VersionNumber     G3VersionType // 3 NUM 4.85 Group III Version Number 038
	MagneticSignature string        // 128 ASCII HEX 4.67 Electro Magnetic Signature Magnetic Stripe Card Data - pass through only
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v038ElectroMagneticSignatureRequest(signature string) *G3v038ElectroMagneticSignatureRequest {
	return &G3v038ElectroMagneticSignatureRequest{G3v038ElectroMagneticSignature, signature}
}

func (g *G3v038ElectroMagneticSignatureRequest) String() string {
	return fmt.Sprintf("%03d%-128s", g.VersionNumber, g.MagneticSignature)
}

func ParseG3v038ElectroMagneticSignatureRequest(s string) *G3v038ElectroMagneticSignatureRequest {
	g := G3v038ElectroMagneticSignatureRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	g.MagneticSignature = s[3:131]
	return &g
}

type CardholderVerifictionMethodType rune

const (
	CardholderVerificationByPin       = 'P'
	CardholderVerificationBySignature = 'S'
)

// This field is used by MasterCard in advice requests and reversals. A value of P indicates that a
// PIN was used, and an S indicates a signature or other method. For auto fuel dispensing advice
// messages, the value must be S. (See Table 3.74 for record format and version number).
// ----------------------------------------------------
type G3v039CardholderVerificationMethodRequest struct {
	VersionNumber               G3VersionType                   // 3 NUM 4.85 Group III Version Number 039
	CardholderVerifictionMethod CardholderVerifictionMethodType // 1 A/N 4.44 Cardholder Verification Method
}

func NewG3v039CardholderVerificationMethodRequest(verificationmethod CardholderVerifictionMethodType) *G3v039CardholderVerificationMethodRequest {
	return &G3v039CardholderVerificationMethodRequest{G3v039CardholderVerificationMethod, verificationmethod}
}

func (g *G3v039CardholderVerificationMethodRequest) String() string {
	return fmt.Sprintf("%03d%c", g.VersionNumber, g.CardholderVerifictionMethod)
}

func ParseG3v039CardholderVerificationMethodRequest(s string) *G3v039CardholderVerificationMethodRequest {
	g := G3v039CardholderVerificationMethodRequest{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.CardholderVerifictionMethod)
	return &g
}

// ----------------------------------------------------
type G3v040VisaISAChargeIndicatorRequest struct {
	VersionNumber G3VersionType
}

func NewG3v040VisaISAChargeIndicatorRequest() *G3v040VisaISAChargeIndicatorRequest {
	return &G3v040VisaISAChargeIndicatorRequest{G3v040VisaISAChargeIndicator}
}

func (g *G3v040VisaISAChargeIndicatorRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v040VisaISAChargeIndicatorRequest(s string) *G3v040VisaISAChargeIndicatorRequest {
	g := G3v040VisaISAChargeIndicatorRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v041NtiaUpcSkuDataRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 041
	UpcSku        string        // 1-34 ANS 4.119 UPC/SKU
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v041NtiaUpcSkuDataRequest(upcorsku string) *G3v041NtiaUpcSkuDataRequest {
	return &G3v041NtiaUpcSkuDataRequest{G3v041NtiaUpcSkuData, upcorsku}
}

func (g *G3v041NtiaUpcSkuDataRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.UpcSku, bytes.FileSeparator)
}

func ParseG3v041NtiaUpcSkuDataRequest(s string) *G3v041NtiaUpcSkuDataRequest {
	g := G3v041NtiaUpcSkuDataRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 37 {
		g.UpcSku = s[3:i]
	}

	return &g
}

// ----------------------------------------------------
type G3v042VisaContactlessRequest struct {
	VersionNumber    G3VersionType // 3 NUM 4.85 Group III Version Number 042
	CryptogramAmount uint64        // 0, 12 NUM 4.55 Cryptogram Amount
	// 1 ASCII 4.80 Field Separator <FS>
	ApplicationCryptogram string // 0, 16 AN 4.178 Application Cryptogram
	// 1 ASCII 4.80 Field Separator <FS>
	TransactionApplicationCounter string // 0, 4 AN 4.22 Transaction Application Counter
	// 1 ASCII 4.80 Field Separator <FS>
	CustomerExclusiveData string // 0, 2-64 AN 4.178 Customer Exclusive Data
	// 1 ASCII 4.80 Field Separator <FS>
	FormFactor string // 0, 2-8 AN 4.178 Form Factor
	// 1 ASCII 4.80 Field Separator <FS>
	IssuerApplicationData string // 0, 2-64 AN 4.94 Issuer Application Data
	// 1 ASCII 4.80 Field Separator <FS>
	UnpredictableNumber string // 0, 8 AN 4.173 Unpredictable Number
	// 1 ASCII 4.80 Field Separator <FS>
	CardSequenceNumber uint64 // 0-3 NUM 4.35 Card Sequence Number
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v042VisaContactlessRequest(amount uint64, cryptogram string, counter string, customerdata string, ffactor string, issuerdata string, unpredno string, cardno uint64) *G3v042VisaContactlessRequest {
	return &G3v042VisaContactlessRequest{G3v042VisaContactless, amount, cryptogram, counter, customerdata, ffactor, issuerdata, unpredno, cardno}
}

func (g *G3v042VisaContactlessRequest) String() string {
	return fmt.Sprintf("%03d%012d%c%s%c%s%c%s%c%s%c%s%c%s%c%03d%c", g.VersionNumber, g.CryptogramAmount, bytes.FileSeparator,
		g.ApplicationCryptogram, bytes.FileSeparator,
		g.TransactionApplicationCounter, bytes.FileSeparator,
		g.CustomerExclusiveData, bytes.FileSeparator,
		g.FormFactor, bytes.FileSeparator,
		g.IssuerApplicationData, bytes.FileSeparator,
		g.UnpredictableNumber, bytes.FileSeparator,
		g.CardSequenceNumber, bytes.FileSeparator)
}

func ParseG3v042VisaContactlessRequest(s string) *G3v042VisaContactlessRequest {
	g := G3v042VisaContactlessRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.CryptogramAmount, _ = strconv.ParseUint(fields[j], 10, 64)
			case 1:
				g.ApplicationCryptogram = fields[j]
			case 2:
				g.TransactionApplicationCounter = fields[j]
			case 3:
				g.CustomerExclusiveData = fields[j]
			case 4:
				g.FormFactor = fields[j]
			case 5:
				g.IssuerApplicationData = fields[j]
			case 6:
				g.UnpredictableNumber = fields[j]
			case 7:
				g.CardSequenceNumber, _ = strconv.ParseUint(fields[j], 10, 32)
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

type NetworkIdentificationISOType uint

const (
	NetworkIdVisaNet       NetworkIdentificationISOType = 2  // 0002 V
	NetworkIdInterlink     NetworkIdentificationISOType = 3  // 0003 G
	NetworkIdPlusATM       NetworkIdentificationISOType = 4  // 0004 B
	NetworkIdCirrusATM     NetworkIdentificationISOType = 6  // 0006 O
	NetworkIdMasterCardATM NetworkIdentificationISOType = 7  // 0007 J
	NetworkIdSTAR          NetworkIdentificationISOType = 8  // 0008 N
	NetworkIdPULSE         NetworkIdentificationISOType = 9  // 0009 S
	NetworkIdSTAR_SE       NetworkIdentificationISOType = 10 // 0010 W
	NetworkIdSTAR_NE       NetworkIdentificationISOType = 11 // 0011 Z
	NetworkIdSTAR_W        NetworkIdentificationISOType = 12 // 0012 Q
	NetworkIdAFFN          NetworkIdentificationISOType = 13 // 0013 U
	NetworkIdSTAR_M        NetworkIdentificationISOType = 15 // 0015 M
	NetworkIdMaestro       NetworkIdentificationISOType = 16 // 0016 8
	NetworkIdPulse         NetworkIdentificationISOType = 17 // 0017 L
	NetworkIdNYCE          NetworkIdentificationISOType = 18 // 0018 Y
	NetworkIdPULSE_H       NetworkIdentificationISOType = 19 // 0019 H
	NetworkIdAccel         NetworkIdentificationISOType = 20 // 0020 E
	NetworkIdNETS          NetworkIdentificationISOType = 23 // 0023 P
	NetworkIdCU24          NetworkIdentificationISOType = 24 // 0024 C
	NetworkIdAlaskaOption  NetworkIdentificationISOType = 25 // 0025 3
	NetworkIdNYCE_F        NetworkIdentificationISOType = 27 // 0027 F
	NetworkIdITSShazam     NetworkIdentificationISOType = 28 // 0028 7
	NetworkIdEBT           NetworkIdentificationISOType = 29 // 0029 K
	NetworkIdEBTATM        NetworkIdentificationISOType = 30 // 0030 T
	NetworkIdAmexATM       NetworkIdentificationISOType = 40 // 0040 A
	NetworkIdDiscoverATM   NetworkIdentificationISOType = 41 // 0041 D
	NetworkIdAFFNATM       NetworkIdentificationISOType = 42 // 0042 1
)

func (x NetworkIdentificationISOType) String() string {
	switch x {
	case NetworkIdVisaNet:
		return "NetworkIdVisaNet"
	case NetworkIdInterlink:
		return "NetworkIdInterlink"
	case NetworkIdPlusATM:
		return "NetworkIdPlusATM"
	case NetworkIdCirrusATM:
		return "NetworkIdCirrusATM"
	case NetworkIdMasterCardATM:
		return "NetworkIdMasterCardATM"
	case NetworkIdSTAR:
		return "NetworkIdSTAR"
	case NetworkIdPULSE:
		return "NetworkIdPULSE"
	case NetworkIdSTAR_SE:
		return "NetworkIdSTAR_SE"
	case NetworkIdSTAR_NE:
		return "NetworkIdSTAR_NE"
	case NetworkIdSTAR_W:
		return "NetworkIdSTAR_W"
	case NetworkIdAFFN:
		return "NetworkIdAFFN"
	case NetworkIdSTAR_M:
		return "NetworkIdSTAR_M"
	case NetworkIdMaestro:
		return "NetworkIdMaestro"
	case NetworkIdPulse:
		return "NetworkIdPulse"
	case NetworkIdNYCE:
		return "NetworkIdNYCE"
	case NetworkIdPULSE_H:
		return "NetworkIdPULSE_H"
	case NetworkIdAccel:
		return "NetworkIdAccel"
	case NetworkIdNETS:
		return "NetworkIdNETS"
	case NetworkIdCU24:
		return "NetworkIdCU24"
	case NetworkIdAlaskaOption:
		return "NetworkIdAlaskaOption"
	case NetworkIdNYCE_F:
		return "NetworkIdNYCE_F"
	case NetworkIdITSShazam:
		return "NetworkIdITSShazam"
	case NetworkIdEBT:
		return "NetworkIdEBT"
	case NetworkIdEBTATM:
		return "NetworkIdEBTATM"
	case NetworkIdAmexATM:
		return "NetworkIdAmexATM"
	case NetworkIdDiscoverATM:
		return "NetworkIdDiscoverATM"
	case NetworkIdAFFNATM:
		return "NetworkIdAFFNATM"
	}
	return ""
}

type SharingGroupType byte

const (
	SharingGroupVisaNet       SharingGroupType = 'V' // 0002
	SharingGroupInterlink     SharingGroupType = 'G' // 0003
	SharingGroupPlusATM       SharingGroupType = 'B' // 0004
	SharingGroupCirrusATM     SharingGroupType = 'O' // 0006
	SharingGroupMasterCardATM SharingGroupType = 'J' // 0007
	SharingGroupSTAR          SharingGroupType = 'N' // 0008
	SharingGroupPULSE         SharingGroupType = 'S' // 0009
	SharingGroupSTAR_SE       SharingGroupType = 'W' // 0010
	SharingGroupSTAR_NE       SharingGroupType = 'Z' // 0011
	SharingGroupSTAR_W        SharingGroupType = 'Q' // 0012
	SharingGroupAFFN          SharingGroupType = 'U' // 0013
	SharingGroupSTAR_M        SharingGroupType = 'M' // 0015
	SharingGroupMaestro       SharingGroupType = '8' // 0016
	SharingGroupPulse         SharingGroupType = 'L' // 0017
	SharingGroupNYCE          SharingGroupType = 'Y' // 0018
	SharingGroupPULSE_H       SharingGroupType = 'H' // 0019
	SharingGroupAccel         SharingGroupType = 'E' // 0020
	SharingGroupNETS          SharingGroupType = 'P' // 0023
	SharingGroupCU24          SharingGroupType = 'C' // 0024
	SharingGroupAlaskaOption  SharingGroupType = '3' // 0025
	SharingGroupNYCE_F        SharingGroupType = 'F' // 0027
	SharingGroupITSShazam     SharingGroupType = '7' // 0028
	SharingGroupEBT           SharingGroupType = 'K' // 0029
	SharingGroupEBTATM        SharingGroupType = 'T' // 0030
	SharingGroupAmexATM       SharingGroupType = 'A' // 0040
	SharingGroupDiscoverATM   SharingGroupType = 'D' // 0041
	SharingGroupAFFNATM       SharingGroupType = '1' // 0042
)

func (x SharingGroupType) String() string {
	switch x {
	case SharingGroupVisaNet:
		return "SharingGroupVisaNet"
	case SharingGroupInterlink:
		return "SharingGroupInterlink"
	case SharingGroupPlusATM:
		return "SharingGroupPlusATM"
	case SharingGroupCirrusATM:
		return "SharingGroupCirrusATM"
	case SharingGroupMasterCardATM:
		return "SharingGroupMasterCardATM"
	case SharingGroupSTAR:
		return "SharingGroupSTAR"
	case SharingGroupPULSE:
		return "SharingGroupPULSE"
	case SharingGroupSTAR_SE:
		return "SharingGroupSTAR_SE"
	case SharingGroupSTAR_NE:
		return "SharingGroupSTAR_NE"
	case SharingGroupSTAR_W:
		return "SharingGroupSTAR_W"
	case SharingGroupAFFN:
		return "SharingGroupAFFN"
	case SharingGroupSTAR_M:
		return "SharingGroupSTAR_M"
	case SharingGroupMaestro:
		return "SharingGroupMaestro"
	case SharingGroupPulse:
		return "SharingGroupPulse"
	case SharingGroupNYCE:
		return "SharingGroupNYCE"
	case SharingGroupPULSE_H:
		return "SharingGroupPULSE_H"
	case SharingGroupAccel:
		return "SharingGroupAccel"
	case SharingGroupNETS:
		return "SharingGroupNETS"
	case SharingGroupCU24:
		return "SharingGroupCU24"
	case SharingGroupAlaskaOption:
		return "SharingGroupAlaskaOption"
	case SharingGroupNYCE_F:
		return "SharingGroupNYCE_F"
	case SharingGroupITSShazam:
		return "SharingGroupITSShazam"
	case SharingGroupEBT:
		return "SharingGroupEBT"
	case SharingGroupEBTATM:
		return "SharingGroupEBTATM"
	case SharingGroupAmexATM:
		return "SharingGroupAmexATM"
	case SharingGroupDiscoverATM:
		return "SharingGroupDiscoverATM"
	case SharingGroupAFFNATM:
		return "SharingGroupAFFNATM"
	}
	return ""
}

// ----------------------------------------------------
type G3v043NetworkIDRequest struct {
	VersionNumber G3VersionType                // 3 NUM 4.85 Group III Version Number 043
	NetworkID     NetworkIdentificationISOType // 4 NUM 4.117 Network ID Must not be “0000”
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v043NetworkIDRequest(netid NetworkIdentificationISOType) *G3v043NetworkIDRequest {
	return &G3v043NetworkIDRequest{G3v043NetworkID, netid}
}

func (g *G3v043NetworkIDRequest) String() string {
	return fmt.Sprintf("%03d%04d", g.VersionNumber, g.NetworkID)
}

func ParseG3v043NetworkIDRequest(s string) *G3v043NetworkIDRequest {
	g := G3v043NetworkIDRequest{}
	fmt.Sscanf(s, "%03d%04d", &g.VersionNumber, &g.NetworkID)
	return &g
}

// ----------------------------------------------------
type G3v044AutomatedTellerMachineRequest struct {
	VersionNumber G3VersionType               // 3 NUM 4.85 Group III Version Number 044
	FromAccount   AdditionalAmountAccountType // 0 or 2 NUM 4.2 Account Type (from)
	// 1 ASCII 4.80 Field Separator <FS>
	ToAccount AdditionalAmountAccountType // 0 or 2 NUM 4.3 Account Type (to)
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v044AutomatedTellerMachineRequest(from AdditionalAmountAccountType, to AdditionalAmountAccountType) *G3v044AutomatedTellerMachineRequest {
	return &G3v044AutomatedTellerMachineRequest{G3v044AutomatedTellerMachine, from, to}
}

func (g *G3v044AutomatedTellerMachineRequest) String() string {
	return fmt.Sprintf("%03d%02d%c%02d%c", g.VersionNumber, g.FromAccount, bytes.FileSeparator, g.ToAccount, bytes.FileSeparator)
}

func ParseG3v044AutomatedTellerMachineRequest(s string) *G3v044AutomatedTellerMachineRequest {
	g := G3v044AutomatedTellerMachineRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				fmt.Sscanf(fields[j], "%02d", &g.FromAccount)
			case 1:
				fmt.Sscanf(fields[j], "%02d", &g.ToAccount)
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v045IntegratedChipCardRequest struct {
	VersionNumber    G3VersionType // 3 NUM 4.85 Group III Version Number 045
	CryptogramAmount uint64        // 12 NUM 4.55 Cryptogram Amount Tag 9F02
	// 1 ASCII 4.80 Field Separator <FS>
	AuthorizationRequestCryptogram string // 0,16 AN 4.25 Authorization Request Cryptogram (ARQC) Tag 9F26
	// 1 ASCII 4.80 Field Separator <FS>
	ApplicationCounter string // 4 AN 4.22 Transaction Application Counter Tag 9F36
	// 1 ASCII 4.80 Field Separator <FS>
	CustomerData string // 0-64 AN 4.178 Customer Exclusive Data Visa USA only. Tag 9F7C
	// 1 ASCII 4.80 Field Separator <FS>
	FormFactor string // 0-10 AN 4.178 Form Factor Tag 9F6E
	// 1 ASCII 4.80 Field Separator <FS>
	UnpredictableNumber string // 8 AN 4.173 Unpredictable Number Tag 9F37
	// 1 ASCII 4.80 Field Separator <FS>
	CardSequenceNumber uint64 // 0-3 NUM 4.35 Card Sequence Number Send up to 3 decimal digits. Tag 5F34

	// 1 ASCII 4.80 Field Separator <FS>
	CryptogramInformationData uint // 0,2 NUM 4.58 Cryptogram Information Data Tag 9F27
	// 1 ASCII 4.80 Field Separator <FS>
	IssuerApplicationData string // 0-64 AN 4.94 Issuer Application Data Tag 9F10
	// 1 ASCII 4.80 Field Separator <FS>
	TerminalCountryCode uint // 3 NUM 4.150 Terminal Country Code Tag 9F1A
	// 1 ASCII 4.80 Field Separator <FS>
	IFDSerialNumber string // 0,8 ASCII 4.90 IFD Serial Number Tag 9F1E
	// 1 ASCII 4.80 Field Separator <FS>
	TerminalCapabilityProfile string // 0,6 AN 4.149 Terminal Capability Profile Tag 9F33
	// 1 ASCII 4.80 Field Separator <FS>
	IssuerScriptResults string // 0-40 AN 4.97 Issuer Script Results Tag 9F5B
	// 1 ASCII 4.80 Field Separator <FS>
	ApplicationInterchangeProfile uint // 0,4 NUM 4.21 Application Interchange Profile Tag 82
	// 1 ASCII 4.80 Field Separator <FS>

	TerminalVerificationResults uint64 // 0,10 NUM 4.155 Terminal Verification Results Tag 95
	// 1 ASCII 4.80 Field Separator <FS>
	CryptogramTransactionType uint // 0,2 NUM 4.59 Cryptogram Transaction Type Tag 9C
	// 1 ASCII 4.80 Field Separator <FS>
	TerminalTransactionTime string // 6 NUM 4.153 Terminal Transaction Time Tag 9F21
	// 1 ASCII 4.80 Field Separator <FS>
	TerminalTransactionDate string // 6 NUM 4.151 Terminal Transaction Date YYMMDD. Tag 9A
	// 1 ASCII 4.80 Field Separator <FS>
	CryptogramCurrencyCode uint // 3 NUM 4.57 Cryptogram Currency Code Tag 5F2A
	// 1 ASCII 4.80 Field Separator <FS>
	CryptogramCashbackAmount uint64 // 0,12 NUM 4.56 Cryptogram Cashback Amount Required when cryptogram amount includes cashback amount. Tag 9F03
	// 1 ASCII 4.80 Field Separator <FS>

	CardholderVerificationMethodResults string // 0,6 AN Cardholder Verification Method Results Tag 9F34
	// 1 ASCII 4.80 Field Separator <FS>
	TerminalType string // 0,2 AN 4.154 Terminal Type Tag 9F35
	// 1 ASCII 4.80 Field Separator <FS>
	TransactionCategoryCode string // 0,1 AN 4.162 Transaction Category Code MC usage only. Tag 9F53
	// 1 ASCII 4.80 Field Separator <FS>
	SecondaryPinBlock string // 0,16 AN 4.139 Secondary PIN Block Visa usage only. Tag C0. Only used when terminal is changing Encrypted PIN to be loaded into card.
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>

}

func NewG3v045IntegratedChipCardRequest() *G3v045IntegratedChipCardRequest {
	return &G3v045IntegratedChipCardRequest{VersionNumber: G3v045IntegratedChipCard}
}

func (g *G3v045IntegratedChipCardRequest) String() string {
	return fmt.Sprintf("%03d%012d%c%s%c%s%c%s%c%s%c%s%c%d%c%d%c%s%c%d%c%s%c%s%c%s%c%d%c%010d%c%02d%c%-06s%c%-06s%c%03d%c%012d%c%s%c%s%c%s%c%s%c", g.VersionNumber,
		g.CryptogramAmount, bytes.FileSeparator, //Tag 9F02
		g.AuthorizationRequestCryptogram, bytes.FileSeparator, //Tag 9F26
		g.ApplicationCounter, bytes.FileSeparator, //Tag 9F36
		g.CustomerData, bytes.FileSeparator, //Tag 9F7C
		g.FormFactor, bytes.FileSeparator, //Tag 9F6E
		g.UnpredictableNumber, bytes.FileSeparator, //Tag 9F37
		g.CardSequenceNumber, bytes.FileSeparator, //Tag 5F34
		g.CryptogramInformationData, bytes.FileSeparator, //Tag 9F27
		g.IssuerApplicationData, bytes.FileSeparator, //Tag 9F10
		g.TerminalCountryCode, bytes.FileSeparator, //Tag 9F1A
		g.IFDSerialNumber, bytes.FileSeparator, //Tag 9F1E
		g.TerminalCapabilityProfile, bytes.FileSeparator, //Tag 9F33
		g.IssuerScriptResults, bytes.FileSeparator, //Tag 9F5B
		g.ApplicationInterchangeProfile, bytes.FileSeparator, //Tag 82
		g.TerminalVerificationResults, bytes.FileSeparator, //Tag 95
		g.CryptogramTransactionType, bytes.FileSeparator, //Tag 9C
		g.TerminalTransactionTime, bytes.FileSeparator, //Tag 9F21
		g.TerminalTransactionDate, bytes.FileSeparator, //Tag 9A
		g.CryptogramCurrencyCode, bytes.FileSeparator, //Tag 5F2A
		g.CryptogramCashbackAmount, bytes.FileSeparator, //Tag 9F03
		g.CardholderVerificationMethodResults, bytes.FileSeparator, //Tag 9F34
		g.TerminalType, bytes.FileSeparator, //Tag 9F35
		g.TransactionCategoryCode, bytes.FileSeparator, //Tag 9F53
		g.SecondaryPinBlock, bytes.FileSeparator) // Tag C0
}

func ParseG3v045IntegratedChipCardRequest(s string) *G3v045IntegratedChipCardRequest {
	g := G3v045IntegratedChipCardRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	//strconv.ParseUint(fields[j], 10, 64)
	//fmt.Sscanf(fields[j], "%03d", &g.xx)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				fmt.Sscanf(fields[j], "%012d", &g.CryptogramAmount) // 12 NUM Tag 9F02
			case 1:
				g.AuthorizationRequestCryptogram = fields[j] //Tag 9F26
			case 2:
				g.ApplicationCounter = fields[j] //Tag 9F36
			case 3:
				g.CustomerData = fields[j] //Tag 9F7C
			case 4:
				g.FormFactor = fields[j] //Tag 9F6E
			case 5:
				g.UnpredictableNumber = fields[j] //Tag 9F37
			case 6:
				g.CardSequenceNumber, _ = strconv.ParseUint(fields[j], 10, 64) // 0-3 NUM Tag 5F34
			case 7:
				fmt.Sscanf(fields[j], "%02d", &g.CryptogramInformationData) // 0,2 NUM Tag 9F27
			case 8:
				g.IssuerApplicationData = fields[j] //Tag 9F10
			case 9:
				fmt.Sscanf(fields[j], "%03d", &g.TerminalCountryCode) // 3 NUM Tag 9F1A
			case 10:
				g.IFDSerialNumber = fields[j] //Tag 9F1E
			case 11:
				g.TerminalCapabilityProfile = fields[j] //Tag 9F33
			case 12:
				g.IssuerScriptResults = fields[j] //Tag 9F5B
			case 13:
				fmt.Sscanf(fields[j], "%04d", &g.ApplicationInterchangeProfile) // 0,4 NUM Tag 82
			case 14:
				fmt.Sscanf(fields[j], "%010d", &g.TerminalVerificationResults) // 0,10 NUM
			case 15:
				fmt.Sscanf(fields[j], "%02d", &g.CryptogramTransactionType) // 0,2 NUM Tag 9C
			case 16:
				g.TerminalTransactionTime = fields[j] // 6 NUM Tag 9F21
			case 17:
				g.TerminalTransactionDate = fields[j] //Tag 9A
			case 18:
				fmt.Sscanf(fields[j], "%03d", &g.CryptogramCurrencyCode) // 3 NUM Tag 5F2A
			case 19:
				fmt.Sscanf(fields[j], "%012d", &g.CryptogramCashbackAmount) //0,12 NUM Tag 9F03
			case 20:
				g.CardholderVerificationMethodResults = fields[j] //Tag 9F34
			case 21:
				g.TerminalType = fields[j] //Tag 9F35
			case 22:
				g.TransactionCategoryCode = fields[j] //Tag 9F53
			case 23:
				g.SecondaryPinBlock = fields[j] // Tag C0

			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v046CardTypeGroupRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v046CardTypeGroupRequest() *G3v046CardTypeGroupRequest {
	return &G3v046CardTypeGroupRequest{G3v046CardTypeGroup}
}

func (g *G3v046CardTypeGroupRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v046CardTypeGroupRequest(s string) *G3v046CardTypeGroupRequest {
	g := G3v046CardTypeGroupRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

//----------------------------------------------------
// type G3v047TSYSInternalUseOnlyRequest struct {
// 	VersionNumber G3VersionType //3 NUM
// }

// func NewG3v047TSYSInternalUseOnlyRequest() *G3v047TSYSInternalUseOnlyRequest {
// 	return &G3v047TSYSInternalUseOnlyRequest{G3v047TSYSInternalUseOnly, ""}
// }

// func (g *G3v047TSYSInternalUseOnlyRequest) String() string {
// 	return fmt.Sprintf("%03d%-4s", g.VersionNumber, g.xxx)
// }

// ----------------------------------------------------
type G3v048AmexCardholderVerificationResultsRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v048AmexCardholderVerificationResultsRequest() *G3v048AmexCardholderVerificationResultsRequest {
	return &G3v048AmexCardholderVerificationResultsRequest{G3v048AmexCardholderVerificationResults}
}

func (g *G3v048AmexCardholderVerificationResultsRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v048AmexCardholderVerificationResultsRequest(s string) *G3v048AmexCardholderVerificationResultsRequest {
	g := G3v048AmexCardholderVerificationResultsRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v049Gen2TerminalAuthenticationRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 049
	GenKey        string        // 0 or 24 ASCII 4.84.4 GenKey ASCII
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v049Gen2TerminalAuthenticationRequest(key string) *G3v049Gen2TerminalAuthenticationRequest {
	return &G3v049Gen2TerminalAuthenticationRequest{G3v049Gen2TerminalAuthentication, key}
}

func (g *G3v049Gen2TerminalAuthenticationRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.GenKey, bytes.FileSeparator)
}

func ParseG3v049Gen2TerminalAuthenticationRequest(s string) *G3v049Gen2TerminalAuthenticationRequest {
	g := G3v049Gen2TerminalAuthenticationRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i && i < 27 {
		g.GenKey = s[3:i]
	}
	return &g
}

// ----------------------------------------------------
type G3v050AssociationTimestampRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v050AssociationTimestampRequest() *G3v050AssociationTimestampRequest {
	return &G3v050AssociationTimestampRequest{G3v050AssociationTimestamp}
}

func (g *G3v050AssociationTimestampRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v050AssociationTimestampRequest(s string) *G3v050AssociationTimestampRequest {
	g := G3v050AssociationTimestampRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v051MasterCardEventMonitoringRealTimeScoringRequest struct {
	VersionNumber           G3VersionType // 3 NUM 4.85 Group III Version Number 051
	ServiceRequestIndicator uint          // 1 NUM 4.74 EMS Service Request Indicator
	// 	Valid values:
	// 	0 - No action required
	// 	1 - Transaction to be scored
	// 	MasterCard usage only
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v051MasterCardEventMonitoringRealTimeScoringRequest(requestind uint) *G3v051MasterCardEventMonitoringRealTimeScoringRequest {
	return &G3v051MasterCardEventMonitoringRealTimeScoringRequest{G3v051MasterCardEventMonitoringRealTimeScoring, requestind}
}

func (g *G3v051MasterCardEventMonitoringRealTimeScoringRequest) String() string {
	return fmt.Sprintf("%03d%d", g.VersionNumber, g.ServiceRequestIndicator)
}

func ParseG3v051MasterCardEventMonitoringRealTimeScoringRequest(s string) *G3v051MasterCardEventMonitoringRealTimeScoringRequest {
	g := G3v051MasterCardEventMonitoringRealTimeScoringRequest{}
	fmt.Sscanf(s, "%03d%1d", &g.VersionNumber, &g.ServiceRequestIndicator)
	return &g
}

type EncryptionMethodType rune

const (
	VoltageTEP2Encryption = 'V'
)

// ----------------------------------------------------
type G3v052VoltageEncryptionTransmissionBlockRequest struct {
	VersionNumber               G3VersionType        // 5 NUM 4.85 Group III Version Number 052
	EncryptionType              EncryptionMethodType // 1 A/N 4.76 Encryption Type "V" (Voltage TEP2 Encryption)
	EncryptionTransmissionBlock string               // 250-300 A/N, "+", "/", "=" 4.75 Encryption Transmission Block (ETB) Base-64 encoded
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v052VoltageEncryptionTransmissionBlockRequest(transblock string) *G3v052VoltageEncryptionTransmissionBlockRequest {
	return &G3v052VoltageEncryptionTransmissionBlockRequest{G3v052VoltageEncryptionTransmissionBlock, VoltageTEP2Encryption, transblock}
}

func (g *G3v052VoltageEncryptionTransmissionBlockRequest) String() string {
	return fmt.Sprintf("%03d%c%s%c", g.VersionNumber, g.EncryptionType, g.EncryptionTransmissionBlock, bytes.FileSeparator)
}

func ParseG3v052VoltageEncryptionTransmissionBlockRequest(s string) *G3v052VoltageEncryptionTransmissionBlockRequest {
	g := G3v052VoltageEncryptionTransmissionBlockRequest{}
	fmt.Sscanf(s, "%03d%c", &g.VersionNumber, &g.EncryptionType)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 4 < i && i < 304 {
		g.EncryptionTransmissionBlock = s[4:i]
	}

	return &g
}

// ----------------------------------------------------
type G3v053TokenRequest struct {
	VersionNumber G3VersionType //3 NUM 4.85 Group III Version Number 053
}

func NewG3v053TokenRequest() *G3v053TokenRequest {
	return &G3v053TokenRequest{G3v053Token}
}

func (g *G3v053TokenRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v053TokenRequest(s string) *G3v053TokenRequest {
	g := G3v053TokenRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type TransitTransactionType uint

const (
	PrefundedTransit                         TransitTransactionType = 1
	RealTimeAuthorizedTransit                                       = 2
	PostAuthorizedAggregatedTransit                                 = 3
	AuthorizedAggregatedSplitClearingTransit                        = 4
	OtherTransit                                                    = 5
	DebtRecoveryTransit                                             = 7
)

type TransportationModeType uint

const (
	Unknown               TransportationModeType = 0
	UrbanBus                                     = 1
	InterUrbanBus                                = 2
	LightTrainMassTransit                        = 3
	Train                                        = 4
	CommuterTrain                                = 5
	WaterBorneVehicle                            = 6
	Toll                                         = 7
	Parking                                      = 8
	Taxi                                         = 9
	HighSpeedTrain                               = 10
	RuralBus                                     = 11
	ExpressCummuterTrain                         = 12
	ParaTransit                                  = 13
	SelfDriveVehicle                             = 14
	Coach                                        = 15
	Locomotive                                   = 16
	PoweredMotorVehicle                          = 17
	Trailer                                      = 18
	RegionalTrain                                = 19
	InterCity                                    = 20
	FunicularTrailer                             = 21
	CableCar                                     = 22

// 23-99 Reserved for future use
)

// ----------------------------------------------------
type G3v054TransitProgramRequest struct {
	VersionNumber                   G3VersionType          // 3 NUM 4.85 Group III Version Number 054
	TransitTransactionTypeIndicator TransitTransactionType // 2 NUM 4.168 Transit Transaction Type Indicator Valid value
	TransportationModeIndicator     TransportationModeType // 2 NUM 4.169 Transportation Mode Indicator Valid value
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v054TransitProgramRequest(transit TransitTransactionType, mode TransportationModeType) *G3v054TransitProgramRequest {
	return &G3v054TransitProgramRequest{G3v054TransitProgram, transit, mode}
}

func (g *G3v054TransitProgramRequest) String() string {
	return fmt.Sprintf("%03d%02d%02d", g.VersionNumber, g.TransitTransactionTypeIndicator, g.TransportationModeIndicator)
}

func ParseG3v054TransitProgramRequest(s string) *G3v054TransitProgramRequest {
	g := G3v054TransitProgramRequest{}
	fmt.Sscanf(s, "%03d%02d%02d", &g.VersionNumber, &g.TransitTransactionTypeIndicator, &g.TransportationModeIndicator)
	return &g
}

// ----------------------------------------------------
type G3v055IntegratedChipCardEmvTlvRequest struct {
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

func NewG3v055IntegratedChipCardEmvTlvRequest(tlv string) *G3v055IntegratedChipCardEmvTlvRequest {
	return &G3v055IntegratedChipCardEmvTlvRequest{G3v055IntegratedChipCardEmvTlv, tlv}
}

func (g *G3v055IntegratedChipCardEmvTlvRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.EmvTlvDatum, bytes.FileSeparator)
}

func ParseG3v055IntegratedChipCardEmvTlvRequest(s string) *G3v055IntegratedChipCardEmvTlvRequest {
	g := G3v055IntegratedChipCardEmvTlvRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := strings.Index(s, string(bytes.FileSeparator))
	if 3 < i {
		g.EmvTlvDatum = s[3:i]
	}

	return &g
}

type MessageReasonCodeType uint

const (
	//Reversals
	ReversalVoidedByCustomer                    MessageReasonCodeType = 2501
	TransactionNotCompletedTimeoutOrMalfunction MessageReasonCodeType = 2502
	NoPOSConfirmation                           MessageReasonCodeType = 2503
	POSPartialReversal                          MessageReasonCodeType = 2504
	PrematureChipCardRemoval                    MessageReasonCodeType = 2516 // (after online request sent before response received)
	ChipDeclinedTransaction                     MessageReasonCodeType = 2517 // after online issuer approved
	//Adjustments
	AdjustmentVoidedByCustomer                    MessageReasonCodeType = 2001
	WrongAmount                                   MessageReasonCodeType = 2002
	PartialReturn                                 MessageReasonCodeType = 2003
	DebitOrCreditAdjustment                       MessageReasonCodeType = 2007 //no previous transaction *This message reason is valid only through Networks 2 (Visa) and 3 (Interlink)
	DebitOrCreditAdjustmentForPreviousTransaction MessageReasonCodeType = 2009 //*Adjustments with this reason code must also include field 90 - Original transaction information
	AcquirerAuthorizationAdvice                   MessageReasonCodeType = 2104
	AccountFundingTransaction                     MessageReasonCodeType = 2140 // debit or credit adjustment
	EBTVoucher                                    MessageReasonCodeType = 5201
)

// ----------------------------------------------------
type G3v056MessageReasonCodeRequest struct {
	VersionNumber     G3VersionType         // 3 NUM 4.85 Group III Version Number 056
	MessageReasonCode MessageReasonCodeType // 4 NUM 4.113 Message Reason Code
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v056MessageReasonCodeRequest(reasoncode MessageReasonCodeType) *G3v056MessageReasonCodeRequest {
	return &G3v056MessageReasonCodeRequest{G3v056MessageReasonCode, reasoncode}
}

func (g *G3v056MessageReasonCodeRequest) String() string {
	return fmt.Sprintf("%03d%04d", g.VersionNumber, g.MessageReasonCode)
}

func ParseG3v056MessageReasonCodeRequest(s string) *G3v056MessageReasonCodeRequest {
	g := G3v056MessageReasonCodeRequest{}
	fmt.Sscanf(s, "%03d%04d", &g.VersionNumber, &g.MessageReasonCode)
	return &g
}

// ----------------------------------------------------
type G3v057DiscoverOrPayPalAdditionalDataRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 057
	// This field contains additional data from a Discover/PayPal AFD authorization response
	// message that must be used in a Discover/PayPal AFD completion advice message.
	AdditionalResponseData string // 0-25 A/N/S 4.10 Additional Response Data Additional data for Discover/ PayPal AFD messages
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v057DiscoverOrPayPalAdditionalDataRequest(data string) *G3v057DiscoverOrPayPalAdditionalDataRequest {
	return &G3v057DiscoverOrPayPalAdditionalDataRequest{G3v057DiscoverOrPayPalAdditionalData, data}
}

func (g *G3v057DiscoverOrPayPalAdditionalDataRequest) String() string {
	return fmt.Sprintf("%03d%s%c", g.VersionNumber, g.AdditionalResponseData, bytes.FileSeparator)
}

func ParseG3v057DiscoverOrPayPalAdditionalDataRequest(s string) *G3v057DiscoverOrPayPalAdditionalDataRequest {
	g := G3v057DiscoverOrPayPalAdditionalDataRequest{}
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
type G3v058AlternateAccountIDRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 058
}

func NewG3v058AlternateAccountIDRequest() *G3v058AlternateAccountIDRequest {
	return &G3v058AlternateAccountIDRequest{G3v058AlternateAccountID}
}

func (g *G3v058AlternateAccountIDRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v058AlternateAccountIDRequest(s string) *G3v058AlternateAccountIDRequest {
	g := G3v058AlternateAccountIDRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// It is recommended that Group 3, version 59 be accompanied by Group 3, version 58 for
// MasterCard transactions.
// ----------------------------------------------------
type G3v059MasterCardPayPassMappingServiceRequest struct {
	VersionNumber G3VersionType //3 NUM
}

func NewG3v059MasterCardPayPassMappingServiceRequest() *G3v059MasterCardPayPassMappingServiceRequest {
	return &G3v059MasterCardPayPassMappingServiceRequest{G3v059MasterCardPayPassMappingService}
}

func (g *G3v059MasterCardPayPassMappingServiceRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v059MasterCardPayPassMappingServiceRequest(s string) *G3v059MasterCardPayPassMappingServiceRequest {
	g := G3v059MasterCardPayPassMappingServiceRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

type MobileDeviceType uint

const (
	CardReadyMobileDevice                         MobileDeviceType = 0
	MobileNetworkOperatorSecureElement            MobileDeviceType = 1 //(MNO) controlledremovable secure element (SIM or UICC) personalized for use with a Mobile Phone or Smartphone
	KeyFob                                        MobileDeviceType = 2
	WatchReadyMobileDevice                        MobileDeviceType = 3
	MobileTagMobileDevice                         MobileDeviceType = 4
	WristbandMobileDevice                         MobileDeviceType = 5
	MobilePhoneCaseofSleeve                       MobileDeviceType = 6
	MobilePhoneOrSmartphoneWithFixedSecureElement MobileDeviceType = 7  //(nonremovable) secure element controlled by the MNO. For example, code division multiple access CDMA.
	RemovableSecureElement                        MobileDeviceType = 8  //not controlled by the MNO. For example, SD Card personalized for use with a Mobile Phone or Smartphone.
	MobilePhoneOrSmartphoneUncontrolled           MobileDeviceType = 9  //with a fixed (nonremovable) secure element not controlled by the MNO.
	TabletOrEbookWithRemovableSecureElementSIM    MobileDeviceType = 10 //MNO controlled removable secure element (SIM or UICC) personalized for use with a Tablet or EBook.
	TabletOrEbookWithNonRemovableSecureElement    MobileDeviceType = 11 // (non-removable) secure element controlled by the MNO.
	TabletOrEbookWithRemovableSecureElement       MobileDeviceType = 12 // Removable secure element not controlled by the MNO. For example, SD Card personalized for use with a Tablet or E-Book.
	TabletOrEbookWithFixedSecureElement           MobileDeviceType = 13 //Tablet or E-Book with a fixed (non-removable) // secure element not controlled by the MNO.
)

// PayPass Mobile (G3v060) - This group should be sent on all MasterCard PayPass transactions.
// ----------------------------------------------------
type G3v060PayPassMobileRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 060
	DomainServer  uint64        // 0, 1 A/N 4.66 Domain Server This optional one-character field identifies the Remote Payments Domain Server field
	//associated with the Mobile Device Type.
	//      0 No domain
	//      1 Issuer domain
	//      2 Acquirer domain
	// 1 ASCII 4.80 Field Separator <FS>
	MobileDevice MobileDeviceType // 0, 2 A/N 4.115 Mobile Device Type
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v060PayPassMobileRequest(domain uint64, device MobileDeviceType) *G3v060PayPassMobileRequest {
	return &G3v060PayPassMobileRequest{G3v060PayPassMobile, domain, device}
}

func (g *G3v060PayPassMobileRequest) String() string {
	return fmt.Sprintf("%03d%d%c%d%c", g.VersionNumber, g.DomainServer, bytes.FileSeparator, g.MobileDevice, bytes.FileSeparator)
}

func ParseG3v060PayPassMobileRequest(s string) *G3v060PayPassMobileRequest {
	g := G3v060PayPassMobileRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	i := 0
	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	for j := 0; j < len(fields); j++ {
		if 0 < len(fields[j]) {
			switch j {
			case 0:
				g.DomainServer, _ = strconv.ParseUint(fields[j], 10, 32)
			case 1:
				fmt.Sscanf(fields[j], "%02d", &g.MobileDevice)
			}
		}
		i += len(fields[j]) + 1
	}

	return &g
}

// ----------------------------------------------------
type G3v061SpendQualifiedIndicatorRequest struct {
	//This group should be sent on all Visa transactions
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 061
}

func NewG3v061SpendQualifiedIndicatorRequest() *G3v061SpendQualifiedIndicatorRequest {
	return &G3v061SpendQualifiedIndicatorRequest{G3v061SpendQualifiedIndicator}
}

func (g *G3v061SpendQualifiedIndicatorRequest) String() string {
	return fmt.Sprintf("%03d", g.VersionNumber)
}

func ParseG3v061SpendQualifiedIndicatorRequest(s string) *G3v061SpendQualifiedIndicatorRequest {
	g := G3v061SpendQualifiedIndicatorRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)
	return &g
}

// ----------------------------------------------------
type G3v200GiftCardRequest struct {
	VersionNumber G3VersionType // 3 NUM 4.85 Group III Version Number 200
	OperatorID    string        // 0-8 NUM 4.120 Operator / Clerk ID
	// 1 ASCII 4.80 Field Separator <FS>
	BarCodeFormat string // 1 A/N 4.31 Bar Code Format (This field is only referenced if the account data source is an “A”)
	// 1 ASCII 4.80 Field Separator <FS>
	// 1 ASCII 4.86 Group Separator <GS>
}

func NewG3v200GiftCardRequest(operator string, barcode string) *G3v200GiftCardRequest {
	return &G3v200GiftCardRequest{G3v200GiftCard, operator, barcode}
}

func (g *G3v200GiftCardRequest) String() string {
	return fmt.Sprintf("%03d%s%c%s%c", g.VersionNumber, g.OperatorID, bytes.FileSeparator, g.BarCodeFormat, bytes.FileSeparator)
}

func ParseG3v200GiftCardRequest(s string) *G3v200GiftCardRequest {
	g := G3v200GiftCardRequest{}
	fmt.Sscanf(s, "%03d", &g.VersionNumber)

	fields := strings.Split(s[3:], string(bytes.FileSeparator))
	g.OperatorID = fields[0][3:]
	g.BarCodeFormat = fields[1]

	return &g
}

//----------------------------------------------------
