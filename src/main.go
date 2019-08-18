package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	flag.Parse()
	fmt.Println("- [OK] Starting server localhost:4040")
	fmt.Println("- [OK] Address: http://localhost:4040")
	/// @todo prevent main thread to stop before children.
	http.HandleFunc("/place", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		fmt.Println("- [OK] New request", r.Host)
		dispatcher(w, r)
	})

	err := http.ListenAndServe("localhost:4040", nil)
	if err != nil {
		fmt.Println(" - [x] Couldn't start server.")
		log.Fatal(err)
	}
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func dispatcher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := Connect().Find(nil)
		var p []Place
		e := q.All(&p)
		bytes, e := json.Marshal(p)

		if e != nil {
			fmt.Println("- [x] Couldn't get places", e)
		} else {
			_, _ = w.Write(bytes)
		}
	}
	if r.Method == http.MethodPost {
		p := Place{}

		decoder := json.NewDecoder(r.Body)
		e := decoder.Decode(&p)

		if e != nil {
			log.Println(" x - [ERROR] ", e)
		}
		saveErr := p.Save()
		if saveErr != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
	}

	if r.Method == "delete" {
		/// @todo call delete on struct.

		p := Place{}
		deleteErr := p.Delete()

		if deleteErr == false {
			w.WriteHeader(500)
			log.Println(deleteErr)
		} else {
			w.WriteHeader(200)
		}
	}

	if r.Method == "put" {
		/// @todo update document

		p := Place{}

		updateErr := p.Update()

		if updateErr == false {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
	}
}

const StateDone = "DONE"
const StateFresh = "FRESH"
const StateInformed = "INFORMED"
const collection = "place"

type Geo struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Place struct {
	ImgUrl string `json:"img_url"`
	State  string `json:"state"`
	Geo    Geo    `json:"geo"`
}

func (p *Place) Done() {
	p.State = StateDone
}

func (p *Place) Fresh() {
	p.State = StateFresh
}

func (p *Place) Informed() {
	p.State = StateInformed
}

func (p *Place) Save() error {
	c := Connect()
	return c.Insert(p)
}
func (p *Place) Delete() bool {
	/// @todo find document
	/// @todo soft delete document if found.
	return false
}

func (p *Place) Update() bool {
	return false
}

func (p *Place) setState(state string) {
	if state == StateFresh || state == StateDone || state == StateInformed {
		p.State = state
	} else {

	}
}

type Entity interface {
	//// Use this to store entity to the datastore.
	Save() error
	//// Use this to delete an entity from the datastore.
	Delete() bool
	//// Use this to update an entity in the data store..
	Update() bool
}
