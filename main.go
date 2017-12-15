package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

// our main function
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/{width}/{height}", GetImage).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//GetImage load lucho image
func GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	width, err := strconv.ParseUint(params["width"], 10, 32)
	height, err := strconv.ParseUint(params["height"], 10, 32)

	file, err := os.Open("LuisJara2013.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, m, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
