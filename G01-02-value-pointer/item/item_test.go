package item_test

import (
	"testing"

	"github.com/sch39/go-series/g0102/valueptr/item"
)

func TestAddQtyValue(t *testing.T) {
	it := item.Item{Name: "Book", Qty: 5}
	item.AddQtyValue(it)
	if it.Qty != 5 {
		t.Errorf("AddQtyValue changed value; got %d, want 5", it.Qty)
	}
}

func TestAddQtyPointer(t *testing.T) {
	it := item.Item{Name: "Book", Qty: 5}
	item.AddQtyPointer(&it)
	if it.Qty != 6 {
		t.Errorf("AddQtyPointer did not change; got %d, want 6", it.Qty)
	}
}

func TestRangeVsIndex(t *testing.T) {
	items := []item.Item{{Qty: 5}, {Qty: 10}}

	item.AddQtyRange(items)
	if items[0].Qty != 5 || items[1].Qty != 10 {
		t.Errorf("AddQtyRange should not change slice; got %+v", items)
	}

	item.AddQtyIndex(items)
	if items[0].Qty != 6 || items[1].Qty != 11 {
		t.Errorf("AddQtyIndex should change slice; got %+v", items)
	}
}
