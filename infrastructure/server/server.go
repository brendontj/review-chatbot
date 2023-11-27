package server

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/brendontj/review-chatbot/core/usecase"
	"github.com/brendontj/review-chatbot/infrastructure/database"
	"github.com/brendontj/review-chatbot/infrastructure/melody"
	"github.com/brendontj/review-chatbot/infrastructure/template"
	mel "github.com/olahol/melody"
)

type Server struct {
	m *melody.MelodyService
}

func New(m *melody.MelodyService) *Server {
	return &Server{
		m: m,
	}
}

func (svr *Server) setInitialDepencies() {
	database := database.New()
	database.Connect()

	workflowTemplate := template.GenerateWorkflowsTemplate()
	swuc := usecase.NewStartWorkflowUseCase(workflowTemplate, database)

	svr.m.SetHandleMessage(func(s *mel.Session, msg []byte) {
		svr.m.Broadcast(msg)
		regex := regexp.MustCompile("> ").Split(string(msg), 2)
		generatedMsg := swuc.Execute(regex[1])

		if generatedMsg == "" {
			return
		}

		svr.m.Broadcast([]byte(generatedMsg))
	})
}

func (s *Server) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.m.HandleRequest(w, r)
	})

	s.setInitialDepencies()

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

	fmt.Println("Server running on port 8000")
}
