package clean

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/thibran/gofn/info"
)

// Cleaner removes the oldest gofn functions.
type Cleaner struct {
	bindir  string
	itemMax int
	sep     string
}

// NewCleaner creates a new Cleaner object.
func NewCleaner(bindir string) *Cleaner {
	max := os.Getenv("GOFN_MAX")
	itemMax, err := strconv.Atoi(max)
	if err != nil || len(max) == 0 {
		itemMax = 200 // default value
	}
	return &Cleaner{
		bindir:  bindir,
		itemMax: itemMax,
		sep:     string(os.PathSeparator),
	}
}

// Clean removes the oldest gofn functions if there are more than itemMax.
// itemMax is specified by the environment variable GOFN_MAX or defaults to 200.
func (c *Cleaner) Clean() {
	var arr cleanItems
	for _, name := range info.ListFunctions(c.bindir) {
		info, err := info.ByName(name, c.bindir, c.sep)
		if err != nil {
			continue
		}
		//arr = append(arr, info.Time.Unix())
		arr = append(arr, cleanItem{
			name: name,
			time: info.Time.Unix(),
		})
	}
	sort.Sort(arr)
	for _, item := range arr.notNeeded(c.itemMax) {
		deleteGofn(item.nameToPath(c.bindir, c.sep))
	}
}

type cleanItem struct {
	name string
	time int64
}

type cleanItems []cleanItem

// deleteGofn with the path p.
func deleteGofn(p string) {
	if err := os.Remove(p); err != nil {
		fmt.Println(err)
	}
}

func (c *cleanItem) nameToPath(bindir, sep string) string {
	return info.NameToPath(c.name, bindir, sep)
}

func (c cleanItems) notNeeded(itemMax int) (notNeeded cleanItems) {
	if itemMax < 0 || len(c) <= itemMax {
		return cleanItems{}
	}
	return c[itemMax:]
}

func (c cleanItems) shrinkSlice(maxItems int) {
	if len(c) <= maxItems {
		return
	}
	c = c[0:maxItems]
}

func (c cleanItems) Len() int {
	return len(c)
}

func (c cleanItems) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c cleanItems) Less(i, j int) bool {
	return c[i].time > c[j].time
}
