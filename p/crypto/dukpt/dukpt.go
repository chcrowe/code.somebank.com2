package dukpt

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
	"strings"

	estrings "code.somebank.com/p/strings"
)

// EXAMPLE
// bdk := ""00000000-0000-0000-0000-000000000000""
// ksn := "FFFF9999999999999999"
// session_key := GenerateDukptSessionKey(bdk, ksn)
// t.Logf("session key:    %s", session_key)

func GenerateInitialPinEncryptionKey(baseDerivationKey, keypadSerialNumber string) string {
	// The IPEK is going to consist of two 8 byte registers that are produced by two separate TripleDES algorithms.
	// Both will encrypt the 8 most significant bytes of the KSN with the counter zeroed out. The difference between
	// the left and right register is that the right register is going to encrypt the KSN with a slightly modified
	// version of the BDK.

	// Let's assume we have a 16 byte BDK represented as the hexadecimal string "0123456789ABCDEFFEDCBA9876543210".
	// bdk16 := "0123456789ABCDEFFEDCBA9876543210"
	// fmt.Printf("bdk 16:               %s\n", bdk16)

	// We are going to form the BDK used to encrypt the left register of our IPEK by appending the first 8 bytes to
	// the end of the BDK to give us the following 24 byte key.
	bdk24 := estrings.RightPadToLength(baseDerivationKey, baseDerivationKey[:16], 48)

	// We also have a 10 byte KSN with an 8 count represented as the hexadecimal string "FFFF9876543210E00008".
	// ksn10 := "FFFF9876543210E00008"

	ksn10 := estrings.LeftPadToLength(keypadSerialNumber, "F", 20)

	// Now if the KSN is not already 10 bytes in length then pad it to 10 bytes with hexadecimal "F" (1111). It is
	// important to note that the IPEK represents the very first key on the device. This means that we want to generate
	// it with the counter portion of the KSN set to 0. In order to get our KSN with a 0 counter we want to bitmask it
	// with the hexadecimal represented by the string "FFFFFFFFFFFFFFE00000".
	// 		    FFFF9876543210E00008
	// 		and FFFFFFFFFFFFFFE00000
	// 		=   FFFF9876543210E00000

	ksn := AndHexKeyPair(ksn10, "FFFFFFFFFFFFFFE00000")
	// fmt.Printf("ksn:                  %X %v\n", ksn, ksn)

	// Great. Now we have our 0 counter KSN. But we only want the most significant 8 bytes of this KSN. We get this by bit
	// shifting this KSN to the right 16 bits.
	// FFFF9876543210E00000 >> 16 = FFFF9876543210E0
	// fmt.Printf("ksn >> 16:            %X     %v\n", ksn[:8], ksn[:8])

	// The next step is to TripleDES encrypt "FFFF9876543210E0" with the 24 byte "0123456789ABCDEFFEDCBA98765432100123456789ABCDEF"
	// BDK. The result of this encryption should give us the left register of our IPEK.
	// 6AC292FAA1315B4D

	ipek_leftregister := TripleDesEncrypt(bdk24, hex.EncodeToString(ksn[:8]))
	// fmt.Printf("ipek left register:   %X     %v\n", ipek_leftregister, ipek_leftregister)

	// the right register will use a slightly modified version of the BDK to encrypt our KSN. In order to do this we want to start
	// with our original 16 byte BDK "0123456789ABCDEFFEDCBA9876543210" and XOR it with the following mask "C0C0C0C000000000C0C0C0C000000000".
	// This seems to be a totally arbitrary mask, but alas, it's required to get the right IPEK.
	//     0123456789ABCDEFFEDCBA9876543210
	// xor C0C0C0C000000000C0C0C0C000000000
	// =   C1E385A789ABCDEF3E1C7A5876543210

	// fmt.Println()
	bdk2_a := XorHexKeyPair(bdk24[:32], "C0C0C0C000000000C0C0C0C000000000")
	// fmt.Printf("bdk #2A:              %X %v\n", bdk2_a, bdk2_a)

	bdk2_b := fmt.Sprintf("%X%X", bdk2_a, bdk2_a[:8])
	// We are going to do the same thing we did with the key for the left register and take the most significant 8 bytes and append it to the
	// end to get the following 24 byte key.
	// C1E385A789ABCDEF3E1C7A5876543210C1E385A789ABCDEF
	// fmt.Printf("bdk #2B:              %s\n", bdk2_b)

	// Go ahead and take the most significant 8 bytes of the KSN with the counter zeroed (as we computed earlier) and TripleDES encrypt it
	// with this new key we just generated. This will give you the right register of your IPEK, producing the following 16 byte IPEK
	// (I've separated the left and right registers for clarity).
	// 6AC292FAA1315B4D 858AB3A3D7D5933A

	ipek_rightregister := TripleDesEncrypt(bdk2_b, hex.EncodeToString(ksn[:8]))
	// fmt.Printf("ipek right register:  %X     %v\n", ipek_rightregister, ipek_rightregister)

	return fmt.Sprintf("%X%X", ipek_leftregister, ipek_rightregister)
}

func GenerateDukptSessionKey(bdk, keypadSerialNumber string) string {

	ipek := GenerateInitialPinEncryptionKey(bdk, keypadSerialNumber)

	ksn10 := estrings.LeftPadToLength(keypadSerialNumber, "F", 20)

	right_sk := GenerateDukptSessionKeyRegister(ipek, ksn10[4:])

	// This value is the least significant 8 bytes of our session key! We now just need to repeat the above operation once more with
	// different inputs.  This time we are going to take current_sk and XOR it with "C0C0C0C000000000C0C0C0C000000000". As far as I
	// can tell this value is arbitrary but is part of the ANSI standard so you will just have to take my word for it.

	//     6AC292FAA1315B4D858AB3A3D7D5933A
	// XOR C0C0C0C000000000C0C0C0C000000000
	// =   AA02523AA1315B4D454A7363D7D5933A

	session_key_mod := XorHexKeyPair(ipek, "C0C0C0C000000000C0C0C0C000000000")

	// If we take this new value "AA02523AA1315B4D454A7363D7D5933A" and use it in place of the current_sk in the operation I described
	// above we should get "27F66D5244FF62E1". This is the most significant 8 bytes of our session key. Combined it should be
	// "27F66D5244FF62E1AA6F6120EDEB4280".

	left_sk := GenerateDukptSessionKeyRegister(fmt.Sprintf("%X", session_key_mod), ksn10[4:])

	sk_wcounter := fmt.Sprintf("%s%s", left_sk, right_sk)

	// There is one last operation we have to perform on this value before we've generated our final permutation of the key that is going to allow us to decrypt our data. We must XOR it with "00000000000000FF00000000000000FF".

	//     27F66D5244FF62E1AA6F6120EDEB4280
	// XOR 00000000000000FF00000000000000FF
	// =   27F66D5244FF621EAA6F6120EDEB427F

	session_key := XorHexKeyPair(sk_wcounter, "00000000000000FF00000000000000FF")
	return fmt.Sprintf("%X", session_key)
}

func GenerateDukptSessionKeyRegister(current_sk, ksn_mod string) string {
	// here is the algorithm that generates our future keys.  This black box takes the current
	// session key, which I will refer to as current_sk, and a modification of the KSN, which I will refer to as ksn_mod.

	// If we begin with the assumption that the IPEK we generated above was passed in as the current_sk and that our ksn_mod is
	// "9876543210E00008" that we also generated above.

	// To begin with we want to take the current_sk and get the most significant 8 bytes and bit shift it 64 bits to the right to
	// get "6AC292FAA1315B4".  This can be done by performing the following bitmask.

	//      6AC292FAA1315B4D858AB3A3D7D5933A
	// AND  FFFFFFFFFFFFFFFF0000000000000000
	// =    6AC292FAA1315B4D0000000000000000

	// mostsignificant8 := AndHexKeyPair(current_sk, "FFFFFFFFFFFFFFFF0000000000000000")

	// At this point we just need to bit shift it 64 bits to the right to get "6AC292FAA1315B4D". We will call this value left_key
	// (seeing as how it's the left side of the current_sk).

	// left_key := mostsignificant8[:16]
	left_key := current_sk[:16]

	// Next we want to get the 8 least significant bytes of the current_sk which can be done by the following bitmask.

	//     6AC292FAA1315B4D858AB3A3D7D5933A
	// AND 0000000000000000FFFFFFFFFFFFFFFF
	// =   0000000000000000858AB3A3D7D5933A
	// leastsignificant8 := AndHexKeyPair(current_sk, "0000000000000000FFFFFFFFFFFFFFFF")
	// right_key := leastsignificant8[16:]
	right_key := current_sk[16:]

	// Let's call this value (you guessed it) right_key. Now we are going to take the right_key and XOR it with the ksn_mod value
	// "9876543210E00008".

	//     858AB3A3D7D5933A
	// XOR 9876543210E00008
	// =   1DFCE791C7359332

	message := XorHexKeyPair(right_key, ksn_mod)

	// This value we will call the message. Next we are going to take this message value DES encrypt it (that's single DES). We are
	// going to pass the message into the DES algo as the content to be encrypted and we are going to pass in left_key as the key to
	// encrypt it with. This should give us "2FE5D2833A3ED1BA". We now need to XOR this value with right_key.

	message_des := SingleDesEncrypt(left_key, hex.EncodeToString(message))

	//     2FE5D2833A3ED1BA
	// XOR 858AB3A3D7D5933A
	// =   AA6F6120EDEB4280
	message_xor := XorHexKeyPair(hex.EncodeToString(message_des), right_key)

	// This value is the least significant 8 bytes of our session key! We now just need to repeat the above operation once more with
	// different inputs.  This time we are going to take current_sk and XOR it with "C0C0C0C000000000C0C0C0C000000000". As far as I
	// can tell this value is arbitrary but is part of the ANSI standard so you will just have to take my word for it.

	//     6AC292FAA1315B4D858AB3A3D7D5933A
	// XOR C0C0C0C000000000C0C0C0C000000000
	// =   AA02523AA1315B4D454A7363D7D5933A

	//	session_key := XorHexKeyPair(current_sk, "C0C0C0C000000000C0C0C0C000000000")

	// If we take this new value "AA02523AA1315B4D454A7363D7D5933A" and use it in place of the current_sk in the operation I described
	// above we should get "27F66D5244FF62E1". This is the most significant 8 bytes of our session key. Combined it should be
	// "27F66D5244FF62E1AA6F6120EDEB4280".

	return fmt.Sprintf("%X", message_xor)
}

func TripleDesEncrypt(key string, plainhex string) []byte {
	data_in, _ := hex.DecodeString(plainhex)
	tdeskey, _ := hex.DecodeString(key)
	c, _ := des.NewTripleDESCipher(tdeskey)
	out := make([]byte, len(data_in))
	c.Encrypt(out, data_in)
	return out
}

func SingleDesEncrypt(key string, plainhex string) []byte {
	data_in, _ := hex.DecodeString(plainhex)
	deskey, _ := hex.DecodeString(key)
	c, _ := des.NewCipher(deskey)
	out := make([]byte, len(data_in))
	c.Encrypt(out, data_in)
	return out
}

func XorHexKeyParts(keypart1 string, keypart2 string, keypart3 string) []byte {
	bkeypart1, _ := hex.DecodeString(strings.Replace(keypart1, " ", "", -1))
	bkeypart2, _ := hex.DecodeString(strings.Replace(keypart2, " ", "", -1))
	bkeypart3, _ := hex.DecodeString(strings.Replace(keypart3, " ", "", -1))
	bkey := make([]byte, len(bkeypart1))
	for i := 0; i < len(bkeypart1); i++ {
		bkey[i] = bkeypart1[i] ^ bkeypart2[i] ^ bkeypart3[i]
	}
	return bkey
}

func XorHexKeyPair(key string, keymask string) []byte {
	bkeypart, _ := hex.DecodeString(strings.Replace(key, " ", "", -1))
	bkeymask, _ := hex.DecodeString(strings.Replace(keymask, " ", "", -1))
	bkey := make([]byte, len(bkeypart))
	for i := 0; i < len(bkeypart); i++ {
		bkey[i] = bkeypart[i] ^ bkeymask[i]
	}
	return bkey
}

func AndHexKeyPair(key string, keymask string) []byte {
	bkeypart, _ := hex.DecodeString(strings.Replace(key, " ", "", -1))
	bkeymask, _ := hex.DecodeString(strings.Replace(keymask, " ", "", -1))
	bkey := make([]byte, len(bkeypart))
	for i := 0; i < len(bkeypart); i++ {
		bkey[i] = bkeypart[i] & bkeymask[i]
	}
	return bkey
}
