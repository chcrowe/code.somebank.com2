package card

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CardIssuerType rune

const (
	EmptyCardType CardIssuerType = ' '
	Amex          CardIssuerType = 'A'
	Delta         CardIssuerType = 'T'
	Diners        CardIssuerType = 'D'
	Ets           CardIssuerType = 'G'
	Discover      CardIssuerType = 'I'
	Electron      CardIssuerType = 'E'
	Jcb           CardIssuerType = 'J'
	Laser         CardIssuerType = 'L'
	Mastercard    CardIssuerType = 'M'
	Maestro       CardIssuerType = 'O'
	Switch        CardIssuerType = 'W'
	Solo          CardIssuerType = 'S'
	Transax       CardIssuerType = 'X'
	Visa          CardIssuerType = 'V'
	Unrecognized  CardIssuerType = 'U'
)

func (ci CardIssuerType) String() string {
	switch ci {
	case EmptyCardType:
		return "Empty"
	case Amex:
		return "American Express"
	case Delta:
		return "Delta"
	case Diners:
		return "Diners"
	case Ets:
		return "SOMEBANK"
	case Discover:
		return "Discover"
	case Electron:
		return "Electron"
	case Jcb:
		return "JCB"
	case Laser:
		return "Laser"
	case Mastercard:
		return "Mastercard"
	case Maestro:
		return "Maestro"
	case Switch:
		return "Switch"
	case Solo:
		return "Solo"
	case Transax:
		return "Transax"
	case Visa:
		return "Visa"
	case Unrecognized:
		return "Unrecognized"
	}
	return "Unknown"
}

type CardEntryType uint

const (
	MagneticStripeReadTrack01Only     CardEntryType = 0
	MagneticStripeReadTrack02Only     CardEntryType = 1
	FullMagneticStripeRead            CardEntryType = 2
	ManualEntryCardPresent            CardEntryType = 3
	ManualEntryCardNotPresent         CardEntryType = 4
	ManualEntryCardNotPresentInternet CardEntryType = 5
)

func (ce CardEntryType) String() string {
	switch ce {
	case MagneticStripeReadTrack01Only:
		return "MagneticStripeReadTrack01Only"
	case MagneticStripeReadTrack02Only:
		return "MagneticStripeReadTrack02Only"
	case FullMagneticStripeRead:
		return "FullMagneticStripeRead"
	case ManualEntryCardPresent:
		return "ManualEntryCardPresent"
	case ManualEntryCardNotPresent:
		return "ManualEntryCardNotPresent"
	case ManualEntryCardNotPresentInternet:
		return "ManualEntryCardNotPresentInternet"
	}
	return ""
}

type CardAccountType uint

const (
	DefaultCard         CardAccountType = 0
	SavingsCard         CardAccountType = 1
	PrimaryCheckingCard CardAccountType = 2
)

func (ca CardAccountType) String() string {
	switch ca {
	case DefaultCard:
		return "DefaultCard"
	case SavingsCard:
		return "SavingsCard"
	case PrimaryCheckingCard:
		return "PrimaryCheckingCard"
	}
	return ""
}

type CreditCard struct {
	CustomerPresent bool
	CardType        CardIssuerType
	AccountType     CardAccountType
	EntryMethod     CardEntryType

	CardholderName  string
	AccountNumber   string
	ExpirationMonth uint
	ExpirationYear  uint

	TrackData string
	/*
		This is a variable length field with a maximum data length of 76-characters.
		The Track 1 data read from the cardholder's card is checked for parity and LRC errors and
		then converted from the six-bit characters encoded on the card to seven bit characters as
		defined in ANSI X3.4. The character set definitions are provided in Section 5.1 for reference.
		As part of the conversion, the terminal must remove the framing characters (start sentinel,
		end sentinel, and LRC characters). The separators must be converted to either an ASCII “^”
		(HEX 5E) or ASCII <US> (HEX 1F) characters. The entire UNALTERED Track
		(excluding framing characters) must be provided in the authorization request message or an
		error condition results.
		For American Express cards, any spaces in the card account number field of Track 1 must be
		removed before the application executes either a Mod-10 check of the card account number
		or transmits the data to the host for authorization or capture.
	*/
	Track01Data string

	/*
		This is a variable length field with a maximum data length of 37-characters.
		The Track 2 data read from the cardholder's card is checked for Parity and LRC errors and
		then converted from the four-bit characters encoded on the card to seven bit characters as
		defined in ANSI X3.4. The character set definitions are provided in Section 5.1 for reference.
		As part of the conversion, the terminal must remove the start sentinel, end sentinel, and LRC
		characters. The separators must be converted to either an ASCII "=" (HEX 3D) or ASCII
		"D" (HEX 44) characters. The entire UNALTERED Track (excluding framing characters)
		must be provided in the authorization request message or an error message is generated.
	*/
	Track02Data string
	Track03Data string

	AVSStreetAddress1 string
	AVSStreetAddress2 string
	AVSZipCode        string
	CVVData           string

	IssueNumber uint
	IssueMonth  uint
	IssueYear   uint

	ServiceCode       uint
	EncryptedPinBlock string
	DUKPTValue        string
}

func (c *CreditCard) Track01DataNoSentinals() string {
	if 0 < len(c.Track01Data) {
		return c.Track01Data[1 : len(c.Track01Data)-2]
	}
	return ""
}

func (c *CreditCard) Track02DataNoSentinals() string {
	if 0 < len(c.Track02Data) {
		re, err := regexp.Compile("[^0-9=]+")
		if err == nil {
			return re.ReplaceAllString(c.Track02Data, "")
		}
	}
	return ""
}

func CreditCardIsExpired(month int, year int) bool {
	now := time.Now()
	return month < int(now.Month()) && year <= now.Year()
}

func (c *CreditCard) SetAccountNumber(s string) {
	if 0 >= len(s) {
		return
	}

	accountno := ReplaceNonDigits(s)

	anlen := len(accountno)
	if 9 <= anlen && anlen < 21 {
		c.AccountNumber = accountno
		c.CardType = GetCardIssuerType(accountno)
	}
}

func (c *CreditCard) Copy(cc *CreditCard) {
	c.EntryMethod = cc.EntryMethod
	c.CardholderName = cc.CardholderName
	c.CardType = cc.CardType
	c.AccountNumber = cc.AccountNumber
	c.ExpirationMonth = cc.ExpirationMonth
	c.ExpirationYear = cc.ExpirationYear
	c.TrackData = cc.TrackData
	c.Track01Data = cc.Track01Data
	c.Track02Data = cc.Track02Data
	c.Track03Data = cc.Track03Data
	c.AVSStreetAddress1 = cc.AVSStreetAddress1
	c.AVSStreetAddress2 = cc.AVSStreetAddress2
	c.AVSZipCode = cc.AVSZipCode
	c.CVVData = cc.CVVData
	c.EncryptedPinBlock = cc.EncryptedPinBlock
	c.DUKPTValue = cc.DUKPTValue
	c.IssueNumber = cc.IssueNumber
	c.IssueMonth = cc.IssueMonth
	c.IssueYear = cc.IssueYear
}

func ParseCreditCard(magneticStripe string) *CreditCard {
	c := CreditCard{}

	c.TrackData = magneticStripe

	alltracks := strings.Split(c.TrackData, "?")
	for _, track := range alltracks {

		if 0 < len(track) {
			if '%' == track[0] {
				c.Track01Data = track + "?"
				c.EntryMethod = MagneticStripeReadTrack01Only

				trackElements := strings.Split(c.Track01Data, "^")
				if 2 < len(trackElements) && 3 < len(trackElements[2]) {
					c.AccountNumber = trackElements[0][2:]
					c.CardholderName = strings.Trim(trackElements[1], " ")
					fmt.Sscanf(trackElements[2][0:7], "%02d%02d%03d", &c.ExpirationYear, &c.ExpirationMonth, &c.ServiceCode)
				} else if 2 == len(trackElements) {
					c.AccountNumber = trackElements[0][2:]

					servCode, _ := strconv.Atoi(trackElements[1][4:7])
					c.ServiceCode = uint(servCode)
				}
			} else if ';' == track[0] {
				c.Track02Data = track + "?"

				if MagneticStripeReadTrack01Only == c.EntryMethod {
					c.EntryMethod = FullMagneticStripeRead
				} else {
					c.EntryMethod = MagneticStripeReadTrack02Only
				}

				trackElements := strings.Split(c.Track02Data, "=")
				if 1 < len(trackElements) && 8 < len(trackElements[0]) && 3 < len(trackElements[1]) {
					c.AccountNumber = trackElements[0][1:]
					fmt.Sscanf(trackElements[1][0:4], "%02d%02d", &c.ExpirationYear, &c.ExpirationMonth)
				}
			}
		}
	}
	return &c
}

func MaskCreditCardAccountNumber(accountno string, maskchar string) string {

	accountno = ReplaceNonDigits(accountno)
	acctlen := len(accountno)
	if 13 > acctlen {
		return ""
	}

	maskedlen := acctlen - 10
	masked := strings.Repeat(maskchar, maskedlen)
	return accountno[0:6] + masked + accountno[maskedlen+6:]
}

func GetCardIssuerDescription(accountno string) string {
	switch GetCardIssuerType(accountno) {
	case Amex:
		return "Amex"
	case Ets:
		return "SOMEBANK Rewards"
	case Diners:
		return "Diners"
	case Jcb:
		return "JCB"
	case Laser:
		return "Laser"
	case Maestro:
		return "Maestro"
	case Mastercard:
		return "M/C"
	case Visa:
		return "Visa"
	case Delta:
		return "Visa Delta"
	case Electron:
		return "Visa Electron"
	case Solo:
		return "Solo"
	case Switch:
		return "Switch"
	case Transax:
		return "Transax"
	default:
		return "Unknown"
	}
}

func GetCardIssuerType(accountno string) CardIssuerType {

	if 6 > len(accountno) {
		return EmptyCardType
	}

	accountno = ReplaceNonDigits(accountno)

	cardNumber, _ := strconv.ParseUint(accountno[0:6], 10, 64)
	if (cardNumber) >= 400000 && cardNumber <= 499999 {
		return Visa
	} else if (cardNumber >= 510000 && cardNumber <= 559999) || (cardNumber >= 360000 && cardNumber <= 369999) {
		return Mastercard
	} else if (cardNumber >= 340000 && cardNumber <= 349999) || (cardNumber >= 370000 && cardNumber <= 379999) {
		return Amex
	} else if cardNumber >= 601100 && cardNumber <= 601199 {
		return Discover
	} else if (cardNumber >= 300000 && cardNumber <= 305999) || (cardNumber >= 380000 && cardNumber <= 389999) {
		return Diners
	} else if cardNumber >= 352800 && cardNumber <= 358999 {
		return Jcb
	} else if cardNumber == 627571 {
		return Ets
	} else {
		return Unrecognized
	}

	//if( false == ConformsToLUHNFormulaMod10( accountno ) )
	//				return INVALIDCARDTYPE

	return Unrecognized
}

func GetMod10CheckSum(cardNumber string) int {
	// Array to contain individual numbers

	CardLength := len(cardNumber)
	CheckNumbers := make([]int, CardLength)
	// So, get length of card

	// Double the value of alternate digits, starting with the second digit
	// from the right, i.e. back to front.
	// Loop through starting at the end
	for i := CardLength - 2; i >= 0; i = i - 2 {
		// Double the numeric value returned
		altdigitval, _ := strconv.ParseUint(cardNumber[i:i+1], 10, 64)
		CheckNumbers[i] = int(altdigitval) * 2
	}

	CheckSum := 0 // Will hold the total sum of all checksum digits

	// Second stage, add separate digits of all products
	for iCount := 0; iCount <= CardLength-1; iCount++ {
		_count := 0 // will hold the sum of the digits

		// determine if current number has more than one digit
		if CheckNumbers[iCount] > 9 {
			_numstr := strconv.Itoa(CheckNumbers[iCount])
			_numLength := len(_numstr)
			// add count to each digit
			for x := 0; x < _numLength; x++ {
				currentnum, _ := strconv.ParseUint(_numstr[x:x+1], 10, 64)
				_count += int(currentnum)
			}
		} else {
			// single digit, just add it by itself
			_count = CheckNumbers[iCount]
		}
		CheckSum += _count // add sum to the total sum
	}
	// Stage 3, add the unaffected digits
	// Add all the digits that we didn't double still starting from the
	// right but this time we'll start from the rightmost number with
	// alternating digits
	OriginalSum := 0
	for y := CardLength - 1; y >= 0; y -= 2 {
		orignum, _ := strconv.ParseUint(cardNumber[y-1:y], 10, 64)
		OriginalSum += int(orignum)
	}

	// Perform the final calculation, if the sum Mod 10 results in 0 then
	// it's valid, otherwise its false.
	return (OriginalSum + CheckSum)
}

func IsValidCardLength(cdtype CardIssuerType, length int) bool {
	switch cdtype {
	case Amex:
		return (15 == length)
	case Diners:
		return (14 == length || 16 == length)

	case Discover:
	case Ets:
	case Electron:
	case Jcb:
	case Maestro:
	case Mastercard:
		return (16 == length)

	case Delta:
	case Visa:
		return (13 == length || 16 == length)

	case Laser:
	case Solo:
	case Switch:
		return (16 == length || 18 == length || 19 == length)

	case Transax:
	default:
		return true
	}
	return false
}

func ValidateCardNumber(cardNumber string) bool {
	cdtype := GetCardIssuerType(cardNumber)
	if Diners == cdtype {
		return false
	}
	return IsValidCardLength(cdtype, len(cardNumber)) && ((GetMod10CheckSum(cardNumber) % 10) == 0)
}

func ReplaceNonDigits(s string) string {
	if 0 < len(s) {
		re, err := regexp.Compile("[^0-9]+")
		if err == nil {
			return re.ReplaceAllString(s, "")
		}
	}
	return s
}
