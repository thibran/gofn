package run

// StringSet set of strings.
type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{
		m: make(map[string]struct{}),
	}
}

func (s *StringSet) Append(arr ...string) {
	for _, v := range arr {
		s.m[v] = struct{}{}
	}
}

func (s *StringSet) Items() []string {
	arr := make([]string, len(s.m))
	var i = 0
	for k := range s.m {
		arr[i] = k
		i++
	}
	return arr
}

func (s *StringSet) Len() int {
	return len(s.m)
}
