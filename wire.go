package main

func InitializeEvent() Event {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}
