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
	"strings"
	"testing"
)

const CacheFile = "./cache.json"
const StringToSave1 = "[{\"elements\":[{\"equipment\":\"Equipment1\",\"name\":\"switch_truck12\",\"point_id\":2813," +
	"\"sld_id\":10,\"type_id\":9},{\"equipment\":\"Equipment2\",\"name\":\"switch_truck12_ia\",\"point_id\":2829," +
	"\"sld_id\":10,\"type_id\":5}],\"file\":\"/assets/sld/flisr-east-grid-2-5.svg\",\"id\":10,\"is_default\":1," +
	"\"is_enable\":1,\"name\":\"SmartGridExample\",\"parent_id\":6,\"type_id\":1}]"
const StringToSave2 = "[{\"elements\":[{\"equipment\":\"Equipment1\",\"name\":\"switch_truck12\",\"point_id\":2813," +
	"\"sld_id\":10,\"type_id\":9},{\"equipment\":\"Equipment2\",\"name\":\"switch_truck12_ia\",\"point_id\":2829," +
	"\"sld_id\":10,\"type_id\":5}],\"file\":\"/assets/sld/flisr-east-grid-2-5.svg\",\"id\":10,\"is_default\":1," +
	"\"is_enable\":1,\"name\":\"SmartGridExample\",\"parent_id\":6,\"type_id\":2}]"

func TestMain(m *testing.M) {
	err := os.Remove(CacheFile)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		os.Exit(1)
	}
	retCode := m.Run()
	os.Exit(retCode)
}

func TestLoadFileNotExist(t *testing.T) {
	cache := New(CacheFile)
	_, err := cache.Load()
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Unexpected error: %v\n", err)
	}
}

func TestSave1(t *testing.T) {
	cache := New(CacheFile)
	if err := cache.Save([]byte(StringToSave1)); err != nil {
		t.Fatalf("Failed to save: %v\n", err)
	}
}

func TestLoad1(t *testing.T) {
	cache := New(CacheFile)
	loadedData, err := cache.Load()
	if err != nil {
		t.Fatalf("Failed to load: %v\n", err)
	}

	if strings.Compare(string(loadedData), StringToSave1) != 0 {
		t.Fatalf("Failed to compare\n")
	}
}

func TestSave2(t *testing.T) {
	cache := New(CacheFile)
	if err := cache.Save([]byte(StringToSave2)); err != nil {
		t.Fatalf("Failed to save: %v\n", err)
	}
}

func TestLoad2(t *testing.T) {

	t.Cleanup(func() {
		_ = os.Remove(CacheFile)
	})

	cache := New(CacheFile)
	loadedData, err := cache.Load()
	if err != nil {
		t.Fatalf("Failed to load: %v\n", err)
	}

	if strings.Compare(string(loadedData), StringToSave2) != 0 {
		t.Fatalf("Failed to compare\n")
	}
}
