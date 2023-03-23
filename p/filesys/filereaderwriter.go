package filesys

import (
	"io/ioutil"
	"os"
	"strings"
)

type FileReaderWriter struct {
	FilePath string
}

func NewFileReaderWriter(s string) *FileReaderWriter {
	return &FileReaderWriter{s}
}

func (a *FileReaderWriter) Write(p []byte) (n int, err error) {
	return a.WriteString(string(p))
}

func (a *FileReaderWriter) Create() (err error) {

	f, e := os.OpenFile(a.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	return e
}

func (a *FileReaderWriter) WriteString(s string) (n int, err error) {

	f, e := os.OpenFile(a.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	n, e = f.WriteString(s)
	if e != nil {
		panic(e)
	}
	f.Sync()
	return n, e
}

func (a *FileReaderWriter) ReadLines() (s []string, err error) {

	dat, e := ioutil.ReadFile(a.FilePath)

	return strings.Split(string(dat), "\n"), e
}

// func (a *FileReaderWriter) Read(p []byte) (n int, err error) {
// 	dat, e := ioutil.ReadFile(a.FilePath)
// 	plen := len(p)
// 	datlen := len(dat)
// 	if plen < datlen {
// 		return copy(p, dat[0:plen]), e
// 	}
// 	return copy(p, dat), e
// }

// var TESTAUTHFILEPATH string = "c:\\Projects\\Go\\src\\code.somebank.com\\p\\tsys\\tsysauths2.dat"
// var TESTSALT string = "D64C3FFECFC340DA8CDD4123AF16E3A4"

// func TestWriteToFile(t *testing.T) {
// 	arw := auth.NewFileReaderWriter(TESTAUTHFILEPATH)
// 	arw.Create()

// 	sendVirtualNetSslHttpRequest(CreateVisaRetailAuthorization(t))
// 	// MastercardRetailAuthorization(t)
// }

// func TestReadFromFile(t *testing.T) {

// 	arw := auth.NewFileReaderWriter(TESTAUTHFILEPATH)
// 	lines, _ := arw.ReadLines()

// 	a := auth.NewAuthorizationRecord(TESTSALT)
// 	for _, line := range lines {
// 		m := a.Decode(line, true)
// 		if 0 < len(m.Get("utc")) {
// 			fmt.Printf("%s\n\n", m.String())
// 		}
// 	}

// }
