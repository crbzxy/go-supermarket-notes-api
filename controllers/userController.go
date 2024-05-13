package controllers

import (
    "encoding/json"
    "net/http"
    "go-supermarket-notes-api/models"
    "go-supermarket-notes-api/services"
    "go-supermarket-notes-api/utils"
    "github.com/gorilla/mux"
)

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.CreateUser(user)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, user)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} utils.ErrorResponse
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := services.GetUsers()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} utils.ErrorResponse
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, err := services.GetUser(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.UpdateUser(params["id"], user)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} utils.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := services.DeleteUser(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
