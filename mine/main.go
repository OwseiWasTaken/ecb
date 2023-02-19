var mainminebody = ReadFile("mine/translate.html")

func MineHandler(w http.ResponseWriter, r *http.Request) {
	fprintf(w, mainminebody)
}
