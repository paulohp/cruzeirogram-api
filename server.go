package main

import "net/http"
import "log"
import "encoding/json"
import "fmt"
import "io/ioutil"
import "os"
import "github.com/gorilla/mux"
import "./models/instagram"

//import "labix.org/v2/mgo"

func InstagramPicsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get("https://api.instagram.com/v1/tags/fechadocomocruzeiro/media/recent?client_id=48b8d07d69a64138bde5a4278e146f91")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		var data interface{}
		err = json.Unmarshal(contents, &data)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		w.Write([]byte(contents))
		fmt.Printf("%s\n", string(contents))
	}
}

func main() {
	log.Println("Starting server")

	r := mux.NewRouter()
	http.Handle("/api/", r)
	r.HandleFunc("/api/instagram/pics", InstagramPicsHandler).Methods("GET")

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
