include "pager"

import (
	"html"
)

func GetCodeHandler(w http.ResponseWriter, sc uint16) {
	paste, err := load(sc)
	if err != nil {
		fprintf(w, "No such Paste %d", sc)
		printf("→ %s[ECB] client%s didn't get the paste %d\n", red, nc, sc)
		return
	}
	paste = strings.Replace(paste, "\n", "<br>", -1)
	paste = strings.Replace(paste, " ", "&nbsp;", -1)
	paste = strings.Replace(paste, "\t", "&nbsp;&nbsp;", -1)
	fprintf(w, "<html><body><p>%s</p></body></html>", paste)
	printf("→ %s[ECB] client%s got paste %d\n", green, nc, sc)
}

func HandlerEcbMainPage(w http.ResponseWriter, r *http.Request) {
	fprintf(w, mainecbbody)
	printf("→ %s[ECB] client%s got the main page\n", green, nc)
}

func HandleShowCode(w http.ResponseWriter, code uint16) {
	fprintf(w, "<html><body><p>code: </p><h1>%d</h1></body></html>", code)
	printf("→ %s[ECB] client%s notified of new paste's code %d\n", green, nc, code)
}

func HandleMakePage(w http.ResponseWriter, paste []string) {
	if len(paste) != 0 {
		code := save(html.EscapeString(paste[0]))
		printf("→ %s[ECB] client%s made a new paste %d\n", green, nc, code)
		HandleShowCode(w, code)
	} else {
		printf("→ %s[ECB] client%s didn't made the paste\nno info recieved", red, nc)
	}
}

func Ecbhandler(w http.ResponseWriter, r *http.Request) {
	printf("→ %s[ECB] client%s requested %s with (%s, %s)\n",
		cyan, nc, url, vars, form)
	if len(r.URL.Path) == 4 {
		HandlerEcbMainPage(w, r)
		return
	}

	if r.URL.Path == "/ecb/make/" {
		HandleMakePage(w, form["cont"])
		return
	}

	if id, ok := vars["id"]; ok { // id, not from link, from input submit
		sc, err := strconv.Atoi(id[0])
		if err == nil {
			GetCodeHandler(w, uint16(sc))
		} else {
			printf("→ %s[ECB]%s can't transform id `%s` into int\n",
			red, nc, id[0])
		}
		return
	}

	if len(r.URL.Path) > 5 {
		sc, err := strconv.Atoi(r.URL.Path[5:])
		if err == nil {
			GetCodeHandler(w, uint16(sc))
		} else {
			printf("→ %s[ECB] can't transform id %s into int\n",
			red, nc, r.URL.Path[5:])
		}
		return
	}

	//TODO: make error page
	printf("→ %s[ECB] client%s can't satisfy request (%s with %s)\n",
		red, nc, r.URL.Path, r.URL.Query())
	HandlerEcbMainPage(w, r)
}

