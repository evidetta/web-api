package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/evidetta/web_api/models"
)

var (
	DB       *sql.DB
	PageSize int
)

type EmptyStruct struct {
}

func Init(db *sql.DB, pageSize int) {
	DB = db
	PageSize = pageSize
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
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	page := 1
	if r.FormValue("page") != "" {
		page, err = strconv.Atoi(r.FormValue("page"))
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
			return
		}
	}

	users, err := models.GetUsers(DB, PageSize, page)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
		return
	}
	defer r.Body.Close()

	if user.Tag == "" {
		respondWithError(w, http.StatusBadRequest, ErrorNoTagFound.Error())
		return
	}

	u, err := models.GetUserByTag(DB, user.Tag)
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			log.Println(err)
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
		return
	}
	defer r.Body.Close()

	if user.Tag == "" {
		respondWithError(w, http.StatusBadRequest, ErrorNoTagFound.Error())
		return
	}

	u, err := models.GetUserByTag(DB, user.Tag)
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	if user.Name != "" {
		u.Name = user.Name
	}

	if user.Address != "" {
		u.Address = user.Address
	}

	emptyTime := time.Time{}
	if user.DateOfBirth != emptyTime {
		u.DateOfBirth = user.DateOfBirth
	}

	err = u.Update(DB)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorInvalidPayload.Error())
		return
	}
	defer r.Body.Close()

	if user.Tag == "" {
		respondWithError(w, http.StatusBadRequest, ErrorNoTagFound.Error())
		return
	}

	u, err := models.GetUserByTag(DB, user.Tag)
	if err != nil {
		if err == models.ErrorEntryNotFound {
			respondWithError(w, http.StatusNotFound, ErrorObjectNotFound.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		}
		return
	}

	err = u.Delete(DB)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, ErrorInternalServerError.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, EmptyStruct{})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotFound, ErrorInvalidEndpoint.Error())
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
