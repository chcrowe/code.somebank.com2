package dukpt

import (
	"testing"
)

type DukptKeyPair struct {
	BaseDerivationKey string
	KeySerialNo       string
	Expected          string
}

var dukptKeyPairTests = []DukptKeyPair{
	{"0123456789ABCDEFFEDCBA9876543210", "FFFF9876543210E00008", "27F66D5244FF621EAA6F6120EDEB427F"},
	{"0123456789ABCDEFFEDCBA98765432100123456789ABCDEF", "9876543210E00008", "27F66D5244FF621EAA6F6120EDEB427F"},
	{"0123456789ABCDEFFEDCBA9876543210", "9876543210E00008", "27F66D5244FF621EAA6F6120EDEB427F"},
}

func TestDukptSession(t *testing.T) {

	for _, test := range dukptKeyPairTests {
		session_key := GenerateDukptSessionKey(test.BaseDerivationKey, test.KeySerialNo)

		if test.Expected != session_key {
			t.Errorf("DUKPT session key was %s, should be %s", session_key, test.Expected)
		} else {
			t.Logf("session key: %-32s bdk: %-48s ksn: %-32s", session_key, test.BaseDerivationKey, test.KeySerialNo)
		}
	}

}
