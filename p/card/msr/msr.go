package msr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type MagneticStripeTrackI struct {
	AccountNumber   string
	ExpirationMonth uint
	ExpirationYear  uint
	ServiceCode     uint
	CardholderName  string
	TrackData       string
}

func MaskMagneticStripeTrack(magstripe string, account string) string {

	start := strings.Index(magstripe, account)
	if 0 <= start {
		maskedacct := MaskMSRAccountNumber(account, "9")
		track := []byte(strings.Replace(magstripe, account, maskedacct, -1))
		end := len(account) + 1
		for i := 0; i < len(track); i++ {
			if start <= i && i < end {
				continue
			}
			if unicode.IsDigit(rune(track[i])) {
				track[i] = '9'
			}
		}

		startexpiration := strings.LastIndexAny(magstripe, "=^")
		if -1 < startexpiration {
			track[startexpiration+3] = '0'
		}
		return string(track)
	}
	return ""
}

func ParseMagneticStripeTrackI(magstripe string) *MagneticStripeTrackI {
	m := MagneticStripeTrackI{}

	m.TrackData = RemoveMSRFraming(magstripe)

	trackElements := strings.Split(m.TrackData, "^")
	if 2 < len(trackElements) && 3 < len(trackElements[2]) {
		m.AccountNumber = trackElements[0][1:]
		m.CardholderName = strings.Trim(trackElements[1], " ")
		fmt.Sscanf(trackElements[2][0:7], "%02d%02d%03d", &m.ExpirationYear, &m.ExpirationMonth, &m.ServiceCode)
	} else if 2 == len(trackElements) {
		m.AccountNumber = trackElements[0][1:]
		servCode, _ := strconv.Atoi(trackElements[1][4:7])
		m.ServiceCode = uint(servCode)
	}

	return &m
}

type MagneticStripeTrackII struct {
	AccountNumber   string
	ExpirationMonth uint
	ExpirationYear  uint
	TrackData       string
}

func ParseMagneticStripeTrackII(magstripe string) *MagneticStripeTrackII {
	m := MagneticStripeTrackII{}

	m.TrackData = RemoveMSRFraming(magstripe)
	trackElements := strings.Split(m.TrackData, "=")
	if 1 < len(trackElements) && 8 < len(trackElements[0]) && 3 < len(trackElements[1]) {
		m.AccountNumber = trackElements[0]
		fmt.Sscanf(trackElements[1][0:4], "%02d%02d", &m.ExpirationYear, &m.ExpirationMonth)
	}

	return &m
}

func ParseMagneticStripe(magstripe string) (*MagneticStripeTrackI, *MagneticStripeTrackII) {

	var msr1 *MagneticStripeTrackI = nil
	var msr2 *MagneticStripeTrackII = nil

	tracks := strings.Split(magstripe, "?")
	for _, track := range tracks {
		if 0 < len(track) {
			if '%' == track[0] {
				msr1 = ParseMagneticStripeTrackI(track + "?")
			} else if ';' == track[0] {
				msr2 = ParseMagneticStripeTrackII(track + "?")
			}
		}
	}

	return msr1, msr2
}

func MaskMSRAccountNumber(accountno string, maskchar string) string {

	acctlen := len(accountno)
	if 13 > acctlen {
		return ""
	}

	account := []byte(accountno)
	for i := 0; i < len(account); i++ {
		if i < 6 || acctlen-4 <= i {
			continue
		}
		if unicode.IsDigit(rune(account[i])) {
			account[i] = byte(maskchar[0])
		}
	}

	return string(account)
}

func RemoveMSRFraming(s string) string {

	start := strings.IndexAny(s, "%;")
	end := strings.LastIndex(s, "?")

	if -1 < start && -1 < end {
		return s[start+1 : end]
	}

	return s
}
