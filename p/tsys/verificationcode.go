package tsys

import (
	"fmt"
	"strconv"
	"strings"
)

// This six-character field is used in assigning a value that assists in authenticating the physical
// presence of a Visa, MasterCard, Discover, PayPal or American Express credit card.

// Card Verification is not contained in the magnetic stripe information
// nor does it appear on sales receipts; it is an additional three to
// four-character value, printed on the front or back of Visa, MasterCard, Discover, PayPal and
// American Express cards. This field is used to support the following card verification programs.
// • Visa - Card Verification Value 2 (CVV2)
// • MasterCard - Card Validation Code 2 (CVC2)
// • American Express - Cardholder Identification Code (CID)
// • Discover/PayPal - Cardholder Identification Code (CID)

type VerificationCodePresenceType uint

const (
	VerificationCodeIntentionallyNotProvided VerificationCodePresenceType = 0
	VerificationCodePresent                  VerificationCodePresenceType = 1
	VerificationCodePresentButIllegible      VerificationCodePresenceType = 2
	NoVerificationCodePresentOnCard          VerificationCodePresenceType = 9
)

func (x VerificationCodePresenceType) String() string {
	switch x {
	case VerificationCodeIntentionallyNotProvided:
		return "VerificationCodeIntentionallyNotProvided"
	case VerificationCodePresent:
		return "VerificationCodePresent"
	case VerificationCodePresentButIllegible:
		return "VerificationCodePresentButIllegible"
	case NoVerificationCodePresentOnCard:
		return "NoVerificationCodePresentOnCard"
	}
	return ""
}

type VerificationCodeDesiredResponseType uint

const (
	NormalResponseRequested   VerificationCodeDesiredResponseType = 0
	EnhancedResponseRequested VerificationCodeDesiredResponseType = 1
)

func (x VerificationCodeDesiredResponseType) String() string {
	switch x {
	case NormalResponseRequested:
		return "NormalResponseRequested"
	case EnhancedResponseRequested:
		return "EnhancedResponseRequested"
	}
	return ""
}

// Verification code (CVV2, CVC2, CID)
// some test values for CVV CVV2 and CID: Visa = 999, MC = 998, Discover = 996

type VerificationCodeInfo struct {
	CodePresenceIndicator    VerificationCodePresenceType        // position 1
	DesiredResponseIndicator VerificationCodeDesiredResponseType // position 2
	VerificationCode         string                              // positions 3-6
	// 						Verification Code as printed on the card,
	// 						right- justify/space-fill entry
	// 						If Position 1 = 0, 2, or 9, positions 3-6 should be spacefilled
	// 						If the Verification Code is part of an encrypted (Record
	// 						Format = 'W') transaction, positions 3-6 should be
	// 						space-filled.
}

// NOTE:
// • If Position 1 = 1, then Position 2 should be set to 0 or 1.
// • If Position 1 = 0, 2, or 9, then Position 2 should be set to 0.

// If Verification Code is sent as part of an encrypted transaction (Record Format = 'W'),
// Positions 1 and 2 still need to be present even though the Verification Code is sent in the
// Customer Data field.

func NewVerificationCodeInfo(desiredresponse VerificationCodeDesiredResponseType, code string) *VerificationCodeInfo {
	v := VerificationCodeInfo{DesiredResponseIndicator: desiredresponse}
	v.SetVerificationCode(code)
	return &v
}

func (v *VerificationCodeInfo) SetVerificationCode(code string) {
	v.CodePresenceIndicator = VerificationCodeIntentionallyNotProvided
	if 3 <= len(code) {
		v.CodePresenceIndicator = VerificationCodePresent
		v.VerificationCode = strings.Trim(code, " ")
	}
}

func (v *VerificationCodeInfo) Copy(vc *VerificationCodeInfo) {
	v.CodePresenceIndicator = vc.CodePresenceIndicator
	v.DesiredResponseIndicator = vc.DesiredResponseIndicator
	v.VerificationCode = vc.VerificationCode
}

func (v *VerificationCodeInfo) String() string {
	return fmt.Sprintf("%d%d%4s", v.CodePresenceIndicator, v.DesiredResponseIndicator, v.VerificationCode)
}

func ParseVerificationCodeInfo(s string) *VerificationCodeInfo {
	v := VerificationCodeInfo{}
	if 5 <= len(s) {
		c1, _ := strconv.ParseUint(s[0:1], 10, 32)
		v.CodePresenceIndicator = VerificationCodePresenceType(c1)
		c2, _ := strconv.ParseUint(s[1:2], 10, 32)
		v.DesiredResponseIndicator = VerificationCodeDesiredResponseType(c2)
		v.VerificationCode = s[2:]
	}
	return &v
}
