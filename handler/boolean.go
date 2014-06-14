package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
	"github.com/zachlatta/boolr/model"
)

func CreateBoolean(w http.ResponseWriter, r *http.Request,
	u *model.User) *AppError {
	if u == nil {
		return ErrNotAuthorized()
	}

	defer r.Body.Close()
	boolean, err := model.NewBoolean(r.Body, u.ID)
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
