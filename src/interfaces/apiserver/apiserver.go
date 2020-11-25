package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	client "s3/src/infrastructure/persistence"
	"s3/src/infrastructure/persistence/storage"
	"strconv"
)

type ApiServer struct {
	storage *storage.S3Storage
}

func NewApiServer(s3ClientConf *client.S3ClientConfig, storageConf *storage.S3StorageConfig) (*ApiServer, error) {
	s3Client, err := client.OpenS3(*s3ClientConf)
	if err != nil {
		return nil, err
	}

	return &ApiServer{storage: storage.NewS3Storage(s3Client, *storageConf)}, nil
}

func (s *ApiServer) UploadObject(w http.ResponseWriter, r *http.Request) {
	size, ok := r.URL.Query()["size"]
	if !ok || len(size) != 1 {
		log.Println("Url Param 'key' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uploadContentLength, err := strconv.ParseInt(size[0], 10, 64)
	if err != nil {
		log.Println("Url Param 'size' is wrong")
		w.WriteHeader(http.StatusBadRequest)
	}
	uuid, ok := r.URL.Query()["uuid"]
	if !ok || len(uuid) != 1 {
		log.Println("Url Param 'uuid' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename, ok := r.URL.Query()["filename"]
	if !ok || len(filename) != 1 {
		log.Println("Url Param 'filename' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := s.storage.GetUploadURL(fmt.Sprintf("%s/%s", uuid[0], filename[0]), uploadContentLength)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
	w.WriteHeader(http.StatusOK)
	return
}

func (s *ApiServer) DownloadObject(w http.ResponseWriter, r *http.Request) {
	uuid, ok := r.URL.Query()["uuid"]
	if !ok || len(uuid) != 1 {
		log.Println("Url Param 'uuid' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename, ok := r.URL.Query()["filename"]
	if !ok || len(filename) != 1 {
		log.Println("Url Param 'filename' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := s.storage.GetDownloadURL(fmt.Sprintf("%s/%s", uuid[0], filename[0]))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
	w.WriteHeader(http.StatusOK)
	return
}

func (s *ApiServer) DeleteObject(w http.ResponseWriter, r *http.Request) {
	uuid, ok := r.URL.Query()["uuid"]
	if !ok || len(uuid) != 1 {
		log.Println("Url Param 'uuid' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename, ok := r.URL.Query()["filename"]
	if !ok || len(filename) != 1 {
		log.Println("Url Param 'filename' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.storage.DeleteFile(fmt.Sprintf("%s/%s", uuid[0], filename[0])); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (s *ApiServer) GetObjectSize(w http.ResponseWriter, r *http.Request) {
	uuid, ok := r.URL.Query()["uuid"]
	if !ok || len(uuid) != 1 {
		log.Println("Url Param 'uuid' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename, ok := r.URL.Query()["filename"]
	if !ok || len(filename) != 1 {
		log.Println("Url Param 'filename' is wrong")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	size, err := s.storage.GetDownloadURL(fmt.Sprintf("%s/%s", uuid[0], filename[0]))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(size)
	w.WriteHeader(http.StatusOK)
	return
}
