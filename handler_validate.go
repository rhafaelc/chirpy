package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	restricted := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned_body := filterProfane(params.Body, restricted)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleaned_body,
	})
}

func filterProfane(msg string, profanes map[string]struct{}) string {
	words := strings.Split(msg, " ")
	for i, word := range words {
		_, exists := profanes[strings.ToLower(word)]
		if exists {
			words[i] = "****"
			continue
		}
	}

	return strings.Join(words, " ")
}
