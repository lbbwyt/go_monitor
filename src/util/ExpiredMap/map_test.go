package ExpiredMap

import (
	"fmt"
	"testing"
	"time"
)

func TestExpiredMap_Get(t *testing.T) {
	cache := NewExpiredMap()
	cache.Set("test", "123", 5)
	time.Sleep(time.Second * 6)
	fmt.Println(cache.Get("test"))
	fmt.Println(cache.TTL("test"))
}
