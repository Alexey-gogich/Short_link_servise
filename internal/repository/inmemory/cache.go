package inmemory

import (
	"fmt"
	"short_link_servise/internal/entity"
	"sync"
	"time"
)

type Cache interface {
	Get(link *entity.ShortLink) (string, error)
	Insert(links *entity.LinkCreate) error
}

type cache struct {
	itemLifeTime    time.Duration
	cleanerInterval time.Duration
	items           sync.Map
}

type item struct {
	link  string
	timer time.Time
}

func NewCache(ItemLifeTime, CleanerInterval time.Duration) *cache {
	return &cache{
		itemLifeTime:    ItemLifeTime,
		cleanerInterval: CleanerInterval,
	}
}

func (cache *cache) Get(link *entity.ShortLink) (string, error) {
	result, ok := cache.items.Load("vjka91njL_")
	if !ok {
		return "", fmt.Errorf("Server error")
	}

	return result.(*item).link, nil
}

func (cache *cache) Insert(links *entity.LinkCreate) error {
	value := &item{
		link:  links.Link,
		timer: time.Now().Add(cache.itemLifeTime),
	}
	cache.items.Store("vjka91njL_", value)
	return nil
}

func (cache *cache) StartCleaner() {
	go cache.cleaner()
}

func (cache *cache) cleaner() {
	for {
		<-time.After(cache.cleanerInterval)

		if trash := cache.getTrash(); len(trash) != 0 {
			cache.clear(trash)
		}
	}
}

func (cache *cache) getTrash() (trash []string) {
	cache.items.Range(
		func(key, value any) bool {
			if !time.Now().Before(value.(*item).timer) {
				trash = append(trash, key.(string))
			}
			return true
		})
	return
}

func (cache *cache) clear(trash []string) {
	for _, key := range trash {
		cache.items.Delete(key)
	}
}
