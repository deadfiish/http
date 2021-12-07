package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Rone struct {
	Code  int
	Msg   string
	Data  string
}

type Rtwo struct {
	Code  int
	Msg   string
	Data  Data
}

type Data struct {
	Name string
	Age int
}

func YHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	rone := Rone{
		Code: 200,
		Msg: "sucess",
		Data: "your name is " + params["name"],
	}
	ro, err := json.Marshal(rone)
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	fmt.Fprintf(w, string(ro))
}

func ZHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	//fmt.Println(w, r.Form)
	name := r.Form["name"][0]
	a := r.Form["age"][0]
	age, _ := strconv.Atoi(a)

	rtwo := Rtwo{
		Code: 200,
		Msg: "sucess",
	}
	data := Data {
		Name: name,
		Age: age,
	}
	rtwo.Data = data
	rt, err := json.Marshal(rtwo)
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	fmt.Fprintf(w, string(rt))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{name}", YHandler).Methods("GET")
	r.HandleFunc("/users", ZHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}