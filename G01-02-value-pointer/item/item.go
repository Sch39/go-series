package item

type Item struct {
	Name string
	Qty  int
}

func AddQtyValue(it Item) {
	it.Qty++
}

func AddQtyPointer(it *Item) {
	it.Qty++
}

func AddQtyRange(items []Item) {
	for _, item := range items {
		item.Qty++ // copy elemen dari slice
	}
}

func AddQtyIndex(items []Item) {
	for i := range items {
		items[i].Qty++ // elemen asli dalam slice
	}
}

func AddQtyMap(items map[string]Item, key string) {
	item := items[key] // copy value dari map
	item.Qty++         // ubah copy
	items[key] = item  // simpan kembali ke map
}
