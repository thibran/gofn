package info

import (
	"testing"
	"time"
)

// func TestDate(t *testing.T) {
// 	// 2015-11-26 13:49:05.180151594 +0100 CET
// 	var unix int64 = 1448542145
// 	ti := time.Unix(unix, 0)
// 	f := ti.Format("2006-01-02 15:04")
// 	fmt.Println(f)
// }

func TestToInfo(t *testing.T) {
	s := "3406308386 1448303934 go1.5.1"
	info, err := ToInfo(s)
	if err != nil {
		t.Error(err)
	}
	var fnvHash uint32 = 3406308386
	var time = time.Unix(1448303934, 0)
	var goversion = "go1.5.1"
	if info.FnvHash != fnvHash {
		t.Fail()
	}
	if info.Time != time {
		t.Fail()
	}
	if info.Goversion != goversion {
		t.Fail()
	}
}

func TestToInfo_invalidString(t *testing.T) {
	s := "3406308386 go1.5.1"
	if _, err := ToInfo(s); err == nil {
		t.Error("not enough arguments")
	}
	s = "3406308386 aaa go1.5.1"
	if _, err := ToInfo(s); err == nil {
		t.Error("2. item is not int64 value")
	}
	s = "3406308386 aaa  "
	if _, err := ToInfo(s); err == nil {
		t.Error("4 []string segments")
	}
}

func TestNewInfo(t *testing.T) {
	info := NewInfo(`fmt.Println("Hi", arr)`)
	if info.FnvHash == 0 || info.Time.Unix() == 0 || info.Goversion == "" {
		t.Fail()
	}
}

// func TestListFunctions(t *testing.T) {
// 	r := run.NewRun("/home/tux/.local/share/gofn")
// 	if arr := listFunctions(r.bindir); len(arr) > 0 {
// 		fmt.Println(arr)
// 	}
// }
