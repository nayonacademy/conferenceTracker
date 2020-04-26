package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateSpeaker(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("create speaker")
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speaker := models.Speaker{}
	err = json.Unmarshal(body, &speaker)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speaker.Prepare()
	err = speaker.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speakerCreated, err := speaker.SaveSpeaker(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, speaker.ID))
	responses.JSON(w, http.StatusCreated, speakerCreated)
}
func (server *Server) GetSpeakers(w http.ResponseWriter, r *http.Request){
	speaker := models.Speaker{}
	speakers, err := speaker.FindAllSpeaker(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, speakers)
}
func (server *Server) GetSpeaker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	speaker := models.Speaker{}
	speakerGotten, err := speaker.FindSpeakerByID(server.DB,uint64(uid))
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	responses.JSON(w, http.StatusOK, speakerGotten)
}
func (server *Server) UpdateSpeaker(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	speaker := models.Speaker{}
	err = server.DB.Debug().Model(models.Speaker{}).Where("id = ?", pid).Take(&speaker).Error
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, errors.New("speaker not found"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speakerUpdate := models.Speaker{}
	err = json.Unmarshal(body, &speakerUpdate)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speakerUpdate.Prepare()
	err = speakerUpdate.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	speakerUpdate.ID = speaker.ID
	speakerUpdated, err := speakerUpdate.UpdateSpeaker(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, speakerUpdated)
}
func (server *Server) DeleteSpeaker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	speaker := models.Speaker{}
	err = server.DB.Debug().Model(models.Speaker{}).Where("id = ?", pid).Take(&speaker).Error
	if err != nil{
		responses.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}
	_, err = speaker.DeleteSpeaker(server.DB, pid)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent,"")
}