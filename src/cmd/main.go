package main

import (
	"log"
	"net/http"
	client "s3/src/infrastructure/persistence"
	"s3/src/infrastructure/persistence/storage"
	"s3/src/interfaces/apiserver"
	router "s3/src/interfaces/router"
	"time"
)

func main() {
	api, err := apiserver.NewApiServer(
		&client.S3ClientConfig{
			Address: "minio:9000",
			Access:  "access_123",
			Secret:  "secret_123",
			Token:   "",
			Region:  "us-east-1",
		},
		&storage.S3StorageConfig{
			Bucket:                 "data",
			UploadFilePartSizeMB:   64,
			DownloadFilePartSizeMB: 64,
			UploadLinkLifetime:     5 * time.Minute,
			DownloadLinkLifetime:   5 * time.Minute,
			DownloadFileTimeout:    6 * time.Hour,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Handler:      router.NewRouter(api),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(srv.ListenAndServe())
}
