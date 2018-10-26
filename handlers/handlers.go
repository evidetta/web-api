package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/evidetta/db_migrations/models"
	"github.com/gorilla/mux"
)

var (
	DB *sql.DB
)

type EmptyStruct struct {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
		return
	}
	defer r.Body.Close()

	u, err := models.CreateUser(DB, &user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidParameter.Error())
		return
	}

	user, err := models.GetUserById(DB, int64(id))
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidParameter.Error())
		return
	}

	user, err := models.GetUserById(DB, int64(id))
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
		return
	}
	defer r.Body.Close()

	user.ID = int64(id)

	err = user.Update(DB)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidParameter.Error())
		return
	}

	user, err := models.GetUserById(DB, int64(id))
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	user.ID = int64(id)

	err = user.Delete(DB)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, EmptyStruct{})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
