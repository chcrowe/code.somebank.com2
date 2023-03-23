package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

// all code based upon TSYS technical reference data found within:
// EXTERNAL INTERFACE SPECIFICATIONS DATA
// CAPTURE RECORD FORMATS
// EIS 1081
// VERSION 13.2
// APRIL 21, 2015

type BatchDetailOptionalDataRecord struct {
	GroupMap *OptionalDataGroupMap //5-9 5 A/N Record Type H@@@@ 4.243

	// Table 3.6 Transaction detail record - optional data groups
	// 1 Cash Back
	CashbackAmount uint64 // 	12 NUM 4.63 Cash Back Amount

	// 2 Restaurant
	GratuityAmount uint64 // 	12 NUM 4.137 Gratuity Amount

	// 3 Direct Marketing
	TotalAuthorizedAmount        uint64               // 	12 NUM 4.306 Total Authorized Amount
	PurchaseIdentifierFormatCode PurchaseIdFormatType // 	1 A/N 4.236 Purchase Identifier Format Code
	PurchaseIdentifier           string               // 	25 A/N 4.235 Purchase Identifier

	// 4 Auto Rental
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	MarketSpecificDataIdentifier MarketSpecificDataType // 	1 A/N 4.181 Market Specific Data Identifier
	NoShowIndicator              NoShowIndicatorType    // 	1 A/N 4.215 No Show Indicator
	ExtraCharges                 string                 // 	6 A/N 4.121.1 Extra Charges (AutoRentalExtraChargeType & HotelExtraChargeType)
	RentalOrCheckinDate          uint64                 // 	6 NUM 4.246 Rental Date (YYMMDD)

	// 5 Hotel
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	// 	1 A/N 4.181 Market Specific Data Identifier
	// 	1 A/N 4.215 4.278 No Show or Special Program Ind.
	// 	6 A/N 4.121 Extra Charges
	// 	6 NUM 4.72 Check-in Date (YYMMDD)

	// 6 Hotel - Not Front Desk (Non-Lodging Purchase)
	MerchantCategoryCodeOverride MerchantCategoryType // 	4 NUM 4.185 Merchant Category Code Override

	// 7 Hotel - American Express
	HotelExtraChargeCode HotelExtraChargeType // 	1 A/N 4.71 Charge Type
	CheckoutDate         uint64               // 	6 NUM 4.73 Checkout Date (YYMMDD)
	StayDuration         uint                 // 	2 NUM 4.281 Stay Duration
	RoomRate             uint64               // 	12 NUM 4.180 Lodging Room Rate

	// 8 Non-T&E Commercial Card Level II (Visa and MasterCard Only)
	OptionalAmountIdentifier LocalTaxIndicatorType // 	1 A/N 4.221 Optional Amount Identifier
	OptionalAmount           uint64                // 	12 NUM 4.220 Optional Amount
	PurchaseOrderNumber      string                // 	16 A/N 4.237 Purchase Order Number

	// 9 Passenger Transport (Visa Only)
	// 	1 A/N 4.254 Restricted Ticket Indicator
	// 	2 NUM 4.207 Multiple Clearing Sequence Number
	// 	2 NUM 4.206 Multiple Clearing Sequence Count
	// 	13 A/N 4.304 Ticket Number
	// 	20 A/N 4.228 Passenger Name
	// 	6 NUM 4.93 Departure Date (MMDDYY)
	// 	3 A/N 4.224 Origination City/Airport Code
	// 	7 A/N 4.323 Trip Leg 1 Information
	// 	7 A/N 4.323 Trip Leg 2 Information
	// 	7 A/N 4.323 Trip Leg 3 Information
	// 	7 A/N 4.323 Trip Leg 4 Information

	// 10 Direct Debit
	RetrievalReferenceNumber  string                    // 	12 A/N 4.259 Retrieval Reference Number
	SystemTraceAuditNumber    uint64                    // 	6 NUM 4.288 System Trace Audit Number
	NetworkIdentificationCode NetworkIdentificationType // 	1 A/N 4.213 Network Identification Code
	SettlementDate            uint                      // 	4 NUM 4.270 Settlement Date (MMDD)

	// 	11 Multi-Terminal Stores
	TerminalNumberOverride uint // 	4 NUM 4.301 Terminal Number Override

	// 12 Direct Marketing - Installment or Recurring Payments
	MultipleClearingSequenceNumber uint                  // 	2 NUM 4.207 Multiple Clearing Sequence Number
	MultipleClearingSequenceCount  uint                  // 	2 NUM 4.206 Multiple Clearing Sequence Count
	MotoEcomIndicator              MotoEcomIndicatorType // 	1 A/N 4.205 MOTO/e-Commerce Indicator

	// 13 Auto Rental
	// 	20 A/N 4.250 Renter Name
	// 	18 A/N 4.247 Rental Return City
	// 	3 A/N 4.249 Rental Return State/Country
	// 	10 A/N 4.248 Rental Return Location ID

	// 14 Service Development Indicator
	ServiceDevelopmentIndicator ServiceDevelopmentType // 	1 NUM 4.267 Service Development Indicator

	// 16 Existing Debt Indicator (Visa Only)
	ExistingDebtIndicator ExistingDebtIndicatorType // 	1 NUM 4.119 Existing Debt Indicator

	// 17 Universal Cardholder Authentication
	UcafCollectionIndicator UcafCollectionType // 	1 A/N 4.324 UCAF Collection Indicator

	// 18 e-Commerce Goods Indicator
	EcommerceGoodsIndicator EcommerceGoodsIndicatorType // 	2 A/N 4.108 e-Commerce Goods Indicator

	// 19 CCS Private Label Restricted for CCS use only (Mandatory for Private Label transactions)
	// 	5 NUM 4.80 Credit Plan Number (Not Null)
	// 	4 A/N 4.233 Product Code 1
	// 	4 A/N 4.233 Product Code 2
	// 	4 A/N 4.233 Product Code 3
	// 	4 A/N 4.233 Product Code 4
	// 	4 A/N 4.233 Product Code 5
	// 	3 NUM 4.223 Org. ID (Not Null)
	// 	9 NUM 4.283 Store Code (Not Null)
	// 	6 NUM 4.241 Receiving Institution ID (“911111”)

	// 20 Merchant Verification Value
	MerchantVerificationValue string // 	10 A/N 4.195 Merchant Verification Value

	// 21 American Express Corporate Purchasing Card
	SupplierReferenceNumber   string // 	9 A/N 4.287 Supplier Reference Number
	CardholderReferenceNumber string // 	17 A/N 4.59 Cardholder Reference Number
	ShippingZipCode           string // 	6 A/N 4.277 Shipped To ZIP Code
	SalesTax                  uint64 // 	6 N 4.262 Sales Tax
	ChargeDescriptor1         string // 	40 A/N 4.65 Charge Descriptor 1

	// 22 U.S. Only Non-T&E Commercial Card Level II (Visa and MasterCard)
	// 	1 A/N 4.221 Optional Amount Identifier
	// 	12 NUM 4.220 Optional Amount
	// 	17 A/N 4.237 Purchase Order No. - Customer Ref. ID

	// 23 Group Map Extension (Group 25 - 48) Reserved
	// 	4 A/N 4.243 Record Type @@@@
	GroupMapExtension *OptionalDataGroupMap

	// 24 Reserved (POS-partner) Use Only

	// EXTENDED BIT MAP SECTION

	// 25 Commercial Card / Level II (Non-T&E and Fleet Non-fuel)
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	// 	1 NUM 4.174 Local Tax Included Flag
	// 	12 NUM 4.173 Local Tax
	// 	1 NUM 4.209 National Tax Included Flag
	// 	12 NUM 4.208 National Tax Amount
	// 	17 A/N 4.237 Purchase Order No. / Cust. Ref. ID

	// 26 Visa Card Enhanced Data (Non-T&E and Fleet Non-fuel)
	// 	20 A/N 4.194 Merchant VAT Registration Number
	// 	13 A/N 4.87 Customer VAT Registration Number
	// 	4 A/N 4.285 Summary Commodity Code
	// 	12 NUM 4.101 Discount Amount
	// 	12 NUM 4.130 Freight Amount
	// 	12 NUM 4.107 Duty Amount
	// 	10 A/N 4.98 Destination Postal / ZIP Code
	// 	10 A/N 4.271 Ship from Postal / ZIP Code
	// 	3 A/N 4.97 Destination Country Code
	// 	15 A/N 4.325 Unique VAT Invoice Reference Number
	// 	6 NUM 4.222 Order Date
	// 	12 NUM 4.333 VAT / Tax Amount (Freight/Shipping)
	// 	4 NUM 4.334 VAT / Tax Rate (Freight/Shipping)
	// 	3 NUM 4.167 Line Item Count

	// 27 MasterCard Enhanced Data (Non-T&E and Fleet Non-fuel)
	// 	12 NUM 4.130 Freight Amount
	// 	12 NUM 4.107 Duty Amount
	// 	10 A/N 4.98 Destination Postal / Zip Code
	// 	10 A/N 4.271 Ship From Postal / Zip Code
	// 	3 A/N 4.97 Destination Country Code
	// 	1 A/N 4.25 Alternate Tax Amount Indicator
	// 	9 NUM 4.24 Alternate Tax Amount
	// 	3 NUM 4.167 Line Item Count

	// 28 Visa Card Enhanced Data (Fleet)
	// 	4 NUM 4.238 Purchase Time
	// 	1 A/N 4.168 Service Type
	// 	2 NUM 4.133 Fuel Type
	// 	12 NUM 4.203 Motor Fuel Unit Price
	// 	1 A/N 4.326 Motor Fuel Unit of Measure
	// 	12 NUM 4.240 Motor Fuel Quantity
	// 	12 NUM 4.138 Gross Fuel Price
	// 	12 NUM 4.210 Net Fuel Price
	// 	12 NUM 4.212 Net Non-fuel Price
	// 	7 A/N 4.217 Odometer Reading
	// 	17 A/N 4.237 Cust. Code / Cust. Ref. Identifier (CRI)
	// 	1 NUM 4.239 Type of Purchase
	// 	12 NUM 4.279 State Fuel Tax Amount
	// 	12 NUM 4.172 Local Fuel Tax Amount
	// 	4 NUM 4.334 VAT Tax Rate
	// 	2 A/N 4.233 Product Code 1
	// 	2 A/N 4.233 Product Code 2
	// 	2 A/N 4.233 Product Code 3
	// 	2 A/N 4.233 Product Code 4
	// 	2 A/N 4.233 Product Code 5
	// 	2 A/N 4.233 Product Code 6
	// 	2 A/N 4.233 Product Code 7
	// 	2 A/N 4.233 Product Code 8
	// 	12 NUM 4.173 Local Sales Tax
	// 	1 A/N 4.174 Local Tax Included
	// 	9 NUM 4.139 Gross Non-fuel price
	// 	1 A/N 4.198 Misc. Fuel Tax Exemption Status
	// 	1 A/N 4.199 Misc. Non-Fuel Tax Exemption Status
	// 	3 NUM 4.167 Line Item Count (Must be 000)

	// 29 MasterCard Enhanced Data (Fleet - Fuel Only)
	// 	4 NUM 4.218 Oil Company Brand Names
	// 	30 A/N 4.193 Merchant Street Address
	// 	9 A/N 4.190 Merchant Postal/ZIP Code and Ext.
	// 	4 NUM 4.238 Purchase Time
	// 	1 A/N 4.268 Motor Fuel Service Type
	// 	3 NUM 4.233 Motor Fuel Product Code
	// 	5 NUM 4.203 Motor Fuel Unit Price
	// 	1 A/N 4.327 Motor Fuel Unit of Measure
	// 	6 NUM 4.240 Motor Fuel Quantity
	// 	9 NUM 4.202 Motor Fuel Sale Amount
	// 	7 A/N 4.217 Odometer Reading
	// 	17 A/N 4.335 Vehicle Number
	// 	17 A/N 4.106 Driver Number / ID No.
	// 	1 NUM 4.234 Product Type Code
	// 	5 NUM 4.101 Coupon/Discount Amount
	// 	6 A/N 4.291 Tax Amount 1
	// 	6 A/N 4.292 Tax Amount 2
	// 	3 NUM 4.167 Line Item Count

	// 30 American Express CPC Charge Descriptors
	// 	40 A/N 4.66 Charge Descriptor 2
	// 	40 A/N 4.67 Charge Descriptor 3
	// 	40 A/N 4.68 Charge Descriptor 4

	// 31 Market Specific Data Indicator
	// 	1 A/N 4.181 Market Specific Data Indicator

	// 32 POS Data Code
	// 	12 A/N 4.231 POS Data Code

	// 33 Retail Transaction Advice Addenda
	// 	40 A/N 4.255 Retail Department Name

	// 34 Lodging Transaction Advice Addenda
	// 	26 A/N 4.179 Lodging Renter Name

	// 35 Transaction Advice Line Item Counts
	// 	3 NUM 4.290 TAD Line Item Count
	// 	3 NUM 4.289 TAA Line Item Count

	// 36 Location Detail Transaction Advice Addenda
	// 	38 A/N 4.177 Location Detail Name
	// 	38 A/N 4.177 Location Detail Address
	// 	21 A/N 4.177 Location Detail City
	// 	3 A/N 4.177 Location Detail Region Code
	// 	3 A/N 4.177 Country Code
	// 	15 A/N 4.177 Location Detail Postal Code

	// 37 MasterCard Miscellaneous Fields
	// 	3 NUM 4.338 Version Number
	// 	3 A/N 4.265 Service Code

	// 38 Visa Miscellaneous Fields
	// 	3 NUM 4.338 Version Number
	// 	2 A/N 4.51 Card product code

	// 39 Amex CAPN Corporate Purchasing Solution Extension
	// 	38 A/N 4.252 Requester Name
	// 	12 Num 4.311 Total Tax Amount
	// 	3 Num 4.294 Tax Type Code

	// 40 Discover/ PayPal Miscellaneous Fields
	// 	3 Num 4.338 Version Number
	// 	13 A/N 4.103 Discover/ PayPal POS Data Code
	// 	1 A/N 4.227 Partial Shipment Indicator

	// 41 Detail Extension
	// 	4 Num 4.141 Group Length
	// 	Variable A/N 4.140 Group Fields

	// 	ODG 41 Tag Description
	// 		TAG Max size Format Reference Content/Description
	// 		OAI 1 A/N 4.221 Optional Amount Identifier
	// 		OA 12 NUM 4.220 Optional Amount
	// 		PON 25 A/N 4.237 Purchase Order Number
	// 		PTI 3 NUM 4.229 Payment Transaction Indicator
	// 		IIA 1 A/N 4.182 MC IIAS Indicator
	// 		SID 20 A/N 4.191 Merchant Seller ID
	// 		CDV 82 A/N 4.88 Custom Identifier Detail
	// 		ATS 10 NUM 4.35 Association Timestamp
	// 		TTT 2 NUM 4.319 Transit Transaction Type Indicator
	// 		TMI 2 NUM 4.320 Transportation Mode Indicator
	// 		TFA 9 A/N 4.313 Transaction Fee Amount with Indicator
	// 		MDO 1 A/N 4.104 Domain Server
	// 		MDE 2 A/N 4.201 Mobile Device Type
	// 		DTC 1 A/N 4.149 Invoice Level Discount Treatment Code
	// 		TTC 1 A/N 4.293 Tax Treatment Code
	// 		SQI 1 A/N 4.279 Spend Qualified Indicator
	// 		TAL 2 A/N 4.306 Token Assurance Level
	// 		TRI 11 A/N 4.307 Token Requestor ID
	// 		ARS 1 A/N 4.2 Account Range Status
	// 		WID 3 A/N 4.183 MasterCard Wallet Identifier
	// 		ISO 11 NUM 4.151 ISO ID
	// 		PFI 11 NUM 4.229 Payment Facilitator ID NOTE: The presence of this tag requires use of the SMI TAG as well as Usage 2 of ODG 36 “Location Detail Transaction Advice Addenda.”
	// 		SMI 15 A/N 4.285 Sub-Merchant ID NOTE: The presence of this tag requires use of the PFI TAG as well as Usage 2 of ODG 36 “Location Detail Transaction Advice Addenda.”

	// 42 Amex Auto Rental
	// 	14 A/N 4.27.1 Amex Auto Rental Agreement Number
	// 	38 A/N 4.27.2 Amex Auto Rental Pickup Location
	// 	18 A/N 4.27.3 Amex Auto Rental Pickup City Name
	// 	3 A/N 4.27.4 Amex Auto Rental Pickup Region Code
	// 	3 A/N 4.27.5 Amex Auto Rental Pickup Country Code
	// 	8 NUM 4.27.6 Amex Auto Rental Pickup Date
	// 	6 NUM 4.27.7 Amex Auto Rental Pickup Time
	// 	18 A/N 4.27.8 Amex Auto Rental Return City Name
	// 	3 A/N 4.27.9 Amex Auto Rental Return Region Code
	// 	3 A/N 4.27.10 Amex Auto Rental Return Country Code
	// 	8 NUM 4.27.11 Amex Auto Rental Return Date
	// 	6 NUM 4.27.12 Amex Auto Rental Return Time
	// 	26 A/N 4.27.13 Amex Auto Rental Renter Name
	// 	4 A/N 4.27.14 Amex Auto Rental Vehicle Class ID
	// 	5 NUM 4.27.15 Amex Auto Rental Distance
	// 	1 A/N 4.27.16 Amex Auto Rental Distance Unit of Measurement
	// 	1 A/N 4.27.17 Amex Auto Rental Audit Adjustment Indicator
	// 	12 NUM 4.27.18 Amex Auto Rental Audit Adjustment Amount

	// 43 Amex Insurance
	// 	23 A/N 4.28.1 Policy Number
	// 	8 NUM 4.28.2 Coverage Start Date
	// 	8 NUM 4.28.3 Coverage End Date
	// 	7 A/N 4.28.4 Policy Premium Frequency
	// 	23 A/N 4.28.5 Additional Insurance Policy Number
	// 	25 A/N 4.28.6 Type of Policy
	// 	30 A/N 4.28.7 Name of Insured

	// 45 Transaction Security Indicator
	TransactionSecurityIndicator TransactionSecurityIndicatorType // 	1 A/N 4.315 Transaction Security Indicator

	// 46 Merchant Soft Descriptor
	MerchantSoftDescriptor string // 	25 A/N 4.192 Merchant Soft Descriptor

}

func NewBatchDetailOptionalDataRecord(
	dataGroupMap *OptionalDataGroupMap,
	g1request *G1AuthorizationMessageRequest,
	g2request *G2AuthorizationMessageRequest,
	g1response *G1AuthorizationMessageResponse,
	settlementAmount uint64) *BatchDetailOptionalDataRecord {

	d := BatchDetailOptionalDataRecord{}
	d.GroupMap = dataGroupMap

	return &d
}

func (d *BatchDetailOptionalDataRecord) CopyToBuffer(buffer *bytes.SafeBuffer) {
	buffer.AppendFormat("%c", d.GroupMap) //1 1 A/N Record Formad K 4.242

	// Table 3.6 Transaction detail record - optional data groups
	// 1 Cash Back
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%012d", d.CashbackAmount) // 	12 NUM 4.63 Cash Back Amount
	}
	// 2 Restaurant
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(2)) {
		buffer.AppendFormat("%012d", d.GratuityAmount) // 	12 NUM 4.137 Gratuity Amount
	}
	// 3 Direct Marketing
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(3)) {
		buffer.AppendFormat("%012d", d.TotalAuthorizedAmount)     // 	12 NUM 4.306 Total Authorized Amount
		buffer.AppendFormat("%c", d.PurchaseIdentifierFormatCode) // 	1 A/N 4.236 Purchase Identifier Format Code
		buffer.AppendFormat("%-25s", d.PurchaseIdentifier)        // 	25 A/N 4.235 Purchase Identifier
	}
	// 4 Auto Rental or 5 Hotel
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(4)) || true == d.GroupMap.GetGroupBit(GroupMapBitType(5)) {
		buffer.AppendFormat("%012d", d.TotalAuthorizedAmount)     // 	12 NUM 4.306 Total Authorized Amount
		buffer.AppendFormat("%c", d.PurchaseIdentifierFormatCode) // 	1 A/N 4.236 Purchase Identifier Format Code
		buffer.AppendFormat("%-25s", d.PurchaseIdentifier)        // 	25 A/N 4.235 Purchase Identifier
		buffer.AppendFormat("%c", d.MarketSpecificDataIdentifier) // 	1 A/N 4.181 Market Specific Data Identifier
		buffer.AppendFormat("%c", d.NoShowIndicator)              // 	1 A/N 4.215 No Show Indicator
		buffer.AppendFormat("%-6s", d.ExtraCharges)               // 	6 A/N 4.121.1 Extra Charges (AutoRentalExtraChargeType & HotelExtraChargeType)
		buffer.AppendFormat("%06d", d.RentalOrCheckinDate)        // 	6 NUM 4.246 Rental Date (YYMMDD)
	}

	// 5 Hotel
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	// 	1 A/N 4.181 Market Specific Data Identifier
	// 	1 A/N 4.215 4.278 No Show or Special Program Ind.
	// 	6 A/N 4.121 Extra Charges
	// 	6 NUM 4.72 Check-in Date (YYMMDD)

	// 6 Hotel - Not Front Desk (Non-Lodging Purchase)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(6)) {
		buffer.AppendFormat("%04d", d.MerchantCategoryCodeOverride) // 	4 NUM 4.185 Merchant Category Code Override
	}
	// 7 Hotel - American Express
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(7)) {
		buffer.AppendFormat("%c", d.HotelExtraChargeCode) // 	1 A/N 4.71 Charge Type
		buffer.AppendFormat("%06d", d.CheckoutDate)       // 	6 NUM 4.73 Checkout Date (YYMMDD)
		buffer.AppendFormat("%02d", d.StayDuration)       // 	2 NUM 4.281 Stay Duration
		buffer.AppendFormat("%012d", d.RoomRate)          // 	12 NUM 4.180 Lodging Room Rate
	}

	// 8 Non-T&E Commercial Card Level II (Visa and MasterCard Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		buffer.AppendFormat("%c", d.OptionalAmountIdentifier) // 	1 A/N 4.221 Optional Amount Identifier
		buffer.AppendFormat("%012d", d.OptionalAmount)        // 	12 NUM 4.220 Optional Amount
		buffer.AppendFormat("%-16s", d.PurchaseOrderNumber)   // 	16 A/N 4.237 Purchase Order Number
	}

	// 9 Passenger Transport (Visa Only)
	// 	1 A/N 4.254 Restricted Ticket Indicator
	// 	2 NUM 4.207 Multiple Clearing Sequence Number
	// 	2 NUM 4.206 Multiple Clearing Sequence Count
	// 	13 A/N 4.304 Ticket Number
	// 	20 A/N 4.228 Passenger Name
	// 	6 NUM 4.93 Departure Date (MMDDYY)
	// 	3 A/N 4.224 Origination City/Airport Code
	// 	7 A/N 4.323 Trip Leg 1 Information
	// 	7 A/N 4.323 Trip Leg 2 Information
	// 	7 A/N 4.323 Trip Leg 3 Information
	// 	7 A/N 4.323 Trip Leg 4 Information

	// 10 Direct Debit
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		buffer.AppendFormat("%-12s", d.RetrievalReferenceNumber) // 	12 A/N 4.259 Retrieval Reference Number
		buffer.AppendFormat("%06d", d.SystemTraceAuditNumber)    // 	6 NUM 4.288 System Trace Audit Number
		buffer.AppendFormat("%c", d.NetworkIdentificationCode)   // 	1 A/N 4.213 Network Identification Code
		buffer.AppendFormat("%04d", d.SettlementDate)            // 	4 NUM 4.270 Settlement Date (MMDD)
	}
	// 	11 Multi-Terminal Stores
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(11)) {
		buffer.AppendFormat("%04d", d.TerminalNumberOverride) // 	4 NUM 4.301 Terminal Number Override
	}
	// 12 Direct Marketing - Installment or Recurring Payments
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(12)) {
		buffer.AppendFormat("%02d", d.MultipleClearingSequenceNumber) // 	2 NUM 4.207 Multiple Clearing Sequence Number
		buffer.AppendFormat("%02d", d.MultipleClearingSequenceCount)  // 	2 NUM 4.206 Multiple Clearing Sequence Count
		buffer.AppendFormat("%c", d.MotoEcomIndicator)                // 	1 A/N 4.205 MOTO/e-Commerce Indicator
	}
	// 13 Auto Rental
	// 	20 A/N 4.250 Renter Name
	// 	18 A/N 4.247 Rental Return City
	// 	3 A/N 4.249 Rental Return State/Country
	// 	10 A/N 4.248 Rental Return Location ID

	// 14 Service Development Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(14)) {
		buffer.AppendFormat("%d", d.ServiceDevelopmentIndicator) // 	1 NUM 4.267 Service Development Indicator
	}
	// 16 Existing Debt Indicator (Visa Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(16)) {
		buffer.AppendFormat("%d", d.ExistingDebtIndicator) // 	1 NUM 4.119 Existing Debt Indicator
	}
	// 17 Universal Cardholder Authentication
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(17)) {
		buffer.AppendFormat("%c", d.UcafCollectionIndicator) // 	1 A/N 4.324 UCAF Collection Indicator
	}
	// 18 e-Commerce Goods Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(18)) {
		buffer.AppendFormat("%c ", d.EcommerceGoodsIndicator) // 	2 A/N 4.108 e-Commerce Goods Indicator
	}
	// 19 CCS Private Label Restricted for CCS use only (Mandatory for Private Label transactions)
	// 	5 NUM 4.80 Credit Plan Number (Not Null)
	// 	4 A/N 4.233 Product Code 1
	// 	4 A/N 4.233 Product Code 2
	// 	4 A/N 4.233 Product Code 3
	// 	4 A/N 4.233 Product Code 4
	// 	4 A/N 4.233 Product Code 5
	// 	3 NUM 4.223 Org. ID (Not Null)
	// 	9 NUM 4.283 Store Code (Not Null)
	// 	6 NUM 4.241 Receiving Institution ID (“911111”)

	// 20 Merchant Verification Value
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(20)) {
		buffer.AppendFormat("%-10s", d.MerchantVerificationValue) // 	10 A/N 4.195 Merchant Verification Value
	}

	// 21 American Express Corporate Purchasing Card
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(21)) {
		buffer.AppendFormat("%-9s", d.SupplierReferenceNumber)   // 	9 A/N 4.287 Supplier Reference Number
		buffer.AppendFormat("%02d", d.CardholderReferenceNumber) // 	17 A/N 4.59 Cardholder Reference Number
		buffer.AppendFormat("%-6s", d.ShippingZipCode)           // 	6 A/N 4.277 Shipped To ZIP Code
		buffer.AppendFormat("%06d", d.SalesTax)                  // 	6 N 4.262 Sales Tax
		buffer.AppendFormat("%-40s", d.ChargeDescriptor1)        // 	40 A/N 4.65 Charge Descriptor 1
	}

	// 22 U.S. Only Non-T&E Commercial Card Level II (Visa and MasterCard)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(22)) {
		buffer.AppendFormat("%c", d.OptionalAmountIdentifier) // 	1 A/N 4.221 Optional Amount Identifier
		buffer.AppendFormat("%012d", d.OptionalAmount)        // 	12 NUM 4.220 Optional Amount
		buffer.AppendFormat("%-16s", d.PurchaseOrderNumber)   // 	17 A/N 4.237 Purchase Order No. - Customer Ref. ID
	}

	// 23 Group Map Extension (Group 25 - 48) Reserved
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(23)) {
		buffer.AppendFormat("%s", d.GroupMapExtension) // 	4 A/N 4.243 Record Type @@@@
	}
}

func (d *BatchDetailOptionalDataRecord) Bytes() []byte {
	buffer := bytes.NewSafeBuffer(128)
	d.CopyToBuffer(buffer)
	return buffer.Bytes()
}

func (d *BatchDetailOptionalDataRecord) VerboseString() string {

	buffer := bytes.NewSafeBuffer(256)

	buffer.AppendFormat("%-32s%c (%[2]s)\n", "GroupMap", d.GroupMap) //1 1 A/N Record Format K 4.242

	// Table 3.6 Transaction detail record - optional data groups
	// 1 Cash Back
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(1)) {
		buffer.AppendFormat("%-32s%012d", "CashbackAmount", d.CashbackAmount) // 	12 NUM 4.63 Cash Back Amount
	}
	// 2 Restaurant
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(2)) {
		buffer.AppendFormat("%-32s%012d", "GratuityAmount", d.GratuityAmount) // 	12 NUM 4.137 Gratuity Amount
	}
	// 3 Direct Marketing
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(3)) {
		buffer.AppendFormat("%-32s%012d", "TotalAuthorizedAmount", d.TotalAuthorizedAmount)                    // 	12 NUM 4.306 Total Authorized Amount
		buffer.AppendFormat("%-32s%c (%[2]s)", "PurchaseIdentifierFormatCode", d.PurchaseIdentifierFormatCode) // 	1 A/N 4.236 Purchase Identifier Format Code
		buffer.AppendFormat("%-32s%-25s", "PurchaseIdentifier", d.PurchaseIdentifier)                          // 	25 A/N 4.235 Purchase Identifier
	}
	// 4 Auto Rental or 5 Hotel
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(4)) || true == d.GroupMap.GetGroupBit(GroupMapBitType(5)) {
		buffer.AppendFormat("%-32s%012d", "TotalAuthorizedAmount", d.TotalAuthorizedAmount)                    // 	12 NUM 4.306 Total Authorized Amount
		buffer.AppendFormat("%-32s%c (%[2]s)", "PurchaseIdentifierFormatCode", d.PurchaseIdentifierFormatCode) // 	1 A/N 4.236 Purchase Identifier Format Code
		buffer.AppendFormat("%-32s%-25s", "PurchaseIdentifier", d.PurchaseIdentifier)                          // 	25 A/N 4.235 Purchase Identifier
		buffer.AppendFormat("%-32s%c (%[2]s)", "MarketSpecificDataIdentifier", d.MarketSpecificDataIdentifier) // 	1 A/N 4.181 Market Specific Data Identifier
		buffer.AppendFormat("%-32s%c (%[2]s)", "NoShowIndicator", d.NoShowIndicator)                           // 	1 A/N 4.215 No Show Indicator
		buffer.AppendFormat("%-32s%-6s", "ExtraCharges", d.ExtraCharges)                                       // 	6 A/N 4.121.1 Extra Charges (AutoRentalExtraChargeType & HotelExtraChargeType)
		buffer.AppendFormat("%-32s%06d", "RentalOrCheckinDate", d.RentalOrCheckinDate)                         // 	6 NUM 4.246 Rental Date (YYMMDD)
	}

	// 5 Hotel
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	// 	1 A/N 4.181 Market Specific Data Identifier
	// 	1 A/N 4.215 4.278 No Show or Special Program Ind.
	// 	6 A/N 4.121 Extra Charges
	// 	6 NUM 4.72 Check-in Date (YYMMDD)

	// 6 Hotel - Not Front Desk (Non-Lodging Purchase)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(6)) {
		buffer.AppendFormat("%-32s%04d", "MerchantCategoryCodeOverride", d.MerchantCategoryCodeOverride) // 	4 NUM 4.185 Merchant Category Code Override
	}
	// 7 Hotel - American Express
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(7)) {
		buffer.AppendFormat("%-32s%c (%[2]s)", "HotelExtraChargeCode", d.HotelExtraChargeCode) // 	1 A/N 4.71 Charge Type
		buffer.AppendFormat("%-32s%06d", "CheckoutDate", d.CheckoutDate)                       // 	6 NUM 4.73 Checkout Date (YYMMDD)
		buffer.AppendFormat("%-32s%02d", "StayDuration", d.StayDuration)                       // 	2 NUM 4.281 Stay Duration
		buffer.AppendFormat("%-32s%012d", "RoomRate", d.RoomRate)                              // 	12 NUM 4.180 Lodging Room Rate
	}

	// 8 Non-T&E Commercial Card Level II (Visa and MasterCard Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		buffer.AppendFormat("%-32s%c (%[2]s)", "OptionalAmountIdentifier", d.OptionalAmountIdentifier) // 	1 A/N 4.221 Optional Amount Identifier
		buffer.AppendFormat("%-32s%012d", "OptionalAmount", d.OptionalAmount)                          // 	12 NUM 4.220 Optional Amount
		buffer.AppendFormat("%-32s%-16s", "PurchaseOrderNumber", d.PurchaseOrderNumber)                // 	16 A/N 4.237 Purchase Order Number
	}

	// 9 Passenger Transport (Visa Only)
	// 	1 A/N 4.254 Restricted Ticket Indicator
	// 	2 NUM 4.207 Multiple Clearing Sequence Number
	// 	2 NUM 4.206 Multiple Clearing Sequence Count
	// 	13 A/N 4.304 Ticket Number
	// 	20 A/N 4.228 Passenger Name
	// 	6 NUM 4.93 Departure Date (MMDDYY)
	// 	3 A/N 4.224 Origination City/Airport Code
	// 	7 A/N 4.323 Trip Leg 1 Information
	// 	7 A/N 4.323 Trip Leg 2 Information
	// 	7 A/N 4.323 Trip Leg 3 Information
	// 	7 A/N 4.323 Trip Leg 4 Information

	// 10 Direct Debit
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		buffer.AppendFormat("%-32s%-12s", "RetrievalReferenceNumber", d.RetrievalReferenceNumber)        // 	12 A/N 4.259 Retrieval Reference Number
		buffer.AppendFormat("%-32s%06d", "SystemTraceAuditNumber", d.SystemTraceAuditNumber)             // 	6 NUM 4.288 System Trace Audit Number
		buffer.AppendFormat("%-32s%c (%[2]s)", "NetworkIdentificationCode", d.NetworkIdentificationCode) // 	1 A/N 4.213 Network Identification Code
		buffer.AppendFormat("%-32s%04d", "SettlementDate", d.SettlementDate)                             // 	4 NUM 4.270 Settlement Date (MMDD)
	}
	// 	11 Multi-Terminal Stores
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(11)) {
		buffer.AppendFormat("%-32s%04d", "TerminalNumberOverride", d.TerminalNumberOverride) // 	4 NUM 4.301 Terminal Number Override
	}
	// 12 Direct Marketing - Installment or Recurring Payments
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(12)) {
		buffer.AppendFormat("%-32s%02d", "MultipleClearingSequenceNumber", d.MultipleClearingSequenceNumber) // 	2 NUM 4.207 Multiple Clearing Sequence Number
		buffer.AppendFormat("%-32s%02d", "MultipleClearingSequenceCount", d.MultipleClearingSequenceCount)   // 	2 NUM 4.206 Multiple Clearing Sequence Count
		buffer.AppendFormat("%-32s%c (%[2]s)", "MotoEcomIndicator", d.MotoEcomIndicator)                     // 	1 A/N 4.205 MOTO/e-Commerce Indicator
	}
	// 13 Auto Rental
	// 	20 A/N 4.250 Renter Name
	// 	18 A/N 4.247 Rental Return City
	// 	3 A/N 4.249 Rental Return State/Country
	// 	10 A/N 4.248 Rental Return Location ID

	// 14 Service Development Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(14)) {
		buffer.AppendFormat("%-32s%d", "ServiceDevelopmentIndicator", d.ServiceDevelopmentIndicator) // 	1 NUM 4.267 Service Development Indicator
	}
	// 16 Existing Debt Indicator (Visa Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(16)) {
		buffer.AppendFormat("%-32s%d", "ExistingDebtIndicator", d.ExistingDebtIndicator) // 	1 NUM 4.119 Existing Debt Indicator
	}
	// 17 Universal Cardholder Authentication
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(17)) {
		buffer.AppendFormat("%-32s%c (%[2]s)", "UcafCollectionIndicator", d.UcafCollectionIndicator) // 	1 A/N 4.324 UCAF Collection Indicator
	}
	// 18 e-Commerce Goods Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(18)) {
		buffer.AppendFormat("%-32s%c  (%[2]s)", "EcommerceGoodsIndicator", d.EcommerceGoodsIndicator) // 	2 A/N 4.108 e-Commerce Goods Indicator
	}
	// 19 CCS Private Label Restricted for CCS use only (Mandatory for Private Label transactions)
	// 	5 NUM 4.80 Credit Plan Number (Not Null)
	// 	4 A/N 4.233 Product Code 1
	// 	4 A/N 4.233 Product Code 2
	// 	4 A/N 4.233 Product Code 3
	// 	4 A/N 4.233 Product Code 4
	// 	4 A/N 4.233 Product Code 5
	// 	3 NUM 4.223 Org. ID (Not Null)
	// 	9 NUM 4.283 Store Code (Not Null)
	// 	6 NUM 4.241 Receiving Institution ID (“911111”)

	// 20 Merchant Verification Value
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(20)) {
		buffer.AppendFormat("%-32s%-10s", "MerchantVerificationValue", d.MerchantVerificationValue) // 	10 A/N 4.195 Merchant Verification Value
	}

	// 21 American Express Corporate Purchasing Card
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(21)) {
		buffer.AppendFormat("%-32s%-9s", "SupplierReferenceNumber", d.SupplierReferenceNumber)     // 	9 A/N 4.287 Supplier Reference Number
		buffer.AppendFormat("%-32s%02d", "CardholderReferenceNumber", d.CardholderReferenceNumber) // 	17 A/N 4.59 Cardholder Reference Number
		buffer.AppendFormat("%-32s%-6s", "ShippingZipCode", d.ShippingZipCode)                     // 	6 A/N 4.277 Shipped To ZIP Code
		buffer.AppendFormat("%-32s%06d", "SalesTax", d.SalesTax)                                   // 	6 N 4.262 Sales Tax
		buffer.AppendFormat("%-32s%-40s", "ChargeDescriptor1", d.ChargeDescriptor1)                // 	40 A/N 4.65 Charge Descriptor 1
	}

	// 22 U.S. Only Non-T&E Commercial Card Level II (Visa and MasterCard)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(22)) {
		buffer.AppendFormat("%-32s%c (%[2]s)", "OptionalAmountIdentifier", d.OptionalAmountIdentifier) // 	1 A/N 4.221 Optional Amount Identifier
		buffer.AppendFormat("%-32s%012d", "OptionalAmount", d.OptionalAmount)                          // 	12 NUM 4.220 Optional Amount
		buffer.AppendFormat("%-32s%-16s", "PurchaseOrderNumber", d.PurchaseOrderNumber)                // 	17 A/N 4.237 Purchase Order No. - Customer Ref. ID
	}

	// 23 Group Map Extension (Group 25 - 48) Reserved
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(23)) {
		buffer.AppendFormat("%-32s%s", "GroupMapExtension", d.GroupMapExtension) // 	4 A/N 4.243 Record Type @@@@
	}

	return buffer.String()
}

func ParseBatchDetailOptionalDataRecord(s string) (*BatchDetailOptionalDataRecord, int) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!
	d := BatchDetailOptionalDataRecord{}

	i := 0
	if 108 > len(s) {
		panic("minimum EIS1081 Batch Detail Record length is 108")
	}

	// //i++ //advance past the STX if not already removed

	// fmt.Sscanf(s[0:4], "%c%1d%c%c",
	// 	&d.GroupMap,         //1 1 A/N Record Format K 4.242
	// 	&d.ApplicationType,  //2 1 NUM Application Type 1 - Single Batch 3 - Multiple Batch 4.32
	// 	&d.MessageDelimiter, //3 1 A/N Message Delimiter . 4.196
	// 	&d.RoutingId)        // 4 1 A/N X.25 Routing ID Z 4.341

	// d.TransactionCode = TransactionCodeType(s[9:11]) // 10-11 2 A/N Transaction Code 54 - Purchase 56 - Card not present CR - Credit/Return 4.312
	// d.CardholderAccountNumber = s[13:35]             // 14-35 22 A/N Cardholder Account Number Left-justified/space-filled 4.55

	// Table 3.6 Transaction detail record - optional data groups
	// 1 Cash Back
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(1)) {
		fmt.Sscanf(s[i:i+4], "%012d", &d.CashbackAmount) // 	12 NUM 4.63 Cash Back Amount
		i += 4
	}
	// 2 Restaurant
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(2)) {
		fmt.Sscanf(s[i:i+12], "%012d", &d.GratuityAmount) // 	12 NUM 4.137 Gratuity Amount
		i += 12
	}
	// 3 Direct Marketing
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(3)) {
		fmt.Sscanf(s[i:i+12], "%012d", &d.TotalAuthorizedAmount) // 	12 NUM 4.306 Total Authorized Amount
		i += 12
		d.PurchaseIdentifierFormatCode = PurchaseIdFormatType(s[i]) // 	1 A/N 4.236 Purchase Identifier Format Code
		i += 1
		d.PurchaseIdentifier = s[i : i+25] // 	25 A/N 4.235 Purchase Identifier
		i += 25
	}
	// 4 Auto Rental or 5 Hotel
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(4)) || true == d.GroupMap.GetGroupBit(GroupMapBitType(5)) {
		fmt.Sscanf(s[i:i+12], "%012d", &d.TotalAuthorizedAmount) // 	12 NUM 4.306 Total Authorized Amount
		i += 12
		d.PurchaseIdentifierFormatCode = PurchaseIdFormatType(s[i]) // 	1 A/N 4.236 Purchase Identifier Format Code
		i += 1
		d.PurchaseIdentifier = s[i : i+25] // 	25 A/N 4.235 Purchase Identifier
		i += 25
		d.MarketSpecificDataIdentifier = MarketSpecificDataType(s[i]) // 	1 A/N 4.181 Market Specific Data Identifier
		i += 1
		d.NoShowIndicator = NoShowIndicatorType(s[i]) // 	1 A/N 4.215 No Show Indicator
		i += 1
		d.ExtraCharges = s[i : i+6] // 	6 A/N 4.121.1 Extra Charges (AutoRentalExtraChargeType & HotelExtraChargeType)
		i += 6
		fmt.Sscanf(s[i:i+6], "%06d", &d.RentalOrCheckinDate) // 	6 NUM 4.246 Rental Date (YYMMDD)
		i += 6
	}

	// 5 Hotel
	// 	12 NUM 4.306 Total Authorized Amount
	// 	1 A/N 4.236 Purchase Identifier Format Code
	// 	25 A/N 4.235 Purchase Identifier
	// 	1 A/N 4.181 Market Specific Data Identifier
	// 	1 A/N 4.215 4.278 No Show or Special Program Ind.
	// 	6 A/N 4.121 Extra Charges
	// 	6 NUM 4.72 Check-in Date (YYMMDD)

	// 6 Hotel - Not Front Desk (Non-Lodging Purchase)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(6)) {
		fmt.Sscanf(s[i:i+4], "%04d", &d.MerchantCategoryCodeOverride) // 	4 NUM 4.185 Merchant Category Code Override
		i += 4
	}
	// 7 Hotel - American Express
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(7)) {
		d.HotelExtraChargeCode = HotelExtraChargeType(s[i]) // 	1 A/N 4.71 Charge Type
		i += 1
		// 	6 NUM 4.73 Checkout Date (YYMMDD)
		// 	2 NUM 4.281 Stay Duration
		// 	12 NUM 4.180 Lodging Room Rate
		fmt.Sscanf(s[i:i+20], "%06d%02d%012d", &d.CheckoutDate, &d.StayDuration, &d.RoomRate)
		i += 20
	}

	// 8 Non-T&E Commercial Card Level II (Visa and MasterCard Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		d.OptionalAmountIdentifier = LocalTaxIndicatorType(s[i]) // 	1 A/N 4.221 Optional Amount Identifier
		i += 1
		fmt.Sscanf(s[i:i+12], "%012d", &d.OptionalAmount) // 	12 NUM 4.220 Optional Amount
		i += 12
		d.PurchaseOrderNumber = s[i : i+16] // 	16 A/N 4.237 Purchase Order Number
		i += 16
	}

	// 9 Passenger Transport (Visa Only)
	// 	1 A/N 4.254 Restricted Ticket Indicator
	// 	2 NUM 4.207 Multiple Clearing Sequence Number
	// 	2 NUM 4.206 Multiple Clearing Sequence Count
	// 	13 A/N 4.304 Ticket Number
	// 	20 A/N 4.228 Passenger Name
	// 	6 NUM 4.93 Departure Date (MMDDYY)
	// 	3 A/N 4.224 Origination City/Airport Code
	// 	7 A/N 4.323 Trip Leg 1 Information
	// 	7 A/N 4.323 Trip Leg 2 Information
	// 	7 A/N 4.323 Trip Leg 3 Information
	// 	7 A/N 4.323 Trip Leg 4 Information

	// 10 Direct Debit
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(8)) {
		d.RetrievalReferenceNumber = s[i : i+12] // 	12 A/N 4.259 Retrieval Reference Number
		i += 12
		fmt.Sscanf(s[i:i+6], "%06d", &d.SystemTraceAuditNumber) // 	6 NUM 4.288 System Trace Audit Number
		i += 6
		d.NetworkIdentificationCode = NetworkIdentificationType(s[i]) // 	1 A/N 4.213 Network Identification Code
		i += 1
		fmt.Sscanf(s[i:i+4], "%04d", &d.SettlementDate) // 	4 NUM 4.270 Settlement Date (MMDD)
		i += 4
	}
	// 	11 Multi-Terminal Stores
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(11)) {
		fmt.Sscanf(s[i:i+4], "%04d", &d.TerminalNumberOverride) // 	4 NUM 4.301 Terminal Number Override
		i += 4
	}
	// 12 Direct Marketing - Installment or Recurring Payments
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(12)) {
		// 2 NUM 4.207 Multiple Clearing Sequence Number
		// 2 NUM 4.206 Multiple Clearing Sequence Count
		// 1 A/N 4.205 MOTO/e-Commerce Indicator
		fmt.Sscanf(s[i:i+4], "%02d%02d%c",
			&d.MultipleClearingSequenceNumber,
			&d.MultipleClearingSequenceCount,
			&d.MotoEcomIndicator)
		i += 5
	}

	// 13 Auto Rental
	// 	20 A/N 4.250 Renter Name
	// 	18 A/N 4.247 Rental Return City
	// 	3 A/N 4.249 Rental Return State/Country
	// 	10 A/N 4.248 Rental Return Location ID

	// 14 Service Development Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(14)) {
		fmt.Sscanf(s[i:i+1], "%d", &d.ServiceDevelopmentIndicator) // 	1 NUM 4.267 Service Development Indicator
		i += 1
	}
	// 16 Existing Debt Indicator (Visa Only)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(16)) {
		fmt.Sscanf(s[i:i+1], "%d", &d.ExistingDebtIndicator) // 	1 NUM 4.119 Existing Debt Indicator
		i += 1
	}
	// 17 Universal Cardholder Authentication
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(17)) {
		d.UcafCollectionIndicator = UcafCollectionType(s[i]) // 	1 A/N 4.324 UCAF Collection Indicator
		i += 1
	}
	// 18 e-Commerce Goods Indicator
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(18)) {
		d.EcommerceGoodsIndicator = EcommerceGoodsIndicatorType(s[i]) // 	2 A/N 4.108 e-Commerce Goods Indicator
		i += 2                                                        //skip past additonal space
	}
	// 19 CCS Private Label Restricted for CCS use only (Mandatory for Private Label transactions)
	// 	5 NUM 4.80 Credit Plan Number (Not Null)
	// 	4 A/N 4.233 Product Code 1
	// 	4 A/N 4.233 Product Code 2
	// 	4 A/N 4.233 Product Code 3
	// 	4 A/N 4.233 Product Code 4
	// 	4 A/N 4.233 Product Code 5
	// 	3 NUM 4.223 Org. ID (Not Null)
	// 	9 NUM 4.283 Store Code (Not Null)
	// 	6 NUM 4.241 Receiving Institution ID (“911111”)

	// 20 Merchant Verification Value
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(20)) {
		d.MerchantVerificationValue = s[i : i+10] // 	10 A/N 4.195 Merchant Verification Value
		i += 10
	}

	// 21 American Express Corporate Purchasing Card
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(21)) {
		d.SupplierReferenceNumber = s[i : i+9] // 	9 A/N 4.287 Supplier Reference Number
		i += 9
		fmt.Sscanf(s[i:i+2], "%02d", &d.CardholderReferenceNumber) // 	17 A/N 4.59 Cardholder Reference Number
		i += 2
		d.ShippingZipCode = s[i : i+6] // 	6 A/N 4.277 Shipped To ZIP Code
		i += 6
		fmt.Sscanf(s[i:i+6], "%06d", &d.SalesTax) // 	6 N 4.262 Sales Tax
		i += 6
		d.ChargeDescriptor1 = s[i : i+40] // 	40 A/N 4.65 Charge Descriptor 1
		i += 40
	}

	// 22 U.S. Only Non-T&E Commercial Card Level II (Visa and MasterCard)
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(22)) {
		fmt.Sscanf(s[i:i+1], "%c", &d.OptionalAmountIdentifier) // 	1 A/N 4.221 Optional Amount Identifier
		i += 1
		fmt.Sscanf(s[i:i+12], "%012d", &d.OptionalAmount) // 	12 NUM 4.220 Optional Amount
		i += 12
		d.PurchaseOrderNumber = s[i : i+17] // 	17 A/N 4.237 Purchase Order No. - Customer Ref. ID
		i += 17
	}

	// 23 Group Map Extension (Group 25 - 48) Reserved
	if true == d.GroupMap.GetGroupBit(GroupMapBitType(23)) {
		d.GroupMapExtension = ParseOptionalDataGroupsExtension(s[i : i+4]) // 	4 A/N 4.243 Record Type @@@@
		i += 4
	}

	return &d, i
}
