package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortner/models"
)

func GetAllShortUrl(w http.ResponseWriter, r *http.Request) {
	urls := models.GetAllUrls()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urls)
}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL_ := models.CreateUrl(data.URL)
	response := struct {
		ShortURL string `json:"short_url,omitempty"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectUrl(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	url, err := models.GetUrl(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
	}

	http.Redirect(w, r, url.Ourl, http.StatusFound)
}
