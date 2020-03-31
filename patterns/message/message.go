package message

// Type stands for chat message types.
type Type string

const (
	MTPing    Type = "ping"
	MTPong    Type = "pong"
	MTMessage Type = "message"
)

// Message stands for message.
type Message struct {
	Type     Type     `json:"type"`
	Channels []string `json:"channels,omitempty"`
	Data     string   `json:"data,omitempty"`
}
