package list

import (
	"strings"

	"code.somebank.com/p/bytes"
	estrings "code.somebank.com/p/strings"
)

type NameValuePairMap map[string]string

func ParseNameValuePairMap(s string) NameValuePairMap {

	m := NameValuePairMap{}

	keyvalpairs := strings.Split(s, "&")
	for _, keyvaluepair := range keyvalpairs {
		j := strings.Index(keyvaluepair, "=")
		if -1 < j {
			m[keyvaluepair[:j]] = keyvaluepair[j+1:]
		}
	}

	return m
}

func (m NameValuePairMap) GetAny(keys []string) string {
	for _, k := range keys {
		v := m.Get(k)
		if 0 < len(v) {
			return v
		}
	}
	return ""
}

func (m NameValuePairMap) Get(key string) string {
	if m == nil {
		return ""
	}
	vs, ok := m[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs
}

func (m NameValuePairMap) GetUint(key string) uint {
	v := m.Get(key)
	return estrings.StringToUint(v)
}

func (m NameValuePairMap) GetUint64(key string) uint64 {
	v := m.Get(key)
	return estrings.StringToUint64(v)
}

func (m NameValuePairMap) Set(key, value string) {
	m[key] = value
}

func (m NameValuePairMap) Delete(key string) {
	delete(m, key)
}

func (m NameValuePairMap) String() string {

	b := bytes.NewSafeBuffer(1024)

	for key, value := range m {
		b.AppendFormat("%s=%s&", key, value)
	}

	return b.String()
}

func (m NameValuePairMap) Clear() {

	for k := range m {
		delete(m, k)
	}
}
