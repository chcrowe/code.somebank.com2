package auth

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"code.google.com/p/go-uuid/uuid"
	ecrypto "code.somebank.com/p/crypto"
	"code.somebank.com/p/list"
)

type AuthorizationRecord struct {
	Utc string
	Id  string
	Out string
	In  string
	Pii string
}

func NewAuthorizationRecord(request, response, sensitivedata, salt string) *AuthorizationRecord {
	ar := AuthorizationRecord{}
	utm := time.Now().UTC()

	ar.Utc = fmt.Sprintf("%04d%02d%02dT%02d%02d%02d:%d", utm.Year(), utm.Month(), utm.Day(), utm.Hour(), utm.Minute(), utm.Second(), utm.Nanosecond())
	ar.Id = uuid.New()
	ar.Out = request
	ar.In = response

	if 0 < len(sensitivedata) {
		ar.Pii = sensitivedata
		encryptPii(&ar, salt)
	}

	return &ar
}

func decryptPii(a *AuthorizationRecord, salt string) {
	arw := ecrypto.NewAesReaderWriter([]string{a.Utc, a.Id, a.In}, []byte(salt))
	encrypted, _ := hex.DecodeString(a.Pii)
	decrypted := arw.Decrypt(encrypted)
	a.Pii = string(decrypted)
}

func encryptPii(a *AuthorizationRecord, salt string) {
	arw := ecrypto.NewAesReaderWriter([]string{a.Utc, a.Id, a.In}, []byte(salt))
	encrypted := arw.Encrypt([]byte(a.Pii))
	a.Pii = hex.EncodeToString(encrypted)
}

func ParseAuthorizationRecordFromValues(s, salt string) *AuthorizationRecord {

	m := list.ParseNameValuePairMap(s)
	a := AuthorizationRecord{m.Get("Utc"), m.Get("Id"), m.Get("Out"), m.Get("In"), m.Get("Pii")}

	if 0 < len(salt) {
		decryptPii(&a, salt)
	}

	return &a
}

func ParseAuthorizationRecordFromJson(b []byte, salt string) *AuthorizationRecord {

	a := AuthorizationRecord{}
	err := json.Unmarshal(b, &a)
	if err != nil {
		panic(err)
	}

	if 0 < len(salt) {
		decryptPii(&a, salt)
	}

	return &a
}

func (a *AuthorizationRecord) Values() string {
	return fmt.Sprintf("Utc=%s&Id=%s&Out=%s&In=%s&Pii=%s&", a.Utc, a.Id, a.Out, a.In, a.Pii)
}

func (a *AuthorizationRecord) Json() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func UnmarshalFromFile(fname string) ([]AuthorizationRecord, error) {
	dat, e := ioutil.ReadFile(fname)
	var ars []AuthorizationRecord
	if nil == e {
		e = json.Unmarshal(dat, &ars)
	}
	return ars, e
}
