package tsys

import (
	"code.google.com/p/go-uuid/uuid"
	"code.somebank.com/p/bytes"

	//	"fmt"
	"strings"
	"time"
)

type AuthorizationMessageRequest struct {
	TransactionUuid     string
	TransactionDateTime time.Time

	G1 *G1AuthorizationMessageRequest
	G2 *G2AuthorizationMessageRequest
	G3 string
}

func NewAuthorizationMessageRequest(id string, g1 *G1AuthorizationMessageRequest, g2 *G2AuthorizationMessageRequest, g3array []string) *AuthorizationMessageRequest {

	t := AuthorizationMessageRequest{}
	t.TransactionDateTime = time.Now()

	if 0 < len(id) {
		t.TransactionUuid = id
	} else {
		t.TransactionUuid = uuid.New()
	}

	t.G1 = g1
	t.G2 = g2
	t.G3 = strings.Join(g3array, string(bytes.GroupSeparator))

	return &t
}

func (t *AuthorizationMessageRequest) Bytes() []byte {

	buf := bytes.NewSafeBuffer(1024)

	buf.Zeroize()
	buf.AppendByte(byte(bytes.StartOfText))
	buf.AppendString(t.G1.String())

	if nil != t.G2 {
		buf.AppendString(t.G2.String())
	}

	buf.AppendString(t.G3)

	buf.AppendByte(byte(bytes.EndOfText))

	return buf.Bytes()
}

func SplitG3AuthorizationRequestRecords(s string) []string {
	return strings.Split(s, string(bytes.GroupSeparator))
}

func ParseAuthorizationMessageRequest(s string) (*G1AuthorizationMessageRequest, *G2AuthorizationMessageRequest, []string) {

	// must remove all packet framing (STX, ETX) prior to parsing!!!
	s = bytes.RemovePacketFraming(s)

	group1, off := ParseG1AuthorizationMessageRequest(s)

	var group2 *G2AuthorizationMessageRequest = nil
	var off2 int
	if DebitEbtRequest == group1.RecordFormat {
		group2, off2 = ParseG2AuthorizationMessageRequest(s[off:])
		off += off2
	} else {
		group2 = nil
	}

	return group1, group2, SplitG3AuthorizationRequestRecords(s[off:])
}
