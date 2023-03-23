package tsys

import (
	"code.somebank.com/p/bytes"
	//	"fmt"
	"strings"
)

type AuthorizationMessageResponse struct {
	Buffer *bytes.SafeBuffer

	G1 *G1AuthorizationMessageResponse
	G2 *G2AuthorizationMessageResponse
	G3 string
}

func NewAuthorizationMessageResponse(g1 *G1AuthorizationMessageResponse, g2 *G2AuthorizationMessageResponse, g3array []string) *AuthorizationMessageResponse {

	t := AuthorizationMessageResponse{}

	t.Buffer = bytes.NewSafeBuffer(1024)
	t.Buffer.AppendByte(byte(bytes.StartOfText))
	t.Buffer.AppendString(g1.String())

	if nil != g2 {
		t.Buffer.AppendString(g2.String())
	}

	for i, g3 := range g3array {
		if 0 < i {
			t.Buffer.AppendByte(byte(bytes.GroupSeparator))
		}
		t.Buffer.AppendString(g3)
	}

	t.Buffer.AppendByte(byte(bytes.EndOfText))

	return &t
}

func SplitG3AuthorizationResponseRecords(s string) []string {
	return strings.Split(s, string(bytes.GroupSeparator))
}

func ParseAuthorizationMessageResponse(s string) (*G1AuthorizationMessageResponse, *G2AuthorizationMessageResponse, []string) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!
	s = bytes.RemovePacketFraming(s)

	group1, off := ParseG1AuthorizationMessageResponse(s)

	var group2 *G2AuthorizationMessageResponse
	var off2 int
	if DebitEbtResponse == group1.RecordFormat {
		group2, off2 = ParseG2AuthorizationMessageResponse(s[off:])
		off += off2
	} else {
		group2 = nil
	}

	return group1, group2, SplitG3AuthorizationResponseRecords(s[off:])
}
