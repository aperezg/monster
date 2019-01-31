package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aperezg/monster"
	"github.com/aperezg/monster/storage"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

type api struct {
	repository *storage.MonsterRepository
	router     *mux.Router
}

func NewApi(repository *storage.MonsterRepository) *api {
	a := &api{repository: repository}

	r := mux.NewRouter()
	r.Use(accessControl)
	r.HandleFunc("/monsters", a.createMonster).Methods(http.MethodPost)
	r.HandleFunc("/monsters", a.fetchMonsters).Methods(http.MethodGet)
	r.HandleFunc("/monsters/{ID:[a-zA-Z0-9_]+}", a.fetchMonster).Methods(http.MethodGet)
	r.HandleFunc("/monsters/{ID:[a-zA-Z0-9_]+}", a.updateMonster).Methods(http.MethodPatch)
	r.HandleFunc("/monsters/{ID:[a-zA-Z0-9_]+}", a.deleteMonster).Methods(http.MethodDelete)
	a.router = r

	return a
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		h.ServeHTTP(w, r)
	})
}

func (a *api) createMonster(w http.ResponseWriter, r *http.Request) {
	m := monster.NewMonster()
	if err := jsonapi.UnmarshalPayload(r.Body, m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}

	if err := m.Validate(); err != nil {
		a.marshalError(w, err, http.StatusForbidden)
		return
	}

	// Is a example api and we knows that the only errors if when the monster already exist
	if err := a.repository.CreateMonster(m); err != nil {
		a.marshalError(w, err, http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := jsonapi.MarshalPayload(w, m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}
}

func (a *api) fetchMonsters(w http.ResponseWriter, r *http.Request) {
	m, err := a.repository.FetchMonsters()
	if err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	if err := jsonapi.MarshalPayload(w, m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}
}

func (a *api) fetchMonster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := a.repository.FetchMonsterByID(vars["ID"])
	if err != nil {
		a.marshalError(w, err, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	if err := jsonapi.MarshalPayload(w, m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}
}

func (a *api) updateMonster(w http.ResponseWriter, r *http.Request) {
	// TODO: Provisional struct, google/jsonapi have an open issue about customTypes
	var m = struct {
		ID      string `jsonapi:"primary,monsters"`
		Name    string `jsonapi:"attr,name"`
		Attack  int    `jsonapi:"attr,attack"`
		Defense int    `jsonapi:"attr,defense"`
		Type    string `json:"type" jsonapi:"attr,type"`
	}{}

	if err := jsonapi.UnmarshalPayload(r.Body, &m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}
	monster := &monster.Monster{
		ID:      m.ID,
		Name:    m.Name,
		Attack:  m.Attack,
		Defense: m.Defense,
		Type:    monster.MonsterType(m.Type),
	}

	if err := a.repository.UpdateMonster(m.ID, monster); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *api) deleteMonster(w http.ResponseWriter, r *http.Request) {
	var m = struct {
		ID string `jsonapi:"primary,monsters"`
	}{}
	vars := mux.Vars(r)
	if err := jsonapi.UnmarshalPayload(r.Body, &m); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}

	if vars["ID"] != m.ID {
		err := fmt.Errorf("The request ID(%s) is not the same of request body(%s)", vars["ID"], m.ID)
		a.marshalError(w, err, http.StatusForbidden)
		return
	}

	if err := a.repository.DeleteMonster(m.ID); err != nil {
		a.marshalError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *api) marshalError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	status := strconv.Itoa(code)
	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Detail: err.Error(),
		Status: status,
	}})
}
