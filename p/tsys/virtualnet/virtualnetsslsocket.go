package virtualnet

import (
	"bufio"
	"crypto/tls"
	"time"

	"code.somebank.com/p/bytes"
)

// SSL 2(ENHANCED) Test
// ssltest2s.tsysacquiring.net   Port 15004
// ssltest.tsysacquiring.net     Port 5004 Attended and Final testing

var VIRTUALNETSOCKET_TESTHOSTS = []string{"ssltest2s.tsysacquiring.net:15004", "ssltest.tsysacquiring.net:5004"}

// SSL 2(ENHANCED) Production (All port 5003 for Production)
// ssl3.vitalps.net
// ssl2.vitalps.net

var VIRTUALNETSOCKET_PRODUCTIONHOSTS = []string{"ssl3.vitalps.net:5003", "ssl2.vitalps.net:5003"}

type VirtualNetSocketResponse struct {
	E    error
	Body []byte
}

func SendTlsSocketAuthorizationRequest(ch chan *VirtualNetSocketResponse, host string, b []byte, readtimeout int64) {

	conn, e := tls.Dial("tcp", host, &tls.Config{InsecureSkipVerify: true})
	if nil != e || false == conn.ConnectionState().HandshakeComplete {
		ch <- &VirtualNetSocketResponse{E: e}
		return
	}

	var c byte
	conn.SetReadDeadline(time.Now().Add(time.Duration(readtimeout) * time.Millisecond))

	reader := bufio.NewReader(conn)
	c, e = reader.ReadByte()
	if nil != e || bytes.Enquiry != bytes.AsciiCharacterType(c) { //false == etsbytes.MatchesEnquiry(c)
		ch <- &VirtualNetSocketResponse{E: e}
		return
	}

	recvbuf := bytes.NewSafeBuffer(1024)
	var n int
	n, e = conn.Write(b)
	for nil == e && 0 < n && false == bytes.MatchesEndOfTextOrDisconnect(c) {
		c, e = reader.ReadByte()
		recvbuf.AppendByte(c)
	}

	ch <- &VirtualNetSocketResponse{E: e, Body: recvbuf.DecodeParity(0)}
}
