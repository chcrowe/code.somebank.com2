package tsys

import (
	"strconv"

	"code.somebank.com/p/bytes"
)

type POSCheckPrivateData struct {
	FreeFormatText                 string //FXnnndddd
	IDTypeAndNumber                string //IC999dddd
	DateOfBirth                    string //DBnnndddd must be mmddccyy format eg. 03311967
	TelephoneNumber                string //PHnnndddd Number format must be AAANNNNNNN
	DialTerminalIdentificationInfo string //DT999@ddDd	Identifies a dial terminal, where 999 equals
	//length of associated data field, @ represents a terminal
	//information flag, and dddd equals data content.
	ReferenceNumber                string //= "11223344" //RN999dddd
	ProprietaryResponseInformation string //= "CHECK ACCEPTED STAMP CHECK RETURN CHECK " //RC999dddd
	ReceiptInformation             string //= "Y" //RP999dddd
	CallbackInformation            string //CI999dddd

}

func (p *POSCheckPrivateData) Copy(pc *POSCheckPrivateData) {
	p.FreeFormatText = pc.FreeFormatText
	p.IDTypeAndNumber = pc.IDTypeAndNumber
	p.DateOfBirth = pc.DateOfBirth
	p.TelephoneNumber = pc.TelephoneNumber
	p.DialTerminalIdentificationInfo = pc.DialTerminalIdentificationInfo

	p.ReferenceNumber = pc.ReferenceNumber
	p.ProprietaryResponseInformation = pc.ProprietaryResponseInformation //= "CHECK ACCEPTED STAMP CHECK RETURN CHECK " //RC999dddd
	p.ReceiptInformation = pc.ReceiptInformation
	p.CallbackInformation = pc.CallbackInformation
}

func (p *POSCheckPrivateData) String() string {
	//The request side Group III Version 16 might look like this example:
	//016IC01104222334444DB006112255PH0108885556666<FS>
	buffer := bytes.NewSafeBuffer(512)

	if 0 < len(p.FreeFormatText) {
		buffer.AppendFormat("FX%03d%s", len(p.FreeFormatText), p.FreeFormatText) //FXnnndddd
	}
	if 0 < len(p.IDTypeAndNumber) {
		buffer.AppendFormat("IC%03d%s", len(p.IDTypeAndNumber), p.IDTypeAndNumber) //IC999dddd
	}
	if 0 < len(p.DateOfBirth) {
		buffer.AppendFormat("DB%03d%s", len(p.DateOfBirth), p.DateOfBirth) //DBnnndddd must be mmddccyy format eg. 03311967
	}
	if 0 < len(p.TelephoneNumber) {
		buffer.AppendFormat("PH%03d%s", len(p.TelephoneNumber), p.TelephoneNumber) //PHnnndddd Number format must be AAANNNNNNN
	}
	if 0 < len(p.DialTerminalIdentificationInfo) {
		buffer.AppendFormat("DT%03d%s", len(p.DialTerminalIdentificationInfo), p.DialTerminalIdentificationInfo) //DT999@ddDd	Identifies a dial terminal, where 999 equals
	}
	//length of associated data field, @ represents a terminal
	//information flag, and dddd equals data content.
	if 0 < len(p.ReferenceNumber) {
		buffer.AppendFormat("RN%03d%s", len(p.ReferenceNumber), p.ReferenceNumber) //RN999dddd
	}
	if 0 < len(p.ProprietaryResponseInformation) {
		buffer.AppendFormat("RC%03d%s", len(p.ProprietaryResponseInformation), p.ProprietaryResponseInformation) //RC999dddd
	}
	if 0 < len(p.ReceiptInformation) {
		buffer.AppendFormat("RP%03d%s", len(p.ReceiptInformation), p.ReceiptInformation) //RP999dddd
	}
	if 0 < len(p.CallbackInformation) {
		buffer.AppendFormat("CI%03d%s", len(p.CallbackInformation), p.CallbackInformation) //CI999dddd
	}

	//#warning please remove (DEBUGGING ONLY)
	//			return "RN00811223344RC048CHECK ACCEPTED STAMP CHECK RETURN CHECK RP001Y";
	return buffer.String()
}

func ParseLengthValuePair(s string, index int) string {
	length, _ := strconv.ParseInt(s[index+2:index+5], 10, 32)
	return s[index+5 : length]
}

func ParsePOSCheckPrivateData(s string) *POSCheckPrivateData {
	p := POSCheckPrivateData{}
	//The request side Group III Version 16 might look like this example:
	//016IC01104222334444DB006112255PH0108885556666<FS>
	for i := 0; i < len(s)-1; i++ {
		switch s[i : i+2] {
		case "FX":
			p.FreeFormatText = ParseLengthValuePair(s, i) //FXnnndddd
		case "IC":
			p.IDTypeAndNumber = ParseLengthValuePair(s, i) //IC999dddd
		case "DB":
			p.DateOfBirth = ParseLengthValuePair(s, i) //DBnnndddd must be mmddccyy format eg. 03311967
		case "PH":
			p.TelephoneNumber = ParseLengthValuePair(s, i) //PHnnndddd Number format must be AAANNNNNNN
		case "DT":
			p.DialTerminalIdentificationInfo = ParseLengthValuePair(s, i) //DT999@ddDd	Identifies a dial terminal, where 999 equals
		case "RN":
			p.ReferenceNumber = ParseLengthValuePair(s, i) //RN999dddd
		case "RC":
			p.ProprietaryResponseInformation = ParseLengthValuePair(s, i) //RC999dddd
		case "RP":
			p.ReceiptInformation = ParseLengthValuePair(s, i) //RP999dddd
		case "CI":
			p.CallbackInformation = ParseLengthValuePair(s, i) //CI999dddd
		default:
			break
		}
	}
	return &p
}
