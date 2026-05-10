package utils

import (
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	COLLY_CACHE_DIR = "./colly_cache"
)

func CreateCollector() *colly.Collector {
	return colly.NewCollector(
		colly.CacheDir(COLLY_CACHE_DIR),
		colly.CacheExpiration(24*time.Hour),
	)
}
