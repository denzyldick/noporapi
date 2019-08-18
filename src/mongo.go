package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

func Connect() *mgo.Collection {

	s, e := mgo.Dial("localhost:27017")

	if e != nil {
		log.Print(e)
	}
	fmt.Println("- [OK] Connection has been made to MongoDB ", "localhost:20710")
	s.SetMode(mgo.Monotonic, true);
	d := s.DB("nopor")

	c := d.C("place")
	return c
}
