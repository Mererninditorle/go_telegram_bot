package main

import (
	"fmt"
	"io"
	"net/http"
)

type NewStruct struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	var tester NewStruct
	tester.Name = "Name1"
	tester.Age = 63

	url := "http://worldtimeapi.org/api/timezone/Europe/Moscow"

	resp, err := http.Get(url)
	respRead, _ := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(respRead))
	}
}

// func NewFunction(name *string) error {

// 	*name = "hello, " + *name

// 	return nil
// }
