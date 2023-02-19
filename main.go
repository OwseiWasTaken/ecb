package main

import (
	"net/http"
)

var (
	inbin bool
	port string
	mainbody string
	mainecbbody string
	err error
	cyan = RGB(70, 160, 255)
	nc = RGB(255, 255, 255)
	green = RGB(0, 255, 0)
	red = RGB(255, 0, 0)
)

var (
	method string
	form map[string][]string
	vars map[string][]string
	url string
)

// download gutil.go and gc.py (preprocessor)
// to easly replace these "include"s
include "gutil"
Include "ecb"
Include "mine"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	url = r.URL.Path
	method = r.Method
	vars = r.URL.Query()
	form = map[string][]string{}
	if method == "POST" {
		r.ParseForm()
		form = r.Form
	}

	printf("\n%s[MAIN] client%s requested %s with (%s, %s) as %v\n",
		cyan, nc, url, vars, form, method)
	if len(r.URL.Path) == 1 {
		printf("→ %s[MAIN] client%s got main page (%s with %s)\n",
			green, nc, r.URL.Path, r.URL.Query())
		fprintf(w, mainbody)
		return
	}

	if len(r.URL.Path) > 3 {
		if r.URL.Path[:4] == "/ecb" {
			EcbHandler(w, r)
		} else if r.URL.Path[:5] == "/mine" {
			MineHandler(w, r)
		}
		return
	}

	printf("→ %s[MAIN] client%s can't satisfy request (%s with %s)\n",
		red, nc, r.URL.Path, r.URL.Query())
	fprintf(w, mainbody)
}

func main(){
	InitGu()

	inbin = !Exists("./main.go")
	if inbin {
		PS("running as inbin")
	} else {
		PS("not running as inbin")
	}
	mainecbbody = ReadFile("ecb/home.html")
	mainbody = ReadFile("home.html")
	http.HandleFunc("/", MainHandler);
	go http.ListenAndServe(":80", nil)
	PS("server started")
	for {
	Input("")
	}

	Exit(0)
}

