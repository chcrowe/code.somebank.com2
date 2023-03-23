package msr

import (
	"fmt"
	"testing"
)

func TestMSRMasking(t *testing.T) {
	fmt.Println()
	fmt.Println()

	track1tests := []string{"%B4123456789012349^MAVERICK INTERNATIONAL VSA^17121010000000000000000?",
		"%B3701 234567 89017^MAVERICK INTERNATIONAL AMX^1712980100000?",
		"%B4534410307155000^JAL TEST CARD             ^15032010000000000000  328      ?",
		"%B4012000010000^SOMEBANK TEST CARD VISA CHK    ^1512101000158000000000000?"}

	for _, t1 := range track1tests {
		msr1 := ParseMagneticStripeTrackI(t1)
		fmt.Printf("%-8s%-24s%-24s%s\n", "MSR1", msr1.AccountNumber, MaskMSRAccountNumber(msr1.AccountNumber, "9"), msr1.TrackData)
		fmt.Printf("%-8s%-48s%s\n", "", "", MaskMagneticStripeTrack(msr1.TrackData, msr1.AccountNumber))
	}
	fmt.Println()
	fmt.Println()

	track2tests := []string{";4123456789012349=171210100000000000?",
		";370123456789017=1712980100000?",
		"4012000010000=1512101000485",
		";4534410307155000=15032010000032800000?",
		";6035718888930012928=49120003152?"}

	for _, t2 := range track2tests {
		msr2 := ParseMagneticStripeTrackII(t2)
		fmt.Printf("%-8s%-24s%-24s%s\n", "MSR2", msr2.AccountNumber, MaskMSRAccountNumber(msr2.AccountNumber, "9"), msr2.TrackData)
		fmt.Printf("%-8s%-48s%s\n", "", "", MaskMagneticStripeTrack(msr2.TrackData, msr2.AccountNumber))
	}
}

func TestMSRParsing(t *testing.T) {

	msrTests := []string{"%B6035718888930012910^^49120002089?;6035718888930012910=49120002089?",
		"%B9999999800002773^PAYMENTECH^18121015432112345678?;9999999800002773=18121015432112345678?",
		"%B6011000995500000^PAYMENTECH TEST CARD^18121015432112345678?;6011000995500000=18121015432112345678?",
		"%B3566002020140006^PAYMENTECH^18121015432112345678?;3566002020140006=18121015432112345678?"}

	for _, msrTest := range msrTests {
		msr1, msr2 := ParseMagneticStripe(msrTest)
		fmt.Printf("%-8s%-24s%s\n", "MSR1", msr1.AccountNumber, msr1.TrackData)
		fmt.Printf("%-8s%-24s%s\n", "MSR2", msr2.AccountNumber, msr2.TrackData)
	}

}
