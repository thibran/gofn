package run

// func newFunction(fn, imports string, args []string, info info.Info) *Function {
// 	return &Function{
// 		Imports: imports,
// 		Fn:      fn,
// 		Info:    info.String(),
// 		Args:    args,
// 		Name:    "myfn",
// 		Debug:   false,
// 		InfoObj: info,
// 	}
// }
//
// func TestExec(t *testing.T) {
// 	r := NewRun("/home/tux/.local/share/gofn")
// 	cmd := `fmt.Println("Hi", arr)`
// 	imp := "fmt"
// 	args := []string{"aaa", "bbb"}
// 	info := info.NewInfo(cmd)
// 	fn := newFunction(cmd, imp, args, info)
// 	r.Exec(fn)
// }
//
// func TestGofn(t *testing.T) {
// 	r := NewRun("/home/tux/.local/share/gofn")
// 	args := []string{"aaa", "bbb"}
// 	r.gofn("myfn", args)
// }
//
// func BenchmarkGofn(b *testing.B) {
// 	r := NewRun("/home/tux/.local/share/gofn")
// 	args := []string{"aaa", "bbb"}
// 	for n := 0; n < b.N; n++ {
// 		r.gofn("myfn", args)
// 	}
// }
//
// func TestCheckHash(t *testing.T) {
// 	r := NewRun("/home/tux/.local/share/gofn")
// 	if ok, err := r.checkHash("myfn", 3406308386); !ok && err == nil {
// 		t.Error("should be ok")
// 	}
// 	if ok, err := r.checkHash("myfn", 74743374); ok && err == nil {
// 		t.Error("wrong hash")
// 	}
// 	if ok, err := r.checkHash("aaa", 3406308386); ok && err == nil {
// 		t.Error("wrong name")
// 	}
// }
//
// func TestFunctionExists(t *testing.T) {
// 	r := NewRun("/home/tux/.local/share/gofn")
// 	if ok := r.functionExists("myfn"); !ok {
// 		t.Fail()
// 	}
// 	if ok := r.functionExists("aaa"); ok {
// 		t.Fail()
// 	}
// }
