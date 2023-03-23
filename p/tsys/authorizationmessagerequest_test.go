package tsys

import (
	// "bufio"
	"bytes"
	"code.somebank.com/p/auth"
	"code.somebank.com/p/bitwise"
	etsbytes "code.somebank.com/p/bytes"
	"code.somebank.com/p/filesys"
	"code.somebank.com/p/list"
	. "code.somebank.com/p/tsys/virtualnet"
	"crypto/sha512"
	// "crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
	// "time"
)

var TESTAUTHFILEPATH string = "c:\\Projects\\Go\\src\\code.somebank.com\\p\\tsys\\tsysauths2.dat"
var TESTSALT string = "D64C3FFECFC340DA8CDD4123AF16E3A4"

type MessageTest struct {
	request, expected, received string
}

var messageTests = []MessageTest{

	{"T4.99999588800000083559991515QR84084020165    007055999N216593KD9999999999999999=999910100000100000000000000004A0000000000000010000CTC TESTING - RETAIL     Sterling     VA      090814173222425163142282142282 V51234567891324000000112233207Z020000136B290034", "U4.N599915155216500TAS384090814173223APPROVAL TAS384 0425163142282 000000000214534    999995        14228250908020034A ", ""},
}

// t.Fatalf("Unexpected error: %v", err)
// t.Errorf("Contents = %v, want %v", buf, expected)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NoTestUrls(t *testing.T) {
	uri, _ := url.Parse("https://pos.somebank.com/POSCP3/main.asp?a=AUTHORIZE_ONLINE&VendorName=ACME Corp&")
	t.Logf("VERBOSE: %#v", uri)
	t.Logf("PATH: %#v", uri.Path)
	t.Logf("RAWQUERY: %#v", uri.RawQuery)

	matched, _ := regexp.MatchString("/POSCP3", uri.Path)
	t.Logf("MATCHES POSCP3: %#v", matched)
}

func NoTestG3v027(t *testing.T) {
	G3v027 := NewG3v027PosDataRequest(TerminalInputICCReaderAndContactlessTerminal,
		CardholderAuthenticationPINEntryCapability,
		CardCaptureCapability,
		EnvironmentOnCardAcceptorPremisesAttendedTerminal,
		CardholderPresent,
		CardPresent,
		CardInputOnlineChip,
		CardholderPINAuthenticated,
		AuthenticationEntityAuthorizingAgentOnlinePIN,
		CardOutputCapabilityICC,
		TerminalOutputCapabilityICC,
		PINCaptureCapability06CharactersMaximum)

	t.Logf("%v", G3v027.String())

	// s := strings.Split(G3v027.VerboseString(), "; ")
	// for _,s1 := range s{
	// 	t.Logf("%s", s1)
	// }

	g1fields := strings.Split(G3v027.VerboseString(), "; ")
	caption := "G3v027 ->"
	for _, g1f := range g1fields {
		t.Logf("%-12s%s\n", caption, g1f)
		caption = ""
	}

	t.Logf("%-42s %c (%[2]v)", "G3v027.TerminalInputCapability", G3v027.TerminalInputCapability)                       //1 A/N
	t.Logf("%-42s %d (%[2]v)", "G3v027.CardholderAuthenticationCapability", G3v027.CardholderAuthenticationCapability) //1 NUM
	t.Logf("%-42s %d (%[2]v)", "G3v027.CardCaptureCapability", G3v027.CardCaptureCapability)                           //1 NUM
	t.Logf("%-42s %c (%[2]v)", "G3v027.TerminalOperatingEnvironment", G3v027.TerminalOperatingEnvironment)             //1 A/N
	t.Logf("%-42s %d (%[2]v)", "G3v027.CardholderPresentData", G3v027.CardholderPresentData)                           //1 NUM
	t.Logf("%-42s %c (%[2]v)", "G3v027.CardPresentData", G3v027.CardPresentData)                                       //1 A/N
	t.Logf("%-42s %c (%[2]v)", "G3v027.CardInputMode", G3v027.CardInputMode)                                           //1 A/N
	t.Logf("%-42s %c (%[2]v)", "G3v027.CardholderAuthenticationMethod", G3v027.CardholderAuthenticationMethod)         //1 A/N
	t.Logf("%-42s %d (%[2]v)", "G3v027.CardholderAuthenticationEntity", G3v027.CardholderAuthenticationEntity)         //1 NUM
	t.Logf("%-42s %c (%[2]v)", "G3v027.CardOutputCapability", G3v027.CardOutputCapability)                             //1 A/N
	t.Logf("%-42s %d (%[2]v)", "G3v027.TerminalOutputCapability", G3v027.TerminalOutputCapability)                     //1 NUM
	t.Logf("%-42s %c (%[2]v)", "G3v027.PINCaptureCapability", G3v027.PINCaptureCapability)                             //1 A/N

}

// func NoTestVirtualNetTls(t *testing.T) {

// 	g1, g2, g3records := CreateVisaRetailAuthorization(t)
// 	g1.ApplicationType = SingleAuthorizationPerConnection
// 	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)

// 	//t.Logf("VirtualNetTls request: %s\n\n", authrequest.Buffer.String())

// 	chresp := make(chan *VirtualNetSocketResponse, 1)
// 	go SendTlsSocketAuthorizationRequest(chresp, VIRTUALNETSOCKET_TESTHOSTS[0], authrequest.Buffer.EncodeParity(0), 3000)

// 	vnetresp := <-chresp
// 	if nil == vnetresp.E {
// 		body := string(vnetresp.Body)
// 		if -1 == strings.Index(body, "APPROVAL") {
// 			//t.Logf("TestVirtualNetTls() PASS: %s\n", body)
// 			t.Errorf("APPROVAL EXPECTED, GOT: %s\n", body)
// 		}

// 	} else {
// 		t.Errorf("ERROR DESCRIPTION: %v (%v)\n\n\n", vnetresp.E, vnetresp.Body)
// 	}
// }

// func NoTestVirtualNetWeb(t *testing.T) {

// 	g1, g2, g3records := CreateVisaRetailAuthorization(t)
// 	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)

// 	//t.Logf("VirtualNetWeb request: %s\n\n", authrequest.Buffer.String())

// 	chresp := make(chan *VirtualNetWebResponse, 1)
// 	go SendVirtualNetSslAuthorizationRequest(chresp, VIRTUALNETSSL_TESTURLS[0], bytes.NewBuffer(authrequest.Buffer.EncodeParity(0)))

// 	vnetresp := <-chresp
// 	if nil == vnetresp.E {

// 		body := string(vnetresp.Body)
// 		//t.Logf("RESPONSE: %s", body)

// 		if -1 == strings.Index(body, "APPROVAL") {
// 			//t.Logf("TestVirtualNetWeb() PASS: %s\n", body)
// 			t.Errorf("APPROVAL EXPECTED, GOT: %s\n", body)
// 		}

// 	} else {
// 		t.Errorf("ERROR DESCRIPTION: %v (%v)\n\n\n", vnetresp.E, vnetresp.Body)
// 	}
// }

// func NoTestTls(t *testing.T) {

// 	conn, err := tls.Dial("tcp", VIRTUALNETSOCKET_TESTHOSTS[0],
// 		&tls.Config{ServerName: "ssltest2s.tsysacquiring.net", InsecureSkipVerify: true})
// 	if err != nil {
// 		t.Fatalf("tls.Dial() ERROR: %v", err.Error())
// 	}
// 	defer conn.Close()

// 	var c byte
// 	reader := bufio.NewReader(conn)

// 	conn.SetReadDeadline(time.Now().Add(2000 * time.Millisecond))

// 	c, err = reader.ReadByte()
// 	if err != nil || false == etsbytes.MatchesEnquiry(c) {
// 		t.Fatalf("error connect: %v", err.Error())
// 	}

// 	g1, g2, g3records := CreateVisaRetailAuthorization(t)
// 	g1.ApplicationType = SingleAuthorizationPerConnection

// 	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)
// 	//t.Logf("request: %s\n\n", authrequest.Buffer.String())

// 	_, err = conn.Write(authrequest.Buffer.EncodeParity(0))
// 	if err != nil {
// 		t.Fatalf("error writing: %v", err)
// 	}

// 	recvbuf := etsbytes.NewSafeBuffer(1024)

// 	c, err = reader.ReadByte()
// 	//t.Logf("c -> %v\n", c)

// 	for false == etsbytes.MatchesEndOfTextOrDisconnect(c) {
// 		c, err = reader.ReadByte()
// 		recvbuf.AppendByte(c)
// 		conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
// 	}

// 	// // decode and parse the response
// 	body := string(recvbuf.DecodeParity(0))
// 	if -1 == strings.Index(body, "APPROVAL") {
// 		t.Errorf("APPROVAL EXPECTED, GOT: %s\n", body)
// 		//t.Logf("response: %s\n", responsepacket)
// 	}
// 	//G1, _, G3arr := ParseAuthorizationMessageResponse(responsepacket)

// }

func NoTestParseNameValuePairMap(t *testing.T) {

	m := list.ParseNameValuePairMap("k1=value one&k2=value two&k3=value three&k4=value four&k5=value five&")
	m.Delete("k3")
	fmt.Println()
	t.Logf("dump-> %s\n", m.String())
	fmt.Println()
	m.Clear()
	m.Set("newkey1", "new value one")
	t.Logf("after clear-> %s\n", m.String())
}

// func NoTestWriteToFile(t *testing.T) {
// 	arw := filesys.NewFileReaderWriter(TESTAUTHFILEPATH)
// 	arw.Create()
// 	sendVirtualNetSslHttpRequest(CreateVisaRetailAuthorization(t))
// 	// MastercardRetailAuthorization(t)
// }

func NoTestReadFromFile(t *testing.T) {

	arw := filesys.NewFileReaderWriter(TESTAUTHFILEPATH)
	lines, _ := arw.ReadLines()

	for _, line := range lines {
		if 0 < len(line) {
			a := auth.ParseAuthorizationRecordFromJson([]byte(line), TESTSALT)
			t.Logf("%v\n", a)
		}
	}

}

func NoTestAuthorizationRecordRead(t *testing.T) {

	authrec := "Utc=20141030T221027:693821100&Id=eeed3662-f1ef-493b-bd08-39a0a05f104c&Out=D4.99999599999999991100119911QR84084022030    007055999Y000154ZD4123459999992349=990999999999999999000000008800IKEA STORE               FAIRFAX      VA001!010020000136B2480223002840C0000000005003002840C0000000007003002840C0000000009003002840C0000000011000211EB2C1772D025D00000150027H11101C13336032034040046061&In=E4.N001199115000100TAS262103014171027APPROVAL TAS262 0430322500472 00000000000035076A600100200223002840C0000000005003002840C0000000007003002840C0000000009003002840C000000001100021025027032840034A 040046VISA061*&Pii=bb4cc473f2122341954b44c47c2615b22665a4a0c92e38f2baa295fa0fff0e2cb708ca"

	a := auth.ParseAuthorizationRecordFromValues(authrec, TESTSALT)

	fmt.Println()
	t.Logf("dump-> %v\n", a)
	fmt.Println()

	g1, g2, g3records := ParseAuthorizationMessageRequest(a.Out)
	DumpAuthRequest(g1, g2, g3records, a.Out)
	fmt.Println()
	fmt.Println()
	g1b, g2b, g3brecords := ParseAuthorizationMessageResponse(a.In)
	DumpAuthResponse(g1b, g2b, g3brecords, a.In)

}

func Hashing(t *testing.T) {
	fmt.Println()
	fmt.Println()
	h512 := sha512.New()
	io.WriteString(h512, "Hello money money!")
	t.Logf("Hash512 : %x\n", h512.Sum(nil))
}

func TestCustomerDataSterilizing(t *testing.T) {

	g1Tests := []string{"D4.99999599999999991100119911QR84084022030    007055999Y000154MT4219876128293505121585284000000008800IKEA STORE               FAIRFAX      VA001!01000711 999020000136B248",
		"D4.99999599999999991100119911QR84084020165    007055999Y003254@X403949293935342710153000VisaNet Merchant         Sterling     VA020000136B248",
		"D4.99999599999999991100119911QR84084020165    007055999Y001954@X601161143003692908149500VisaNet Merchant         Sterling     VA020000136B248",
		"D4.99999599999999991100119911QR84084020165    007055999Y002354@X521882636840220709142000VisaNet Merchant         Sterling     VA020000136B248",
		"D4.99999599999999991100119911QR84084020165    007055999Y001754@D5499740000000057=151210100000512400VisaNet Merchant         Sterling     VA020000136B248",
		"D4.99999599999999991100119911QR84084020165    007055999Y002554@D4012000010000=15121010004852000VisaNet Merchant         Sterling     VA020000136B248"}

	for _, g1Test := range g1Tests {
		g1, _ := ParseG1AuthorizationMessageRequest(g1Test)

		t.Logf("%-8s%s\n", "G1", g1.String())
		g1.CustomerData.Sterilize()
		t.Logf("%-8s%s\n\n", "G1SAFE", g1.String())
	}
}

// func sendHttpPost(ch chan []byte, b *bytes.Buffer) {
// 	url := VIRTUALNETSSL_TESTURLS[0]
// 	resp, err := http.Post(url, "x-Visa-II/x-auth", b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t.Logf("--> resp.ContentLength: %d\n", resp.ContentLength)
// 	go readHttpResponseBody(ch, resp)
// }

// func readHttpResponseBody(ch chan []byte, resp *http.Response) {
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()

// 	t.Logf("<-- body length: %d\n", len(body))
// 	ch <- body
// }

func sendVirtualNetSslHttpRequestAsync(t *testing.T, g1 *G1AuthorizationMessageRequest, g2 *G2AuthorizationMessageRequest, g3records []string) {

	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)
	buf := etsbytes.NewSafeBuffer(1024)
	buf.AppendBytes(authrequest.Bytes())

	DumpAuthRequest(g1, nil, g3records, string(buf.Bytes()))

	// send http post request
	chresp := make(chan *VirtualNetWebResponse, 1)
	go SendVirtualNetSslAuthorizationRequest(chresp, VIRTUALNETSSL_TESTURLS[0], bytes.NewBuffer(buf.EncodeParity(0)))

	vnetresp := <-chresp
	if nil == vnetresp.E {
		t.Logf("VirtualNetWeb response: %s\n\n\n", vnetresp.Body)
	} else {
		t.Logf("VirtualNetWeb response error: %v (%v)\n\n\n", vnetresp.E, vnetresp.Body)
	}

	// decode and parse the response
	responsepacket := string(vnetresp.Body)
	G1, _, G3arr := ParseAuthorizationMessageResponse(responsepacket)

	DumpAuthResponse(G1, nil, G3arr, responsepacket)

	customerdata := g1.CustomerData.String()
	authrequest.G1.CustomerData.Sterilize()

	a := auth.NewAuthorizationRecord(string(authrequest.Bytes()), responsepacket, customerdata, TESTSALT)
	arw := filesys.NewFileReaderWriter(TESTAUTHFILEPATH)
	arw.WriteString(a.Json() + "\n")

}

func sendVirtualNetSslHttpRequest(g1 *G1AuthorizationMessageRequest, g2 *G2AuthorizationMessageRequest, g3records []string) {

	authrequest := NewAuthorizationMessageRequest("", g1, g2, g3records)
	buf := etsbytes.NewSafeBuffer(1024)
	buf.AppendBytes(authrequest.Bytes())
	DumpAuthRequest(g1, nil, g3records, string(buf.Bytes()))

	// send http post request
	resp, err := http.Post(VIRTUALNETSSL_TESTURLS[0], "x-Visa-II/x-auth", bytes.NewBuffer(buf.EncodeParity(0)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read http post response
	body, _ := ioutil.ReadAll(resp.Body)
	recvbuf := etsbytes.NewSafeBuffer(1024)
	recvbuf.Write(body[0:])

	// decode and parse the response
	responsepacket := string(recvbuf.DecodeParity(0))
	G1, _, G3arr := ParseAuthorizationMessageResponse(responsepacket)

	DumpAuthResponse(G1, nil, G3arr, responsepacket)

	customerdata := g1.CustomerData.String()
	authrequest.G1.CustomerData.Sterilize()

	a := auth.NewAuthorizationRecord(string(authrequest.Bytes()), responsepacket, customerdata, TESTSALT)
	arw := filesys.NewFileReaderWriter(TESTAUTHFILEPATH)
	arw.WriteString(a.Json() + "\n")

}

// func TestParse(t *testing.T) {
// 	for _, msgtest := range messageTests {

// 		// must remove all packet framing (STX, ETX) prior to parsing!!!
// 		s := RemovePacketFraming(msgtest.request)

// 	}
// }

func NoTestOptionalDataGroups(t *testing.T) {

	odg := NewOptionalDataGroupMap(DetailRecord)
	t.Logf("OptionalDataGroupMap: %s", odg)

	odg.SetGroupBit(3, true)
	odg.SetGroupBit(12, true)
	t.Logf("OptionalDataGroupMap (Groups 3 and 12): %s\n", odg)

	var i uint = 1

	for ; i < 25; i++ {
		t.Logf("Group %02d: %-7v (%s)", i, odg.GetGroupBit(GroupMapBitType(i)), GroupMapBitType(i))
	}

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBit(2, true)
	t.Logf("OptionalDataGroupMap (Group 2): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBit(5, true)
	t.Logf("OptionalDataGroupMap (Group 5): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBit(5, true)
	odg.SetGroupBit(7, true)
	t.Logf("OptionalDataGroupMap (Groups 5 and 7): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBit(4, true)
	t.Logf("OptionalDataGroupMap (Group 4): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBits([]GroupMapBitType{DetailRecordBit_AutoRental, DetailRecordBit_AutoRentalReturn}, true)
	t.Logf("OptionalDataGroupMap (Groups 4 and 13): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBit(9, true)
	t.Logf("OptionalDataGroupMap (Group 9): %s\n", odg)

	odg = NewOptionalDataGroupMap(DetailRecord)
	odg.SetGroupBits([]GroupMapBitType{DetailRecordBit_Hotel, DetailRecordBit_HotelAmericanExpress}, true)
	t.Logf("OptionalDataGroupMap NEW (Groups 5 and 7): %s\n", odg)

	t.Log("")
	t.Log("Bit " + bitwise.PrintBitLegend() + "            TOGGLED GROUPS")
	t.Logf("    %s   %s", PrintBits(odg.GroupByteIV), PrintGroupBits(odg.GroupByteIV, 1))
	t.Logf("    %s   %s", PrintBits(odg.GroupByteIII), PrintGroupBits(odg.GroupByteIII, 7))
	t.Logf("    %s   %s", PrintBits(odg.GroupByteII), PrintGroupBits(odg.GroupByteII, 13))
	t.Logf("    %s   %s", PrintBits(odg.GroupByteI), PrintGroupBits(odg.GroupByteI, 19))

	// "D@@@@", "Base Group"
	// "DP@@@", "Encrypted Detail Record"
	// "D@@@B", "(Group 2) Detail Record (restaurant)"
	// "D@@‘D", "(Groups 3 and 12) Detail Record (direct marketing)"
	// "D@@@P",  "(Group 5) Detail Record (Hotel)"
	// "D@@AP", "(Groups 5 and 7) Detail Record (Hotel/American Express)"
	// "D@@@H", "(Group 4) Detail Record (auto rental)"
	// "D@A@H", "(Groups 4 and 13) Detail Record (auto rental/MasterCard)"
	// "D@@D@", "(Group 9) Detail Record (Passenger Transport)"
	// "D@@H@", "(Group 10) Detail Record (Direct Debit)"
	// "D@@H@", "(Group 10) Detail Record (EBT)"
	// "D@@’D", "Detail Record (e-Commerce)"

}

func TestAuthorizationRequestParse(t *testing.T) {
	s := "T4.99999588800000083559991515QR84084020165    007055999N209493KG9999999999999999=9999201175898938900000000000000004A000000000000007000CTC TESTING - RETAIL 2   Sterling     VA      0 71715104453519838693474693474 000000000000434Z020000136B2900340554F07A0000000031010500 B564953412043524544495482025C008407A0000000031010950502000080009A031507159B02E8009C010010050230 30100E02454E5F201A56495341204143515549524552205445535420434152442030315F280201245F2A0208405F340 10190000142900101439F02060000000010009F03060000000000009F0607A00000000310109F0702FF009F0902008C 9F100706010A03A400009F1101019F120F4352454449544F20444520564953419F1A0208409F1E08534330313039353 99F21030901519F2608FACFCA366378280A9F2701809F3303E0D8C89F34034103029F3501229F360200019F3704397A CC7E9F3901059F4005F000F0A0019F4104000001409F530152DF7807A0000000041010DF7908F969CD8ABACD22F2"
	g1, g2, g3records := ParseAuthorizationMessageRequest(s)
	DumpAuthRequest(g1, g2, g3records, s)
	fmt.Println()
	fmt.Println()
}

// func TestRetailAuthorization(t *testing.T) {
// 	G1 := NewG1AuthorizationMessageRequest(999995, Purchase, 1, 1000, 0)
// 	G1.SetTerminalIdentification(999999999911, 1515, 5999, PosPartner, Retail, MiscellaneousAndSpecialtyRetailStores)
// 	G1.SetTerminalLocale(UsdUnitedStatesUsDollar, UnitedStates, English, "22030", NewTimeZoneDifferential(HoursBehindGMTWithDaylightSavings, 5))

// 	cc := ParseCreditCard(";4388570000000006=1512101000001?")
// 	customeraccountdata := NewCustomerDataField(FullMagneticStripeReadAndTransmitTrack2, "", cc.Track02DataNoSentinals())
// 	cardholderidentification := NewCardholderIdentificationDataField(CardholderSignatureTerminalHasPinPad, "", "22031", "", "")
// 	G1.SetCardholderData(customeraccountdata, cardholderidentification)
// 	G1.CardAcceptorData = NewCardAcceptorDataField("GENERAL RETAIL STORE", "FAIRFAX", "VA", "", 8664445555)

// 	WrapAndPrintText(fmt.Sprintf("%+v\n\n", G1.String()), "G1 packet ->", 128)

// 	g1fields := strings.Split(G1.VerboseString(), "; ")
// 	caption := "G1 fields ->"
// 	for _, g1f := range g1fields {
// 		t.Logf("%-32s%s\n", caption, g1f)
// 		caption = ""
// 	}

// 	G3v001 := NewG3v001CommercialCardRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v001 ->", G3v001.String(), *G3v001)

// 	G3v007 := NewG3v007CardVerificationCodeRequest(GetCardIssuerType(cc.AccountNumber), "456")
// 	t.Logf("%-32s%s %+v\n\n", "G3v007 ->", G3v007.String(), *G3v007)

// 	G3v008 := NewG3v008FleetFuelingCardRequest(G3v008IdNumber, "T65334499")
// 	t.Logf("%-32s%s %+v\n\n", "G3v008 ->", G3v008.String(), *G3v008)

// 	G3v009 := NewG3v009SeteCommerceRequest("E219A12A0DDA42E1802EA655DBE50AAF", "6A419A5915994058921AEF271855ADD9", "0ADA33A6CE104A5CA3B5874D75B725F73DD4D800", "11AD8DCA8F7A43AD8C52A2D9AB914F2B9EA7A6ED")
// 	WrapAndPrintText(fmt.Sprintf("%-32s%s\n", G3v009.String()), "G3v009 ->", 128)
// 	WrapAndPrintText(fmt.Sprintf("%+v\n\n", *G3v009), "", 128)

// 	G3v011 := NewG3v011ChipConditionRequest(ServiceCodeDoesNotBeginWith2Or6)
// 	t.Logf("%-32s%s %+v\n\n", "G3v011 ->", G3v011.String(), *G3v011)

// 	G3v013 := NewG3v013ElectronicBenefitsTransferRequest("10F79ED", "BA0F1CAC203163B", "0687D1")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v013.String(), *G3v013), "G3v013 ->", 128)

// 	G3v014 := NewG3v014MotoEcommerceRequest(SingleTransaction)
// 	t.Logf("%-32s%s %+v\n\n", "G3v014 ->", G3v014.String(), *G3v014)

// 	G3v015 := NewG3v015ServiceDevelopmentIndicatorRequest(Transponder)
// 	t.Logf("%-32s%s %+v\n\n", "G3v015 ->", G3v015.String(), *G3v015)

// 	G3v017 := NewG3v017Visa3dSecurEcomVerificationRequest("D0EBB6E4AA3B4830868B3974148817007861728C", "950E546A7ECF451E86C5F03DA9FFD7474DA89201")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v017.String(), *G3v017), "G3v017 ->", 128)

// 	G3v018 := NewG3v018ExistingDebtIndicatorRequest(ExistingDebtTransaction)
// 	t.Logf("%-32s%s %+v\n\n", "G3v018 ->", G3v018.String(), *G3v018)

// 	G3v019 := NewG3v019MastercardUniversalCardholderAuthenticationRequest(UcafDataWasPopulated, "FD546A6609D840CBA90076F0B6B8BCBC")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v019.String(), *G3v019), "G3v019 ->", 128)

// 	G3v020 := NewG3v020DeveloperInformationRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v020 ->", G3v020.String(), *G3v020)

// 	G3v021 := NewG3v021MerchantVerificationValueRequest("1EB2C1772D")
// 	t.Logf("%-32s%s %+v\n\n", "G3v021 ->", G3v021.String(), *G3v021)

// 	additionalAmounts := [...]AdditionalAmountInfo{*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UsdUnitedStatesUsDollar, PositiveBalance, 500)),
// 		*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UsdUnitedStatesUsDollar, PositiveBalance, 700)),
// 		*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UsdUnitedStatesUsDollar, PositiveBalance, 900)),
// 		*(NewAdditionalAmountInfo(CreditCardAccount, DepositAccountAvailableBalance, UsdUnitedStatesUsDollar, PositiveBalance, 1100))}

// 	G3v022 := NewG3v022AdditionalAmountsRequest(&additionalAmounts)
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v022.String(), *G3v022), "G3v022 ->", 128)

// 	G3v023 := NewG3v023VisaMastercardHealthcareRequest("WeServeEveryone Clinic	 ;1111 First Street California;111-111-11111   Fax: 111-111-1111	Chart Summary;Monica Latte;Home: 444-444-4444;Female  DOB: 04/04/1950	0000-44444	Ins: Commercial xxxxx")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v023.String(), *G3v023), "G3v023 ->", 128)

// 	G3v024 := NewG3v024MerchantAdviceCodeRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v024 ->", G3v024.String(), *G3v024)

// 	G3v025 := NewG3v025TransactionFeeAmountRequest(DebitTransactionFee, 150)
// 	t.Logf("%-32s%s %+v\n\n", "G3v025 ->", G3v025.String(), *G3v025)

// 	p := [...]ProductParticipationGroupType{MerchandiseSalesCanBePartiallyApproved, MerchandiseOnlyCanBePartiallyApproved}

// 	G3v026 := NewG3v026ProductParticipationGroupRequest(fmt.Sprintf("%s", p))
// 	t.Logf("%-32s%s %+v\n\n", "G3v026 ->", G3v026.String(), *G3v026)

// 	G3v027 := NewG3v027PosDataRequest(TerminalInputICCReaderAndContactlessTerminal,
// 		CardholderAuthenticationPINEntryCapability,
// 		CardCaptureCapability,
// 		EnvironmentOnCardAcceptorPremisesAttendedTerminal,
// 		CardholderPresent,
// 		CardPresent,
// 		CardInputOnlineChip,
// 		CardholderPINAuthenticated,
// 		AuthenticationEntityAuthorizingAgentOnlinePIN,
// 		CardOutputCapabilityICC,
// 		TerminalOutputCapabilityICC,
// 		PINCaptureCapability06CharactersMaximum)
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v027.String(), *G3v027), "G3v027 ->", 128)

// 	G3v028 := NewG3v028AmericanExpressAdditionalDataRequest("AXITDCE~24CFFROST@EMAILADDRESS.COMCH~14PHX.QW.AOL.COMHBT46MOZILLA/4.0~(COMPATIBLE;~MSIE~5.0;~WINDOWS~95)STC03840SM~0202MPS08TKDC315U127.142.005.056602555121200")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v028.String(), *G3v028), "G3v028 ->", 128)

// 	billing := NewContactAddressType("21207", "4408 Belle Ave", "Gwynn Oak", "MD", UnitedStates, "James", "Smith", 4103859378)
// 	shipping := NewContactAddressType("21201", "200 W Lombard St", "Baltimore", "MD", UnitedStates, "James", "Smith", 4435732800)

// 	G3v029 := NewG3v029ExtendedAVSdataRequest(billing, shipping)
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v029.String(), *G3v029), "G3v029 ->", 128)

// 	G3v030 := NewG3v030AmericanExpressMerchantNameLocationCATRequest("OILCATREQ0378256", "21201")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v030.String(), *G3v030), "G3v030 ->", 128)

// 	G3v030B := NewG3v030AmericanExpressMerchantNameLocationAggregatorRequest("AGGREGATOR462905", "200 W Lombard St", "Baltimore", "21201", 4435732800, "jmsmith@aol.com")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v030B.String(), *G3v030B), "G3v030B ->", 128)

// 	G3v031 := NewG3v031AgentIdentificationServiceRequest("F2A11", "1ABA2AEC790B")
// 	t.Logf("%-32s%s %+v\n\n", "G3v031 ->", G3v031.String(), *G3v031)

// 	G3v032 := NewG3v032CurrencyConversionDataRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v032 ->", G3v032.String(), *G3v032)

// 	G3v033 := NewG3v033ReversalRequestCodeRequest(CustomerDirectedReversal)
// 	t.Logf("%-32s%s %+v\n\n", "G3v033 ->", G3v033.String(), *G3v033)

// 	G3v034 := NewG3v034CardProductCodeRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v034 ->", G3v034.String(), *G3v034)

// 	G3v035 := NewG3v035PromotionalCodeRequest("PAPAL PROMO ID: 3C5V3UR-AE4LS")
// 	t.Logf("%-32s%s %+v\n\n", "G3v035 ->", G3v035.String(), *G3v035)

// 	G3v036 := NewG3v036PaymentTransactionIdentifierRequest(OtherPaymentTransaction)
// 	t.Logf("%-32s%s %+v\n\n", "G3v036 ->", G3v036.String(), *G3v036)

// 	G3v037 := NewG3v037RealTimeSubstantiationRequest(ExemptFromIIAS90PercentRule)
// 	t.Logf("%-32s%s %+v\n\n", "G3v037 ->", G3v037.String(), *G3v037)

// 	G3v038 := NewG3v038ElectroMagneticSignatureRequest("2CAE0AE1D7964C719E4B8E335D126FD15346B63BC3FE42F28028BE140FFAA4B9F3A81EFD9E844E0FADADBBBFBA896C081E62B583E4C64F7CBE18EE986A07701B")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v038.String(), *G3v038), "G3v038 ->", 128)

// 	G3v039 := NewG3v039CardholderVerificationMethodRequest(CardholderVerificationByPin)
// 	t.Logf("%-32s%s %+v\n\n", "G3v039 ->", G3v039.String(), *G3v039)

// 	G3v040 := NewG3v040VisaISAChargeIndicatorRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v040 ->", G3v040.String(), *G3v040)

// 	G3v041 := NewG3v041NtiaUpcSkuDataRequest("01234-001-F10-XL")
// 	t.Logf("%-32s%s %+v\n\n", "G3v041 ->", G3v041.String(), *G3v041)

// 	//9F2600084A3FD0EA4B18C557008200021C009F36000200799F3400030100029F270001809F330003E0F0C89F350001220095000580000000009F10000706040A03A400009F370004B4C950679F0200060000000001419F0300060000000000009F1A000201245F2A00020124009A0003140925009C0001005F34000101005700114506445006931933D2012220016290740F9F060007A00000027710109F410003000164
// 	G3v042 := NewG3v042VisaContactlessRequest(2995, "00084A3F", "0002", "4506445006931933D", "E0F0C89F", "000706040A03A400009F370004B4C950679F020006000000", "0004B4C9", 5700114506445006)
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v042.String(), *G3v042), "G3v042 ->", 128)

// 	G3v043 := NewG3v043NetworkIDRequest(NetworkIdVisaNet)
// 	t.Logf("%-32s%s %+v\n\n", "G3v043 ->", G3v043.String(), *G3v043)

// 	G3v044 := NewG3v044AutomatedTellerMachineRequest(CheckingAccount, CreditCardAccount)
// 	t.Logf("%-32s%s %+v\n\n", "G3v044 ->", G3v044.String(), *G3v044)

// 	G3v045 := NewG3v045IntegratedChipCardRequest()
// 	G3v045.CryptogramAmount = 600                                            //Tag 9F02
// 	G3v045.AuthorizationRequestCryptogram = "00084A3FD0EA4B18"               //Tag 9F26
// 	G3v045.ApplicationCounter = "0021"                                       //Tag 9F36
// 	G3v045.CustomerData = "000706040A03A400009F370004B4C950679F020006000000" //Tag 9F7C
// 	G3v045.FormFactor = "E0F0C89F"                                           //Tag 9F6E
// 	G3v045.UnpredictableNumber = "C6FB967F"                                  //Tag 9F37
// 	G3v045.CardSequenceNumber = 63                                           //Tag 5F34
// 	G3v045.CryptogramInformationData = 5                                     //Tag 9F27
// 	G3v045.IssuerApplicationData = "02"                                      //Tag 9F10
// 	G3v045.TerminalCountryCode = 840                                         //Tag 9F1A
// 	G3v045.IFDSerialNumber = "4422874A"                                      //Tag 9F1E
// 	G3v045.TerminalCapabilityProfile = "7F9137"                              //Tag 9F33
// 	G3v045.IssuerScriptResults = "3EFDEE0E59F74B799848A39DE4E54EDB6E9AA486"  //Tag 9F5B
// 	G3v045.ApplicationInterchangeProfile = 34                                //Tag 82
// 	G3v045.TerminalVerificationResults = 5800000                             //Tag 95
// 	G3v045.CryptogramTransactionType = 79                                    //Tag 9C
// 	G3v045.TerminalTransactionTime = "060823"                                //Tag 9F21
// 	G3v045.TerminalTransactionDate = "101314"                                //Tag 9A
// 	G3v045.CryptogramCurrencyCode = 840                                      //Tag 5F2A
// 	G3v045.CryptogramCashbackAmount = 0                                      //Tag 9F03
// 	G3v045.CardholderVerificationMethodResults = "000301"                    //Tag 9F34
// 	G3v045.TerminalType = "EB"                                               //Tag 9F35
// 	G3v045.TransactionCategoryCode = " "                                     //Tag 9F53
// 	G3v045.SecondaryPinBlock = ""                                            // Tag C0
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v045.String(), *G3v045), "G3v045 ->", 128)

// 	G3v048 := NewG3v048AmexCardholderVerificationResultsRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v048 ->", G3v048.String(), *G3v048)

// 	G3v049 := NewG3v049Gen2TerminalAuthenticationRequest("7A4DD6E75A4359BAE10B38E2")
// 	t.Logf("%-32s%s %+v\n\n", "G3v049 ->", G3v049.String(), *G3v049)

// 	G3v050 := NewG3v050AssociationTimestampRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v050 ->", G3v050.String(), *G3v050)

// 	G3v051 := NewG3v051MasterCardEventMonitoringRealTimeScoringRequest(1)
// 	t.Logf("%-32s%s %+v\n\n", "G3v051 ->", G3v051.String(), *G3v051)

// 	transblock := "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd"
// 	G3v052 := NewG3v052VoltageEncryptionTransmissionBlockRequest(transblock)
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v052.String(), *G3v052), "G3v052 ->", 128)

// 	G3v053 := NewG3v053TokenRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v053 ->", G3v053.String(), *G3v053)
// 	G3v054 := NewG3v054TransitProgramRequest(RealTimeAuthorizedTransit, CommuterTrain)
// 	t.Logf("%-32s%s %+v\n\n", "G3v054 ->", G3v054.String(), *G3v054)
// 	G3v055 := NewG3v055IntegratedChipCardEmvTlvRequest("004F0007A0000000041010005000104465626974204D617374657243617264005700115457210002001065D1412201032580745F00820002580000840007A0000000041010009500050200008000009A0003141015009B0002E800009C0001001005000230305F2A000201245F34000100900000014290010001439F02000")
// 	WrapAndPrintText(fmt.Sprintf("%s %+v\n", G3v055.String(), *G3v055), "G3v055 ->", 128)

// 	G3v056 := NewG3v056MessageReasonCodeRequest(DebitOrCreditAdjustment)
// 	t.Logf("%-32s%s %+v\n\n", "G3v056 ->", G3v056.String(), *G3v056)

// 	G3v057 := NewG3v057DiscoverOrPayPalAdditionalDataRequest("22994F6F968DBA53")
// 	t.Logf("%-32s%s %+v\n\n", "G3v057 ->", G3v057.String(), *G3v057)

// 	G3v058 := NewG3v058AlternateAccountIDRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v058 ->", G3v058.String(), *G3v058)

// 	G3v059 := NewG3v059MasterCardPayPassMappingServiceRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v059 ->", G3v059.String(), *G3v059)

// 	G3v060 := NewG3v060PayPassMobileRequest(2, MobileNetworkOperatorSecureElement)
// 	t.Logf("%-32s%s %+v\n\n", "G3v060 ->", G3v060.String(), *G3v060)

// 	G3v061 := NewG3v061SpendQualifiedIndicatorRequest()
// 	t.Logf("%-32s%s %+v\n\n", "G3v061 ->", G3v061.String(), *G3v061)

// 	G3v200 := NewG3v200GiftCardRequest("JENNIFER10", "11680-SMG7-E010-M1-0001%NH167L")
// 	t.Logf("%-32s%s %+v\n\n", "G3v200 ->", G3v200.String(), *G3v200)

// }
