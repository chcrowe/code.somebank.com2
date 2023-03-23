package strings

import (
	"regexp"
	"strconv"
	"strings"
)

func removeNonDigits(s string) string {
	if 0 < len(s) {
		re, err := regexp.Compile("[^0-9]+")
		if err == nil {
			return re.ReplaceAllString(s, "")
		}
	}
	return s
}

func StringToUint64(s string) uint64 {
	n, e := strconv.ParseUint(removeNonDigits(s), 10, 64)
	if nil != e {
		return 0
	}
	return n
}

func StringToUint(s string) uint {
	return uint(StringToUint64(s))
}

func StringToInt(s string) int {
	return int(StringToUint64(s))
}

func StringDecimalAmountToUint64(s string) uint64 {

	i := strings.LastIndex(s, ".")
	if -1 < i {
		return StringToUint64(s)
	} else {
		amount := StringToUint64(s)
		return amount * 100
	}
}

func LeftPadToLength(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func RightPadToLength(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
