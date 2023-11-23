package melody

import (
    "fmt"
    "github.com/olahol/melody"
)

type MelodyService struct {
    *melody.Melody
}

func New() *MelodyService {
    m := melody.New()

    m.HandleMessage(func(s *melody.Session, msg []byte) {
        fmt.Println("Received message from client: ", string(msg))
        m.Broadcast(msg)
        m.Broadcast([]byte("Hello from server"))
    }) 

    return &MelodyService{m}
}

