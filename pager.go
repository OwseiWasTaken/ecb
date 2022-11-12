import (
	"hash/fnv"
)

func MakeHash(s string) (uint16) {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint16(h.Sum32())
}

func CodeToString (b uint16) (string) {
	return strconv.Itoa(int(b))
}

func CodeToFilename (b uint16) (string) {
	return "saved/"+CodeToString(b)+".txt"
}

func save (paste string) (uint16) {
	code := MakeHash(paste)
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

func load(code uint16) (string, error) {
	flname := CodeToFilename(code)
	p, err := os.ReadFile(flname)
	return string(p), err
}

