package info

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Info hosts the values 'gofn -info' returns.
type Info struct {
	FnvHash   uint32
	Time      time.Time
	Goversion string
}

// NewInfo creates a new Info object using the passed function code fnCode.
func NewInfo(fnCode string) Info {
	return Info{
		FnvHash:   fnvHash(fnCode),
		Time:      time.Now(),
		Goversion: runtime.Version(),
	}
}

func (info Info) String() string {
	return fmt.Sprintf("%v %v %v", info.FnvHash, info.Time.Unix(), info.Goversion)
}

// ByName returns the Info object for passed gofn-function name.
func ByName(name, bindir, sep string) (Info, error) {
	fn := NameToPath(name, bindir, sep)
	b, err := exec.Command(fn, "-info").Output()
	if err != nil {
		return Info{}, err
	}
	return ToInfo(string(b))
}

// NameToPath returns the path to the give gofn function name.
func NameToPath(name, bindir, sep string) string {
	return fmt.Sprintf("%s%sgofn-%s", bindir, sep, name)
}

// ToInfo converts the passed string s to an Info object.
func ToInfo(s string) (Info, error) {
	arr := strings.Split(s, " ")
	if len(arr) != 3 {
		return Info{}, fmt.Errorf("E: passed info string is invalid: %s", s)
	}
	hash, err := strconv.ParseUint(arr[0], 10, 32)
	if err != nil {
		return Info{}, err
	}
	t, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return Info{}, err
	}
	time := time.Unix(t, 0)
	return Info{
		FnvHash:   uint32(hash),
		Time:      time,
		Goversion: arr[2],
	}, nil
}

func fnvHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// ListFunctions returns a list of all gofn function in the bindir directory.
func ListFunctions(bindir string) []string {
	dirlist, err := ioutil.ReadDir(bindir)
	if err != nil {
		return []string{}
	}
	var arr []string
	for _, f := range dirlist {
		// exclude folders
		if f.IsDir() {
			continue
		}
		// exclude other files
		if !strings.HasPrefix(f.Name(), "gofn-") {
			continue
		}
		name := strings.Split(f.Name(), "-")[1]
		arr = append(arr, name)
	}
	return arr
}
