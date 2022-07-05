package main

import (
	"fmt"
	"github.com/pusher/pusher-http-go/v5"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/webhook", pusherWebhook)
	http.ListenAndServe(":8090", nil)
}

func pusherWebhook(res http.ResponseWriter, req *http.Request) {

	pusherClient := pusher.Client{
		AppID:   "1431028",
		Key:     "addkeyhere",
		Secret:  "addSecrethere",
		Cluster: "ap1",
	}
	body, _ := ioutil.ReadAll(req.Body)

	webhook, err := pusherClient.Webhook(req.Header, body)

	if err == nil {
		for _, event := range webhook.Events {
			fmt.Println("Channel occupied: " + event.Channel)

			switch event.Name {
			case "my-channel":
				AddLogIntoFile("Channel occupied: " + event.Channel)
			case "channel_vacated":
				AddLogIntoFile("Channel occupied: " + event.Channel)
			}
		}
	}

	AddLogIntoFile("test")
}

func AddLogIntoFile(text string) {
    fmt.Println("Channel occupied: " + event.Channel)
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(text + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("done")

}
