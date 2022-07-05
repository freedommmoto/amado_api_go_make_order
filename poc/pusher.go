package main

import (
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	// instantiate a client
	pusherClient := pusher.Client{
		AppID:   "1431028",
		Key:     "3b3bee31bf863d7fa58d",
		Secret:  "a2d42c87b123c2b6abb0",
		Cluster: "ap1",
	}

	/**
	used
	"github.com/pusher/pusher-http-go/v5"
	*/

	data := map[string]string{"message": "hello world"}

	// trigger an event on a channel, along with a data payload
	err := pusherClient.Trigger("my-channel", "my_event", data)

	// All trigger methods return an error object, it's worth at least logging this!
	if err != nil {
		panic(err)
	}
}
