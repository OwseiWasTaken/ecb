package main

import (
	"net/http"
	"net/url"
	"html"
)

var (
	inbin bool
	mainbody string
	mainecbbody string
	err error
	cyan = RGB(70, 160, 255)
	nc = RGB(255, 255, 255)
	green = RGB(0, 255, 0)
	red = RGB(255, 0, 0)
)

// download gutil.go and gc.py (preprocessor)
// to easly replace these "include"s
include "gutil"
Include "ecb"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	printf("\n%s[MAIN] client%s requested %s with %s\n", cyan, nc, r.URL.Path, r.URL.Query())
	if len(r.URL.Path) > 3 {
		if r.URL.Path[:4] == "/ecb" {
			ecbhandler(w, r)
			return
		}
	}

	printf("\n%s[Main] client%s can't satisfy requeste (%s with %s)\n",
		red, nc, r.URL.Path, r.URL.Query())
	fprintf(w, mainbody)
}

func main(){
	InitGu()
	inbin = !exists("./main.go")
	if inbin {
		PS("running as inbin")
	} else {
		PS("not running as inbin")
	}
	mainecbbody = ReadFile("ecb/home.html")
	mainbody = ReadFile("home.html")
	http.HandleFunc("/", MainHandler)
	PS("server started")
	http.ListenAndServe(":80", nil)
	exit(0)
}

