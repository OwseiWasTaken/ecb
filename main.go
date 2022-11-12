package main

import (
	"net/http"
	"syscall"
	"time"
)

include "gutil"

include "pager"

var (
	body = `
<html>
	<body>
		<h1>
			%s
		</h1>
		<form>
			<button>
			</button>
		</form>
	</body>
</html>
`
)

func handler(w http.ResponseWriter, r *http.Request) {
	sc, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fprintf(w, "Wrong Code Formatting")
		return
	}
	pg, err := load(uint32(sc))
	if err != nil {
		fprintf(w, "No such Paste %d", sc)
		return
	}
	fprintf(w, body, string(pg.paste))
	printf("client requested paste %d\n", sc)
}

func main(){
	InitGu()
	http.HandleFunc("/", handler)
	PS("server started")
	http.ListenAndServe(":6969", nil)
	exit(0)
}

