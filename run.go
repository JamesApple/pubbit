package main

import (
	"log"
	"time"

	"github.com/lib/pq"
)

func readMessages(messages chan (string)) {
	for {
		data := <-messages
		log.Println(data)
	}
}

func errorReporter(ev pq.ListenerEventType, err error) {
	if err != nil {
		log.Print(err)
	}
}

func run(config Config) {
	listener := pq.NewListener(config.PostgresURL, 10*time.Second, time.Minute, errorReporter)

	err := listener.Listen("event_change")

	if err != nil {
		log.Fatal(err)
	}

	message := make(chan string, 100)
	go readMessages(message)

	for {
		select {
		case notification := <-listener.Notify:
			message <- notification.Extra
		case <-time.After(90 * time.Second):
			go func() {
				log.Println("Here again")
				err := listener.Ping()
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
	}

}
