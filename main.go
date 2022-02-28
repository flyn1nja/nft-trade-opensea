package main

import (
	"flag"
	"fmt"
	"strconv"
	"syscall/js"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5050"
	CONN_TYPE = "tcp"
	CONN_ADDR = "http://localhost:5050"

	GET_COL_URL        = "https://api.opensea.io/api/v1/collection/"
	GET_COL_ASSETS_URL = "https://api.opensea.io/api/v1/assets?order_by=sale_price&order_direction=asc&offset=%d&limit=%d&collection=%s"

	PRICE_MULT = 1000000000000000000

	API_KEY = "b3e478286c074bbc9d3fe40bc2e3cf50"
)

// var rtr = mux.NewRouter()
// var code = ""
// var browser *exec.Cmd
// var client *http.Client

var collection string // Nom de la collection à demander
var max int           // Nombre d'assets à considérer
var start int         // Commencer à partir du __ième asset de la collection

func init() {
	flag.StringVar(&collection, "col", "", "Nom de la collection à demander")
	flag.IntVar(&max, "max", 10, "Nombre d'assets à considérer")
	flag.IntVar(&start, "start", 0, "Commencer à partir du __ième asset de la collection")
}

func main() {

	flag.Parse()

	getColCB := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		col := js.Global().Get("document").Call("getElementById", "ColName").Get("value").String()
		count, _ := strconv.Atoi(js.Global().Get("document").Call("getElementById", "ACVal").Get("value").String())
		begin, _ := strconv.Atoi(js.Global().Get("document").Call("getElementById", "SOVal").Get("value").String())

		fmt.Println("Fetch de la collection " + col)
		// switch len(args) {
		// case 2:
		// 	go GetCollection(0, args[0].Int(), args[1].String())
		// case 3:
		// 	go GetCollection(args[0].Int(), args[1].Int(), args[2].String())
		// default:
		// 	go GetCollection(start, max, collection)
		// }

		go GetCollection(begin, count, col)

		// cb.Release() // release the function if the button will not be clicked again
		return nil
	})

	js.Global().Get("document").Call("getElementById", "GetColBtn").Call("addEventListener", "click", getColCB)
	// c := make(chan struct{}, 0)
	// js.Global().Set("GetCol", getColCB()))
	// <-c

	go GetCollection(start, max, collection)

	// fmt.Println("hello, webassembly!")
	// document := js.Global().Get("document")
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", "Hello WASM from Go!")
	// document.Get("body").Call("appendChild", p)
	// js.Global().Set("add", js.FuncOf(add))

	Host()
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// func add(this js.Value, arg []js.Value) interface{} {
// 	value1 := js.Global().Get("document").Call("getElementById", arg[0].String()).Get("value").String()
// 	value2 := js.Global().Get("document").Call("getElementById", arg[1].String()).Get("value").String()
// 	numTwo, _ := strconv.Atoi(value2)
// 	numOne, _ := strconv.Atoi(value1)
// 	document := js.Global().Get("document")
// 	div := document.Call("createElement", "div")
// 	div.Set("innerHTML", numOne+numTwo)
// 	document.Get("body").Call("appendChild", div)
// 	println(numOne + numTwo)
// 	return nil
// }
