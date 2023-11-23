package main 

import ( 
    "github.com/brendontj/review-chatbot/infrastructure/server" 
    "github.com/brendontj/review-chatbot/infrastructure/melody"
)

func main() {
    m := melody.New()   
    server := server.New(m)
    server.Run()
}
