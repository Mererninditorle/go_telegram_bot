package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const apiUrl = "https://api.telegram.org/" + "bot5623754964:AAGQ0ZOl4db56Itqked3Im3SXTx19-Q03S0"

func main() {
	go UpdateLoop()
	router := mux.NewRouter()
	router.HandleFunc("/api", IndexHandler)
	router.PathPrefix("/").Handler(http.FileServer((http.Dir("./static/"))))
	http.ListenAndServe("localhost:8000", router)
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	var R MainStru

	Ping() /// - Страница посещена

	resp, err := http.Get(apiUrl + "/getMe")

	if err != nil {
		fmt.Println(err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	err = json.Unmarshal(respBody, &R) // заполнили перемнную р
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	R.Result.Abilities = append(R.Result.Abilities, "reacting to commands")

	respReady, err := json.Marshal(R.Result)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(respReady))

	println("НАШИ ДАННЫЕ ПРОЧИТАНЫ! ПОЛНАЯ ГОТОВНОСТЬ У НАС ГОСТИ!")

	w.Write([]byte("Вывод успешно произведён!"))
}

// Обращение//////////////////////////////////
var appeal = ("мой бот")

func UpdateLoop() {
	lastId := 0
	for {
		lastId = Update(lastId)
		time.Sleep(5 * time.Second)
	}
}

func Update(lastId int) int {
	raw, err := http.Get(apiUrl + "/getUpdates?offset=" + strconv.Itoa(lastId))
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(raw.Body)

	var v UpdateResponse
	err = json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}

	if len(v.Result) > 0 {
		ev := v.Result[len(v.Result)-1]
		txt := strings.ToLower(ev.Message.Text)
		if txt == "/privet" {
			txtmsg := SendMessage{
				ChId:                ev.Message.Chat.Id,
				Text:                "Hello!",
				Reply_To_Message_ID: ev.Message.Id,
			}

			bytemsg, _ := json.Marshal(txtmsg)
			_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
			if err != nil {
				fmt.Println(err)
				return lastId
			} else {
				return ev.Id + 1
			}

		}
		/////////////////////////// 22.10.22
		if strings.Split(txt, ", ")[0] == appeal {

			switch strings.Split(strings.Split(txt, ", ")[1], ": ")[0] {
			case "расскажи анекдот":
				{
					return Anek(lastId, ev)
				}
			case "сгенерируй число":
				{
					return RandGen(lastId, ev, txt)
				}
			case "измени обращение на":
				{
					if strings.Contains(txt, ": ") {
						return ChangeName(lastId, ev, txt)
					} else {
						fmt.Println("error")
					}
				}
			}

		}
	}
	return lastId
}

func Anek(lastId int, ev UpdateStruct) int {
	txtmsg := SendMessage{
		ChId: ev.Message.Chat.Id,
		Text: "When goods are getting damaged, are they becoming bads?",
	}

	bytemsg, _ := json.Marshal(txtmsg)
	_, err := http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
	if err != nil {
		fmt.Println(err)
		return lastId
	} else {
		return ev.Id + 1
	}
}

func RandGen(lastId int, ev UpdateStruct, txt string) int {
	fmt.Println("Randgen")
	retotal := strings.Split(txt, "до ")[1]
	s, err := strconv.Atoi(retotal)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	num := strconv.Itoa(rand.Intn(s))
	txtmsg := SendMessage{
		ChId: ev.Message.Chat.Id,
		Text: "Сгенерированное число: " + num,
	}

	bytemsg, _ := json.Marshal(txtmsg)
	_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))

	if err != nil {
		fmt.Println(err)
		return lastId
	} else {
		return ev.Id + 1
	}
}

func ChangeName(lastId int, ev UpdateStruct, txt string) int {
	newap := strings.Split(txt, "измени обращение на: ")
	appeal = newap[1]
	fmt.Println(appeal)
	txtmsg := SendMessage{
		ChId: ev.Message.Chat.Id,
		Text: "Обращение изменено на: " + appeal,
	}

	bytemsg, _ := json.Marshal(txtmsg)
	_, err := http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))

	if err != nil {
		fmt.Println(err)
		return lastId
	} else {
		return ev.Id + 1
	}
}

func Ping() {
	txtmsg := SendMessage{
		ChId: 690215801,
		Text: "Страница посещена",
	}

	bytemsg, _ := json.Marshal(txtmsg)
	_, err := http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
	if err != nil {
		fmt.Println(err)
	}
}
