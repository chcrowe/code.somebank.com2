package auth

import (
	"encoding/xml"
	"strings"

	"code.somebank.com/p/list"
)

// Child - Version identification of the Core ( SecureVault30Core or SecureVault255Core)
// CoreType - The type of Core being requested (i.e. SecureVault30Core or SecureVault255Core)
// Parent - Version identification of the Api ( SecureVault30Api or SecureVault255Api)
// SettlementBatchAmount

type RequestMessage struct {
	AVSStreet1                  string `xml:"AVSStreet1"` //
	AVSStreet2                  string `xml:"AVSStreet2"` //
	AVSZip                      string `xml:"AVSZip"`     //
	AccountDataSource           string `xml:"AccountDataSource"`
	AccountNumber               string `xml:"AccountNumber"` //4012111111111111/AccountNumber
	Action                      string `xml:"Action"`        //AUTHORIZE_ONLINE/Action
	ApprovalCode                string `xml:"ApprovalCode"`
	CAVV                        string `xml:"CAVV"`
	CVV                         string `xml:"CVV"`      //
	CardType                    string `xml:"CardType"` //VISA/CardType
	CashbackAmount              string `xml:"CashbackAmount"`
	CheckInDate                 string `xml:"CheckInDate"`      //
	CheckOutDate                string `xml:"CheckOutDate"`     //
	ClientID                    string `xml:"ClientID"`         //5562/ClientID
	ComponentVersion            string `xml:"ComponentVersion"` //4.75.2.0/ComponentVersion
	CustomData                  string `xml:"CustomData"`       //
	CustomerEmail               string `xml:"CustomerEmail"`
	CustomerPhoneNumber         string `xml:"CustomerPhoneNumber"`
	DCCData                     string `xml:"DCCData"` //
	DeliveryDate                string `xml:"DeliveryDate"`
	DeviceSerialNumber          string `xml:"DeviceSerialNumber"`
	EncryptedDeviceData         string `xml:"EncryptedDeviceData"`
	EncryptedPINData            string `xml:"EncryptedPINData"`
	EncryptedTrack1Data         string `xml:"EncryptedTrack1Data"`
	EncryptedTrack2Data         string `xml:"EncryptedTrack2Data"`
	EncryptedTrack3Data         string `xml:"EncryptedTrack3Data"`
	EntryMethod                 string `xml:"EntryMethod"`    //KEYED_CARD_PRESENT/EntryMethod
	ExpirationDate              string `xml:"ExpirationDate"` //1215/ExpirationDate
	GratuityAmount              string `xml:"GratuityAmount"` //
	InstallmentNumber           string `xml:"InstallmentNumber"`
	InstallmentCount            string `xml:"InstallmentCount"`
	LocalTransactionDate        string `xml:"LocalTransactionDate"`
	MSRKeySerialNumber          string `xml:"MSRKeySerialNumber"`
	MasterKey                   string `xml:"MasterKey"` //535061/MasterKey
	MerchantVerificationValue   string `xml:"MerchantVerificationValue"`
	MessageAuthenticationCode   string `xml:"MessageAuthenticationCode"`
	NameOnCard                  string `xml:"NameOnCard"`      //
	NoShowIndicator             string `xml:"NoShowIndicator"` //
	OperatorID                  string `xml:"OperatorID"`      //32550BE3AC/OperatorID
	OriginalTransactionDateTime string `xml:"OriginalTransactionDateTime"`
	PINPadSequenceNumber        string `xml:"PINPadSequenceNumber"`
	PurchaseIdentifier1         string `xml:"PurchaseIdentifier1"` //98C6A118/PurchaseIdentifier1
	PurchaseIdentifier2         string `xml:"PurchaseIdentifier2"` //61225A/PurchaseIdentifier2
	PurchaseOrderNumber         string `xml:"PurchaseOrderNumber"`
	PurchaserEMail              string `xml:"PurchaserEMail"`
	PurchaserName               string `xml:"PurchaserName"`
	PurchaserNote               string `xml:"PurchaserNote"`
	PurchaserPhoneNumber        string `xml:"PurchaserPhoneNumber"`
	ReferenceTransactionID      string `xml:"ReferenceTransactionID"` //3B6D7730-EA76-4036-BB1F-9B66180F1C3B/ReferenceTransactionID
	ResendIndicator             string `xml:"ResendIndicator"`        //0/ResendIndicator
	RoomRateAmount              string `xml:"RoomRateAmount"`         //
	SettlementAmount            string `xml:"SettlementAmount"`
	SignatureData               string `xml:"SignatureData"` //
	StayDuration                string `xml:"StayDuration"`  //
	TaxAmount                   string `xml:"TaxAmount"`
	TaxPercentage               string `xml:"TaxPercentage"`
	TerminalID                  string `xml:"TerminalID"` //303/TerminalID
	TrackData                   string `xml:"TrackData"`
	TrackRead                   string `xml:"TrackRead"`
	TransactionAmount           string `xml:"TransactionAmount"` //55.99/TransactionAmount
	TransactionID               string `xml:"TransactionID"`     //3B6D7730-EA76-4036-BB1F-9B66180F1C3B/TransactionID
	TransactionSequenceNumber   string `xml:"TransactionSequenceNumber"`
	UCAF                        string `xml:"UCAF"`
	VendorID                    string `xml:"VendorID"`      //S34/VendorID
	VendorName                  string `xml:"VendorName"`    //ACME BOOK STORE/VendorName
	VendorProduct               string `xml:"VendorProduct"` //[TV 2.55.17]ACME .Net Point of Sale 1.0/VendorProduct
	VendorVersion               string `xml:"VendorVersion"`
	VoucherApprovalCode         string `xml:"VoucherApprovalCode"`
	VoucherSerialNumber         string `xml:"VoucherSerialNumber"`
	XID                         string `xml:"XID"`
}

func (r *RequestMessage) Xml() string {
	b, _ := xml.Marshal(r)
	return string(b)
}

func (r *RequestMessage) Values() string {

	m := list.NameValuePairMap{}

	m.Set("AccountDataSource", r.AccountDataSource)
	m.Set("AccountNumber", r.AccountNumber)
	m.Set("Action", r.Action)
	m.Set("ApprovalCode", r.ApprovalCode)
	m.Set("AVSStreet1", r.AVSStreet1)
	m.Set("AVSStreet2", r.AVSStreet2)
	m.Set("AVSZip", r.AVSZip)
	m.Set("CardType", r.CardType)
	m.Set("CashbackAmount", r.CashbackAmount)
	m.Set("CheckInDate", r.CheckInDate)
	m.Set("CheckOutDate", r.CheckOutDate)
	m.Set("ClientID", r.ClientID)
	m.Set("ComponentVersion", r.ComponentVersion)
	m.Set("CustomData", r.CustomData)
	m.Set("CustomerEmail", r.CustomerEmail)
	m.Set("CustomerPhoneNumber", r.CustomerPhoneNumber)
	m.Set("CAVV", r.CAVV)
	m.Set("CVV", r.CVV)
	m.Set("DCCData", r.DCCData)
	m.Set("DeliveryDate", r.DeliveryDate)
	m.Set("DeviceSerialNumber", r.DeviceSerialNumber)
	m.Set("EncryptedDeviceData", r.EncryptedDeviceData)
	m.Set("EncryptedPINData", r.EncryptedPINData)
	m.Set("EncryptedTrack1Data", r.EncryptedTrack1Data)
	m.Set("EncryptedTrack2Data", r.EncryptedTrack2Data)
	m.Set("EncryptedTrack3Data", r.EncryptedTrack3Data)
	m.Set("EntryMethod", r.EntryMethod)
	m.Set("ExpirationDate", r.ExpirationDate)
	m.Set("GratuityAmount", r.GratuityAmount)
	m.Set("LocalTransactionDate", r.LocalTransactionDate)
	m.Set("MasterKey", r.MasterKey)
	m.Set("MerchantVerificationValue", r.MerchantVerificationValue)
	m.Set("MSRKeySerialNumber", r.MSRKeySerialNumber)
	m.Set("NameOnCard", r.NameOnCard)
	m.Set("NoShowIndicator", r.NoShowIndicator)
	m.Set("OperatorID", r.OperatorID)
	m.Set("OriginalTransactionDateTime", r.OriginalTransactionDateTime)
	m.Set("PINPadSequenceNumber", r.PINPadSequenceNumber)
	m.Set("PurchaseIdentifier1", r.PurchaseIdentifier1)
	m.Set("PurchaseIdentifier2", r.PurchaseIdentifier2)
	m.Set("PurchaserEMail", r.PurchaserEMail)
	m.Set("PurchaserName", r.PurchaserName)
	m.Set("PurchaserNote", r.PurchaserNote)
	m.Set("PurchaserPhoneNumber", r.PurchaserPhoneNumber)
	m.Set("ReferenceTransactionID", r.ReferenceTransactionID)
	m.Set("RoomRateAmount", r.RoomRateAmount)
	m.Set("SignatureData", r.SignatureData)
	m.Set("StayDuration", r.StayDuration)
	m.Set("TaxAmount", r.TaxAmount)
	m.Set("TaxPercentage", r.TaxPercentage)
	m.Set("TerminalID", r.TerminalID)
	m.Set("TrackRead", r.TrackRead)
	m.Set("TrackData", r.TrackData)
	m.Set("TransactionAmount", r.TransactionAmount)
	m.Set("TransactionID", r.TransactionID)
	m.Set("UCAF", r.UCAF)
	m.Set("VendorID", r.VendorID)
	m.Set("VendorName", r.VendorName)
	m.Set("VendorProduct", r.VendorProduct)
	m.Set("VendorVersion", r.VendorVersion)
	m.Set("VoucherApprovalCode", r.VoucherApprovalCode)
	m.Set("VoucherSerialNumber", r.VoucherSerialNumber)
	m.Set("XID", r.XID)

	return m.String()
}

func UnmarshallFromXml(s string, r *RequestMessage) {
	xml.Unmarshal([]byte(s), &r)
}

func UnmarshallFromNameValuePairs(s string, r *RequestMessage) {

	m := list.ParseNameValuePairMap(s)

	r.AccountDataSource = m.Get("AccountDataSource")
	r.AccountNumber = m.Get("AccountNumber")
	r.Action = m.Get("Action")
	r.ApprovalCode = m.Get("ApprovalCode")
	r.AVSStreet1 = m.Get("AVSStreet1")
	r.AVSStreet2 = m.Get("AVSStreet2")
	r.AVSZip = m.Get("AVSZip")
	r.CardType = m.Get("CardType")
	r.CashbackAmount = m.Get("CashbackAmount")
	r.CheckInDate = m.GetAny([]string{"CHID", "CheckInDate"})
	r.CheckOutDate = m.GetAny([]string{"CHOD", "CheckOutDate"})
	r.ClientID = m.Get("ClientID")
	r.ComponentVersion = m.Get("ComponentVersion")
	r.CustomData = m.Get("CustomData")
	r.CustomerEmail = m.Get("CustomerEmail")
	r.CustomerPhoneNumber = m.Get("CustomerPhoneNumber")
	r.CVV = m.Get("CVV")
	r.CAVV = m.Get("CAVV")
	r.DCCData = m.Get("DCCData")
	r.DeliveryDate = m.Get("DeliveryDate")
	r.DeviceSerialNumber = m.Get("DeviceSerialNumber")
	r.EncryptedDeviceData = m.Get("EncryptedDeviceData")
	r.EncryptedPINData = m.Get("EncryptedPINData")
	r.EncryptedTrack1Data = m.Get("EncryptedTrack1Data")
	r.EncryptedTrack2Data = m.Get("EncryptedTrack2Data")
	r.EncryptedTrack3Data = m.Get("EncryptedTrack3Data")
	r.EntryMethod = m.Get("EntryMethod")
	r.ExpirationDate = m.Get("ExpirationDate")
	r.GratuityAmount = m.Get("GratuityAmount")
	r.LocalTransactionDate = m.Get("LocalTransactionDate")
	r.MasterKey = m.Get("MasterKey")
	r.MerchantVerificationValue = m.Get("MerchantVerificationValue")
	r.MSRKeySerialNumber = m.Get("MSRKeySerialNumber")
	r.NameOnCard = m.Get("NameOnCard")
	r.NoShowIndicator = m.Get("NoShowIndicator")
	r.OperatorID = m.Get("OperatorID")
	r.OriginalTransactionDateTime = m.GetAny([]string{"OTDT", "OriginalTransactionDateTime"})
	r.PINPadSequenceNumber = m.Get("PINPadSequenceNumber")
	r.PurchaseIdentifier1 = m.Get("PurchaseIdentifier1")
	r.PurchaseIdentifier2 = m.Get("PurchaseIdentifier2")
	r.PurchaserEMail = m.Get("PurchaserEMail")
	r.PurchaserName = m.Get("PurchaserName")
	r.PurchaserNote = m.Get("PurchaserNote")
	r.PurchaserPhoneNumber = m.Get("PurchaserPhoneNumber")
	r.ReferenceTransactionID = m.Get("ReferenceTransactionID")
	r.RoomRateAmount = m.GetAny([]string{"RRA", "RoomRateAmount"})
	r.SignatureData = m.Get("SignatureData")
	r.StayDuration = m.GetAny([]string{"STYD", "StayDuration"})
	r.TaxAmount = m.Get("TaxAmount")
	r.TerminalID = m.Get("TerminalID")
	r.TrackData = m.Get("TrackData")
	r.TransactionAmount = m.Get("TransactionAmount")
	r.TransactionID = m.Get("TransactionID")
	r.UCAF = m.Get("UCAF")
	r.VendorID = m.Get("VendorID")
	r.VendorName = m.Get("VendorName")
	r.VendorProduct = m.Get("VendorProduct")
	r.VendorVersion = m.Get("VendorVersion")
	r.VoucherApprovalCode = m.GetAny([]string{"VAC", "VoucherApprovalCode"})
	r.VoucherSerialNumber = m.GetAny([]string{"VSN", "VoucherSerialNumber"})
	r.XID = m.Get("XID")
}

func ParseRequestMessageFromPOSCP3Values(s string) *RequestMessage {

	r := RequestMessage{}
	m := list.ParseNameValuePairMap(s)

	r.AccountNumber = m.Get("AN")
	r.Action = strings.ToUpper(m.Get("A"))
	r.ApprovalCode = m.Get("AC")
	r.AVSStreet1 = m.Get("AS1")
	r.AVSStreet2 = m.Get("AS2")
	r.AVSZip = m.Get("AZ")
	r.CardType = m.Get("CT")
	r.ClientID = m.Get("CID")
	r.EntryMethod = m.Get("DS")
	r.ExpirationDate = m.Get("ED")
	r.GratuityAmount = m.Get("GA")
	r.MasterKey = m.Get("PWD")
	r.NameOnCard = m.Get("NOC")
	r.PurchaseIdentifier1 = m.Get("PID1")
	r.PurchaseIdentifier2 = m.Get("PID2")
	r.SettlementAmount = m.Get("SA")
	r.TrackData = m.Get("TD")
	r.TrackRead = m.Get("TR")
	r.TransactionAmount = m.Get("TA")
	r.TransactionSequenceNumber = m.Get("TSN")

	return &r
}

func ParseRequestMessageFromPOSCP308Values(s string) *RequestMessage {

	r := RequestMessage{}
	m := list.ParseNameValuePairMap(s)

	r.AccountNumber = m.Get("AN")
	r.Action = strings.ToUpper(m.Get("A"))
	r.ApprovalCode = m.Get("AC")
	r.AVSStreet1 = m.Get("AS1")
	r.AVSZip = m.Get("AZ")
	r.CardType = m.Get("CT")
	r.ClientID = m.Get("CID")
	r.EncryptedTrack1Data = m.Get("ET1D")
	r.EncryptedTrack2Data = m.Get("ET2D")
	r.EncryptedTrack3Data = m.Get("ET3D")
	r.EntryMethod = m.Get("DS")
	r.ExpirationDate = m.Get("ED")
	r.GratuityAmount = m.Get("GA")
	r.MasterKey = m.Get("PWD")
	r.MSRKeySerialNumber = m.Get("MKSN")
	r.NameOnCard = m.Get("NOC")
	r.PurchaseIdentifier1 = m.Get("PID1")
	r.PurchaseIdentifier2 = m.Get("PID2")
	r.SettlementAmount = m.Get("SA")
	r.TrackData = m.Get("TD")
	r.TrackRead = m.Get("TR")
	r.TransactionAmount = m.Get("TA")
	r.TransactionSequenceNumber = m.Get("TSN")

	return &r
}

func ParseRequestMessageFromPOSCP4Values(s string) *RequestMessage {

	r := RequestMessage{}
	m := list.ParseNameValuePairMap(s)

	r.AccountNumber = m.Get("AN")
	r.Action = strings.ToUpper(m.Get("A"))
	r.ApprovalCode = m.Get("AC")
	r.AVSStreet1 = m.Get("AS1")
	r.AVSStreet2 = m.Get("AS2")
	r.AVSZip = m.Get("AZ")
	r.CashbackAmount = m.Get("CA")
	r.CheckInDate = m.GetAny([]string{"CHID", "CheckInDate"})
	r.CheckOutDate = m.GetAny([]string{"CHOD", "CheckOutDate"})
	r.ClientID = m.Get("CID")
	r.CVV = m.Get("CVV")
	r.DeviceSerialNumber = m.Get("DeviceSerialNumber")
	r.EncryptedPINData = m.Get("EPD")
	r.EncryptedTrack1Data = m.Get("ET1D")
	r.EncryptedTrack2Data = m.Get("ET2D")
	r.EncryptedTrack3Data = m.Get("ET3D")
	r.EntryMethod = m.Get("EM")
	r.ExpirationDate = m.Get("ED")
	r.GratuityAmount = m.Get("GA")
	r.MasterKey = m.Get("MK")
	r.MSRKeySerialNumber = m.Get("MKSN")
	r.NameOnCard = m.Get("NOC")
	r.OperatorID = m.Get("OID")
	r.OriginalTransactionDateTime = m.GetAny([]string{"OTDT", "OriginalTransactionDateTime"})
	r.PINPadSequenceNumber = m.Get("PPSN")
	r.PurchaseIdentifier1 = m.Get("PID1")
	r.PurchaseIdentifier2 = m.Get("PID2")
	r.ReferenceTransactionID = m.Get("RTID")
	r.ResendIndicator = m.Get("RI")
	r.RoomRateAmount = m.GetAny([]string{"RRA", "RoomRateAmount"})
	r.SettlementAmount = m.Get("SA")
	r.SignatureData = m.Get("SD")
	r.StayDuration = m.GetAny([]string{"STYD", "StayDuration"})
	r.TaxAmount = m.Get("TXA")
	r.TaxPercentage = m.Get("TXP")
	r.TerminalID = m.Get("TLID")
	r.TrackData = m.Get("TD")
	r.TransactionAmount = m.Get("TA")
	r.TransactionID = m.Get("TID")
	r.TransactionSequenceNumber = m.Get("TSN")
	r.VendorProduct = m.Get("VP")
	r.VoucherApprovalCode = m.GetAny([]string{"VAC", "VoucherApprovalCode"})
	r.VoucherSerialNumber = m.GetAny([]string{"VSN", "VoucherSerialNumber"})

	return &r
}
