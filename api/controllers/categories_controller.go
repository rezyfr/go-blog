package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rezyfr/go-blog/api/auth"
	"github.com/rezyfr/go-blog/api/models"
	"github.com/rezyfr/go-blog/api/responses"
	"github.com/rezyfr/go-blog/api/utils/formaterror"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	category := models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	category.Prepare()
	err = category.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	categoryCreated, err := category.SaveCategory(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, categoryCreated.ID))
	responses.JSON(w, http.StatusCreated, categoryCreated)
}

func (server *Server) GetCategories(w http.ResponseWriter, r *http.Request) {

	category := models.Category{}

	categories, err := category.FindAllCategories(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, categories)
}

func (server *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid category id given to us?
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// Is this user authenticated?
	_, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the category exist
	category := models.Category{}
	err = server.DB.Debug().Model(models.Category{}).Where("id = ?", id).Take(&category).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = category.DeleteCategory(server.DB, strconv.FormatUint(id, 10))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
