package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type People struct {
	People []Person `json:"people`
}

type Person struct {
	Name    string `json:"name`
	Age     int    `json:"age`
	Comment string `json:"comment`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Handler)
	router.HandleFunc("/summary", SummaryHandler)
	http.Handle("/", router)
	http.ListenAndServe("localhost:8000", router)
}

func Handler(writer http.ResponseWriter, req *http.Request) {
	var mapa map[string]int = make(map[string]int)
	mapa["Паша"] = 2 // паша: 2
	mapa["Алина"] = 10
	mapa["Anon"] = -42

	var resp People

	for i, n := range mapa {
		// resp += i + ": " + strconv.Itoa(n) + "\n"
		pers := Person{
			Name:    i,  // pers.Name = i
			Age:     n,  // pers.Age = n
			Comment: "", // pers.Comment = ""
		}

		resp.People = append(resp.People, pers)
	}
	respReady, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	// writer.WriteHeader(http.StatusOK)
	// writer.Write([]byte("Hello, World!"))
	// writer.Write([]byte(resp))
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Content-Disposition", "list")
	writer.Write(respReady)
}

func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	var p People

	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	newResp := "Количество человек: " + strconv.Itoa(len(p.People)) + "\n" + "Имена: "

	for _, n := range p.People {
		newResp += n.Name + ", "
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Write([]byte(newResp))
}
