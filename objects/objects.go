package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Content string
}

func (this *Message) ToJSON() []byte {
	b, err := json.Marshal(this)
	if err != nil {
		return []byte("<error>")
	}
	return b
}

func (this *Message) FromJSON(b []byte) error {
	err := json.Unmarshal(b, this)
	return err
}

func MessageFromJSON(b []byte) *Message {
	m := &Message{}
	m.FromJSON(b)
	return m
}

func main() {

	m1 := &Message {"Al", "Hello"}
	bs := m1.ToJSON()
	fmt.Printf("%s\n", bs)

	m2 := &Message{}
	m2.FromJSON(bs)
	fmt.Printf("%s\n", m2.ToJSON())

	m3 := MessageFromJSON(bs)
	fmt.Printf("%s\n", m3.ToJSON())
}
