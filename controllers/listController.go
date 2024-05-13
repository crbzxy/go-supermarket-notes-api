package controllers

import (
    "encoding/json"
    "net/http"
    "go-supermarket-notes-api/models"
    "go-supermarket-notes-api/services"
    "go-supermarket-notes-api/utils"
    "github.com/gorilla/mux"
)

// CreateList godoc
// @Summary Create a list
// @Description Create a new list
// @Tags lists
// @Accept  json
// @Produce  json
// @Param list body models.List true "List"
// @Success 200 {object} models.List
// @Failure 400 {object} utils.ErrorResponse
// @Router /lists [post]
func CreateList(w http.ResponseWriter, r *http.Request) {
    var list models.List
    if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.CreateList(list)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, list)
}

// GetLists godoc
// @Summary Get all lists
// @Description Get all lists
// @Tags lists
// @Accept  json
// @Produce  json
// @Success 200 {array} models.List
// @Failure 500 {object} utils.ErrorResponse
// @Router /lists [get]
func GetLists(w http.ResponseWriter, r *http.Request) {
    lists, err := services.GetLists()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, lists)
}

// GetList godoc
// @Summary Get a list
// @Description Get a list by ID
// @Tags lists
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Success 200 {object} models.List
// @Failure 500 {object} utils.ErrorResponse
// @Router /lists/{id} [get]
func GetList(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    list, err := services.GetList(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, list)
}

// UpdateList godoc
// @Summary Update a list
// @Description Update a list by ID
// @Tags lists
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Param list body models.List true "List"
// @Success 200 {object} models.List
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /lists/{id} [put]
func UpdateList(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var list models.List
    if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.UpdateList(params["id"], list)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, list)
}

// DeleteList godoc
// @Summary Delete a list
// @Description Delete a list by ID
// @Tags lists
// @Accept  json
// @Produce  json
// @Param id path string true "List ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} utils.ErrorResponse
// @Router /lists/{id} [delete]
func DeleteList(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := services.DeleteList(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
