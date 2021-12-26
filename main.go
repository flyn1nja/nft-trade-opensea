package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"

	"rogchap.com/v8go"

	mux "github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "5050"
	CONN_TYPE = "tcp"
	CONN_ADDR = "http://localhost:5050"
)

var rtr = mux.NewRouter()
var code = ""
var browser *exec.Cmd
var client *http.Client

func main() {

	flag.Parse()

	// go func() {
	// 	StartClient()
	// }()
	// source := "const multiply = (a, b) => a * b"
	// // source := "const multiply = getValx2(1)"
	// iso1 := v8go.NewIsolate()                                                         // creates a new JavaScript VM
	// ctx1 := v8go.NewContext(iso1)                                                     // new context within the VM
	// ctx1.RunScript("const add = (a, b) => a + b", "math.js")
	// script1, _ := iso1.CompileUnboundScript(source, "test.js", v8go.CompileOptions{}) // compile script to get cached data
	// val, _ := script1.Run(ctx1)
	// fmt.Printf("addition result: %s \n", val)

	// cachedData := script1.CreateCodeCache()

	// iso2 := v8go.NewIsolate()     // create a new JavaScript VM
	// ctx2 := v8go.NewContext(iso2) // new context within the VM

	// script2, _ := iso2.CompileUnboundScript(source, "test.js", v8go.CompileOptions{CachedData: cachedData}) // compile script in new isolate with cached data
	// scval, _ := script2.Run(ctx2)
	// fmt.Printf("Other value time 2: %s \n", scval)

	// jsscript := readScript("./test.js")

	// println(jsscript)

	ctx := v8go.NewContext(nil)                             // creates a new V8 context with a new Isolate aka VM
	ctx.RunScript("const add = (a, b) => a + b", "math.js") // executes a script on the global context
	ctx.RunScript("const result = add(3, 4)", "main.js")    // any functions previously added to the context can be called
	val, _ := ctx.RunScript("result", "value.js")           // return a value in JavaScript back to Go
	fmt.Printf("Addition result: %s \n", val)

	ctx.RunScript("./test.js", "cool.js")
	ctx.RunScript("const result2 = testVal()", "main.js")
	val2, _ := ctx.RunScript("result2", "value.js")
	// ctx.RunScript("const result = add(3, 4)", "main.js")
	// fmt.Printf("addition result: %s \n", val)
	fmt.Printf("Test result: %s \n", val2)
}

// func readScript(fileToRead string) string {
// 	file, err := os.Open(fileToRead)

// 	var res strings.Builder

// 	buf := make([]byte, 8)

// 	for err == nil {
// 		res.Write(buf)
// 		_, err = file.Read(buf)
// 	}

// 	return res.String()
// }

// func StartClient() {

// 	// browser = openbrowser(fmt.Sprintf("https://osu.ppy.sh/oauth/authorize?client_id=%d&redirect_uri=%s&response_type=code", CLIEND_ID, CONN_ADDR))
// 	client = &http.Client{}
// }
