package main

import (
	"net/http"
	"syscall"
	"time"
)

include "gutil"

include "pager"

var (
	codebody string
	mainbody string
	err error
)

func GetCodeHandler(w http.ResponseWriter, sc uint32) {
	pg, err := load(sc)
	if err != nil {
		fprintf(w, "No such Paste %d", sc)
		return
	}
	fprintf(w, codebody, string(pg.paste))
	//fprintf(w, codebody)
	printf("client got paste %d\n", sc)
}

func HandlerMainPage(w http.ResponseWriter, r *http.Request) {
	fprintf(w, mainbody)
	printf("client got the main page\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
	printf("client requested %s\n", r.URL.Path)
	if len(r.URL.Path) == 0 {
		HandlerMainPage(w, r)
		return
	}
	sc, err := strconv.Atoi(r.URL.Path[1:])
	if err == nil {
		GetCodeHandler(w, uint32(sc))
		return
	}
	HandlerMainPage(w, r)
}

func main(){
	InitGu()
	mainbody = ReadFile("home.html")
	codebody = ReadFile("code.html")
	http.HandleFunc("/", handler)
	PS("server started")
	http.ListenAndServe(":6969", nil)
	exit(0)
}

