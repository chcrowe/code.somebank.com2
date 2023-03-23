package main

import (
	"code.somebank.com/p/auth"
	//	"code.somebank.com/p/list"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"time"
)

func main() {

	// memStats := &runtime.MemStats{}
	// runtime.ReadMemStats(memStats)
	// fmt.Printf("Memory: %+v\n", memStats)

	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

	done := make(chan bool, 1) //this is the main service channel

	go signalCtrlC_handler(done) //register the CTRL+C handler

	go listenAndServe() //start the HTTP server

	<-done //wait for the CTRL+C interupt
}

func listenAndServe() {
	dir, f := filepath.Split(os.Args[0])
	fmt.Printf("%q started [:8080] in folder %q\n", f, dir)

	http.HandleFunc("/", url_handler)
	http.ListenAndServe(":8080", nil)

}

func signalCtrlC_handler(done chan bool) {

	sig := make(chan os.Signal, 1) //signal channel to capture CTRL+C
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Print(" shutting down ")
	for j := 0; j < 3; j++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	fmt.Println(" [done]")

	done <- true
}

func url_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "time: %v\n\n", time.Now())

	if true == auth.UrlMatchesPOSCP4(r.URL.Path) {
		fmt.Fprintf(w, "UrlMatchesPOSCP4: %v\n\n", r.URL.Path)
	} else if true == auth.UrlMatchesPOSCP308(r.URL.Path) {
		fmt.Fprintf(w, "UrlMatchesPOSCP308: %v\n\n", r.URL.Path)
	} else if true == auth.UrlMatchesPOSCP30(r.URL.Path) {
		fmt.Fprintf(w, "UrlMatchesPOSCP30: %v\n\n", r.URL.Path)
	}

	body, _ := ioutil.ReadAll(r.Body)

	s := string(body)

	if true == auth.MatchesXml(s) {
		fmt.Fprintf(w, "MatchesXml body: %v\n\n", s)
	} else if true == auth.MatchesNameValuePairs(s) {
		//nvp := list.ParseNameValuePairMap(string(body))
		fmt.Fprintf(w, "MatchesNameValuePairs body: %v\n\n", s)
	} else if true == auth.MatchesTsys1080(s) {
		fmt.Fprintf(w, "Matches1080 body: %v\n\n", s)
	}
}
