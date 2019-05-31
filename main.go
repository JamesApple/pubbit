// Package main provides ...
package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("pubbit", "A go pub/sub postgres change event app")
	runcommand = app.Command("run", "Run the listener")
	addcommand = app.Command("add", "Add an event to the table for the example")
	urlarg     = addcommand.Arg("event", "event name to add").String()
)

type Config struct {
	PostgresURL  string
	GCPProjectID string
	GCPTopicID   string
}

func main() {
	config := Config{}

	pgurl, present := os.LookupEnv("PG_URL")
	if !present {
		log.Fatal("No postgres url provided")
	}
	config.PostgresURL = pgurl

	projectid, present := os.LookupEnv("PROJECT_ID")

	if !present {
		log.Fatal("No project id provided")
	}
	config.GCPProjectID = projectid

	topicid, present := os.LookupEnv("TOPIC_ID")
	if !present {
		log.Fatal("No topic provided")
	}
	config.GCPTopicID = topicid

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case runcommand.FullCommand():
		run(config)
	case addcommand.FullCommand():
		add(config)

	}
}
