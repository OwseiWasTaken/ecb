func CodeToString (b uint32) (string) {
	return strconv.Itoa(int(b))
}

func CodeToFilename (b uint32) (string) {
	return "saved/"+CodeToString(b)+".txt"
}

//TODO: assert unused
func GetUnusedCode () (uint32) {
	return uint32(rint(0, int(U32MAX)))
}

func save (paste string) (uint32) {
	code := GetUnusedCode()
	flname := CodeToFilename(code)
	err := os.WriteFile(flname, []byte(paste), 0644) // 1X 2W 4R
	panic(err)
	return code
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

func load(code uint32) (string, error) {
	flname := CodeToFilename(code)
	p, err := os.ReadFile(flname)
	return string(p), err
}

