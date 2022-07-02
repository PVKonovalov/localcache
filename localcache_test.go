// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
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
