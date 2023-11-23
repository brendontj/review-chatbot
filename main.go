package main

import (
	"github.com/brendontj/review-chatbot/infrastructure/melody"
	"github.com/brendontj/review-chatbot/infrastructure/server"
)

func main() {
	m := melody.New()
	server := server.New(m)
	server.Run()
	if err := server.Run(); err != nil {
		panic(err)
	}
}
