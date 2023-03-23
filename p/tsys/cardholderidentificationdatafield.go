package tsys

import (
	"fmt"
	"strings"

	"code.somebank.com/p/bytes"
)

type CardholderIdentificationDataField struct {
	CardholderIdCode CardholderIdCodeType
	AVSStreetAddress string
	AVSZipCode       string

	EncryptedPinBlock string
	DUKPTValue        string
}

func NewCardholderIdentificationDataField(idcode CardholderIdCodeType, street string, zipcode string, pin string, dukpt string) *CardholderIdentificationDataField {
	return &CardholderIdentificationDataField{idcode, street, zipcode, pin, dukpt}
}

func ParseCardholderIdentificationDataField(idcode CardholderIdCodeType, s string) (*CardholderIdentificationDataField, int) {

	cd := CardholderIdentificationDataField{CardholderIdCode: idcode}

	charsread := strings.Index(s, string(bytes.FileSeparator))
	if 0 >= charsread {
		charsread = len(s)
	}

	if charsread >= 128+1 {
		panic("maximum cardholder identification data field size cannot exceed 128 characters")
	}

	switch idcode {
	case CardPresentUnableToReadMagStripeSendingAvsData: // 'M'
		cd.AVSZipCode = s

	case CardNotPresentIncludesEcomAndFullAvsData: // 'N'
		if 0 < len(s) {
			AVSElements := strings.Split(s, " ")
			if nil != AVSElements && 0 < len(AVSElements) {
				if 1 == len(AVSElements) {
					cd.AVSZipCode = AVSElements[0]
				} else if 1 < len(AVSElements) {
					cd.AVSStreetAddress = AVSElements[0]
					cd.AVSZipCode = AVSElements[1]
				}
			}
		}

	case PersonalIdentificationNumber32CharDukpt:
		if 32 <= len(s) {
			cd.EncryptedPinBlock = s[0:16]
			cd.DUKPTValue = s[16:]
		}
	}

	return &cd, charsread
}

func (cd *CardholderIdentificationDataField) String() string {

	switch cd.CardholderIdCode {
	case CardPresentUnableToReadMagStripeSendingAvsData: // 'M'
		return cd.AVSZipCode

	case CardNotPresentIncludesEcomAndFullAvsData: // 'N'
		return fmt.Sprintf("%s %s", cd.AVSStreetAddress, cd.AVSZipCode)

	case PersonalIdentificationNumber32CharDukpt:
		return fmt.Sprintf("%s%s", cd.EncryptedPinBlock, cd.DUKPTValue)

	}
	return ""
}

func (cd *CardholderIdentificationDataField) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	switch cd.CardholderIdCode {
	case CardPresentUnableToReadMagStripeSendingAvsData: // 'M'
		buffer.AppendFormat("%8s%-32s%s\n", " ", "AVSZipCode", cd.AVSZipCode)

	case CardNotPresentIncludesEcomAndFullAvsData: // 'N'
		buffer.AppendFormat("%8s%-32s%s\n", " ", "AVSStreetAddress1", cd.AVSStreetAddress)
		buffer.AppendFormat("%8s%-32s%c (%[3]d)\n", " ", "Space (Field Delimeter)", ' ')
		buffer.AppendFormat("%8s%-32s%s\n", " ", "AVSZipCode", cd.AVSZipCode)

	case PersonalIdentificationNumber32CharDukpt:
		buffer.AppendFormat("%8s%-32s%s\n", " ", "EncryptedPinBlock", cd.EncryptedPinBlock)
		buffer.AppendFormat("%8s%-32s%s\n", " ", "DUKPTValue", cd.DUKPTValue)

	}
	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}
