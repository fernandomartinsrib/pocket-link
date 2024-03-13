package api

import (
	"fmt"
	"net/http"
)

type createUrlRequest struct {
	LongUrl string `json:"long_url"`
}

func (server *Server) CreateUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post a new url")
}
