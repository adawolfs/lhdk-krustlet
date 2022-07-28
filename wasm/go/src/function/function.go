package function

import (
	"fmt"
)

type Values struct {
	Message string
}

func Message(value string) string {
	message := "This is a message sent from GO: " + value
	fmt.Println(message)
	return message
}
