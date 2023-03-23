package virtualnet

import (
	"bytes"
	"io/ioutil"
	"net/http"

	etsbytes "code.somebank.com/p/bytes"
)

// SSL 1(NOT ENHANCED - HTTPS) Test (All Port 443)
// https://ssltest2h.tsysacquiring.net/scripts/gateway.dll?transact
// https://ssltest.tsysacquiring.net/scripts/gateway.dll?transact *

// SSL 1(NOT ENHANCED - HTTPS) Production (All Port 443)
// https://ssl1.tsysacquiring.net/scripts/gateway.dll?transact
// https://ssl1.vitalps.net/scripts/gateway.dll?transact

var VIRTUALNETSSL_TESTURLS = []string{"https://ssltest2h.tsysacquiring.net/scripts/gateway.dll?transact", "https://ssltest.tsysacquiring.net/scripts/gateway.dll?transact"}

var VIRTUALNETSSL_PRODUCTIONURLS = []string{"https://ssl1.tsysacquiring.net/scripts/gateway.dll?transact", "https://ssl1.vitalps.net/scripts/gateway.dll?transact"}

type VirtualNetWebResponse struct {
	E    error
	Body []byte
}

func SendVirtualNetSslAuthorizationRequest(ch chan *VirtualNetWebResponse, url string, b *bytes.Buffer) {
	//url := "https://ssltest2h.tsysacquiring.net/scripts/gateway.dll?transact"
	resp, e := http.Post(url, "x-Visa-II/x-auth", b)
	if nil != e {
		ch <- &VirtualNetWebResponse{E: e}
		return
	}
	go readHttpResponseBody(ch, resp)
}

func readHttpResponseBody(ch chan *VirtualNetWebResponse, resp *http.Response) {
	body, e := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if nil != e || (nil != body && 0 < len(body)) {
		recvbuf := etsbytes.NewSafeBuffer(1024)
		recvbuf.Write(body)
		ch <- &VirtualNetWebResponse{E: e, Body: recvbuf.DecodeParity(0)}
	} else {
		ch <- &VirtualNetWebResponse{E: e}
	}

}

// func TestVirtualNetWeb(t *testing.T) {

// 	g1, g2, g3records := CreateVisaRetailAuthorization(t)
// 	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)

// 	fmt.Printf("VirtualNetWeb request: %s\n\n", authrequest.Buffer.String())

// 	chresp := make(chan *VirtualNetWebResponse, 1)
// 	go SendVirtualNetSslAuthorizationRequest(chresp, VIRTUALNETSSL_TESTURLS[0], bytes.NewBuffer(authrequest.Buffer.EncodeParity(0)))

// 	vnetresp := <-chresp
// 	if nil == vnetresp.E {
// 		fmt.Printf("VirtualNetWeb response: %s\n\n\n", vnetresp.Body)
// 	}
// }
