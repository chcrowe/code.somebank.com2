package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
	"code.somebank.com/p/card/msr"

	//	"regexp"
	"strings"
)

/// <summary>
///
/*
	4.31 Customer data field, Group I, EIS 1080 Authorization (Request)
	This is a variable length field containing the customer account or check acceptance ID data in
	one of three formats. The cardholder account information can be read from one of two
	magnetic stripes on the card or it can contain data which has been entered manually.


	Security features include the use of encryption and tokens. Encryption allows the transaction
	to hold encrypted data which ensures transmission security. Tokens can be requested in a
	transaction and returned with a response. The use of tokens does not enhance transmission
	security but it allows a receiver to store unique card-type data that then reduces PCI Audit
	Scope. Gen 2 Terminal Authentication (G3V049) is required in order to use either security
	feature.

	Encryption
	When including encrypted data in the Customer Data Field (4.62), this needs to be identified
	in the Record Format (4.128). For Credit transactions, use Record Format 'W'. For Debit
	transactions, use Record Format 'X'. Group III Version 052 (3.2.48) will be required to include
	the Encryption Type and Encryption Transmission Block so that the data may be decrypted.
	Refer to details later in this document.

	Tokens
	When requesting a token, Group III Version 053 (3.2.49) is required and the token will be sent
	in the response. Use of Transaction Code '5T' will send a token in the response message
	without performing any card verification. Use of G3V053 with any other Transaction Code
	(4.163) performs the usual card verification or authorization request as specified. Refer to
	details later in this document.

	Purchase with a Token
	In general, a purchase with a token is used for processing recurring payments for existing debt
	or repeat business. The token must have been requested previously with a valid PAN. When
	using a token in place of the PAN, use Record Format 'Y' for credit transactions and Record
	Format 'Z' for debit transactions. This functionality is not available for card present scenarios.
	For these types of authorization requests, the source of the customer data entered must be
	'manually keyed', the Cardholder ID Code must indicate the transaction is CNP including e-
	Commerce and Full Address Verification Data (CID = 'N'). The token should be passed in the
	Customer Data Field. Please refer to Section 6.5 for an example. The transaction codes
	acceptable for use in a token purchase are indicated as such in Table 4.73. Record Format 'Y'
	and 'Z' transactions should never contain a request for token (Token Request/Response
	(G3v053)). This type of transaction would be rejected.
*/
/// </summary>
type CustomerDataField struct {
	AccountDataSource AccountDataSourceType

	MSRTrack01DataWithoutSentinals string
	MSRTrack02DataWithoutSentinals string

	ManuallyEnteredAccountNumber    string
	ManuallyEnteredExpirationMonth  uint
	ManuallyEnteredExpirationYear   uint
	ManuallyEnteredVerificationCode string

	EncryptedMSRTrackData string
	MICRData              string
}

func NewManuallyEnteredCustomerDataField(datasrc AccountDataSourceType, pan string, month uint, year uint, securitycode string) *CustomerDataField {
	return &CustomerDataField{AccountDataSource: datasrc, ManuallyEnteredAccountNumber: pan, ManuallyEnteredExpirationMonth: month, ManuallyEnteredExpirationYear: year, ManuallyEnteredVerificationCode: securitycode}
}

func NewCustomerDataField(datasrc AccountDataSourceType, track1 string, track2 string) *CustomerDataField {
	return &CustomerDataField{AccountDataSource: datasrc, MSRTrack01DataWithoutSentinals: track1, MSRTrack02DataWithoutSentinals: track2}
}

func NewEncryptedCustomerDataField(datasrc AccountDataSourceType, trackdata string) *CustomerDataField {
	return &CustomerDataField{AccountDataSource: datasrc, EncryptedMSRTrackData: trackdata}
}

func (cd *CustomerDataField) Sterilize() {
	if 0 < len(cd.ManuallyEnteredAccountNumber) {
		cd.ManuallyEnteredAccountNumber = msr.MaskMSRAccountNumber(cd.ManuallyEnteredAccountNumber, "9")
	}

	if 0 < cd.ManuallyEnteredExpirationMonth {
		cd.ManuallyEnteredExpirationMonth = 9
	}
	if 0 < cd.ManuallyEnteredExpirationYear {
		cd.ManuallyEnteredExpirationYear = 99
	}
	cd.ManuallyEnteredVerificationCode = ""
	// cd.EncryptedMSRTrackData = ""

	if 0 < len(cd.MSRTrack01DataWithoutSentinals) {
		msr1 := msr.ParseMagneticStripeTrackI(cd.MSRTrack01DataWithoutSentinals)
		cd.MSRTrack01DataWithoutSentinals = msr.MaskMagneticStripeTrack(cd.MSRTrack01DataWithoutSentinals, msr1.AccountNumber)
	}
	if 0 < len(cd.MSRTrack02DataWithoutSentinals) {
		msr2 := msr.ParseMagneticStripeTrackII(cd.MSRTrack02DataWithoutSentinals)
		cd.MSRTrack02DataWithoutSentinals = msr.MaskMagneticStripeTrack(cd.MSRTrack02DataWithoutSentinals, msr2.AccountNumber)
	}

}

func (cd *CustomerDataField) AccountNumber() string {

	if 0 < len(cd.ManuallyEnteredAccountNumber) {
		return cd.ManuallyEnteredAccountNumber
	} else if 0 < len(cd.MSRTrack01DataWithoutSentinals) {
		msr1 := msr.ParseMagneticStripeTrackI(cd.MSRTrack01DataWithoutSentinals)
		return msr1.AccountNumber
	} else if 0 < len(cd.MSRTrack02DataWithoutSentinals) {
		msr2 := msr.ParseMagneticStripeTrackII(cd.MSRTrack02DataWithoutSentinals)
		return msr2.AccountNumber
	}
	return ""
}

func (cd *CustomerDataField) String() string {
	switch cd.AccountDataSource {
	case MicrDataWasAcquiredByOcrReader:
		//MICR Data
		return fmt.Sprintf("%c%c%s", bytes.FileSeparator, bytes.FileSeparator, cd.MICRData)

	case FullMagneticStripeReadAndTransmitTrack1:
		//full track 1
		return cd.MSRTrack01DataWithoutSentinals

	case FullMagneticStripeReadAndTransmitTrack2, ChipCardReadDataCCPS, ProximityPaymentUsingEmvRules:
		//full track 2
		return cd.MSRTrack02DataWithoutSentinals

	// case ManuallyKeyedTerminalHasNoCardReadingCapability:
	// case ManuallyKeyedTrack1Capable:
	// case ManuallyKeyedTrack2Capable:
	default:
		return fmt.Sprintf("%s%c%02d%02d%c", cd.ManuallyEnteredAccountNumber, bytes.FileSeparator, cd.ManuallyEnteredExpirationMonth, cd.ManuallyEnteredExpirationYear, bytes.FileSeparator)
	}

	return ""
}

func (cd *CustomerDataField) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	switch cd.AccountDataSource {
	case MicrDataWasAcquiredByOcrReader:
		//MICR Data

		buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "FS", bytes.FileSeparator)
		buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "FS", bytes.FileSeparator)
		buffer.AppendFormat("%8s%-32s%s\n", " ", "MICRData", cd.MICRData)

	case FullMagneticStripeReadAndTransmitTrack1:
		//full track 1
		buffer.AppendFormat("%8s%-32s%s\n", " ", "MSRTrack01DataWithoutSentinals", cd.MSRTrack01DataWithoutSentinals)

	case FullMagneticStripeReadAndTransmitTrack2, ChipCardReadDataCCPS, ProximityPaymentUsingEmvRules:
		//full track 2
		buffer.AppendFormat("%8s%-32s%s\n", " ", "MSRTrack02DataWithoutSentinals", cd.MSRTrack02DataWithoutSentinals)

	// case ManuallyKeyedTerminalHasNoCardReadingCapability:
	// case ManuallyKeyedTrack1Capable:
	// case ManuallyKeyedTrack2Capable:
	default:

		buffer.AppendFormat("%8s%-32s%s\n", " ", "ManuallyEnteredAccountNumber", cd.ManuallyEnteredAccountNumber)
		buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "FS", bytes.FileSeparator)
		buffer.AppendFormat("%8s%-32s%02d\n", " ", "ManuallyEnteredExpirationMonth", cd.ManuallyEnteredExpirationMonth)
		buffer.AppendFormat("%8s%-32s%02d\n", " ", "ManuallyEnteredExpirationYear", cd.ManuallyEnteredExpirationYear)
		buffer.AppendFormat("%8s%-32s%c (%[3]s)\n", " ", "FS", bytes.FileSeparator)
	}

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}

func ParseCustomerDataField(datasrc AccountDataSourceType, s string) (*CustomerDataField, int) {

	cd := CustomerDataField{AccountDataSource: datasrc}

	customerdatafields := strings.Split(s, string(bytes.FileSeparator))
	charsread := len(customerdatafields[0])
	fmt.Printf("\n\nParseCustomerDataField-> customerdatafields[0]: %s\n\n", customerdatafields[0])

	switch cd.AccountDataSource {
	case MicrDataWasAcquiredByOcrReader:
		//MICR Data
		//cd.MICRData = DataParts[2]
		//iSeparators = 3 //including the two leading <FS>'s
		break
	case FullMagneticStripeReadAndTransmitTrack1:
		//full track 1
		//cd.ParseTrackData("%" + DataParts[0] + "?")
		cd.MSRTrack01DataWithoutSentinals = customerdatafields[0]
		break
	case FullMagneticStripeReadAndTransmitTrack2:
		//full track 2
		//cd.ParseTrackData(";" + DataParts[0] + "?")
		cd.MSRTrack02DataWithoutSentinals = customerdatafields[0]
		break
	case ChipCardReadDataCCPS, ProximityPaymentUsingEmvRules:
		//result := regexp.MustCompile("[=]|D").FindString(customerdatafields[0])
		//if 0 == len(result) {
		//full track 2
		cd.MSRTrack02DataWithoutSentinals = customerdatafields[0]
		//}
		break

		//	case ManuallyKeyedTerminalHasNoCardReadingCapability:
		//	case ManuallyKeyedTrack1Capable:
		//	case ManuallyKeyedTrack2Capable:
	default:
		if 1 < len(customerdatafields) {
			cd.ManuallyEnteredAccountNumber = customerdatafields[0]
			fmt.Sscanf(customerdatafields[1], "%02d%02d", &cd.ManuallyEnteredExpirationMonth, &cd.ManuallyEnteredExpirationYear)
			charsread += len(customerdatafields[1]) + 2
		}
		break
	}

	return &cd, charsread
}
