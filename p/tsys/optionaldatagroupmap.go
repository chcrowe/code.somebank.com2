package tsys

import (
	"code.somebank.com/p/bitwise"
	//	"code.somebank.com/p/card"
	"fmt"
	"strings"
)

// Bit position seven will always be a "1" to exclude the range of control character.
const BITSEVEN_VALUE = '@'
const DATAGROUPRECORDSIZE = 5

type OptionalDataGroupMap struct {
	RecordTypeCode RecordIndicatorType
	GroupByteI     byte
	GroupByteII    byte
	GroupByteIII   byte
	GroupByteIV    byte
}

func NewOptionalDataGroupMap(recordInd RecordIndicatorType) *OptionalDataGroupMap {
	o := OptionalDataGroupMap{recordInd, BITSEVEN_VALUE, BITSEVEN_VALUE, BITSEVEN_VALUE, BITSEVEN_VALUE}

	return &o
}

func (o *OptionalDataGroupMap) InitializeDetailRecordGroups(
	g1request *G1AuthorizationMessageRequest,
	g2request *G2AuthorizationMessageRequest,
	g3RequestRecords []string,
	g3ResponseRecords []string) {

	//cardIssuerCode := card.GetCardIssuerType(g1request.CustomerData.AccountNumber())

	switch g1request.IndustryCode {
	// case AutoRental:
	// case BankFinancialInstitution:
	// case GroceryStoreSupermarket:
	// case LimitedAmountTerminal:
	// case OilCompanyAutomatedFuelingSystem:
	// case PassengerTransport:
	// case Retail:
	case DirectMarketing:
		o.SetGroupBit(DetailRecordBit_DirectMarketing, true)
	case FoodRestaurant:
		o.SetGroupBit(DetailRecordBit_Restaurant, true)
	case HotelAndLodging:
		o.SetGroupBit(DetailRecordBit_Hotel, true)
	case Internet:
		o.SetGroupBit(DetailRecordBit_ElectronicCommerceGoodsIndicator, true)
	}

	switch g1request.TransactionCode {
	case DirectDebitPurchase, DirectDebitPurchaseReturn:
		o.SetGroupBit(DetailRecordBit_DirectDebit, true)
	case CashBenefitsCashWithdrawalEbt, FoodStampPurchaseEbt:
		o.SetGroupBit(DetailRecordBit_DirectDebit, true)
	}

	// var g3ver G3VersionType
	// for _, g3 := range g3RequestRecords {

	// 	fmt.Sscanf(g3, "%03d", &g3ver)
	// 	switch g3ver {

	// 	// case G3v007CardVerificationCode:
	// 	// 	G3v007 := ParseG3v007CardVerificationCodeRequest(g3)

	// 	case G3v008FleetFuelingCard:
	// 		G3v008 := ParseG3v008FleetFuelingCardRequest(g3)
	// 		if card.Visa == cardIssuerCode {
	// 			o.SetGroupBit(DetailRecordBit_VisaCardEnhancedDataFleet, true)
	// 		} else if card.Mastercard == cardIssuerCode {
	// 			o.SetGroupBit(DetailRecordBit_MasterCardEnhancedDataFleetFuelOnly, true)
	// 		}

	// 	case G3v015ServiceDevelopmentIndicator:
	// 		G3v015 := ParseG3v015ServiceDevelopmentIndicatorRequest(g3)
	// 		o.SetGroupBit(DetailRecordBit_ServiceDevelopmentIndicator, true)

	// 	case G3v018ExistingDebtIndicator:
	// 		G3v018 := ParseG3v018ExistingDebtIndicatorRequest(g3)
	// 		o.SetGroupBit(DetailRecordBit_ExistingDebtIndicatorVisaOnly, true)

	// 	case G3v019MastercardUniversalCardholderAuthentication:
	// 		G3v019 := ParseG3v019MastercardUniversalCardholderAuthenticationRequest(g3)
	// 		o.SetGroupBit(DetailRecordBit_UniversalCardholderAuthentication, true)

	// 	case G3v020DeveloperInformation:
	// 		G3v020 := ParseG3v020DeveloperInformationRequest(g3)

	// 	case G3v021MerchantVerificationValue:
	// 		G3v021 := ParseG3v021MerchantVerificationValueRequest(g3)
	// 		o.SetGroupBit(DetailRecordBit_MerchantVerificationValue, true)

	// 	case G3v022AdditionalAmounts:
	// 		G3v022 := ParseG3v022AdditionalAmountsRequest(g3)

	// 	case G3v025TransactionFeeAmount:
	// 		G3v025 := ParseG3v025TransactionFeeAmountRequest(g3)

	// 	case G3v027PosData:
	// 		G3v027 := ParseG3v027PosDataRequest(g3)
	// 		o.SetGroupBit(DetailRecordBit_POSDataCode, true)

	// 	case G3v032CurrencyConversionData:
	// 		G3v032 := ParseG3v032CurrencyConversionDataRequest(g3)

	// 	case G3v034CardProductCode:
	// 		G3v034 := ParseG3v034CardProductCodeRequest(g3)

	// 	case G3v040VisaISAChargeIndicator:
	// 		G3v040 := ParseG3v040VisaISAChargeIndicatorRequest(g3)

	// 	case G3v046CardTypeGroup:
	// 		G3v046 := ParseG3v046CardTypeGroupRequest(g3)

	// 	case G3v050AssociationTimestamp:
	// 		G3v050 := ParseG3v050AssociationTimestampRequest(g3)

	// 	case G3v051MasterCardEventMonitoringRealTimeScoring:
	// 		G3v051 := ParseG3v051MasterCardEventMonitoringRealTimeScoringRequest(g3)

	// 	case G3v058AlternateAccountID:
	// 		G3v058 := ParseG3v058AlternateAccountIDRequest(g3)

	// 	case G3v059MasterCardPayPassMappingService:
	// 		G3v059 := ParseG3v059MasterCardPayPassMappingServiceRequest(g3)

	// 	case G3v061SpendQualifiedIndicator:
	// 		G3v061 := ParseG3v061SpendQualifiedIndicatorRequest(g3)

	// 	}
	// }

	// for _, g3 := range g3ResponseRecords {

	// 	fmt.Sscanf(g3, "%03d", &g3ver)
	// 	switch g3ver {

	// 	case G3v001CommercialCard:
	// 		G3v001 := ParseG3v001CommercialCardResponse(g3)
	// 		switch G3v001.CommercialCardResponseIndicator {
	// 		//case BusinessCard,NonCommercialCard,InvalidRequestIndicatorReceived:
	// 		case PurchasingCard, CorporateCard, VisaCommerceReserved:
	// 			o.SetGroupBit(DetailRecordBit_CommercialCardLevelIINonTEFleetNonFuel, true)
	// 		}
	// 		// DetailRecordBit_VisaCardEnhancedDataNonTEFleetNonFuel        GroupMapBitType = 26
	// 		// DetailRecordBit_MasterCardEnhancedDataNonTEFleetNonFuel      GroupMapBitType = 27
	// 		// DetailRecordBit_VisaCardEnhancedDataFleet                    GroupMapBitType = 28
	// 		// DetailRecordBit_MasterCardEnhancedDataFleetFuelOnly          GroupMapBitType = 29

	// 	case G3v007CardVerificationCode:
	// 		G3v007 := ParseG3v007CardVerificationCodeResponse(g3)

	// 	case G3v018ExistingDebtIndicator:
	// 		G3v018 := ParseG3v018ExistingDebtIndicatorResponse(g3)
	// 		o.SetGroupBit(DetailRecordBit_ExistingDebtIndicatorVisaOnly, true)

	// 	case G3v019MastercardUniversalCardholderAuthentication:
	// 		G3v019 := ParseG3v019MastercardUniversalCardholderAuthenticationResponse(g3)
	// 		o.SetGroupBit(DetailRecordBit_UniversalCardholderAuthentication, true)

	// 	case G3v020DeveloperInformation:
	// 		G3v020 := ParseG3v020DeveloperInformationResponse(g3)

	// 	case G3v021MerchantVerificationValue:
	// 		G3v021 := ParseG3v021MerchantVerificationValueResponse(g3)
	// 		o.SetGroupBit(DetailRecordBit_MerchantVerificationValue, true)

	// 	case G3v022AdditionalAmounts:
	// 		G3v022 := ParseG3v022AdditionalAmountsResponse(g3)

	// 	case G3v025TransactionFeeAmount:
	// 		G3v025 := ParseG3v025TransactionFeeAmountResponse(g3)

	// 	case G3v032CurrencyConversionData:
	// 		G3v032 := ParseG3v032CurrencyConversionDataResponse(g3)

	// 	case G3v034CardProductCode:
	// 		G3v034 := ParseG3v034CardProductCodeResponse(g3)

	// 	case G3v040VisaISAChargeIndicator:
	// 		G3v040 := ParseG3v040VisaISAChargeIndicatorResponse(g3)

	// 	case G3v046CardTypeGroup:
	// 		G3v046 := ParseG3v046CardTypeGroupResponse(g3)

	// 	case G3v050AssociationTimestamp:
	// 		G3v050 := ParseG3v050AssociationTimestampResponse(g3)

	// 	case G3v051MasterCardEventMonitoringRealTimeScoring:
	// 		G3v051 := ParseG3v051MasterCardEventMonitoringRealTimeScoringResponse(g3)

	// 	case G3v058AlternateAccountID:
	// 		G3v058 := ParseG3v058AlternateAccountIDResponse(g3)

	// 	case G3v059MasterCardPayPassMappingService:
	// 		G3v059 := ParseG3v059MasterCardPayPassMappingServiceResponse(g3)

	// 	case G3v061SpendQualifiedIndicator:
	// 		G3v061 := ParseG3v061SpendQualifiedIndicatorResponse(g3)

	// 	}
	// }
}

func setGroupBit(x *byte, pos uint, enable bool) {
	var y byte
	if true == enable {
		y = bitwise.SetBit(*x, pos)
	} else {
		y = bitwise.ClearBit(*x, pos)
	}
	*x = y
}

func ParseOptionalDataGroups(s string) *OptionalDataGroupMap {
	o := OptionalDataGroupMap{}

	if DATAGROUPRECORDSIZE > len(s) {
		return &o
	}

	o.RecordTypeCode = RecordIndicatorType(s[0])
	o.GroupByteI = s[1]
	o.GroupByteII = s[2]
	o.GroupByteIII = s[3]
	o.GroupByteIV = s[4]

	return &o
}

func ParseOptionalDataGroupsExtension(s string) *OptionalDataGroupMap {
	o := OptionalDataGroupMap{}

	if DATAGROUPRECORDSIZE-1 > len(s) {
		return &o
	}

	o.RecordTypeCode = GroupMapExtensionRecord
	o.GroupByteI = s[0]
	o.GroupByteII = s[1]
	o.GroupByteIII = s[2]
	o.GroupByteIV = s[3]

	return &o
}

func (o *OptionalDataGroupMap) SetGroupBits(groups []GroupMapBitType, enable bool) {
	for _, grp := range groups {
		o.SetGroupBit(grp, enable)
	}
}

func (o *OptionalDataGroupMap) GetGroupBit(groupbit GroupMapBitType) bool {
	var group uint = uint(groupbit)

	//check to see if the field number is present in the Bit Map
	if 1 <= group && group <= 6 {
		return bitwise.HasBit(o.GroupByteIV, group-1)
	} else if 7 <= group && group <= 12 {
		group -= 6
		return bitwise.HasBit(o.GroupByteIII, group-1)
	} else if 13 <= group && group <= 18 {
		group -= 12
		return bitwise.HasBit(o.GroupByteII, group-1)
	} else if 19 <= group && group <= 24 {
		group -= 18
		return bitwise.HasBit(o.GroupByteI, group-1)
	}
	return false
}

func (o *OptionalDataGroupMap) SetGroupBit(groupbit GroupMapBitType, enable bool) {

	var group uint = uint(groupbit)
	//check to see if the field number is present in the Bit Map
	if 1 <= group && 6 >= group {
		setGroupBit(&o.GroupByteIV, group-1, enable)
	} else if 7 <= group && 12 >= group {
		group -= 6
		setGroupBit(&o.GroupByteIII, group-1, enable)
	} else if 13 <= group && 18 >= group {
		group -= 12
		setGroupBit(&o.GroupByteII, group-1, enable)
	} else if 19 <= group && 24 >= group {
		group -= 18
		setGroupBit(&o.GroupByteI, group-1, enable)
	}

}

func (o *OptionalDataGroupMap) String() string {
	return fmt.Sprintf("%c%c%c%c%c",
		o.RecordTypeCode,
		o.GroupByteI,
		o.GroupByteII,
		o.GroupByteIII,
		o.GroupByteIV)
}

func PrintBits(b byte) string {

	bitstates := []string{"P", "1", "0", "0", "0", "0", "0", "0"}
	var pos, n uint = 5, 2

	for ; ; pos-- {
		if true == bitwise.HasBit(b, pos) {
			bitstates[n] = "1"
		}
		n++
		if pos == 0 {
			break
		}

	}
	return strings.Join(bitstates[0:], " ")
}

func PrintGroupBits(b byte, startgroup uint) string {

	switch startgroup {
	case 1, 7, 13, 19, 25, 31, 37, 42, 28:
		break
	default:
		panic("valid startgroup values include groups 1, 7, 13, 19, 25, 31, 37, 42, 28")
	}

	bitstates := []string{"0", "0", "0", "0", "0", "0"}

	var pos uint = 0
	var groupno uint = 6

	for ; pos < 6; pos++ {
		if true == bitwise.HasBit(b, pos) {
			bitstates[groupno-1] = fmt.Sprintf("(%02d)", startgroup+pos)
		} else {
			bitstates[groupno-1] = fmt.Sprintf(" %02d ", startgroup+pos)
		}
		groupno--
	}
	return "[ " + strings.Join(bitstates[0:], " ") + " ]"
}
