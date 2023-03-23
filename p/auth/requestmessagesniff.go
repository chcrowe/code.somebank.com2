package auth

import (
	"regexp"
	"strings"
)

func MatchesXml(s string) bool {
	return strings.Contains(s, "<RequestMessage>")
}

func MatchesNameValuePairs(s string) bool {
	matched, _ := regexp.MatchString("[a-zA-Z_]+=[a-zA-Z_]+&", s)
	return matched
}

func MatchesTsys1080(s string) bool {
	return strings.HasPrefix(s, "D4.")
}

func UrlMatchesPOSCP30(urlpath string) bool {
	matched, _ := regexp.MatchString("/POSCP3", urlpath)
	return matched
}

func UrlMatchesPOSCP308(urlpath string) bool {
	matched, _ := regexp.MatchString("/POSCP308", urlpath)
	return matched
}

func UrlMatchesPOSCP4(urlpath string) bool {
	matched, _ := regexp.MatchString("/POSCP4", urlpath)
	return matched
}
