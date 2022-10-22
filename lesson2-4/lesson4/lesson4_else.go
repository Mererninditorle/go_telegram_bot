package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// const apiURL = "https://api.telegram.org/" + "bot5623754964:AAGQ0ZOl4db56Itqked3Im3SXTx19-Q03S0"

// func main() {
// 	go UpdateLoop()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", IndexHandler)
// 	http.ListenAndServe("localhost:8000", router)
// }

// type UpdateResponse struct {
// 	Ok     bool           `json:"ok"`
// 	Result []UpdateStruct `json:"result"`
// }

// type User struct {
// 	Id       int    `json:"id"`
// 	Is_Bot   bool   `json:"is_bot"`
// 	Username string `json:"username"`
// 	Is_Prem  bool   `json:"is_prem"`
// }

// type Chat struct {
// 	Id   int    `json:"id"`
// 	Type string `json:"type"`
// }

// type Message struct {
// 	Id   int    `json:"message_id"`
// 	User User   `json:"from"`
// 	Date int    `json:"date"`
// 	Chat Chat   `json:"chat"`
// 	Text string `json:"text"`
// }

// type SendMessage struct {
// 	ChId int    `json:"chat_id"`
// 	Text string `json:"txt"`
// }

// type UpdateStruct struct {
// 	Id                  int     `json:"update_id"`           // update_id
// 	Message             Message `json:"message"`             // message
// 	Edited_Message      Message `json:"edited_message"`      // edited_message
// 	Channel_Post        Message `json:"channel_post"`        // channel_post
// 	Edited_Channel_Post Message `json:"edited_channel_post"` // edited_channel_post
// }

// type MainStru struct {
// 	Ok     bool   `json:"ok"`
// 	Result Result `json:"result"`
// }

// type Result struct {
// 	Id                          int    `json:"id"`
// 	Is_bot                      bool   `json:"is_bot"`
// 	First_Name                  string `json:"first_name"`
// 	Username                    string `json:"username"`
// 	Can_Join_Groups             bool   `json:"can_join_groups"`
// 	Can_Read_All_Group_Messages bool   `json:"can_read_all_group_messages"`
// 	Supports_Inline_Requests    bool   `json:"supports_inline_requests"`
// }

// func IndexHandler(w http.ResponseWriter, _ *http.Request) {

// 	var М MainStru

// 	tgtoken := "bot5623754964:AAGQ0ZOl4db56Itqked3Im3SXTx19-Q03S0"
// 	apiUrl := "https://api.telegram.org/" + tgtoken
// 	resp, err := http.Get(apiUrl + "/getMe")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	respBody, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(respBody))

// 	err = json.Unmarshal(respBody, &М)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	respReady, err := json.Marshal(М.Result)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Write([]byte(respReady))
// }

// func UpdateLoop() {
// 	lastId := 0
// 	for {
// 		lastId = Update(lastId)
// 		time.Sleep(5 * time.Second)
// 	}
// }

// func Update(lastId int) int {
// 	raw, err := http.Get(apiURL + "/getUpdates?offset=" + strconv.Itoa(lastId))
// 	if err != nil {
// 		panic(err)
// 	}
// 	body, _ := io.ReadAll(raw.Body)

// 	var v UpdateResponse
// 	err = json.Unmarshal(body, &v)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(v.Result) > 0 {
// 		ev := v.Result[len(v.Result)-1]
// 		txt := ev.Message.Text
// 		if txt == "/privet" {
// 			txtMsg := SendMessage{
// 				ChId: ev.Message.Chat.Id,
// 				Text: "Привет, и тебя туда же",
// 			}

// 			bytemsg, _ := json.Marshal(txtMsg)

// 			_, err = http.Post(apiURL+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
// 			if err != nil {
// 				fmt.Println(err)
// 				return lastId
// 			} else {
// 				return ev.Id + 1
// 			}
// 		}
// 	}
// 	return lastId
// }

// if strings.Contains(strings.ToLower(txt), appeal) {
// 1. Анекдот
// if strings.Contains(strings.ToLower(txt), "расскажи анекдот") {
// 	txtmsg := SendMessage{
// 		ChId: ev.Message.Chat.Id,
// 		Text: "When goods are getting damaged, are they becoming bads?",
// 	}

// 	bytemsg, _ := json.Marshal(txtmsg)
// 	_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
// 	if err != nil {
// 		fmt.Println(err)
// 		return lastId
// 	} else {
// 		return ev.Id + 1
// 	}

// }
// 2. Генерация числа
// if strings.Contains(strings.ToLower(txt), "сгенерируй число") {
// 	re := regexp.MustCompile("[0-9]+")
// 	retotal := re.FindAllString(txt, -1)
// 	fmt.Println(retotal)
// 	s, err := strconv.Atoi(retotal[0])
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(s)
// 	num := strconv.Itoa(rand.Intn(s))
// 	txtmsg := SendMessage{
// 		ChId: ev.Message.Chat.Id,
// 		Text: "Сгенерированное число: " + num,
// 	}

// 	bytemsg, _ := json.Marshal(txtmsg)
// 	_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))
// 	if err != nil {
// 		fmt.Println(err)
// 		return lastId
// 	} else {
// 		return ev.Id + 1
// 	}
// }
// 3. Изменение обращения
// if strings.Contains(txt, "измени обращение на:") {
// 	newap := strings.Split(txt, "измени обращение на: ")
// 	appeal = newap[1]
// 	fmt.Println(appeal)
// 	txtmsg := SendMessage{
// 		ChId: ev.Message.Chat.Id,
// 		Text: "Обращение изменено на: " + appeal,
// 	}

// 	bytemsg, _ := json.Marshal(txtmsg)
// 	_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))

// 	if err != nil {
// 		fmt.Println(err)
// 		return lastId
// 	} else {
// 		return ev.Id + 1
// 	}
// }
// } else {
// 	txtmsg := SendMessage{
// 		ChId: ev.Message.Chat.Id,
// 		Text: "Неверное обращение. Обращайтесь ко мне на: " + appeal,
// 	}

// 	bytemsg, _ := json.Marshal(txtmsg)
// 	_, err = http.Post(apiUrl+"/sendMessage", "application/json", bytes.NewReader(bytemsg))

// 	if err != nil {
// 		fmt.Println(err)
// 		return lastId
// 	} else {
// 		return ev.Id + 1
// 	}
// }
