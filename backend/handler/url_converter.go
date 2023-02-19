package handler

import (
	"errors"
	"log"

	"github.com/ongaaron96/url-shortener/backend/util"
)

const Base62Elements = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type UrlConverter struct {
	counter            *util.Counter
	shortToLongStorage map[string]string
	longToShortStorage map[string]string
}

func NewUrlConverter(counter *util.Counter) *UrlConverter {
	return &UrlConverter{
		counter:            counter,
		shortToLongStorage: make(map[string]string),
		longToShortStorage: make(map[string]string),
	}
}

func (uc *UrlConverter) longToShort(url string) (string, error) {
	if shortUrl, exists := uc.longToShortStorage[url]; exists {
		return shortUrl, nil
	}

	count := uc.counter.GetNextCount()
	urlShort := uc.base10To62(count)
	uc.shortToLongStorage[urlShort] = url
	uc.longToShortStorage[url] = urlShort
	return urlShort, nil
}

func (uc *UrlConverter) base10To62(num uint64) string {
	base62Str := ""
	for num > 0 {
		base62Str += string(Base62Elements[num%62])
		num /= 62
	}

	return base62Str
}

func (uc *UrlConverter) shortToLong(url string) (string, error) {
	if longUrl, exists := uc.shortToLongStorage[url]; !exists {
		log.Println("short url not found in storage")
		return "", errors.New("short url not found in storage")
	} else {
		return longUrl, nil
	}
}
