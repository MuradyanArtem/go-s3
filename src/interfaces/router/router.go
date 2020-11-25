package router

import (
	"net/http"
	"s3/src/interfaces/apiserver"

	"github.com/gorilla/mux"
)

func NewRouter(s *apiserver.ApiServer) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/file/upload", s.UploadObject).
		Methods("GET").
		Queries(
			"filename", "{filename=[a-zA-Z0-9]+.[a-zA-Z0-9]+}",
			"uuid", "{uuid=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
			"size", "{size=[0-9]+}",
		)

	router.HandleFunc("/file/download", s.DownloadObject).
		Methods("GET").
		Queries(
			"filename", "{filename=[a-zA-Z0-9]+.[a-zA-Z0-9]+}",
			"uuid", "{uuid=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
		)

	router.HandleFunc("/file", s.DeleteObject).
		Methods("DELETE").
		Queries(
			"filename", "{filename=[a-zA-Z0-9]+.[a-zA-Z0-9]+}",
			"uuid", "{uuid=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
		)

	router.HandleFunc("/file/size", s.GetObjectSize).
		Methods("GET").
		Queries(
			"filename", "{filename=[a-zA-Z0-9]+.[a-zA-Z0-9]+}",
			"uuid", "{uuid=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}",
		)

	http.Handle("/", router)
	return router
}
