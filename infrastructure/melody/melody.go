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
    return &MelodyService{m}
}

func (m *MelodyService) SetHandleMessage(handler func(*melody.Session, []byte)) {
	m.HandleMessage(handler)
}

