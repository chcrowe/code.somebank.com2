package tsys

// Mandatory EMV tags

type EmvTlvIdType string

const (
	TlvDF79_KernalVersionNumber EmvTlvIdType = "DF79" // The version number of the kernel used to process the chip data in the transaction. (variable 1 - 32)
	TlvDF78_DeviceSerialNumber  EmvTlvIdType = "DF78" // The manufacturer's unique serial number of the device that interacts with the chip card. (variable 1 - 20)

	// Summary chip card transaction data elements

	Tlv9F40_AdditionalTerminalCapabilities EmvTlvIdType = "9F40" // Indicates the data input and output capabilities of the terminal.
	Tlv9F02_AuthorizedAmount               EmvTlvIdType = "9F02" // Authorized amount of the transaction (excluding adjustments).
	Tlv9F03_SecondaryAmount                EmvTlvIdType = "9F03" // Other Secondary amount associated with the transaction representing a cashback amount
	Tlv9F26_ApplicationCryptogram          EmvTlvIdType = "9F26" // Cryptogram returned by the ICC in response of the GENERATE AC command.
	Tlv4F_ICCApplicationIdentifier         EmvTlvIdType = "4F"   // (AID) - ICC Identifies the application as described in ISO/IEC 7816-5.
	Tlv9F06_TerminalApplicationIdentifier  EmvTlvIdType = "9F06" // (AID) - Terminal Identifies the application as described in ISO/IEC 7816-5.
	Tlv82_ApplicationInterchangeProfile    EmvTlvIdType = "82"   // Mnemonic associated with the AID according to ISO/IEC 7816-5.

	Tlv9F36_ApplicationTransactionCounter EmvTlvIdType = "9F36" // Counter maintained by the application in the ICC(incrementing the ATC is managed by the ICC).
	Tlv9F34_CardholderVerificationMethod  EmvTlvIdType = "9F34" // (CVM) Results Identifies a method of verification of the cardholder supported by the application.
	Tlv9F27_CryptogramInformationData     EmvTlvIdType = "9F27" // Indicates the type of cryptogram and the actions to be performed by the terminal.
	Tlv9F39_POSEntryMode                  EmvTlvIdType = "9F39" // Indicates the method by which the PAN was entered, according to the first two digits of the ISO 8583:1987 POS Entry Mode.
	Tlv9F33_TerminalCapabilities          EmvTlvIdType = "9F33" // Indicates the card data input, CVM, and security capabilities of the terminal.
	Tlv9F1A_TerminalCountryCode           EmvTlvIdType = "9F1A" // Indicates the country of the terminal, represented according to ISO 3166.
	Tlv9F35_TerminalType                  EmvTlvIdType = "9F35" // Indicates the environment of the terminal, its communications capability, and its operational control.
	Tlv95_TerminalVerificationResults     EmvTlvIdType = "95"   // Status of the different functions as seen from the terminal.
	Tlv5F2A_TransactionCurrencyCode       EmvTlvIdType = "5F2A" // Indicates the currency code of the transaction according to ISO 4217.
	Tlv9A_TransactionDate                 EmvTlvIdType = "9A"   // Local date that the transaction was authorized.
	Tlv9B_TransactionStatus               EmvTlvIdType = "9B"   // Information Indicates the functions performed in a transaction.
	Tlv9F21_TransactionTime               EmvTlvIdType = "9F21" // Local time that the transaction was authorized.
	Tlv9C_TransactionType                 EmvTlvIdType = "9C"   // Indicates the type of financial transaction, represented by the first two digits of ISO 8583:1987 Processing Code.
	Tlv9F37_UnpredictableNumber           EmvTlvIdType = "9F37" // Value to provide variability and uniqueness to the generation of a cryptogram.
	Tlv5F2D_LanguagePreference            EmvTlvIdType = "5F2D" // 1-4 languages stored in order of preference, each represented by 2 alphabetical characters according to ISO 639.
	Tlv91_IssuerAuthenticationData        EmvTlvIdType = "91"   // Data sent to the ICC for online issuer authentication.
	Tlv5F34_PrimaryAccountNumber          EmvTlvIdType = "5F34" // (PAN) Sequence Number Identifies and differentiates cards with the same PAN.
	Tlv84_DedicatedFile                   EmvTlvIdType = "84"   // (DF) Name Identifies the name of the DF as described in ISO/ IEC 7816-4.
	Tlv9F10_IssuerApplicationData         EmvTlvIdType = "9F10" // Contains proprietary application data for transmission to the issuer in an online transaction.
	Tlv9F5B_IssuerScriptResults           EmvTlvIdType = "9F5B" // Indicates the result of the terminal script processing.
	Tlv9F0D_IssuerActionCodeDefault       EmvTlvIdType = "9F0D" // Specifies the issuer's conditions that cause a transaction to be rejected if it might have been approved online, but the terminal is unable to process the transaction online. (Capture only)
	Tlv9F0E_IssuerActionCodeDefault       EmvTlvIdType = "9F0E" // Specifies the issuer's conditions that cause the denial of a transaction without attempt to go online. (Capture only)
	Tlv9F0F_IssuerActionCodeOnline        EmvTlvIdType = "9F0F" // Specifies the issuer's conditions that cause a transaction to be transmitted online. (Capture only)

	// EMV tags that should not be sent

	Tlv5A_PrimaryAccountNumber        EmvTlvIdType = "5A"
	Tlv57_Track2EquivalentData        EmvTlvIdType = "57"
	Tlv5F24_ApplicationExpirationDate EmvTlvIdType = "5F24"
	Tlv5F30_ServiceCode               EmvTlvIdType = "5F30"
	Tlv9F1F_Track1DiscretionaryData   EmvTlvIdType = "9F1F"
	Tlv9F20_Track2DiscretionaryData   EmvTlvIdType = "9F20"
)

func (x EmvTlvIdType) String() string {
	switch x {
	case TlvDF79_KernalVersionNumber:
		return "TlvDF79_KernalVersionNumber"
	case TlvDF78_DeviceSerialNumber:
		return "TlvDF78_DeviceSerialNumber"
	case Tlv9F40_AdditionalTerminalCapabilities:
		return "Tlv9F40_AdditionalTerminalCapabilities"
	case Tlv9F02_AuthorizedAmount:
		return "Tlv9F02_AuthorizedAmount"
	case Tlv9F03_SecondaryAmount:
		return "Tlv9F03_SecondaryAmount"
	case Tlv9F26_ApplicationCryptogram:
		return "Tlv9F26_ApplicationCryptogram"
	case Tlv4F_ICCApplicationIdentifier:
		return "Tlv4F_ICCApplicationIdentifier"
	case Tlv9F06_TerminalApplicationIdentifier:
		return "Tlv9F06_TerminalApplicationIdentifier"
	case Tlv82_ApplicationInterchangeProfile:
		return "Tlv82_ApplicationInterchangeProfile"
	case Tlv9F36_ApplicationTransactionCounter:
		return "Tlv9F36_ApplicationTransactionCounter"
	case Tlv9F34_CardholderVerificationMethod:
		return "Tlv9F34_CardholderVerificationMethod"
	case Tlv9F27_CryptogramInformationData:
		return "Tlv9F27_CryptogramInformationData"
	case Tlv9F39_POSEntryMode:
		return "Tlv9F39_POSEntryMode"
	case Tlv9F33_TerminalCapabilities:
		return "Tlv9F33_TerminalCapabilities"
	case Tlv9F1A_TerminalCountryCode:
		return "Tlv9F1A_TerminalCountryCode"
	case Tlv9F35_TerminalType:
		return "Tlv9F35_TerminalType"
	case Tlv95_TerminalVerificationResults:
		return "Tlv95_TerminalVerificationResults"
	case Tlv5F2A_TransactionCurrencyCode:
		return "Tlv5F2A_TransactionCurrencyCode"
	case Tlv9A_TransactionDate:
		return "Tlv9A_TransactionDate"
	case Tlv9B_TransactionStatus:
		return "Tlv9B_TransactionStatus"
	case Tlv9F21_TransactionTime:
		return "Tlv9F21_TransactionTime"
	case Tlv9C_TransactionType:
		return "Tlv9C_TransactionType"
	case Tlv9F37_UnpredictableNumber:
		return "Tlv9F37_UnpredictableNumber"
	case Tlv5F2D_LanguagePreference:
		return "Tlv5F2D_LanguagePreference"
	case Tlv91_IssuerAuthenticationData:
		return "Tlv91_IssuerAuthenticationData"
	case Tlv5F34_PrimaryAccountNumber:
		return "Tlv5F34_PrimaryAccountNumber"
	case Tlv84_DedicatedFile:
		return "Tlv84_DedicatedFile"
	case Tlv9F10_IssuerApplicationData:
		return "Tlv9F10_IssuerApplicationData"
	case Tlv9F5B_IssuerScriptResults:
		return "Tlv9F5B_IssuerScriptResults"
	case Tlv9F0D_IssuerActionCodeDefault:
		return "Tlv9F0D_IssuerActionCodeDefault"
	case Tlv9F0E_IssuerActionCodeDefault:
		return "Tlv9F0E_IssuerActionCodeDefault"
	case Tlv9F0F_IssuerActionCodeOnline:
		return "Tlv9F0F_IssuerActionCodeOnline"
	case Tlv5A_PrimaryAccountNumber:
		return "Tlv5A_PrimaryAccountNumber"
	case Tlv57_Track2EquivalentData:
		return "Tlv57_Track2EquivalentData"
	case Tlv5F24_ApplicationExpirationDate:
		return "Tlv5F24_ApplicationExpirationDate"
	case Tlv5F30_ServiceCode:
		return "Tlv5F30_ServiceCode"
	case Tlv9F1F_Track1DiscretionaryData:
		return "Tlv9F1F_Track1DiscretionaryData"
	case Tlv9F20_Track2DiscretionaryData:
		return "Tlv9F20_Track2DiscretionaryData"
	}
	return ""
}

// Tlv4F_ICCApplicationIdentifier
// Tlv5F2A_TransactionCurrencyCode
// Tlv5F2D_LanguagePreference
// Tlv5F34_PrimaryAccountNumber
// Tlv82_ApplicationInterchangeProfile
// Tlv84_DedicatedFile
// Tlv91_IssuerAuthenticationData
// Tlv95_TerminalVerificationResults
// Tlv9A_TransactionDate
// Tlv9B_TransactionStatus
// Tlv9C_TransactionType
// Tlv9F02_AuthorizedAmount
// Tlv9F03_SecondaryAmount
// Tlv9F06_TerminalApplicationIdentifier
// Tlv9F0D_IssuerActionCodeDefault
// Tlv9F0E_IssuerActionCodeDefault
// Tlv9F0F_IssuerActionCodeOnline
// Tlv9F10_IssuerApplicationData
// Tlv9F1A_TerminalCountryCode
// Tlv9F21_TransactionTime
// Tlv9F26_ApplicationCryptogram
// Tlv9F27_CryptogramInformationData
// Tlv9F33_TerminalCapabilities
// Tlv9F34_CardholderVerificationMethod
// Tlv9F35_TerminalType
// Tlv9F36_ApplicationTransactionCounter
// Tlv9F37_UnpredictableNumber
// Tlv9F39_POSEntryMode
// Tlv9F40_AdditionalTerminalCapabilities
// Tlv9F5B_IssuerScriptResults
// TlvDF78_DeviceSerialNumber
// TlvDF79_KernalVersionNumber

// // EMV tags that should not be sent

// Tlv5A_PrimaryAccountNumber
// Tlv57_Track2EquivalentData
// Tlv5F24_ApplicationExpirationDate
// Tlv5F30_ServiceCode
// Tlv9F1F_Track1DiscretionaryData
// Tlv9F20_Track2DiscretionaryData
