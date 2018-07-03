package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"abc@gmail.com"},
		From:    []string{"xyz@gmail.com"},
		Content: "Hey there!!",
	}

	failedMessage := FailedMessage{
		ErrorMessage:   "Message intercepted",
		OrignalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
	case receivedMessage := <-msgCh:
		fmt.Println(receivedMessage)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages to read")
	}
}

type Message struct {
	To      []string
	From    []string
	Content string
}

type FailedMessage struct {
	ErrorMessage   string
	OrignalMessage Message
}
