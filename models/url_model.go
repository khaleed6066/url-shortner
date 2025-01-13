package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
)

type URL struct {
	Id        string    `json:"id"`
	Ourl      string    `json:"ourl"`
	Surl      string    `json:"surl"`
	CreatedAt time.Time `json:"created_at"`
}

var UrlDB = make(map[string]URL)

func generateUrl(OrginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OrginalUrl))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func CreateUrl(originalUrl string) string {
	shortUrl := generateUrl(originalUrl)
	id := shortUrl
	UrlDB[id] = URL{
		Id:        id,
		Ourl:      originalUrl,
		Surl:      shortUrl,
		CreatedAt: time.Now(),
	}
	return shortUrl
}

func GetUrl(id string) (URL, error) {
	url, ok := UrlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

func GetAllUrls() []URL {

	urls := make([]URL, 0, len(UrlDB)) // Preallocate the slice
	for _, url := range UrlDB {
		urls = append(urls, url)
	}
	return urls
}
