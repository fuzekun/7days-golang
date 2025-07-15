package data_structure

import "container/list"

type LRUCacheX struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}
