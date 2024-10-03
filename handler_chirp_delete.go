package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/messdev072/goChirpy/internal/auth"
	"github.com/messdev072/goChirpy/internal/database"
)

func (cfg *apiConfig) handleChirpDelete(w http.ResponseWriter, r *http.Request) {

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}
	type response struct {
		Chirp
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't found chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Unauthorized", err)
	}
	cfg.db.DeleteChirpByID(r.Context(), database.DeleteChirpByIDParams{
		ID:     chirpID,
		UserID: userID,
	})
	w.WriteHeader(http.StatusNoContent)
}
