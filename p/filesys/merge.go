package filesys

import (
	//	"code.google.com/p/go-uuid/uuid"
	"io/ioutil"
	"sort"
	"strings"

	ebytes "code.somebank.com/p/bytes"
	//    "fmt"
)

func MergeFiles(destfile string, filenames []string) {

	mergebuf := ebytes.NewSafeBuffer(4096)

	for _, fname := range filenames {

		b, e := ioutil.ReadFile(fname)
		if nil != e {
			panic(e)
		}
		mergebuf.AppendBytes(b)
	}

	totlines := strings.Split(mergebuf.String(), "\n")
	sort.Strings(totlines)

	ioutil.WriteFile(destfile, copyUnique(totlines), 0666)
}

func removeDuplicates(src []string) []string {
	dest := make([]string, len(src))
	var last string
	var j int = 0
	for i := 0; i < len(src); i++ {
		if false == strings.EqualFold(src[i], last) {
			dest[j] = src[i]
			j++
		}
		last = src[i]
	}
	return dest[0:j]
}

func copyUnique(src []string) []byte {
	dest := ebytes.NewSafeBuffer(4096)
	var last string
	for i := 0; i < len(src); i++ {
		if false == strings.EqualFold(src[i], last) {
			dest.AppendString(src[i] + "\n")
		}
		last = src[i]
	}
	return dest.Bytes()
}

// package main

// import (
// 	"code.somebank.com/p/filesys"
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {

// 	mergefile := "all.txt"
// 	filesys.MergeFiles(mergefile, []string{"dump1.txt", "dump2.txt", "dump3.txt"})
// 	b, _ := ioutil.ReadFile(mergefile)
// 	fmt.Printf("%s\n", string(b))
// }
