package texter

import "fmt"

type Type string

const (
	STATIC Type = "static"
	SCROLL Type = "scroll"
)

type Alignment string

const (
	LEFT   Alignment = "left"
	RIGHT  Alignment = "right"
	CENTER Alignment = "center"
)

type Message struct {
	Content   string    `json:"content"`
	Type      Type      `json:"type"`
	Time      int       `json:"time"`
	Alignment Alignment `json:"alignment"`
}

func (m Message) String() string {
	return fmt.Sprintf("Message{%#v, %v, %v, %#v}", m.Content, m.Type, m.Time, m.Alignment)
}
