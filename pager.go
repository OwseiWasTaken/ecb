type Page struct {
	code uint32
	paste []byte
}

func CodeToString (b uint32) (string) {
	return strconv.Itoa(int(b))
}

func CodeToFilename (b uint32) (string) {
	return "pastes/"+CodeToString(b)+".txt"
}

func (p *Page) save () {
	flname := CodeToFilename(p.code)
	err := os.WriteFile(flname, p.paste, 0644) // 1X 2W 4R
	panic(err)
}

func GetFileCtime(name string) (ctime int64, err error) {
    fi, err := os.Stat(name)
    if err != nil {
        return
    }
    stat := fi.Sys().(*syscall.Stat_t)
    ctime = int64(stat.Ctim.Sec)
    return ctime, nil
}

func load(code uint32) (*Page, error) {
	flname := CodeToFilename(code)
	paste, err := os.ReadFile(flname)
	if err != nil {
		return &Page{0, []byte("can't find code")}, err
	}
	return &Page{code, paste}, err
}

func GetUnusedCode () (uint32) {
	return uint32(rint(0, int(U32MAX)))
}

func MakePage (paste string) (*Page) {
	code:=GetUnusedCode()
	return &Page{code, []byte(paste)}
}

func (p *Page) repr () {
	printf("code: %d\npaste: %s\n", p.code, string(p.paste))
}

