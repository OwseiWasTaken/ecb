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
	if inbin {
		return "../saved/"+CodeToString(b)+".txt"
	}
	return "saved/"+CodeToString(b)+".txt"
}

func save (paste string) (uint16) {
	code := MakeHash(paste)
	flname := CodeToFilename(code)
	WriteFile(flname, paste) // 1X 2W 4R
	return code
}

func load(code uint16) (string, error) {
	flname := CodeToFilename(code)
	p, err := os.ReadFile(flname)
	return string(p), err
}

