include "pager"

func GetCodeHandler(w http.ResponseWriter, sc uint16) {
	paste, err := load(sc)
	if err != nil {
		fprintf(w, "No such Paste %d", sc)
		printf("%s[ECB] client%s didn't get the paste %d\n", red, nc, sc)
		return
	}
	paste = strings.Replace(paste, "\n", "<br>", -1)
	paste = strings.Replace(paste, " ", "&nbsp;", -1)
	paste = strings.Replace(paste, "\t", "&nbsp;&nbsp;", -1)
	fprintf(w, "<html><body><p>%s</p></body></html>", paste)
	printf("%s[ECB] client%s got paste %d\n", green, nc, sc)
}

func HandlerEcbMainPage(w http.ResponseWriter, r *http.Request) {
	fprintf(w, mainecbbody)
	printf("%s[ECB] client%s got the main page\n", green, nc)
}

func HandleShowCode(w http.ResponseWriter, code uint16) {
	fprintf(w, "<html><body><p>code: </p><h1>%d</h1></body></html>", code)
	printf("%s[ECB] client%s notified of new paste's code %d\n", green, nc, code)
}

func HandleMakePage(w http.ResponseWriter, v string) {
	vm, err := url.ParseQuery(v)
	panic(err)
	paste, ok := vm["paste"]
	if ok {
		code := save(html.EscapeString(paste[0]))
		printf("%s[ECB] client%s made a new paste %d\n", green, nc, code)
		HandleShowCode(w, code)
	} else {
		printf("%s[ECB] client%s didn't made the paste\n", red, nc)
		printf("vars: %v\n", vm)
	}
}

func ecbhandler(w http.ResponseWriter, r *http.Request) {
	printf("\n%s[ECB] client%s requested %s with %s\n", cyan, nc, r.URL.Path, r.URL.Query())
	if len(r.URL.Path) == 0 {
		HandlerEcbMainPage(w, r)
		return
	}

	if r.URL.Path == "/ecb/make" {
		HandleMakePage(w, r.URL.RawQuery)
		return
	}

	if len(r.URL.Path) > 5 {
		sc, err := strconv.Atoi(r.URL.Path[5:])
		if err == nil {
			GetCodeHandler(w, uint16(sc))
			return
		}
	}
	//TODO: make error page
	printf("\n%s[ECB] client%s can't satisfy requeste (%s with %s)\n",
		red, nc, r.URL.Path, r.URL.Query())
	HandlerEcbMainPage(w, r)
}

