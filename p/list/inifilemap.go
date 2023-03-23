package list

import (
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	"code.somebank.com/p/bytes"
	estrings "code.somebank.com/p/strings"
)

type IniFileMap map[string]string

func ParseIniFileMap(s string) *IniFileMap {

	m := IniFileMap{}

	regex := regexp.MustCompile(`(?P<key>^[^#].+) ?= ?(?P<value>.+)`)
	section := regexp.MustCompile(`^\[[^].]+\]`)
	prefix := ""

	lines := strings.Split(s, "\n")
	for _, line := range lines {

		text := strings.TrimSpace(line)
		if text == "" {
			prefix = ""
			continue
		}

		if strings.HasPrefix(text, "#") {
			continue
		}

		if strings.HasPrefix(text, "[") {
			s := section.FindString(text)
			prefix = s[1:len(s)-1] + "."
		}

		matchs := regex.FindAllStringSubmatch(text, -1)
		if len(matchs) > 0 {
			for _, val := range matchs {
				key := prefix + val[1]
				m[key] = val[2]
			}
		}
	}

	return &m
}

func LoadIniFileMap(fileName string) *IniFileMap {
	b, _ := ioutil.ReadFile(fileName)
	return ParseIniFileMap(string(b))
}

func (m *IniFileMap) Save(fname string) (n int, err error) {

	f, e := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	n, e = f.WriteString(m.Serialize())
	if e != nil {
		panic(e)
	}
	f.Sync()
	return n, e
}

func (m IniFileMap) Get(key string) string {
	if m == nil {
		return ""
	}
	vs, ok := m[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs
}

func (m IniFileMap) GetUint(key string) uint {
	v := m.Get(key)
	return estrings.StringToUint(v)
}

func (m IniFileMap) GetUint64(key string) uint64 {
	v := m.Get(key)
	return estrings.StringToUint64(v)
}

func (m IniFileMap) Set(key, value string) {
	m[key] = value
}

func (m IniFileMap) Delete(key string) {
	delete(m, key)
}

func (m IniFileMap) Serialize() string {

	b := bytes.NewSafeBuffer(1024)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	var last string
	sort.Strings(keys)
	for _, k := range keys {
		keyparts := strings.Split(k, ".")

		if false == strings.EqualFold(last, keyparts[0]) {
			b.AppendFormat("\n\n[%s]\n", keyparts[0])
		}

		b.AppendFormat("%s=%s\n", keyparts[1], m[k])
		last = keyparts[0]
	}

	return b.String()
}

func (m IniFileMap) String() string {

	b := bytes.NewSafeBuffer(1024)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		b.AppendFormat("%s=%s\n", k, m[k])
	}

	return b.String()
}

func (m IniFileMap) Clear() {

	for k := range m {
		delete(m, k)
	}
}
