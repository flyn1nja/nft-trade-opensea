package main

import (
	"fmt"
	"net/http"

	// "syscall/js"
	"time"
	// wbrowser "github.com/toqueteos/webbrowser"
)

// var browser *wbrowser.Browser

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// func sendETHFromMetaMask(this js.Value, args []js.Value) interface{} {
// 	println("Hello")
// 	return js.TypeNull
// }

func send(w http.ResponseWriter, req *http.Request) {
	// cmd, err := wbrowser.Candidates[0].Command("window.ethereum")
	// if err != nil {
	// 	w.Write([]byte(fmt.Sprint("Command failed to run:\n%s", err.Error())))
	// 	return
	// }
	// cmd.Run()
}

func Host() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/send", send)
	http.HandleFunc("/headers", headers)

	// go wbrowser.Open("http://localhost:8090/send")
	time.Sleep(100 * time.Millisecond)
	http.ListenAndServe(":8090", nil)
}
