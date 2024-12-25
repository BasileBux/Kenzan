package cache

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	f "github.com/basileb/kenzan/fonts"
)

type cache struct {
	FontPath string `json:"font_path"`
	FontName string `json:"font_name"`
}

type CachePayload struct {
	FontName string
}

type CacheData struct {
	FontPath string
}

func Cache(payload CachePayload) CacheData {
	cached := CacheData{}
	oldCache, err := readCache()
	if newCache := verifyCache(payload); err != nil || newCache != nil {
		cached.FontPath = newCache.FontPath
	} else {
		cached.FontPath = oldCache.FontPath
	}
	return cached
}

func verifyCache(payload CachePayload) *cache {
	cache, err := readCache()
	if err != nil {
		return payload.updateCache()
	}

	if cache.FontName != payload.FontName {
		return payload.updateCache()
	}
	return nil
}

func (p *CachePayload) updateCache() *cache {
	newCache := &cache{FontPath: f.GetFontPath(p.FontName), FontName: p.FontName}
	if err := writeCache(*newCache); err != nil {
		log.Println("Error: Could not write cache")
	}
	return newCache
}

func readCache() (cache, error) {
	cachePath := getCachePath()
	data, err := os.ReadFile(cachePath)
	if err != nil {
		return cache{}, err
	}

	var currentCache cache
	if err := json.Unmarshal(data, &currentCache); err != nil {
		return cache{}, err
	}

	return currentCache, nil
}

func writeCache(cache cache) error {
	cacheJson, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}
	cachePath := getCachePath()
	err = os.WriteFile(cachePath, cacheJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getCachePath() string {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		panic("Could not find default cache folder")
	}
	cachePath = filepath.Join(cachePath, "kenzan", "cache.json")
	return cachePath
}
