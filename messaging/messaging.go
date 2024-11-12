// pkg/messaging/message.go
package messaging

import "fmt"

type Message struct {
	Sender    string
	Recipient string
	Content   string
}

func (m *Message) Send() {
	fmt.Printf("Sending message from %s to %s: %s\n", m.Sender, m.Recipient, m.Content)
}
