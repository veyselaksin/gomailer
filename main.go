package main

import (
	"fmt"
	"gomailler/pkg/message"
)

func main() {
	m := message.NewMessage("Test", "Body message.")
	m.To = []string{"veyselaksn@gmail.com"}
	m.AttachFile("test.txt")
	fmt.Println(m)
	m.ToBytes()

}
