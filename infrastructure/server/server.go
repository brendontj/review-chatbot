package server 

import (
    "fmt"
    "net/http"
    "github.com/brendontj/review-chatbot/infrastructure/melody"
)

type Server struct {
    m *melody.MelodyService
}

func New(m *melody.MelodyService) *Server {
    return &Server{
        m: m,
    }
}

func (s *Server) Run() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/index.html")
    })

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        s.m.HandleRequest(w, r)
    })
    
    if err := http.ListenAndServe(":8000", nil); err != nil {
        panic(err)
    }

    fmt.Println("Server running on port 8000")
}   
