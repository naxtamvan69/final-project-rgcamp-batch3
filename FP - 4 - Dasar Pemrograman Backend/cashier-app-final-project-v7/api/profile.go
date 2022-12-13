package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	filename := path.Join("assets", "images", "img-avatar.png")
	file, err := os.Open(filename)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	fileLocation := path.Join("assets", "images", "img-avatar.png")
	uploadFile, _, err := r.FormFile("file-avatar")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	defer uploadFile.Close()
	targetFile, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	defer targetFile.Close()
	_, err = io.Copy(targetFile, uploadFile)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	api.dashboardView(w, r)
}
