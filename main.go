package main

import (
	"net/http"
	"net/url"
	"html"
)

var (
	inbin bool
	mainbody string
	err error
	cyan = RGB(70, 160, 255)
	nc = RGB(255, 255, 255)
	green = RGB(0, 255, 0)
	red = RGB(255, 0, 0)
)

// download gutil.go and gc.py (preprocessor)
// to easly replace these "include"s
include "gutil"
include "pager"

func GetCodeHandler(w http.ResponseWriter, sc uint16) {
	paste, err := load(sc)
	if err != nil {
		fprintf(w, "No such Paste %d", sc)
		printf("%sclient%s didn't get the paste %d\n", red, nc, sc)
		return
	}
	fprintf(w, "<html><body><p>%s</p></body></html>", paste)
	printf("%sclient%s got paste %d\n", green, nc, sc)
}

func HandlerMainPage(w http.ResponseWriter, r *http.Request) {
	fprintf(w, mainbody)
	printf("%sclient%s got the main page\n", green, nc)
}

func HandleShowCode(w http.ResponseWriter, code uint16) {
	fprintf(w, "<html><body><p>code: </p><h1>%d</h1></body></html>", code)
	printf("%sclient%s notified of new paste's code %d\n", green, nc, code)
}

func HandleMakePage(w http.ResponseWriter, v string) {
	vm, err := url.ParseQuery(v)
	panic(err)
	paste, ok := vm["paste"]
	if ok {
		code := save(html.EscapeString(paste[0]))
		printf("%sclient%s made a new paste %d\n", green, nc, code)
		HandleShowCode(w, code)
	} else {
		printf("%sclient%s didn't made the paste\n", red, nc)
		printf("vars: %v\n", vm)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	printf("\n%sclient%s requested %s with %s\n", cyan, nc, r.URL.Path, r.URL.Query())
	if len(r.URL.Path) == 0 {
		HandlerMainPage(w, r)
		return
	}

	if r.URL.Path == "/make" {
		HandleMakePage(w, r.URL.RawQuery)
		return
	}

	sc, err := strconv.Atoi(r.URL.Path[1:])
	if err == nil {
		GetCodeHandler(w, uint16(sc))
		return
	}
	//TODO: make error page
	HandlerMainPage(w, r)
}

func main(){
	InitGu()
	inbin = !exists("./saved/")
	if inbin {
		PS("running as inbin")
	} else {
		PS("not running as inbin")
	}
	mainbody = ReadFile("home.html")
	http.HandleFunc("/", handler)
	PS("server started")
	http.ListenAndServe(":6969", nil)
	exit(0)
}

