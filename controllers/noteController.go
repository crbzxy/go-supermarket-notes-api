package controllers

import (
    "encoding/json"
    "net/http"
    "go-supermarket-notes-api/models"
    "go-supermarket-notes-api/services"
    "go-supermarket-notes-api/utils"
    "github.com/gorilla/mux"
)

// CreateNote godoc
// @Summary Create a note
// @Description Create a new note
// @Tags notes
// @Accept  json
// @Produce  json
// @Param note body models.Note true "Note Example" 
// @Success 201 {object} models.Note
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notes [post]
func CreateNote(w http.ResponseWriter, r *http.Request) {
    var note models.Note
    if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.CreateNote(note)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, note)
}

// GetNotes godoc
// @Summary Get all notes
// @Description Get all notes
// @Tags notes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Note
// @Failure 500 {object} utils.ErrorResponse
// @Router /notes [get]
func GetNotes(w http.ResponseWriter, r *http.Request) {
    notes, err := services.GetNotes()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, notes)
}

// GetNote godoc
// @Summary Get a note
// @Description Get a note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Success 200 {object} models.Note
// @Failure 500 {object} utils.ErrorResponse
// @Router /notes/{id} [get]
func GetNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    note, err := services.GetNote(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, note)
}

// UpdateNote godoc
// @Summary Update a note
// @Description Update a note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Param note body models.Note true "Note Example"
// @Success 200 {object} models.Note
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /notes/{id} [put]
func UpdateNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var note models.Note
    if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := services.UpdateNote(params["id"], note)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, note)
}

// DeleteNote godoc
// @Summary Delete a note
// @Description Delete a note by ID
// @Tags notes
// @Accept  json
// @Produce  json
// @Param id path string true "Note ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} utils.ErrorResponse
// @Router /notes/{id} [delete]
func DeleteNote(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := services.DeleteNote(params["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
