package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/erwinhermanto31/go-upload-image-to-bucket/cloudinary"
	"github.com/erwinhermanto31/go-upload-image-to-bucket/module"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cloudinary.Init()

	http.HandleFunc("/upload_image", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	// declare method for this endpoint
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get email from FormFile
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get file nmae
	filename := handler.Filename

	// save temporary image in this files folder
	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send data to module
	ctx := context.Background()
	err = module.NewUploadImage().Handler(ctx, fileLocation)
	if err != nil {
		return
	}

	// delete temporary email
	e := os.Remove(fileLocation)
	if e != nil {
		log.Fatal(e)

	}

	//return if done
	w.Write([]byte("done"))
}
