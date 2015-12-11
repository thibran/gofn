package run

import "testing"

func TestCheckName(t *testing.T) {
	if checkName("aaa") {
		t.Fail()
	}
	if checkName("aAa") {
		t.Fail()
	}
	if checkName("0aa") {
		t.Fail()
	}
	if !checkName("aa ") {
		t.Fail()
	}
	if !checkName("a-a") {
		t.Fail()
	}
	if !checkName("aa.") {
		t.Fail()
	}
}

func BenchmarkPrintFunctions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		printFunctions()
	}
}
