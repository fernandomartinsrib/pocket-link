package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	db "pocket-link/db/sqlc"
	"pocket-link/pkg/shortener"

	"github.com/rs/zerolog/log"
)

type createUrlRequest struct {
	LongUrl string `json:"long_url"`
}

func (server *Server) CreateUrl(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal().Msgf("cannot connect to db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	var message createUrlRequest
	if err = json.Unmarshal(bodyBytes, &message); err != nil {
		log.Fatal().Msgf("failed to unmarshall body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl, err := shortener.Encode(message.LongUrl)
	if err != nil {
		log.Fatal().Msgf("failed to crete short_url: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	arg := db.CreateUrlParams{
		LongUrl:  message.LongUrl,
		ShortUrl: shortUrl,
	}

	url, err := server.store.CreateUrl(context.Background(), arg)
	if err != nil {
		log.Fatal().Msgf("failed to store url: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(url)
	if err != nil {
		log.Fatal().Msgf("failed to marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		log.Fatal().Msgf("failed to return response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
