package types

import "fmt"

type MyMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (m MyMessage) String() string {
	return fmt.Sprintf("A message of type %q containing the message %q", m.Type, m.Text)
}
