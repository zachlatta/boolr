package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
	"github.com/zachlatta/boolr/model"
)

func CreateUserBoolean(w http.ResponseWriter, r *http.Request,
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

	defer r.Body.Close()
	boolean, err := model.NewBoolean(r.Body, id)
	if err != nil {
		return ErrCreatingModel(err)
	}

	err = database.SaveBoolean(boolean)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, boolean, http.StatusOK)
}

func GetBoolean(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID(err)
	}

	boolean, err := database.GetBoolean(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound(err)
		}

		return ErrDatabase(err)
	}

	if boolean.UserID != u.ID {
		return ErrForbidden()
	}

	return renderJSON(w, boolean, http.StatusOK)
}

func SwitchBoolean(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return ErrInvalidID(err)
	}

	boolean, err := database.GetBoolean(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound(err)
		}

		return ErrDatabase(err)
	}

	if boolean.UserID != u.ID {
		return ErrForbidden()
	}

	boolean.Switch()

	err = database.SaveBoolean(boolean)
	if err != nil {
		return ErrDatabase(err)
	}

	return renderJSON(w, boolean, http.StatusOK)
}
