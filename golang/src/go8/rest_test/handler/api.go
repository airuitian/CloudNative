package handler

import (
	"encoding/json"
	"go8/rest_test/model"
	"io/ioutil"
	"log"
	"net/http"
)

func parseJson(r *http.Request, payload interface{}) error {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(bytes, payload)
	if err != nil {
		return err
	}
	return nil
}

func respondJson(w http.ResponseWriter, status int, payload interface{}) {
	result, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Json marshal failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJson(w, status, map[string]string{"error": message})
}

func Create(w http.ResponseWriter, r *http.Request) {
	var api model.Api
	err := parseJson(r, &api)
	if err != nil {
		log.Printf("Json parse failed: %v", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Create", api.TenantId, api.Path+":"+api.Method, api)

	respondJson(w, http.StatusOK, api)

}

func Update(w http.ResponseWriter, r *http.Request) {
	var api model.Api
	err := parseJson(r, &api)
	if err != nil {
		log.Printf("Json parse failed: %v", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Update", api.TenantId, api.Path+":"+api.Method, api)

	respondJson(w, http.StatusOK, api)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	tenantId := r.URL.Query().Get("tenantId")
	if tenantId == "" {
		log.Printf("tenantId is empty")
		respondError(w, http.StatusBadRequest, "tenantId is empty")
		return
	}
	path := r.URL.Query().Get("path")
	if path == "" {
		log.Printf("path is empty")
		respondError(w, http.StatusBadRequest, "path is empty")
		return
	}
	method := r.URL.Query().Get("method")
	if method == "" {
		log.Printf("method is empty")
		respondError(w, http.StatusBadRequest, "method is empty")
		return
	}

	log.Println("Delete", tenantId, path+":"+method)

	respondJson(w, http.StatusOK, "")

}

func Read(w http.ResponseWriter, r *http.Request) {
	tenantId := r.URL.Query().Get("tenantId")
	if tenantId == "" {
		log.Printf("tenantId is empty")
		respondError(w, http.StatusBadRequest, "tenantId is empty")
		return
	}
	path := r.URL.Query().Get("path")
	if path == "" {
		log.Printf("path is empty")
		respondError(w, http.StatusBadRequest, "path is empty")
		return
	}
	method := r.URL.Query().Get("method")
	if method == "" {
		log.Printf("method is empty")
		respondError(w, http.StatusBadRequest, "method is empty")
		return
	}

	log.Println("Read", tenantId, path+":"+method)

	respondJson(w, http.StatusOK, "")
}
