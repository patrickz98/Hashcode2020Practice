package index

import "sort"

type IndexStruct struct {
	items []int
	maps  map[int]bool
}

func (index *IndexStruct) add(item int) {
	index.items = append(index.items, item)
	index.maps[item] = true
	sort.Ints(index.items)
}
