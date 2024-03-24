package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hits reset to 0"))
	if err != nil {
		return
	}
}
