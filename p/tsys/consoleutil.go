package tsys

import (
	"fmt"
	"strings"

	"code.somebank.com/p/bytes"
)

func WrapAndPrintText(text string, caption string, width int) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if 0 < len(line) {
			wrappedtextarray := WrapString(line, width)
			for _, wrappedtext := range wrappedtextarray {
				fmt.Printf("%-32s%s\n", caption, wrappedtext)
				caption = ""
			}
		}
	}
}

func WrapString(s string, maxline int) []string {
	totallength := len(s)
	linesneeded := totallength / maxline
	extralines := totallength % maxline
	if 0 < extralines {
		linesneeded++
	}

	wrapped := make([]string, linesneeded)

	for i := 0; i < linesneeded; i++ {
		start := i * maxline
		end := start + maxline
		if end < totallength {
			wrapped[i] = s[start:end]
		} else {
			wrapped[i] = s[start:]
		}
	}

	return wrapped
}

func DumpAuthRequest(group1 *G1AuthorizationMessageRequest, group2 *G2AuthorizationMessageRequest, gp3records []string, s string) {

	fmt.Printf("AUTHORIZATION REQUEST ANALYSIS\n\n")
	WrapAndPrintText(s, "Actual packet ->", 128)
	fmt.Println("")
	WrapAndPrintText(fmt.Sprintf("%+v\n\n", group1.String()), "G1 packet ->", 128)

	g1fields := strings.Split(group1.VerboseString(), "\n")
	caption := "G1 (expanded) ->"
	for _, g1f := range g1fields {
		fmt.Printf("%-32s%s\n", caption, g1f)
		caption = ""
	}

	WrapAndPrintText(strings.Join(gp3records, string(bytes.GroupSeparator)), "G3 packet ->", 128)
	fmt.Println("")

	caption = "G3 (expanded) ->"

	var g3ver G3VersionType
	for k, g3 := range gp3records {

		if 0 < k {
			caption = ""
			fmt.Printf("%-32s%c (GS=%[2]d)\n", "", bytes.GroupSeparator)
		}

		fmt.Sscanf(g3, "%03d", &g3ver)
		switch g3ver {

		case G3v001CommercialCard:
			G3v001 := ParseG3v001CommercialCardRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v001), caption, 128)

		case G3v007CardVerificationCode:
			G3v007 := ParseG3v007CardVerificationCodeRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v007), caption, 128)

		case G3v019MastercardUniversalCardholderAuthentication:
			G3v019 := ParseG3v019MastercardUniversalCardholderAuthenticationRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v019), caption, 128)

		case G3v020DeveloperInformation:
			G3v020 := ParseG3v020DeveloperInformationRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v020), caption, 128)

		case G3v021MerchantVerificationValue:
			G3v021 := ParseG3v021MerchantVerificationValueRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v021), caption, 128)

		case G3v022AdditionalAmounts:
			G3v022 := ParseG3v022AdditionalAmountsRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v022), caption, 128)
		case G3v025TransactionFeeAmount:
			G3v025 := ParseG3v025TransactionFeeAmountRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v025), caption, 128)
		case G3v027PosData:
			G3v027 := ParseG3v027PosDataRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s\n%s", g3, G3v027.VerboseString()), caption, 128)
		case G3v032CurrencyConversionData:
			G3v032 := ParseG3v032CurrencyConversionDataRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v032), caption, 128)

		case G3v034CardProductCode:
			G3v034 := ParseG3v034CardProductCodeRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v034), caption, 128)

		case G3v040VisaISAChargeIndicator:
			G3v040 := ParseG3v040VisaISAChargeIndicatorRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v040), caption, 128)
		case G3v046CardTypeGroup:
			G3v046 := ParseG3v046CardTypeGroupRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v046), caption, 128)
		case G3v050AssociationTimestamp:
			G3v050 := ParseG3v050AssociationTimestampRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v050), caption, 128)

		case G3v051MasterCardEventMonitoringRealTimeScoring:
			G3v051 := ParseG3v051MasterCardEventMonitoringRealTimeScoringRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v051), caption, 128)

		case G3v055IntegratedChipCardEmvTlv:
			G3v055 := ParseG3v055IntegratedChipCardEmvTlvRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v055), caption, 128)

		case G3v058AlternateAccountID:
			G3v058 := ParseG3v058AlternateAccountIDRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v058), caption, 128)
		case G3v059MasterCardPayPassMappingService:
			G3v059 := ParseG3v059MasterCardPayPassMappingServiceRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v059), caption, 128)

		case G3v061SpendQualifiedIndicator:
			G3v061 := ParseG3v061SpendQualifiedIndicatorRequest(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v061), caption, 128)

		default:
			WrapAndPrintText(fmt.Sprintf("%s %s", g3, g3ver), caption, 128)

		}

	}
	fmt.Println()
	fmt.Println()

}

func DumpAuthResponse(group1 *G1AuthorizationMessageResponse, group2 *G2AuthorizationMessageResponse, gp3records []string, s string) {

	fmt.Printf("AUTHORIZATION RESPONSE ANALYSIS\n\n")
	WrapAndPrintText(s, "Actual packet ->", 128)
	fmt.Println("")
	WrapAndPrintText(fmt.Sprintf("%+v\n\n", group1.String()), "G1 packet ->", 128)

	g1fields := strings.Split(group1.VerboseString(), "\n")
	caption := "G1 fields ->"
	for _, g1f := range g1fields {
		fmt.Printf("%-32s%s\n", caption, g1f)
		caption = ""
	}

	WrapAndPrintText(strings.Join(gp3records, string(bytes.GroupSeparator)), "G3 packet ->", 128)
	fmt.Println("")

	caption = "G3 records ->"

	var g3ver G3VersionType
	for k, g3 := range gp3records {

		if 0 < k {
			caption = ""
			fmt.Printf("%-32s%c (GS=%[2]d)\n", "", bytes.GroupSeparator)
		}

		fmt.Sscanf(g3, "%03d", &g3ver)
		switch g3ver {

		case G3v001CommercialCard:
			G3v001 := ParseG3v001CommercialCardResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v001), caption, 128)

		case G3v007CardVerificationCode:
			G3v007 := ParseG3v007CardVerificationCodeResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v007), caption, 128)

		case G3v019MastercardUniversalCardholderAuthentication:
			G3v019 := ParseG3v019MastercardUniversalCardholderAuthenticationResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v019), caption, 128)

		case G3v020DeveloperInformation:
			G3v020 := ParseG3v020DeveloperInformationResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v020), caption, 128)

		case G3v021MerchantVerificationValue:
			G3v021 := ParseG3v021MerchantVerificationValueResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v021), caption, 128)

		case G3v022AdditionalAmounts:
			G3v022 := ParseG3v022AdditionalAmountsResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v022), caption, 128)
		case G3v025TransactionFeeAmount:
			G3v025 := ParseG3v025TransactionFeeAmountResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v025), caption, 128)
		case G3v027PosData:
			G3v027 := ParseG3v027PosDataResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v027), caption, 128)
		case G3v032CurrencyConversionData:
			G3v032 := ParseG3v032CurrencyConversionDataResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v032), caption, 128)

		case G3v034CardProductCode:
			G3v034 := ParseG3v034CardProductCodeResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v034), caption, 128)

		case G3v040VisaISAChargeIndicator:
			G3v040 := ParseG3v040VisaISAChargeIndicatorResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v040), caption, 128)
		case G3v046CardTypeGroup:
			G3v046 := ParseG3v046CardTypeGroupResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v046), caption, 128)
		case G3v050AssociationTimestamp:
			G3v050 := ParseG3v050AssociationTimestampResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v050), caption, 128)

		case G3v051MasterCardEventMonitoringRealTimeScoring:
			G3v051 := ParseG3v051MasterCardEventMonitoringRealTimeScoringResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v051), caption, 128)

		case G3v058AlternateAccountID:
			G3v058 := ParseG3v058AlternateAccountIDResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v058), caption, 128)
		case G3v059MasterCardPayPassMappingService:
			G3v059 := ParseG3v059MasterCardPayPassMappingServiceResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v059), caption, 128)

		case G3v061SpendQualifiedIndicator:
			G3v061 := ParseG3v061SpendQualifiedIndicatorResponse(g3)
			WrapAndPrintText(fmt.Sprintf("%s %+v", g3, *G3v061), caption, 128)

		default:
			WrapAndPrintText(fmt.Sprintf("%s %s", g3, g3ver), caption, 128)

		}

	}
	fmt.Println()
	fmt.Println()

}
