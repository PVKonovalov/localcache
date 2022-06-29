// Created by Pavel Konovalov pkonovalov@orxagrid.com
//
// localcache is a local cache library
//

package localcache

import (
	"hash/crc32"
	"os"
	"path"
)

type LocalCache struct {
	ChecksumIEEE uint32
	PathToCache  string
	IsChanged    bool
}

func New(pathToCache string) *LocalCache {
	return &LocalCache{
		PathToCache:  pathToCache,
		ChecksumIEEE: 0,
		IsChanged:    false,
	}
}

// Save configuration to local cache file
func (s *LocalCache) Save(profileData []byte) error {

	if _, err := os.Stat(s.PathToCache); err == nil {
		checksumIEEE := crc32.ChecksumIEEE(profileData)

		if s.ChecksumIEEE == 0 {
			_, _ = s.Load()
		}

		if checksumIEEE == s.ChecksumIEEE {
			return nil
		} else {
			s.ChecksumIEEE = checksumIEEE
			s.IsChanged = true
		}
	}

	dir, _ := path.Split(s.PathToCache)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(s.PathToCache, profileData, 0644)
}

// Load from local cache
func (s *LocalCache) Load() ([]byte, error) {
	buf, err := os.ReadFile(s.PathToCache)

	if err != nil {
		s.ChecksumIEEE = 0
		return buf, err
	}

	s.ChecksumIEEE = crc32.ChecksumIEEE(buf)

	return buf, err
}
