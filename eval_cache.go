package decexpr

import "sync"

type cacheItem struct {
	Items []*RPNItem
}

type EvalCache interface {
	Put(key string, items []*RPNItem)
	Get(key string) (items []*RPNItem, found bool)
}

var _ EvalCache = (*EvalMapCache)(nil)

type EvalMapCache struct {
	cache map[string]cacheItem
	mutex sync.RWMutex
}

func NewEvalMapCache() *EvalMapCache {
	return &EvalMapCache{
		cache: make(map[string]cacheItem),
	}
}

func (ec *EvalMapCache) Put(key string, items []*RPNItem) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	ec.cache[key] = cacheItem{
		Items: items,
	}
}

func (ec *EvalMapCache) Get(key string) (items []*RPNItem, found bool) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()

	item, found := ec.cache[key]
	if !found {
		return nil, false
	}

	return item.Items, true
}

var _ EvalCache = (*EvalMapCache)(nil)

type EvalNoopCache struct{}

func NewEvalNoopCache() *EvalNoopCache {
	return &EvalNoopCache{}
}

func (*EvalNoopCache) Put(key string, items []*RPNItem) {}

func (*EvalNoopCache) Get(key string) (items []*RPNItem, found bool) {
	return nil, false
}
