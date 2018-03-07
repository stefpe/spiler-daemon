package main

import (
	"github.com/franela/goreq"
)

type Item struct {
	Payload string
}

func main() {

	//item := Item{Payload: "herzilein"}

	_, _ = goreq.Request{
		Method:    "POST",
		Uri:       "http://127.0.0.1:8000/server.php",
		Body:      "test",
		ShowDebug: true,
	}.Do()

	//log.Println(err.Error())
}
