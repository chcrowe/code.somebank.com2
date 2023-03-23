package bytes

import (
	"fmt"
	"strings"
)

type AsciiCharacterType uint

const (
	Null                   AsciiCharacterType = 0
	StartOfHeading         AsciiCharacterType = 1
	StartOfText            AsciiCharacterType = 2
	EndOfText              AsciiCharacterType = 3
	EndOfTransmission      AsciiCharacterType = 4
	Enquiry                AsciiCharacterType = 5
	Acknowledged           AsciiCharacterType = 6
	Bell                   AsciiCharacterType = 7
	BackSpace              AsciiCharacterType = 8
	HorizontalTab          AsciiCharacterType = 9
	LineFeed               AsciiCharacterType = 10
	VerticalTab            AsciiCharacterType = 11
	FormFeed               AsciiCharacterType = 12
	CarriageReturn         AsciiCharacterType = 13
	ShiftOut               AsciiCharacterType = 14
	ShiftIn                AsciiCharacterType = 15
	DatalinkEscape         AsciiCharacterType = 16
	DeviceControl1         AsciiCharacterType = 17
	DeviceControl2         AsciiCharacterType = 18
	DeviceControl3         AsciiCharacterType = 19
	DeviceControl4         AsciiCharacterType = 20
	NotAcknowledged        AsciiCharacterType = 21
	SynchronousIdle        AsciiCharacterType = 22
	EndOfTransmissionBlock AsciiCharacterType = 23
	Cancel                 AsciiCharacterType = 24
	EndOfMedium            AsciiCharacterType = 25
	Substitute             AsciiCharacterType = 26
	Escape                 AsciiCharacterType = 27
	FileSeparator          AsciiCharacterType = 28
	GroupSeparator         AsciiCharacterType = 29
	RecordSeparator        AsciiCharacterType = 30
	UnitSeparator          AsciiCharacterType = 31
	Space                  AsciiCharacterType = 32
	MessageDelimiter       AsciiCharacterType = 46
)

func (x AsciiCharacterType) String() string {
	switch x {
	case Null:
		return "Null"
	case StartOfHeading:
		return "StartOfHeading"
	case StartOfText:
		return "STX-2"
	case EndOfText:
		return "ETX-3"
	case EndOfTransmission:
		return "EOT-4"
	case Enquiry:
		return "ENQ-Enquiry-5"
	case Acknowledged:
		return "ACK-Acknowledged-6"
	case Bell:
		return "BEL-Bell-7"
	case BackSpace:
		return "BS-BackSpace-8"
	case HorizontalTab:
		return "HorizontalTab"
	case LineFeed:
		return "LineFeed"
	case VerticalTab:
		return "VerticalTab"
	case FormFeed:
		return "FormFeed"
	case CarriageReturn:
		return "CarriageReturn"
	case ShiftOut:
		return "SO-ShiftOut"
	case ShiftIn:
		return "SI-ShiftIn"
	case DatalinkEscape:
		return "DatalinkEscape"
	case DeviceControl1:
		return "DC1-DeviceControl1"
	case DeviceControl2:
		return "DC2-DeviceControl2"
	case DeviceControl3:
		return "DC3-DeviceControl3"
	case DeviceControl4:
		return "DC4-DeviceControl4"
	case NotAcknowledged:
		return "NAK-NotAcknowledged"
	case SynchronousIdle:
		return "SynchronousIdle"
	case EndOfTransmissionBlock:
		return "ETB-EndOfTransmissionBlock"
	case Cancel:
		return "Cancel"
	case EndOfMedium:
		return "EndOfMedium"
	case Substitute:
		return "Substitute"
	case Escape:
		return "ESC-Escape"
	case FileSeparator:
		return "FS-FileSeparator-28"
	case GroupSeparator:
		return "GS-GroupSeparator-29"
	case RecordSeparator:
		return "RS-RecordSeparator-30"
	case UnitSeparator:
		return "US-UnitSeparator-31"
	case Space:
		return "Space-32"
	case MessageDelimiter:
		return ".-MessageDelimiter-46"
	}
	return ""
}

func MatchesEndOfTextOrDisconnect(c byte) bool {
	switch AsciiCharacterType(c) {
	case EndOfText, EndOfMedium, EndOfTransmission, EndOfTransmissionBlock:
		return true
	case NotAcknowledged, Cancel:
		return true
	}
	return false
}

func MatchesEnquiry(c byte) bool {
	return AsciiCharacterType(c) == Enquiry
}

func Search(packet []byte, startIndex int, endchars []byte) int {
	for i := startIndex; i < len(packet) && 0 <= startIndex; i++ {
		for _, endchar := range endchars {
			if endchar == packet[i] {
				return i
			}
		}
	}

	return -1
}

func ReverseSearch(packet []byte, startIndex int, endchars []byte) int {
	for i := startIndex; -1 < i && -1 < startIndex; i-- {
		for _, endchar := range endchars {
			if endchar == packet[i] {
				return i
			}
		}
	}

	return -1
}

func GetParityBitCount(x byte) uint {
	var count uint = 0
	//for first 7 bits of each byte, Left shift by one bit position
	//count num of ones

	var a uint
	for a = 0; a < 7; a++ {
		if 1 == ((x >> a) & 01) {
			count++
		}
	}

	return count
}

func CalculateLRC(packet []byte, start int, length int) byte {
	var lrc byte = 0x00
	i := start

	for nLength := len(packet); i < nLength && i < length; i++ {
		lrc ^= packet[i]
	}
	lrc ^= packet[i]
	return lrc
}

func SetBitParity(x byte, evenParity bool) byte {

	if 0 != GetParityBitCount(x)%2 {

		if false == evenParity {
			//x &= byte(0x01 >> 7)

			if 128 < uint(x) {
				x = byte(uint(x) - 128)
			}

		} else {
			x |= byte(0x01 << 7)
		}

		//uneven number means insert one in 8th bit
		//packet[i] += Convert.ToByte(128)
	}

	return x
}

func SetPacketParity(packet []byte, start int, length int, evenParity bool) {
	if nil == packet || 0 >= len(packet) {
		panic("packet cannot be nil or empty")
	}

	for i := start; i < length; i++ {
		packet[i] = SetBitParity(packet[i], evenParity)
	}
}

func RemovePacketFraming(s string) string {

	if 3 > len(s) {
		return s
	}

	iStx := strings.Index(s, string(StartOfText))
	iEtx := strings.LastIndex(s, string(EndOfText))
	if -1 == iStx {
		iStx = 0
	} else if iStx+1 < iEtx && iEtx < len(s) {
		return s[iStx+1 : iEtx]
	}

	return s
}

func FramePacket(s string, start byte, end byte) string {
	return fmt.Sprintf("%c%s%c", start, s, end)
}

// func DecodePacketParity(bufsize uint, packet []byte) []byte {
// 	recordbuf := NewSafeBuffer(bufsize)
// 	recordbuf.AppendBytes(packet)
// 	return recordbuf.DecodeParity(0)
// }

// func EncodePacketParity(bufsize uint, packet []byte) []byte {
// 	recordbuf := NewSafeBuffer(bufsize)
// 	recordbuf.AppendBytes(packet)
// 	return recordbuf.EncodeParity(0)
// }

func FrameAndEncodePacket(bufsize uint, packet []byte, endbyte byte) []byte {

	recordbuf := NewSafeBuffer(bufsize)

	recordbuf.AppendByte(byte(StartOfText))
	recordbuf.AppendBytes(packet)
	recordbuf.AppendByte(endbyte)

	return recordbuf.EncodeParity(0)
}
