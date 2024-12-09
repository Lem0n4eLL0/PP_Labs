package sixtask

type Message struct {
	Info string `json:"info"`
}

func NewMessage(info string) *Message {
	return &Message{info}
}
