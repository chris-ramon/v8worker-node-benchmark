package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ry/v8worker"
)

var (
	JSCode string
)

func init() {
	JSCode = `
        $send("hello from v8worker");
    `
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	V8Worker := v8worker.New(func(msg string) {
		log.Printf("v8worker response. msg: %v", msg)
		w.Write([]byte(msg))
		return
	})
	if err := V8Worker.Load("code.js", JSCode); err != nil {
		log.Printf("failed to load js file. error: %v", err)
		w.Write([]byte("something went wrong!"))
		return
	}
}

func noJS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok nojs!"))
}

func hitNode(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://localhost:8081", nil)
	res, _ := http.DefaultClient.Do(req)
	b, _ := ioutil.ReadAll(res.Body)
	w.Write([]byte(string(b)))
}

func main() {
	log.Print("Server running on :8080")
	http.HandleFunc("/v8worker", rootHandler)
	http.HandleFunc("/nojs", noJS)
	http.HandleFunc("/hitnode", hitNode)
	http.ListenAndServe(":8080", nil)
}
