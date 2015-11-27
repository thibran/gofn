package clean

// func TestNotNeeded(t *testing.T) {
// 	arr := cleanItems{
// 		cleanItem{"4", 1448545758},
// 		cleanItem{"1", 1448303934},
// 		cleanItem{"3", 1448545374},
// 		cleanItem{"2", 1448541394},
// 	}
// 	sort.Sort(arr)
// 	maxItems := 2
// 	arr = arr.notNeeded(maxItems)
// 	if len(arr) != 2 {
// 		t.Fail()
// 	}
// 	if arr[0].name != "2" {
// 		t.Fail()
// 	}
// 	if arr[1].name != "1" {
// 		t.Fail()
// 	}
// }
//
// func TestNotNeededNothingToDo(t *testing.T) {
// 	arr := cleanItems{
// 		cleanItem{"4", 1448545758},
// 		cleanItem{"1", 1448303934},
// 	}
// 	sort.Sort(arr)
// 	maxItems := 2
// 	arr = arr.notNeeded(maxItems)
// 	if len(arr) != 0 {
// 		t.Fail()
// 	}
// }
//
// func TestDeleteGofn(t *testing.T) {
// 	p := "/home/tux/.local/share/gofn/gofn-myfn4"
// 	deleteGofn(p)
// }
//
// func TestClean(t *testing.T) {
// 	c := NewCleaner("/home/tux/.local/share/gofn")
// 	c.Clean()
// }
