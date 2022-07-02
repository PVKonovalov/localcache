// Created by Pavel Konovalov pkonovalov@orxagrid.com
//
// localcache is a local cache library
//
package localcache

import (
	"errors"
	"os"
	"testing"
)

func TestLoadFileNotExist(t *testing.T) {
	cache := New("./cache.json")
	_, err := cache.Load()
	if !errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}
}
