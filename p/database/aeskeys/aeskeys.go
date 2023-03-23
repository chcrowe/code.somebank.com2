package aeskeys

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"fmt"

	ecrypto "code.somebank.com/p/crypto"
	"code.somebank.com/p/database"
	"code.somebank.com/p/list"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetMerchantAesCipher(dsn string, cid uint) (cipher.Block, []byte) {

	db_s, err_sec := sql.Open("mssql", dsn)
	if err_sec != nil {
		//fmt.Println("Cannot connect to security db: ", err_sec.Error())
		return nil, nil
	}
	defer db_s.Close()

	k, err2 := getKeys(db_s, cid)
	if err2 != nil {
		// fmt.Println(err2)
		return nil, nil
	}

	merchantkey, merchantiv := ExpandMerchantKeyset(k["base64Key"], k["base64Vector"])

	// fmt.Printf("\nKEY EXPANDED: %v, %s (%d)", merchantkey, base64.StdEncoding.EncodeToString(merchantkey), len(merchantkey))
	// fmt.Printf("\nIV EXPANDED:  %v, %s (%d)", merchantiv, base64.StdEncoding.EncodeToString(merchantiv), len(merchantiv))

	tek_block, err_tek := aes.NewCipher(merchantkey)
	if err_tek != nil {
		//fmt.Printf("\nCannot initialize TEK: %s", err_tek.Error())
		return nil, nil
	}

	// acctno := ecrypto.DecryptFromBase64String(tek_block, merchantiv, t["AccountNumber"])
	// nameoncard := ecrypto.DecryptFromBase64String(tek_block, merchantiv, t["NameOnCard"])
	return tek_block, merchantiv
}

func ExpandMerchantKeyset(b64EncryptedKey, b64EncryptedVector string) ([]byte, []byte) {

	b64KeyEncryptingKey := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX="
	b64KeyEncryptingVector := "XXXXXXXXXXXXXXXXXXXXXXXX"
	ivbytes, _ := base64.StdEncoding.DecodeString(b64KeyEncryptingVector)

	mode, err_kek := ecrypto.CreateAesCipherFromBase64Key(b64KeyEncryptingKey)
	if err_kek != nil {
		return nil, nil
	}

	tekKey := ecrypto.DecryptFromBase64String(mode, ivbytes, b64EncryptedKey)
	tekVector := ecrypto.DecryptFromBase64String(mode, ivbytes, b64EncryptedVector)
	tekKeyBytes, _ := base64.StdEncoding.DecodeString(string(tekKey))
	tekVectorBytes, _ := base64.StdEncoding.DecodeString(string(tekVector))

	return tekKeyBytes, tekVectorBytes
}

func getKeys(db *sql.DB, cid uint) (list.NameValuePairMap, error) {

	cmd := fmt.Sprintf("exec dbo._spXXXXXXXXXXX_Select %d ", cid)

	rows, err := db.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}

	m := make(map[string]string)

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			continue
		}
		for i := 0; i < len(vals); i++ {
			m[cols[i]] = database.GetValue(vals[i].(*interface{}))
		}
	}

	if rows.Err() != nil {
		return m, rows.Err()
	}

	return m, nil
}
