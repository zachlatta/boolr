package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
	"github.com/zachlatta/boolr/model"
)

func Login(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	defer r.Body.Close()

	var requestUser model.RequestUser
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		return ErrUnmarshalling(err)
	}

	userFromDB, err := database.GetUserByUsername(requestUser.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound(err)
		}
		return ErrDatabase(err)
	}

	err = userFromDB.ComparePassword(requestUser.Password)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return &AppError{err, "invalid password", http.StatusBadRequest}
	} else if err != nil {
		return &AppError{err, "error checking password",
			http.StatusInternalServerError}
	}

	token, err := model.NewToken(userFromDB)
	if err != nil {
		return &AppError{err, "error creating jwt token",
			http.StatusInternalServerError}
	}

	return renderJSON(w, token, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	defer r.Body.Close()
	user, err := model.NewUser(r.Body)
	if err != nil {
		return ErrCreatingModel(err)
	}

	err = database.SaveUser(user)
	if err != nil {
		if err == model.ErrInvalidUserUsername {
			return ErrCreatingModel(err)
		}
		return ErrDatabase(err)
	}

	return renderJSON(w, user, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request, u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	stringID := vars["id"]

	var id int64
	if stringID == "me" {
		id = u.ID
	} else {
		var err error
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			return ErrInvalidID(err)
		}
	}

	if id == u.ID {
		return renderJSON(w, u, http.StatusOK)
	}

	return ErrForbidden()
}

func GetUserBooleans(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	stringID := vars["id"]

	var id int64
	if stringID == "me" {
		id = u.ID
	} else {
		var err error
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			return ErrInvalidID(err)
		}
	}

	if id != u.ID {
		return ErrForbidden()
	}

	booleans, err := database.GetBooleansForUser(u.ID)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, booleans, http.StatusOK)
}
