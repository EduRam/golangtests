package efabus

import (
	"fmt"
	"testing"
)

// Make the types prettyprint.
var itemName = map[itemType]string{
	itemError: "error",
	itemEOF:   "EOF",
}

func (i itemType) String() string {
	s := itemName[i]
	if s == "" {
		return fmt.Sprintf("item%d", int(i))
	}
	return s
}

type lexTest struct {
	name  string
	input string
	items []item
}

var (
	tEOF = item{itemEOF, 0, ""}
)

var lexTests = []lexTest{
	{"Even1",
		"[abc=[def=ghi]]",
		[]item{
			{itemEventName, 0, "abc"},
			{itemAttrKey, 0, "def"},
			{itemAttrValue, 0, "ghi"},
			tEOF,
		}},
	{"Event2",
		"[a=[d=g]]",
		[]item{
			{itemEventName, 0, "a"},
			{itemAttrKey, 0, "d"},
			{itemAttrValue, 0, "g"},
			tEOF,
		}},
	{"Event3",
		"[a=[b=c][d=e]]",
		[]item{
			{itemEventName, 0, "a"},
			{itemAttrKey, 0, "b"},
			{itemAttrValue, 0, "c"},
			{itemAttrKey, 0, "d"},
			{itemAttrValue, 0, "e"},
			tEOF,
		}},
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest, left, right string) (items []item) {
	l := lex(t.name, t.input, left, right)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func isEqual(i1, i2 []item, checkPos bool) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {

		if i1[k].typ != i2[k].typ {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
		if checkPos && i1[k].pos != i2[k].pos {
			return false
		}
	}
	return true
}

func TestLex(t *testing.T) {
	_ = "breakpoint"
	for _, test := range lexTests {
		items := collect(&test, "", "")
		if !isEqual(items, test.items, false) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, items, test.items)
		}
	}
}
